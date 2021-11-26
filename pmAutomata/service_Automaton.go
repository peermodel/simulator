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
	. "cca/debug"
	. "cca/framework"
	. "cca/helpers"
	. "cca/pmModel"
)

// continuously running machine for the service on its in and out containers:
func NewAutomaton_Service(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
	// ------------------------------------------------------------
	// declare local variables (LVS) interface struct:
	// ------------------------------------------------------------
	// Local variables
	type localVariables struct {
		incid  string
		outcid string
		fu     ServiceFunc
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

			// cast <-:
			return (interface{})(newLvs)
		}

		// ------------------------------------------------------------
		// STATE INIT:
		// ------------------------------------------------------------
		a.AddState("init", "init", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			/**/
			m.PrintlnS(TRACE0, TAB, "Peer", ctx.Pid)
			/**/ m.PrintlnS(TRACE0, TAB, "Wiring", ctx.Wid)
			/**/ m.PrintlnS(TRACE0, TAB, "Sid", ctx.Sid)
			// get right service wrapper (from wiring of peer)
			// (nb: go map access returns null valuesof expected type, if key was not found in map)
			p := s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid]
			if "" == p.Id {
				m.SystemError("ill. peer id")
			}
			w := p.Wirings[ctx.Wid]
			if "" == w.Id {
				m.SystemError("ill. wiring id")
			}
			sw := w.ServiceWrappers[ctx.Sid]
			if "" == sw.InCid { // use just any field to check invalidity
				m.UserError("service does not exist")
			}

			// set lvs:
			// nb: convert cids
			incidptr := ConvertCtoM(sw.InCid, ctx.WMNo)
			lvs.incid = *incidptr
			outcidptr := ConvertCtoM(sw.OutCid, ctx.WMNo)
			lvs.outcid = *outcidptr
			/**/ m.PrintlnSS(TRACE0, TAB, "service: incid", lvs.incid, "outcid", lvs.outcid)
			lvs.fu = sw.Fu

			m.CurrentState = "1"

			return OK
		})
		// ------------------------------------------------------------
		// STATE 1:
		// ------------------------------------------------------------
		a.AddState("1", "execute service", func(s *Status, m *Machine) StateRetEnum {
			lvs := m.LocalVariables.(*localVariables)
			ctx := m.Context.(*Context)

			// call service
			/**/
			m.PrintlnSS(TRACE0, TAB, "call service: incid", lvs.incid, "outcid", lvs.outcid)
			lvs.fu(s.MetaContext.(*MetaContext).PeerSpace, ctx.Wfid, ctx.Vars, &s.Scheduler, lvs.incid, lvs.outcid, s.ControllerChannel)

			m.CurrentState = "1"

			// leave for a while
			// @@@ ??? why
			// CAUTION: set new state *before* wait!
			// s.Wait4NoEvent(m)

			return OK
		})
	}
	return a, m
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
