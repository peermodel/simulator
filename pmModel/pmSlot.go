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
// Date: 2015, 2016
////////////////////////////////////////

package pmModel

import (
	. "cca/debug"
	"fmt"
)

// time slot for scheduler
type PMSlot struct {
	// ETTS, ETTL, LTTS, LTTL, WTTS, WTTL, STTL, MEGA_HUNT
	// - nb: can expanded for Peer, WIID, ...
	Type PMSlotTypeEnum
	// if ETTS or ETTL
	Eid string
	// if LTTS, LTTL
	Wiid   string
	LinkNo int
	// if WTTS, WTTL
	// nb: also used for LTTS, LTTL: but for link, wid serves only for docu/trace
	Wid string
	// for entry hunter:
	Pid            string
	RepeatInterval int
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
// private
func newPMSlot(slotType PMSlotTypeEnum, eid string, wiid string, LinkNo int, wid string, pid string, repeatInterval int) *PMSlot {
	slot := new(PMSlot)

	slot.Type = slotType
	slot.Eid = eid
	slot.Wiid = wiid
	slot.LinkNo = LinkNo
	slot.Wid = wid
	slot.Pid = pid
	slot.RepeatInterval = repeatInterval

	return slot
}

// ----------------------------------------
func NewEttsSlot(eid string) *PMSlot {
	return newPMSlot(ETTS, eid, "" /* wiid */, 0 /* LinkNo */, "" /* wid */, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
func NewEttlSlot(eid string) *PMSlot {
	return newPMSlot(ETTL, eid, "" /* wiid */, 0 /* LinkNo */, "" /* wid */, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
// nb: wid serves only for docu/trace
func NewLttsSlot(wid string, wiid string, linkNo int) *PMSlot {
	return newPMSlot(LTTS, "" /* eid */, wiid, linkNo, wid, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
// nb: wid serves only for docu/trace
func NewLttlSlot(wid string, wiid string, linkNo int) *PMSlot {
	return newPMSlot(LTTL, "" /* eid */, wiid, linkNo, wid, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
func NewWttsSlot(wid string) *PMSlot {
	return newPMSlot(WTTS, "" /* eid */, "" /* wiid */, 0 /* LinkNo */, wid, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
func NewWttlSlot(wid string) *PMSlot {
	return newPMSlot(WTTL, "" /* eid */, "" /* wiid */, 0 /* LinkNo */, wid, "" /* pid */, 0 /* repeatInterval */)
}

// ----------------------------------------
// TBD: improve names...
func NewPeerEntriesHuntSlot(pid string, repeatInterval int) *PMSlot {
	return newPMSlot(WIRING_ENTRIES_HUNT, "" /* eid */, "" /* wiid */, 0 /* LinkNo */, "" /* wid */, pid, 0 /* repeatInterval */)
}

////////////////////////////////////////
// methods
////////////////////////////////////////

// ----------------------------------------
// deep copy
func (slot *PMSlot) Copy() interface{} {
	newSlot := new(PMSlot)

	// copy all fields:
	// - Type:
	newSlot.Type = slot.Type
	// - Eid:
	newSlot.Eid = slot.Eid
	// - Wiid:
	newSlot.Wiid = slot.Wiid
	// - LinkNo:
	newSlot.LinkNo = slot.LinkNo
	// - Wid:
	newSlot.Wid = slot.Wid

	return newSlot
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (slot *PMSlot) IsEmpty() bool {
	if nil == slot {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (slot *PMSlot) String() string {
	tmpS := ""
	switch slot.Type {
	case ETTS:
		tmpS = fmt.Sprintf("%s<eid=%s>", slot.Type, slot.Eid)
	case ETTL:
		tmpS = fmt.Sprintf("%s<eid=%s>", slot.Type, slot.Eid)
	case LTTS:
		tmpS = fmt.Sprintf("%s<wid=%s, wiid=%s, linkNo=%d>", slot.Type, slot.Wid, slot.Wiid, slot.LinkNo)
	case LTTL:
		tmpS = fmt.Sprintf("%s<wid=%s, wiid=%s, linkNo=%d>", slot.Type, slot.Wid, slot.Wiid, slot.LinkNo)
	case WTTS:
		tmpS = fmt.Sprintf("%s<wid=%s>", slot.Type, slot.Wid)
	case WTTL:
		tmpS = fmt.Sprintf("%s<wid=%s>", slot.Type, slot.Wid)
	default:
		Panic(fmt.Sprintf("ill. pm slot type = %s", slot.Type))
	}
	return tmpS
}

// ----------------------------------------
func (slot *PMSlot) Print(tab int) {
	s := NBlanksToString("", tab)
	/**/ String2TraceFile(fmt.Sprintf("%s%s", s, slot))
}

// ----------------------------------------
// should not be used!
// - is called by enclosing event type
func (slot *PMSlot) Println(tab int) {
	/**/ slot.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
