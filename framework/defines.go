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
// Date: 2015, 2016
//------------------------------------------------------------
// framework vars, consts, enums
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package framework

import (
	"math/rand"
)

//////////////////////////////////////////////////////////////
// vars
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// counts the number of simulation or model checking runs (paths)
var RUN_COUNT int

//------------------------------------------------------------
// random number generator
// - see: https://golang.cafe/blog/golang-random-number-generator.html
// - Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
// - Use the Seed function to initialize the default Source if different behavior is required for each run.
// - Typically a non-fixed seed should be used, such as time.Now().UnixNano().
// - https://golang.org/pkg/time/#Time.UnixNano
var RANDOM_GENERATOR *rand.Rand = rand.New(rand.NewSource(99))

// static var that signals if the random generator has already been "seeded"
var RANDOM_GENERATOR_WAS_SEEDED_FLAG = false

//////////////////////////////////////////////////////////////
// consts
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// default channel size; caution must be > 0 -> so that sender is not blocked
// - eg if sending a message to own machine at stop etc.
const CHAN_SIZE int = 100

////////////////////////////////////////
// enums
////////////////////////////////////////

//============================================================
// machine message types
//============================================================

//------------------------------------------------------------
type MachineMessageTypeEnum int

//------------------------------------------------------------
const (
	EXCEPTION MachineMessageTypeEnum = iota
	SYSTEM_ERROR
	SYSTEM_INFO
	SYSTEM_WARNING
	USER_ERROR
	USER_WARNING
)

//------------------------------------------------------------
// output is normalized through padding -> try to stay with 11 chars...
func (t MachineMessageTypeEnum) String() string {
	switch t {
	case EXCEPTION:
		return "EXCEPTION"
	case SYSTEM_ERROR:
		return "SYS ERROR"
	case SYSTEM_INFO:
		return "SYS INFO"
	case SYSTEM_WARNING:
		return "SYS WARNING"
	case USER_ERROR:
		return "USR ERROR"
	case USER_WARNING:
		return "USR WARNING"
	default:
		return "ill. machine message type"
	}
}

//============================================================
// machine start types
//============================================================

//------------------------------------------------------------
type MachineStartTypeEnum int

//------------------------------------------------------------
const (
	SYNC MachineStartTypeEnum = iota
	ASYNC
)

//------------------------------------------------------------
func (t MachineStartTypeEnum) String() string {
	switch t {
	case SYNC:
		return "sync"
	case ASYNC:
		return "async"
	default:
		return "ill. machine start type"
	}
}

//////////////////////////////////////////////////////////////
// EOF
////////////////////////////////////////
//////////////////////////////////////////////////////////////
