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
// machine control: data type and logic
// - as the name suggests: needed for each machine in order to control it
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package framework

import (
	. "github.com/peermodel/simulator/controller"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/scheduler"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// machine control structure - used by peer model controler
type MachineControl struct {
	//------------------------------------------------------------
	// this channel is used to send a control signal (ENTER, LEAVE, STOP) to the machine
	CtrlChan chan ChanSig
	//------------------------------------------------------------
	// this channel is used to send the mutex signal ENTER to the machine
	MutexChan chan ChanSig
	//------------------------------------------------------------
	// wait-for condition to get mutex again:
	// - shared pointer with wait-for event structure!!!
	// - there is exactly one event for which a machine waits at a time
	// -- which can also be the NO_EVENT or CLEARED_EVENT
	// CAUTION: do not set to nil
	Condition *Event
	//------------------------------------------------------------
	// pointer back to my machine
	M *Machine
	//------------------------------------------------------------
	// last execution time of the machine
	LastExecutionTime int
	//------------------------------------------------------------
	// how often did machine enter the critical section?
	NCriticalSections int
	//------------------------------------------------------------
	// debug
	// - go routine id
	// - nb: sync sub machines share Gid with caller
	Gid uint64
	//============================================================
	// for model checking (complicated...):
	//============================================================
	//------------------------------------------------------------
	// if a choice point is recovered all wait4 user events are reset to NO_EVENT
	// - because the wake up event could have happened before the CP;
	// - this helps that the machine is waked up, BUT: the wait 4 user state of the
	// - machine will be renetered and again call the wait 4 user event with the current event time
	// - and overwrite the wait 4 condition on the machine control; so the missed event is still a problem...
	// - therefore this flag is needed to indicate that the machine
	// - was waiting for a user event *before* the choice point was created;
	// - caution: reset the flag after the first wake up of machine after the CP!
	UserConditionResettedFlag bool
	// - ALTERNATIVE WAY to UserConditionResettedFlag:
	// skip the checking of the condition in the next wait state of an automaton
	WaitInterruptedByChoicePointFlag bool
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// alloc and init a new machine control
// - with new channels
func NewMachineControl(m *Machine, condition *Event) *MachineControl {
	//------------------------------------------------------------
	// alloc
	mc := new(MachineControl)
	//------------------------------------------------------------
	// init
	// - create channels
	mc.CtrlChan = make(chan ChanSig, CHAN_SIZE)
	mc.MutexChan = make(chan ChanSig, CHAN_SIZE)
	// - set condition
	mc.Condition = condition
	// - set pointer to machine
	mc.M = m
	//------------------------------------------------------------
	// return
	return mc
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// clone for a new Model Checking path run
// - share, renew or copy respective parts
func (mc *MachineControl) Clone4NewRun() *MachineControl {
	//------------------------------------------------------------
	// assertions
	if nil == mc.Condition {
		Panic(fmt.Sprintf("machine %s is in machine controls but has empty condition (g)", mc.M.Key()))
	}
	if mc.Condition.Type == USER_EVENT && nil == mc.Condition.UserEvent {
		Panic(fmt.Sprintf("Clone4NewRun: machine %s has empty user event struct in user condition", mc.M.Key()))
	}
	//------------------------------------------------------------
	// first: copy condition:
	copiedCondition := mc.Condition.Copy()
	//------------------------------------------------------------
	// second: copy M:
	// - M: deep copy of LVs
	// -- nb: fus and automata are shared
	copiedM := mc.M.Clone4NewRun()
	//------------------------------------------------------------
	// alloc
	newMc := NewMachineControl(copiedM, copiedCondition)
	//------------------------------------------------------------
	// treat all fields:
	//............................................................
	// - CtrlChan (new)
	//............................................................
	// - MutexChan (new)
	//............................................................
	// - Condition (set above)
	//............................................................
	// - M (set above)
	//............................................................
	// - LastExecutionTime:
	newMc.LastExecutionTime = mc.LastExecutionTime
	//............................................................
	// - NCriticalSections:
	newMc.NCriticalSections = mc.NCriticalSections
	//------------------------------------------------------------
	// Gid (new; must be set by caller)
	//------------------------------------------------------------
	// return
	return newMc
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// debug
// - generated debug string for modelcheckingtrace
func (mc *MachineControl) KeyInfo() string {
	//------------------------------------------------------------
	// assertion
	if mc.Condition == nil {
		Panic(fmt.Sprintf("machine %s is in machine controls but has empty condition (h)", mc.M.Key()))
	}
	if mc.M == nil {
		Panic(fmt.Sprintf("empty machine %s in machine controls", mc.M.Key()))
	}
	//------------------------------------------------------------
	retS := fmt.Sprintf("%s (state=%s): %s, t=%d, et=%d", mc.M.Key(), mc.M.CurrentState, mc.Condition.String(), CLOCK, EVENT_CLOCK)
	if MODEL_CHECKING_DETAILS1_TRACE.DoTrace() {
		retS = fmt.Sprintf("%s, genCP=%t, resetted=%t", retS, mc.Condition.GenerateChoiceFlag, mc.UserConditionResettedFlag)
	}
	return retS
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
