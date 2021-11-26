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
	. "github.com/peermodel/simulator/pmModel"
)

func NewAutomaton_SpaceUndoWrite(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
	// ------------------------------------------------------------
	// declare local variables (LVS) interface struct:
	// ------------------------------------------------------------
	type localVariables struct {
		// alias: shared with status -> do not deep copy but recompute!
		l *Link
		// alias: shared with status -> do not deep copy but recompute!
		wtx *Tx
		// alias: shared with status -> do not deep copy but recompute!
		c *Container

		eid string
	}

	// --------------------------------------
	// create new automaton:
	// --------------------------------------
	if createAutomatonFlag {
		a = NewAutomaton(automatonName)
	}

	// --------------------------------------
	// create new machine:
	// --------------------------------------
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
			// c
			newLvs.c = GetContainerAlias(theM, s)

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
			// c:
			lvs.c = GetContainerAlias(m, s)
			// return values:
			ctx.RetErr = nil

			m.CurrentState = "1"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 1:
		// ------------------------------------------------------------
		a.AddState("1", "check link type", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			switch lvs.l.Type {
			case ACTION:
				m.CurrentState = "2"
			case GUARD, SERVICE_IN, SERVICE, SERVICE_OUT:
				m.CurrentState = "exit"
				return EXIT
			default:
				m.SystemError("ill. link type")
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 2:
		// ------------------------------------------------------------
		a.AddState("2", "Action link", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			/**/
			m.PrintlnY(TRACE0, TAB, "Txcc", lvs.wtx.Txcc)
			if lvs.wtx.Txcc == PCC {
				m.CurrentState = "3"
			} else {
				m.SystemError("OCC not yet implemented.") // @@@
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 3:
		// ------------------------------------------------------------
		a.AddState("3", "get next entry from Es", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			/**/
			m.PrintlnX(TRACE0, TAB, "Es", ctx.Es)
			if 0 < len(ctx.Es) {
				lvs.eid = ctx.Es[0].Id
				// ----------------------------------------
				// remove e from Es:
				ctx.Es = ctx.Es.RemoveEntry(lvs.eid)
				/**/ m.PrintlnX(TRACE0, TAB, "Es", ctx.Es)
			} else {
				lvs.eid = ""
			}
			m.CurrentState = "5"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 5:
		// ------------------------------------------------------------
		a.AddState("5", "entry found?", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			if "" == lvs.eid {
				m.CurrentState = "exit"
				return EXIT
			} else {
				m.CurrentState = "6"
			}

			return OK
		})
		// ------------------------------------------------------------
		// STATE 6:
		// ------------------------------------------------------------
		a.AddState("6", "remove e from Es", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)

			/**/
			m.PrintlnS(TRACE0, TAB, "eid", lvs.eid)
			for _, e := range lvs.c.Entries {
				if e.Id == lvs.eid {
					// ----------------------------------------
					// remove entry from c:
					/**/
					m.PrintlnX(TRACE0, TAB, "", lvs.c)
					/**/ m.PrintlnS(TRACE0, TAB, "eid", lvs.eid)
					/**/ m.PrintlnS(TRACE0, TAB, "", "remove e from c")

					// nb: entry must exist, otherwise error
					lvs.c.RemoveEntry(lvs.eid)

					/**/
					m.PrintlnX(TRACE0, TAB, "", lvs.c)
					break
				}
			}

			m.CurrentState = "3"

			return OK
		})
	}
	return a, m
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
