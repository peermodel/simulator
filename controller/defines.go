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
//////////////////////////////////////////////////////////////// System: Peer Model State Machine
// Author: eva Kühn
// Date: 2015, 2016
//------------------------------------------------------------
// signal type that can be sent on the controller and mutex channel
// - caution: only ENTER and LEAVE make sense on a mutex channel
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// TBD:
// - extra signal type for mutex channel, namely solely for ENTER and LEAVE signals
//////////////////////////////////////////////////////////////

package controller

//////////////////////////////////////////////////////////////
// enums
//////////////////////////////////////////////////////////////

//============================================================
// signal types
// - used on controller and mutex channels
//============================================================

//------------------------------------------------------------
type SignalTypeEnum int

//------------------------------------------------------------
const (
	//------------------------------------------------------------
	// controller to machine:
	// - machine is permitted to enter critical section
	ENTER SignalTypeEnum = iota
	//------------------------------------------------------------
	// machine to controller:
	// - info that machine leaves critical section
	LEAVE
	//------------------------------------------------------------
	// system/controller to controller:
	// - kick controller to let next machine in
	// - says that there is no machine currently in the critical section
	KICK
	//------------------------------------------------------------
	// machine to controller:
	// - controller shall stop the entire system and tell all machines to stop
	// controller to machine:
	// - machine must stop (only it may finish the execution of its current state)
	STOP
	//------------------------------------------------------------
	// machine to controller:
	// - info that machine has terminated
	TERMINATED
)

//------------------------------------------------------------
func (t SignalTypeEnum) String() string {
	switch t {
	case ENTER:
		return "ENTER"
	case KICK:
		return "KICK"
	case LEAVE:
		return "LEAVE"
	case STOP:
		return "STOP"
	case TERMINATED:
		return "TERMINATED"
	default:
		return "ill. signal type"
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
