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
// "Slot" data type & methods
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package scheduler

import (
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/slotInterface"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// time slot for scheduler
type Slot struct {
	//------------------------------------------------------------
	// slot type:
	// - STTL or USER_SLOT
	Type SlotTypeEnum
	//------------------------------------------------------------
	// system time, when the slot becomes "enabled"
	Time int
	//------------------------------------------------------------
	// model specific slot info
	// - interface
	UserSlot ISlot
}

//////////////////////////////////////////////////////////////
// constructors
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new slot
// - private fu
func allocSlot() *Slot {
	//------------------------------------------------------------
	// alloc
	slot := new(Slot)
	//------------------------------------------------------------
	// return
	return slot
}

//------------------------------------------------------------
// create new sttl slot
func NewSttlSlot(time int) *Slot {
	//------------------------------------------------------------
	// alloc
	slot := allocSlot()
	//------------------------------------------------------------
	// set
	// - nb: user slot is not needed
	slot.Type = STTL
	slot.Time = time
	//------------------------------------------------------------
	// return
	return slot
}

//------------------------------------------------------------
// create new user slot
func NewUserSlot(time int, userSlot ISlot) *Slot {
	//------------------------------------------------------------
	// alloc
	slot := allocSlot()
	//------------------------------------------------------------
	// set
	slot.Type = USER_SLOT
	slot.Time = time
	slot.UserSlot = userSlot
	//------------------------------------------------------------
	// return
	return slot
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// copy
func (slot *Slot) Copy() *Slot {
	//------------------------------------------------------------
	// alloc
	newSlot := allocSlot()
	//------------------------------------------------------------
	// copy all fields
	// - Type
	newSlot.Type = slot.Type
	//------------------------------------------------------------
	// - Time
	newSlot.Time = slot.Time
	//------------------------------------------------------------
	// - UserSlot
	newSlot.UserSlot = slot.UserSlot.Copy().(ISlot)
	//------------------------------------------------------------
	// return
	return newSlot
}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
func (slot *Slot) ToString(ind int) string {
	s := NBlanksToString("", ind)
	s = fmt.Sprintf("%s[time=%d, type=%s", s, slot.Time, slot.Type)
	if USER_SLOT == slot.Type {
		s = fmt.Sprintf("%s, %s", s, slot.UserSlot)
	}
	s = fmt.Sprintf("%s]", s)
	return s
}

//------------------------------------------------------------
func (slot *Slot) Print(ind int) {
	/**/ String2TraceFile(slot.ToString(ind))
}

// --------------------------------------------
func (slot *Slot) Println(ind int) {
	/**/ slot.Print(ind)
	/**/ String2TraceFile("\n")
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
