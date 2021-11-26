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
// "Event" data struct & methods
// - used for both: event and wait4 condition
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// TBD:
// - review debug fus
//////////////////////////////////////////////////////////////

package scheduler

import (
	. "github.com/peermodel/simulator/config"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/eventInterface"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data struct
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// used for both condition and event
type Event struct {
	//------------------------------------------------------------
	// NO_EVENT, TIME_EVENT, USER_EVENT
	Type EventTypeEnum
	//------------------------------------------------------------
	// time when event or condition is issued
	// - caution: use event clock (EVENT_CLOCK)
	IssueEventTime int
	//------------------------------------------------------------
	// only used if Type == TIME_EVENT; this is the time (absolut) that must be reached
	// - caution: this refers to the system clock (CLOCK)
	Wait4Time int
	//------------------------------------------------------------
	// user event data
	UserEvent IEvent
	//------------------------------------------------------------
	// for model checking, if event struct is used as condition:
	// - shall the event-waiting point generate a choice?
	// NB: if a choice point is recovered all wait4 user events are reset to NO_EVENT
	// - because the wake up event could have happened before the CP;
	// - but: NO_EVENT might generate another choice point (depending on automaton spec);
	// - therefore in addition this flag is needed
	GenerateChoiceFlag bool
	//------------------------------------------------------------
	// for model checking
	// - "reset" the event in order to go to the "enter again" part of an interrupter wait state
	// - DEPRECATED
	EventResettedForNextModelCheckingPathFlag bool
	// ALTERNATIVE WAY
	// the waiting for the event was interrupted by the model checker
	// - namely the wait is receovered after a choice point;
	// - in this case the original wait4 condition is used to resume the machine (in the wait state),
	// -- but caution must be taken that the wait does no call leave again but skips the leave
	// -- and continues with the execution;
	// -- the flag indicates that;
	WaitInterruptedByCP_Flag bool
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new event without precondition for machine start
func NewEvent(evt_or_cond_type EventTypeEnum) *Event {
	//------------------------------------------------------------
	// alloc
	evt := new(Event)
	//------------------------------------------------------------
	// init:
	evt.Type = evt_or_cond_type
	// - use big bang time, because also events that happend before now can wake me up
	evt.IssueEventTime = 0
	evt.Wait4Time = SYSTEM_TTL
	// TBD... depends on user model
	// - by default create a cp ... better than none...
	evt.GenerateChoiceFlag = true
	evt.WaitInterruptedByCP_Flag = false
	//------------------------------------------------------------
	// return
	return evt
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (evt *Event) Copy() *Event {
	//------------------------------------------------------------
	// alloc
	newEvt := NewEvent(evt.Type)
	//------------------------------------------------------------
	// copy all fields:
	// - Type: done above
	// - IssueTime:
	newEvt.IssueEventTime = evt.IssueEventTime
	// - TargetSystemTime:
	newEvt.Wait4Time = evt.Wait4Time
	// - UserEvent
	if nil != evt.UserEvent {
		newEvt.UserEvent = evt.UserEvent.Copy().(IEvent)
	}
	//------------------------------------------------------------
	// - ChoicePointFlag
	newEvt.GenerateChoiceFlag = evt.GenerateChoiceFlag
	newEvt.EventResettedForNextModelCheckingPathFlag = evt.EventResettedForNextModelCheckingPathFlag
	//------------------------------------------------------------
	// return
	return newEvt
}

//------------------------------------------------------------
// clear
func (evt *Event) Clear() {
	evt.Type = EMPTY_CONDITION
	evt.IssueEventTime = 0
	evt.Wait4Time = SYSTEM_TTL
	evt.UserEvent = nil
	evt.GenerateChoiceFlag = false
}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// - optionally add: ... ", fulfilled=%t", e.FulfilledFlag
func (evt *Event) String() string {
	//------------------------------------------------------------
	retS := fmt.Sprintf("<%s", evt.Type)
	//------------------------------------------------------------
	switch evt.Type {
	case EMPTY_CONDITION:
	// no extra info
	case NO_EVENT:
	// no extra info
	case TIME_EVENT:
	// no extra info
	case USER_EVENT:
		// assertion
		if nil == evt.UserEvent {
			Panic(fmt.Sprintf("empty user event struct in USER_EVENT"))
		}
		// to string
		retS = fmt.Sprintf("%s,%s", retS, evt.UserEvent.String())
	default:
		Panic(fmt.Sprintf("ill. evt.Type=%s", evt.Type))
	}
	if USER_EVENT == evt.Type {
	}
	//------------------------------------------------------------
	retS = fmt.Sprintf("%s,wait4t=%d", retS, evt.Wait4Time)
	retS = fmt.Sprintf("%s,issueEt=%d>", retS, evt.IssueEventTime)
	return retS
}

//------------------------------------------------------------
func (evt *Event) Print(tab int) {
	/**/ NBlanks2TraceFile(tab)
	/**/ String2TraceFile(fmt.Sprintf("%s", evt.String()))
}

//------------------------------------------------------------
func (evt *Event) Println(tab int) {
	/**/ evt.Print(tab)
	/**/ String2TraceFile("\n")
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
