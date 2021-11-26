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
// "Scheduler" data type & logic
// - insert only slots whose time <= system ttl
//------------------------------------------------------------
//////////////////////////////////////////////////////////////
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package scheduler

import (
	. "cca/config"
	. "cca/debug"
	. "cca/helpers"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// scheduler is a sorted slice of slots
// - sorted by slot time
// - nb: the index is *not* the slot time!
type Scheduler []*Slot

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new scheduler
func NewScheduler() Scheduler {
	return Scheduler{}
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// shallow copy
// - nb: slots can be shared as they are never changed
func (scheduler Scheduler) Copy() Scheduler {
	// create new slice of same length and capacity
	newScheduler := make(Scheduler, len(scheduler), cap(scheduler))
	// copy
	copy(newScheduler, scheduler)
	// return
	return newScheduler
}

//------------------------------------------------------------
// sorted insert
// - cf: https://play.golang.org/p/iFzojVHSpq
func (scheduler Scheduler) SortedInsert(slot *Slot) Scheduler {
	//------------------------------------------------------------
	// search for index where element shall be inserted
	i := 0
	for _, nextSlot := range scheduler {
		if slot.Time <= nextSlot.Time {
			break
		}
		i++
	}
	//------------------------------------------------------------
	// create one dummy place at the end
	scheduler = append(scheduler, nil)
	//------------------------------------------------------------
	// shift everything to the right after the insert place
	// - copy(dest, source)
	copy(scheduler[i+1:], scheduler[i:])
	//------------------------------------------------------------
	// insert element
	scheduler[i] = slot
	//------------------------------------------------------------
	// return
	return scheduler
}

//------------------------------------------------------------
// return next slot provided that it is "ripe" and remove it from scheduler
// - return nil if no slot is there or no slot is ripe
// - in any case return also the (changed) scheduler list
// - private fu
func (scheduler Scheduler) pickFirstSlotIfRipe() (Scheduler, *Slot) {
	//------------------------------------------------------------
	// is there a first slot (ie at index 0) whose time is ripe (ie has reached CLOCK)?
	if 0 < len(scheduler) && CLOCK >= scheduler[0].Time {
		//------------------------------------------------------------
		// get the slot
		slot := scheduler[0]
		//------------------------------------------------------------
		// remove the slot and return everything
		// return append(scheduler[:0], scheduler[1:]...), slot
		// fmt.Println(fmt.Sprintf("BEFORE = %d, AFTER = %d", len(scheduler), len(scheduler[1:]))) // DEBUG
		return scheduler[1:], slot
	}
	//------------------------------------------------------------
	// no slot found
	return scheduler, nil
}

//------------------------------------------------------------
// get and remove next ripe scheduler slot;
// returns:
// - next ripe slot, or nil if there is no ripe slot
// - (changed) scheduler with ripe slot removed
// - flag whether we shall stop, because STTL was reached
func (scheduler Scheduler) GetAndRemoveNextRipeSlot() (*Slot, Scheduler, bool) {
	// fmt.Println(fmt.Sprintf("XBEFORE = %d", len(scheduler))) // DEBUG
	//------------------------------------------------------------
	// return vars
	stopFlag := false
	var slot *Slot = nil
	//------------------------------------------------------------
	// debug
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ scheduler.Println(0)
	}
	//------------------------------------------------------------
	// get and remove first slot, if ripe
	scheduler, slot = scheduler.pickFirstSlotIfRipe()
	//------------------------------------------------------------
	// if there is no ripe slot, just return
	if nil == slot {
		return slot, scheduler, stopFlag
	}
	//------------------------------------------------------------
	// debug
	if SCHEDULER_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("Scheduler: treat %s\n", slot.Type))
	}
	//------------------------------------------------------------
	// there is a ripe slot:
	// - compute value for stop flag
	switch slot.Type {
	// last slot to be treated?
	// - nb: no problem if behind it other slots with same time (= system ttl) exist, as the system will be stopped now!
	// - s.SystemInfo(fmt.Sprintf("SYSTEM TTL %d exceeded", SYSTEM_TTL))
	case STTL:
		stopFlag = true
	case USER_SLOT:
		break
	}
	//------------------------------------------------------------
	return slot, scheduler, stopFlag
}

//////////////////////////////////////////////////////////////
// helpers
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// check that a tts is not infinite
// - if so -> user error
func CheckInfiniteTts(tts int) int {
	if INFINITE == tts {
		UserError("TTS must not be infinite")
	}
	return tts
}

//------------------------------------------------------------
// convert a ttl, if infinite
// - namely, set it to system ttl plus 1
// - e.g., used to avoid that an entry expires at run time end and becomes an exception
func ConvertInfiniteTtl(ttl int) int {
	if INFINITE == ttl {
		return SYSTEM_TTL + 1
	}
	return ttl
}

// //------------------------------------------------------------
// // convert relative time into absolute one (depending on CLOCK);
// // - if infinite -> set it to system ttl
// // - unused
// func ConvertRelativeToAbsoluteTime(relativeTime int) int {
// 	if INFINITE != relativeTime {
// 		return CLOCK + relativeTime
// 	}
// 	return SYSTEM_TTL
// }

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
func (scheduler Scheduler) ToString(ind int) string {
	// indentation
	s := NBlanksToString("", ind)
	// add header
	s = fmt.Sprintf("%sScheduler Slots at t=%d:\n", s, CLOCK)
	// add all slots:
	for i := 0; i < len(scheduler); i++ {
		s = fmt.Sprintf("%s%s\n", s, scheduler[i].ToString(ind+TAB))
	}
	// return
	return s
}

//------------------------------------------------------------
func (scheduler Scheduler) Print(ind int) {
	/**/ String2TraceFile(scheduler.ToString(ind))
}

//------------------------------------------------------------
func (scheduler Scheduler) Println(ind int) {
	/**/ scheduler.Print(ind)
}

//------------------------------------------------------------
// trick: just for code generator to guarantee that this package is used...
func DummySchedulerFu() {
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
