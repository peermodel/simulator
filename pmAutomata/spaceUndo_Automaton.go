//////////////////////////////////////////////////////////////
// Peer Model Tool Chain
// Copyright (C) 2021 Eva Maria Kuehn
//////////////////////////////////////////////////////////////
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
////////////////////////////////////////
// System: Peer Model State Machine
// Author: eva Kühn
// Date: 2015
////////////////////////////////////////

package pmAutomata

import (
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/framework"
	. "github.com/peermodel/simulator/helpers"
	. "github.com/peermodel/simulator/pmModel"
)

func NewAutomaton_SpaceUndo(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
	// ------------------------------------------------------------
	// declare local variables (LVS) interface struct:
	// ------------------------------------------------------------
	type localVariables struct {
		// alias: shared with status -> do not deep copy but recompute!
		l *Link
		// alias: shared with status -> do not deep copy but recompute!
		wtx *Tx

		eid string
		cid string
		x   string
	}

	// --------------------------------------
	// create new automaton:
	// --------------------------------------
	if createAutomatonFlag {
		a = NewAutomaton(automatonName)
	}

	// ------------------------------------------------------------
	// create new machine:
	// ------------------------------------------------------------
	m := NewMachine(a)

	// ------------------------------------------------------------
	// alloc LVS:
	// ------------------------------------------------------------
	m.LocalVariables = new(localVariables)

	if createAutomatonFlag {

		// ------------------------------------------------------------
		// define LVS copy function:
		// ------------------------------------------------------------
		a.LocalVariablesCopyFunction = func(theM *Machine, lvs interface{}) interface{} {
			// cast ->:
			tmpOrigLvs := lvs.(*localVariables)
			// alloc LVS:
			tmpNewLvs := new(localVariables)

			// copy static fields:
			*tmpNewLvs = *tmpOrigLvs

			// copy dynamic fields:

			// cast <-:
			return (interface{})(tmpNewLvs)
		}

		// ------------------------------------------------------------
		// define LVS alias function:
		// ------------------------------------------------------------
		a.CompleteLocalVariablesAliasFunction = func(s *Status, theM *Machine, lvs interface{}) interface{} {
			// cast ->:
			newLvs := lvs.(*localVariables)

			// l
			newLvs.l = GetLinkAlias(theM, s)
			// wtx
			newLvs.wtx = GetWtxAlias(theM, s)

			// cast <-:
			return (interface{})(newLvs)
		}

		// ------------------------------------------------------------
		// STATE INIT:
		// ------------------------------------------------------------
		a.AddState("init", "init", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			// l:
			lvs.l = GetLinkAlias(m, s)
			// wtx:
			lvs.wtx = GetWtxAlias(m, s)
			// return values:
			ctx.RetErr = nil

			m.CurrentState = "2"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 2:
		// ------------------------------------------------------------
		a.AddState("2", "check running state of tx", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			/**/
			m.PrintlnS(TRACE0, TAB, "State", lvs.wtx.State)
			if RUNNING == lvs.wtx.State {
				m.CurrentState = "3"
			} else {
				// do nothing, if tx has already terminated:
				m.CurrentState = "exit"
				return EXIT
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 3:
		// ------------------------------------------------------------
		a.AddState("3", "check Txcc", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			/**/
			m.PrintlnY(TRACE0, TAB, "Txcc", lvs.wtx.Txcc)
			if PCC == lvs.wtx.Txcc {
				m.CurrentState = "6"
			} else {
				m.CurrentState = "4"
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 4:
		// ------------------------------------------------------------
		a.AddState("4_@@@", "OCC", func(s *Status, m *Machine) StateRetEnum {
			// lvs := m.LocalVariables.(*localVariables)

			m.SystemError("OCC not yet implemented.") // @@@
			m.CurrentState = "exit"

			return EXIT
		})
		// ------------------------------------------------------------
		// STATE 5:
		// ------------------------------------------------------------
		a.AddState("5", "set state of wtx to ROLLEDBACK", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			lvs.wtx.State = ROLLEDBACK

			m.CurrentState = "exit"

			return EXIT
		})
		// ------------------------------------------------------------
		// STATE 6:
		// ------------------------------------------------------------
		a.AddState("6", "is there any further cid locked by wtx", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			if 0 < len(lvs.wtx.Pcc.LockedCids) {
				lvs.cid = lvs.wtx.Pcc.LockedCids[0]
				m.CurrentState = "7"
			} else {
				lvs.cid = ""
				m.CurrentState = "5"
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 7:
		// ------------------------------------------------------------
		a.AddState("7", "remove cid from wtx's pcc locked cids; set c to cid-container", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			// ----------------------------------------
			// remove cid from wtx's pcc locked cids: it is on position 0:
			len := len(lvs.wtx.Pcc.LockedCids)
			/**/ m.PrintlnI(TRACE0, TAB, "len of locked cids", len)
			// @@@ why is this necessary? len was asserted above...?! otherwise out of bounds if [1 : 0]
			if len > 1 {
				lvs.wtx.Pcc.LockedCids = lvs.wtx.Pcc.LockedCids[1 : len-1]
			} else {
				lvs.wtx.Pcc.LockedCids = Strings{}
			}
			// ----------------------------------------
			// signal container change event:
			s.MetaContext.(*MetaContext).PeerSpace.ContainerChangeEvent(lvs.cid)
			/**/ m.PrintlnS(TRACE0, TAB, "cid", lvs.cid)

			m.CurrentState = "8"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 8:
		// ------------------------------------------------------------
		a.AddState("8", "is there an entry e in c that is locked by wtx?", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			c := s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
			if nil == c {
				m.SystemError("ill. container")
			}

			lvs.eid = ""
			for _, e := range c.Entries {
				// set eid:
				lvs.eid = e.Id
				// check for READ lock of wtx:
				if 0 < e.RLocks[ctx.Wtxid] {
					m.CurrentState = "9"
					return OK
				}
				// check for DELETE lock of wtx:
				if 0 < e.DLocks[ctx.Wtxid] {
					m.CurrentState = "10"
					return OK
				}
				// check for WRITE lock of wtx:
				if 0 < e.WLocks[ctx.Wtxid] {
					m.CurrentState = "11"
					return OK
				}
			}
			m.CurrentState = "6"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 9:
		// ------------------------------------------------------------
		a.AddState("9", "undo READ lock", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			/**/
			m.PrintlnS(TRACE0, TAB, "", lvs.cid)
			/**/ m.PrintlnS(TRACE0, TAB, "eid", lvs.eid)
			/**/ m.PrintlnS(TRACE0, TAB, "", "remove 1 READ-lock on e")

			c := s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
			if nil == c {
				m.SystemError("ill. container")
			}

			// remove lock and remember all entries on which a lock was removed im tmpEs:
			tmpEs := Entries{}
			for _, e := range c.Entries {
				if e.Id == lvs.eid {
					e.RemoveLock(READ, ctx.Wtxid)
					tmpEs = append(tmpEs, e)
					/**/ m.PrintlnX(TRACE0, TAB, "e", &e)
					/**/ m.PrintlnX(TRACE0, TAB, "", c)
					break
				}
			}

			// check if entry's ttl has expired and wrap it into exception if so, writing the exception into the right poc
			// do it in extra loop, because c is changed by check ettl
			for _, e := range tmpEs {
				s.MetaContext.(*MetaContext).PeerSpace.EntryHunter(e.Id, &s.Scheduler)
			}

			m.CurrentState = "8"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 10:
		// ------------------------------------------------------------
		a.AddState("10", "undo DELETE lock", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			/**/
			m.PrintlnS(TRACE0, TAB, "", lvs.cid)
			/**/ m.PrintlnS(TRACE0, TAB, "eid", lvs.eid)
			/**/ m.PrintlnS(TRACE0, TAB, "", "remove 1 DELETE-lock on e")

			c := s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
			if nil == c {
				m.SystemError("ill. container")
			}

			// remove lock and collect all entries on which a lock was removed
			tmpEs := Entries{}
			for _, e := range c.Entries {
				if e.Id == lvs.eid {
					e.RemoveLock(DELETE, ctx.Wtxid)
					tmpEs = append(tmpEs, e)
					/**/ m.PrintlnX(TRACE0, TAB, "e", &e)
					/**/ m.PrintlnX(TRACE0, TAB, "", c)
					break
				}
			}

			// check if entry's ttl has expired and wrap it into exception if so, writing the exception into the right poc
			// do it in extra loop, because c is changed by check ettl
			// @@@ probably this is redundant, because every delete locked entry must also be read locked
			//     -> so the ettl treatment was done already in state 9 above
			for _, e := range tmpEs {
				s.MetaContext.(*MetaContext).PeerSpace.EntryHunter(e.Id, &s.Scheduler)
			}

			m.CurrentState = "8"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 11:
		// ------------------------------------------------------------
		a.AddState("11", "undo WRITE lock", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			/**/
			m.PrintlnS(TRACE0, TAB, "", lvs.cid)
			/**/ m.PrintlnS(TRACE0, TAB, "eid", lvs.eid)
			/**/ m.PrintlnS(TRACE0, TAB, "", "remove e")

			c := s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
			if nil == c {
				m.SystemError("ill. container")
			}

			// nb: entry must exist, otherwise error
			c.RemoveEntry(lvs.eid)

			/**/
			m.PrintlnX(TRACE0, TAB, "", c)

			m.CurrentState = "8"

			return OK
		})

	}

	return a, m
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
