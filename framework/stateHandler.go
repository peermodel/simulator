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
// Author: eva Kühn
// Date: 2015
//------------------------------------------------------------
// data type "StateHandler"
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package framework

//////////////////////////////////////////////////////////////
// enum
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
type StateRetEnum int

//------------------------------------------------------------
// return value of an automaton state handler executed by a machine
const (
	//------------------------------------------------------------
	// ok end of a non-exit state
	// - nb: still the critical section is hold by the machine
	OK StateRetEnum = iota
	//------------------------------------------------------------
	// ok end of automaton, ie end of last state that sets to current state to "exit"
	// - nb: still the critical section is hold by the machine
	EXIT
	//------------------------------------------------------------
	// state execution was stopped
	// - ie STOP signal received while waiting4 an event, ie waiting to enter the critical section again
	// - nb: the critical section is not hold any more by the machine
	STOPPED
)

//------------------------------------------------------------
func (t StateRetEnum) String() string {
	switch t {
	case OK:
		return "OK"
	case EXIT:
		return "EXIT"
	case STOPPED:
		return "STOPPED"
	default:
		return "ill. state ret type"
	}
}

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// data type state handler is a function that takes a status pointer as arg
type StateHandler func(*Status, *Machine) StateRetEnum

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
