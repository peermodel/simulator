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
//////////////////////////////////////////////////////////////
// System: Peer Model State Machine
// Copyright: eva Kuehn
//////////////////////////////////////////////////////////////

package framework

// . "github.com/peermodel/simulator/contextInterface"
// . "github.com/peermodel/simulator/debug"
// . "github.com/peermodel/simulator/framework"
// . "github.com/peermodel/simulator/helpers"
// . "github.com/peermodel/simulator/pmModel"
// . "github.com/peermodel/simulator/scheduler"
// "errors"
// "fmt"

//////////////////////////////////////////////////////////////
// data types
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// automaton specification
// - defines the program (= code) to be executed by machines (= instances) of this automaton
type Automaton struct {
	//------------------------------------------------------------
	// automaton name
	// - must be unique
	Name string
	//------------------------------------------------------------
	// local variables (LVS) interface
	// - function that deep copies all local vars
	LocalVariablesCopyFunction LVSCopyHandler
	// - function that resolves alias vars contained in LVS
	CompleteLocalVariablesAliasFunction LVSAliasHandler
	//------------------------------------------------------------
	// code to be performed per state
	// - key = state id
	StateHandlers map[string]StateHandler
	//------------------------------------------------------------
	// debug info: map with all states' comments
	// - one comment per state
	// - key = state id
	StateComments map[string]string
	//------------------------------------------------------------
	// debug: for statistics only
	// - count how often a machine executing this automation did enter a critical section
	// -- caution: concurrent access is avoided as machine must only inc the counter if it is in its CS
	// -- only relevant for async machines
	nCriticalSections int
	// - count how many machines were used (ie started and ended) for this automaton
	// -- caution: concurrent access is avoided as machine can only inc the counter if
	// --- its calling async top parent machine is in its CS
	nMachinesUsedCount int
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new automaton
func NewAutomaton(name string) *Automaton {
	//------------------------------------------------------------
	// alloc
	a := new(Automaton)
	//------------------------------------------------------------
	// init:
	a.Name = name
	a.StateHandlers = make(map[string]StateHandler)
	a.StateComments = make(map[string]string)
	//------------------------------------------------------------
	// return
	return a
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// add state handler and comment
// - key is the state-id
func (a *Automaton) AddState(stateId, stateComment string, stateHandlerFunction StateHandler) {
	a.StateHandlers[stateId] = stateHandlerFunction
	a.StateComments[stateId] = stateComment
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
