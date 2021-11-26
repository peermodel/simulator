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
// consts, vars and enums for the scheduler and for events/conditions
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package scheduler

import (
	. "github.com/peermodel/simulator/config"
	"fmt"
)

//////////////////////////////////////////////////////////////
// consts
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// INFINITE
// - is set to system time to live (= SYSTEM_TTL)
// - alternatively it can be set to MAX_INT
const INFINITE int = SYSTEM_TTL

//////////////////////////////////////////////////////////////
// vars
//////////////////////////////////////////////////////////////

//============================================================
// different clocks
//============================================================
//------------------------------------------------------------
// CLOCK: system time
// - user's view of the time
// - controls TTLs and TTSs
var CLOCK int

//------------------------------------------------------------
// EVENT CLOCK:
var EVENT_CLOCK int

//////////////////////////////////////////////////////////////
// enums
//////////////////////////////////////////////////////////////

//============================================================
// slot type enum
//============================================================
//------------------------------------------------------------
type SlotTypeEnum int

//------------------------------------------------------------
// slot types:
// - system ttl slot (i.e. slot for system end)
// - user slot
const (
	STTL SlotTypeEnum = iota
	USER_SLOT
)

//------------------------------------------------------------
func (t SlotTypeEnum) String() string {
	switch t {
	case STTL:
		return "STTL"
	case USER_SLOT:
		return "USER_SLOT"
	default:
		return fmt.Sprintf("ill. slot type = %s", t)
	}
}

//============================================================
// event type enum
//============================================================
//------------------------------------------------------------
type EventTypeEnum int

//------------------------------------------------------------
const (
	// machine is not waiting (used at machine start, and by model checking to recover a wait state between leave & enter")
	EMPTY_CONDITION EventTypeEnum = iota
	// machine waits for "no event": used to give up critical section and then enter it again
	NO_EVENT
	// machine waits for clock to approach a given time
	TIME_EVENT
	// machine waits for a user event
	USER_EVENT
)

//------------------------------------------------------------
func (t EventTypeEnum) String() string {
	switch t {
	case EMPTY_CONDITION:
		return "EMPTY_CONDITION"
	case NO_EVENT:
		return "NO_EVENT"
	case TIME_EVENT:
		return "TIME_EVENT"
	case USER_EVENT:
		return "USER_EVENT"
	default:
		return "ill. event type"
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
