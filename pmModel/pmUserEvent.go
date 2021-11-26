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
// Code Review: 2021 Apr, Eva Maria Kuehn
//------------------------------------------------------------
//////////////////////////////////////////////////////////////

package pmModel

import (
	. "github.com/peermodel/simulator/debug"
	"fmt"
)

//////////////////////////////////////////////////////////////
// type and enum
//////////////////////////////////////////////////////////////

//============================================================
// pm event type
//============================================================

//------------------------------------------------------------
// used for condition and event
type PMUserEvent struct {
	// for which container event we shall wait?
	// - see enum CONTAINER_EVENT
	Type PMUserEventTypeEnum
	// cid of container that is concerned
	Cid string
	// debug
	// - if condition type == CONTAINER_EVENT -> for which event to wait
	EntryType string
}

//============================================================
// enum
//============================================================

//------------------------------------------------------------
type PMUserEventTypeEnum int

//------------------------------------------------------------
// TBD: container change could be improved:
// - what can be now become fulfillable? read, take or none?
const (
	CONTAINER_CHANGE_EVENT PMUserEventTypeEnum = iota
	TBD_EVENT                                  // TBD
)

//------------------------------------------------------------
func (t PMUserEventTypeEnum) String() string {
	switch t {
	case CONTAINER_CHANGE_EVENT:
		return "CONTAINER_EVENT"
	default:
		return "ill. event type"
	}
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

// ----------------------------------------
func NewPMUserEvent(evtType PMUserEventTypeEnum, cid string, entryType string) *PMUserEvent {
	//------------------------------------------------------------
	// alloc
	e := new(PMUserEvent)
	//------------------------------------------------------------
	// init
	e.Type = evtType
	e.Cid = cid
	e.EntryType = entryType
	//------------------------------------------------------------
	// return
	return e
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (evt *PMUserEvent) Copy() interface{} {
	//------------------------------------------------------------
	// alloc
	newEvt := new(PMUserEvent)
	//------------------------------------------------------------
	// copy all fields:
	// - Type:
	newEvt.Type = evt.Type
	// - Cid:
	newEvt.Cid = evt.Cid
	// - EntryType:
	newEvt.EntryType = evt.EntryType
	//------------------------------------------------------------
	// return
	return newEvt
}

// ----------------------------------------
func (e *PMUserEvent) IsEmpty() bool {
	if nil == e {
		return true
	} else {
		return false
	}
}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
func (e *PMUserEvent) String() string {
	tmpS := ""
	if CONTAINER_CHANGE_EVENT == e.Type {
		tmpS = fmt.Sprintf("%s,%s,%s", e.Type, e.Cid, e.EntryType)
	}
	return tmpS
}

//------------------------------------------------------------
// just print the fields seperated by comma, strating with ", "
// - is called by enclosing event type
// -- tab is not needed
func (e *PMUserEvent) Print(tab int) {
	if CONTAINER_CHANGE_EVENT == e.Type {
		/**/ String2TraceFile(fmt.Sprintf(", %s", e.Type))
		/**/ String2TraceFile(fmt.Sprintf(", %s", e.Cid))
		/**/ String2TraceFile(fmt.Sprintf(", %s", e.EntryType))
	}
}

//------------------------------------------------------------
// should not be used!
// - is called by enclosing event type
func (e *PMUserEvent) Println(tab int) {
	/**/ e.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
