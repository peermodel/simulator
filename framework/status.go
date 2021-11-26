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
// Date: 2015
//------------------------------------------------------------
// status data type and logic
// - status is shared between all concurrent machines
// - the controller fu controls the execution of all re-entrant concurrent machines in a sequential way
// -- by coordinating configurable forms of interleaving of them
// - there is exactly one critical section coordinated by status
// - every concurrent (= asynchronous) machine must
// -- enter the critical section before it can start
// -- leave the critical section when it stops or interrupts its execution (ie on wait or exit)
// status maintains:
// - the global system state (= context)
// - all communication channels
// - info about all machines
// - the system scheduler
// status manages:
// - wait and wake up of machines
// -- incl. enter and leaving of the exactly one critical section
// --- which is needed for a machine in order to be allowed to execute
// -- incl. stopping and resuming of machines
// - main control of the system
// -- treating all signals sent between machines and controller
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// - TBD: no code review yet for print/debug fus
//////////////////////////////////////////////////////////////

package framework

import (
	. "github.com/peermodel/simulator/config"
	. "github.com/peermodel/simulator/controller"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/eventInterface"
	. "github.com/peermodel/simulator/helpers"
	. "github.com/peermodel/simulator/metaContextInterface"
	. "github.com/peermodel/simulator/scheduler"
	"fmt"
	"runtime"
	"sync"
	"syscall"
)

//////////////////////////////////////////////////////////////
// global vars
//////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// fu that initializes the use case
// nb: app specific
type InitAppUseCaseFuType func() *Status

//------------------------------------------------------------
// reflects the exactly one shared status on which all machines operate
// - nb: the used, unbounded channels perform a *synchonous* handshake
type Status struct {
	//------------------------------------------------------------
	// function that refers to app program
	// - inits the entire use case
	// - needed for simulation runs that must start from scratch every time
	InitAppUseCaseFu InitAppUseCaseFuType
	//------------------------------------------------------------
	// mutex to secure Status against concurrent access
	// - nb: needed, because otherwise the 'fatal error: concurrent map read and map write' might occur
	// - nb: Lock and Unlock are for write, and RLock and RUnlock for read access
	// caution: use it for MCS andAtomata access
	StatusMutex sync.RWMutex
	//------------------------------------------------------------
	// context shared with all machines
	// - holds the meta runtime model and data
	// - is an interface
	MetaContext IMetaContext
	//------------------------------------------------------------
	// controller's channel
	// - here a machine must send its requests to the controller
	ControllerChannel chan ChanSig
	//------------------------------------------------------------
	// map with control structs for *all* machines (ie both SYNC and ASYNC)
	// - the key in the map denotes the machine and is created with the function Key() of machine
	// - caution: use its mutex (above) to lock and unlock it when accessing it!
	// -- be careful to keep the lock/unlcok bracket and not to return between lock and unlock...
	// -- TBD: a bit weak code; mutex hidden in machine controls data struct would be nicer;
	// --- but complicated, because of many range loops;
	// -- TBD: usage of sync.Map instead (albeit probably slower)
	MachineControls map[string]*MachineControl
	//------------------------------------------------------------
	// map with automata specifications
	// - the key in the map is the name of the automaton
	Automata map[string]*Automaton
	//------------------------------------------------------------
	// for statistics only:
	// - archive of terminated machine names
	DoneMachines Strings
	//------------------------------------------------------------
	// scheduler
	// - i.e. list of time slots indicating work to be done by the system
	Scheduler Scheduler
	//------------------------------------------------------------
	// for debug only:
	// - which machine is currently in the critical section
	// - to assert that machine is allowed to leave the critical section
	// - "" ... no machine is currently in the critical section
	CurMachineKey string
	//------------------------------------------------------------
	// some counters for statistics
	// - critical section counter
	CriticalSectionCounter int
	// - critical section counter
	MachineTerminationCounter int
	//------------------------------------------------------------
	// trick:
	// - needed only by the code generator (written in Java) that transforms visio automata into go code
	// -- so that fmt include is needed by every automaton -> in init state just sprintf machine name here
	DummyString string
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new status that is shared among all machines!
// - args: end time when to stop the system, and the meta context of the use case
func NewStatus(systemTtl int, metaContext IMetaContext) *Status {
	//------------------------------------------------------------
	// alloc
	s := new(Status)
	//------------------------------------------------------------
	// init
	s.MetaContext = metaContext
	//............................................................
	s.ControllerChannel = make(chan ChanSig, CHAN_SIZE)
	//............................................................
	s.MachineControls = make(map[string]*MachineControl)
	//............................................................
	s.Automata = make(map[string]*Automaton)
	//............................................................
	s.DoneMachines = Strings{}
	//............................................................
	s.CurMachineKey = ""
	//............................................................
	// - create scheduler
	s.Scheduler = NewScheduler()
	// - create and append new slot for system end; its time is the system ttl
	s.Scheduler = append(s.Scheduler, NewSttlSlot(systemTtl))
	//............................................................
	s.DummyString = ""
	//------------------------------------------------------------
	// return
	return s
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create a clone of status for a new  Model Checking path run
// - share everything that can be shared
// - renew *all* channels incl. those of the MCS
// - copy the rest
func (s *Status) Clone4NewRun(systemTtl int) *Status {
	//------------------------------------------------------------
	// create new status
	// - meta context is copied below
	newS := NewStatus(systemTtl, nil /* meta context */)
	//------------------------------------------------------------
	// copy all required fields
	//............................................................
	// - HelpFu (share)
	newS.InitAppUseCaseFu = s.InitAppUseCaseFu
	//............................................................
	// - StatusMutex (new)
	//............................................................
	// - MetaContext (deep copy)
	// -- eg for Peer Model this is the peer space and the transactions
	newS.MetaContext = s.MetaContext.Copy().(IMetaContext)
	//............................................................
	// - ControllerChan (new)
	//............................................................
	// - MachineControls (share, renew or copy respective parts)
	s.StatusMutex.RLock() // LOCK FOR READ //
	for key, mc := range s.MachineControls {
		newS.MachineControls[key] = mc.Clone4NewRun()
	}
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//............................................................
	// - Automata (share!!)
	s.StatusMutex.RLock() // LOCK FOR READ //
	for aName, _ := range s.Automata {
		newS.Automata[aName] = s.Automata[aName]
	}
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//............................................................
	// - DoneMachineControls
	newS.DoneMachines = s.DoneMachines.Copy()
	//............................................................
	// - Scheduler
	newS.Scheduler = s.Scheduler.Copy()
	//............................................................
	// - CurMachineKey
	newS.CurMachineKey = s.CurMachineKey
	//------------------------------------------------------------
	// - DummyString
	newS.DummyString = s.DummyString
	//------------------------------------------------------------
	// return
	return newS
}

//------------------------------------------------------------
// a machine wants to enter *the" critical section (called by machine (from its thread))
// - wait for ENTER signal from controller (at machine's mutex channel)
// - also check other signals from controller (at machine's ctrl channel)
//............................................................
// returns
// - true if critical section could be entered
// - false otherwise;
// !!! CAUTION: caller must check retval and exit the machine if false (because stop was sent to machine) !!!
//............................................................
func (s *Status) EnterCriticalSection(m *Machine, mc *MachineControl) bool {
	//------------------------------------------------------------
	// debug
	// String2TraceFile(fmt.Sprintf("try to EnterCriticalSection: m = %s, GID=%d\n", m.Key(), GetGoRoutineID())) // DEBUG
	//------------------------------------------------------------
	// wait for ENTER message on the machine's mutex channel (that allows the machine to proceed) and return true
	// - if meanwhile a STOP message arrives on the machine's ctrl channel, then return false
	// - if any other message arrives on either channel, then system error
	for {
		select {
		case mutexMsg := <-mc.MutexChan:
			//------------------------------------------------------------
			// message from mutex channel received: check if it is ENTER
			// - if ok, set also debug info  which machine is now in the critical section
			//------------------------------------------------------------
			switch mutexMsg.Sig {
			case ENTER:
				//------------------------------------------------------------
				// ENTER
				//------------------------------------------------------------
				// assert that no machine is in the critical section
				if s.CurMachineKey != "" {
					/**/ m.SystemError(fmt.Sprintf("EnterCriticalSection: m = '%s' gets enter permit for non empty critical section, occupied by m = %s", m.Key(), s.CurMachineKey))
				}
				s.CurMachineKey = m.Key()
				// fmt.Println(fmt.Sprintf("ENTER Critical Section (m = %s)", m.Key())) // DEBUG
				//------------------------------------------------------------
				// statistics
				// - overall counter
				s.CriticalSectionCounter++
				// - per automaton counter
				// -- nb: no lock needed as there is exactly only one machine in the CS and therefore executing
				m.A.nCriticalSections++
				//------------------------------------------------------------
				// debug
				// String2TraceFile(fmt.Sprintf(" EnterCriticalSection OK: m = %s, GID=%d\n", m.Key(), GetGoRoutineID())) // DEBUG
				//------------------------------------------------------------
				// machine has now the critical section permit
				// - return ok so that calling machine handler can proceed to execute
				return true
			default:
				//------------------------------------------------------------
				// DEFAULT
				//------------------------------------------------------------
				m.SystemError(fmt.Sprintf("EnterCriticalSection: ill. mutex msg = %s", mutexMsg.Sig))
			}

		case ctrlMsg := <-mc.CtrlChan:
			//------------------------------------------------------------
			// message from ctrl channel received: check if it is STOP
			//------------------------------------------------------------
			switch ctrlMsg.Sig {
			case STOP:
				//------------------------------------------------------------
				// debug
				// fmt.Println(fmt.Sprintf("EnterCriticalSection: STOP received while waiting for ENTER (m = %s)", m.Key())) // DEBUG
				//------------------------------------------------------------
				// debug
				// String2TraceFile(fmt.Sprintf(" EnterCriticalSection NOT OK: m = %s, GID=%d\n", m.Key(), GetGoRoutineID())) // DEBUG
				//------------------------------------------------------------
				// terminate this machine: ie  tell handler that there is nothing to be done any more
				// - it will just set the next state to "exit" and terminate the current state
				return false
			default:
				m.SystemError("EnterCriticalSection: ill. ctrl msg")
			}
		}
	}
}

//------------------------------------------------------------
// general wait-for-an-event function for a machine:
// - called via any Wait4 func via the machine's handler code
// - if wait returns true, the machine will continue on exactly the place where it issued the wait
// CAUTION: in model checking mode a choice point may be generated for M when it was between leave and enter
// - in this case the fulfillment of the orig condition will cause the m to be waked up, to execute again
// -- the wait state and thus to reset the wait condition which will cause an event that happened before
// -- the CP to be lost, because the newly updated condition will have a too young issue time...
// - solution:
// -- (1) if there is still a condition set, do not leave
// -- (2) clear the condition upon ending this fu
// CAUTION: must not be called by syncronous machine! (TBD why?)
//............................................................
// algorithm:
// - the machine gives up its critical section
// -- so that any other machine can get it (maybe itself again)
// -- because as for now the machine cannot proceed as long as the wait4-condition is not fulfilled
// - indicates for which event it waits
// - and immediately tries to get the critical section again
// -- which is possible if the condition is meanwhile filfilled
// NB: the function blocks until the above algorithm is successful !!!
//............................................................
// args:
// - machine
// - wait4 event (nb: every event must be bounded with time), consisting of:
// -- event type
// -- targetSystemTime (= absolute system time; if infinite -> take system ttl)
// -- optional: user event description (it eventType == USER_EVENT)
// - choicePointFlag: for model checking (if choice point shall be created at this point)
//............................................................
// returns:
// - true if successful
// - false, if irregular event happened, eg STOP was sent to machine;
// -- CAUTION: caller must exit the machine in this case!
//............................................................
// (private fu)
//............................................................
func (s *Status) wait(m *Machine, eventType EventTypeEnum, targetSystemTime int, userEvent IEvent, choicePointFlag bool) bool {
	//------------------------------------------------------------
	// assertion
	// - wait4 is only allowed in async machine
	if m.StartType == SYNC {
		m.SystemError(fmt.Sprintf("wait4 function called in sync machine, automaton = %s, eventType = %s", m.A.Name, eventType))
	}
	//------------------------------------------------------------
	// increment event clock
	EVENT_CLOCK++
	//------------------------------------------------------------
	s.StatusMutex.RLock() // LOCK FOR READ //
	//------------------------------------------------------------
	// get machine control from the machine (from status)
	mc := s.MachineControls[m.Key()]
	//------------------------------------------------------------
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	// assertion
	if mc.Condition == nil {
		s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (i)", mc.M.Key()))
	}
	//------------------------------------------------------------
	// skip the leave and enter phases... see comment above
	// TBD: && mc.Condition.Type != CONDITION_DONE
	if VERIFICATION_MODE == MODEL_CHECKING && mc.Condition.WaitInterruptedByCP_Flag {
		//------------------------------------------------------------
		// debug
		if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
			String2TraceFile(fmt.Sprintf("SKIP LEAVE IN WAIT STATE (m=%s)", mc.M.Key())) // DEBUG
		}
		//------------------------------------------------------------
		// clear the condition
		mc.Condition.Clear()
		//------------------------------------------------------------
		// clear the flag
		mc.Condition.WaitInterruptedByCP_Flag = false
		//------------------------------------------------------------
		// done, ie do not leave and enter again, as machine was between enter and leave when CP was created
		return true
	}
	//------------------------------------------------------------
	// update with new wait-4-condition
	// - the event condition is used to wake up the machine (depending on execution mode)
	// - CAUTION: do it before leaving!
	// - set:
	// -- condition type
	mc.Condition.Type = eventType
	// -- issue event time
	mc.Condition.IssueEventTime = EVENT_CLOCK
	// -- convert infinite to system ttl
	if INFINITE == targetSystemTime {
		mc.Condition.Wait4Time = SYSTEM_TTL
	} else {
		mc.Condition.Wait4Time = targetSystemTime
	}
	// user event struct
	mc.Condition.UserEvent = userEvent
	// -- shall this wait state generate a choice for model checking?
	mc.Condition.GenerateChoiceFlag = choicePointFlag
	//------------------------------------------------------------
	// debug
	if EVENT_CONDITION_TRACE.DoTrace() { // DEBUG
		// String2TraceFile(fmt.Sprintf("    Wait for %s, entryType=%s, cid=%s, targetTime=%d (current EVENT_CLOCK=%d, CLOCK=%d)\n", evtType, entryType, cid, targetSystemTime, EVENT_CLOCK, CLOCK))  // DEBUG
		String2TraceFile(fmt.Sprintf("    Wait for %s", eventType)) // DEBUG
		if nil != userEvent {                                       // DEBUG
			/**/ userEvent.Print(0 /* tab - unused */) // DEBUG
		} // DEBUG
		/**/ String2TraceFile(fmt.Sprintf(", targetTime=%d (current EVENT_CLOCK=%d, CLOCK=%d)\n", targetSystemTime, EVENT_CLOCK, CLOCK)) // DEBUG
	} // DEBUG
	//------------------------------------------------------------
	// LEAVE
	//------------------------------------------------------------
	// debug:
	// - fmt.Println(fmt.Sprintf("m = %s: wait: (1) call LeaveCritical Section: eventType = %s, targetSystemTime = %d", m.Key(), eventType, targetSystemTime)) // DEBUG
	//-  m.SystemInfo(fmt.Sprintf("waiting %s calls leave critical section, GID=%d", m.Key(), GetGoRoutineID())) // DEBUG
	//------------------------------------------------------------
	// give free *the* critical section (maintained by status)
	// - so another machine might work inbetween
	s.LeaveCriticalSection(m)
	//------------------------------------------------------------
	// ENTER
	//------------------------------------------------------------
	// debug:
	// - fmt.Println(fmt.Sprintf("m = %s: wait: (2) call EnterCriticalSection Section: eventType = %s, targetSystemTime = %d", m.Key(), eventType, targetSystemTime)) // DEBUG
	//------------------------------------------------------------
	// try to get again *the* critical section (maintained by status)
	// - nb: blocks !!!
	ret := s.EnterCriticalSection(m, mc)
	//------------------------------------------------------------
	// debug
	// - String2TraceFile(fmt.Sprintf("%s waked up\n",m.key(), )) // DEBUG
	//------------------------------------------------------------
	// return flag whether critical section has been successfully entered
	// - nb: returns false, if a stop signal was received while waiting
	//------------------------------------------------------------
	s.StatusMutex.RLock() // LOCK FOR READ //
	//------------------------------------------------------------
	// get machine control from the machine (from status)
	mc = s.MachineControls[m.Key()]
	//------------------------------------------------------------
	// clear condition *here*, ie if the enter fu has been passed
	// - in both cases: successfully or not;
	// - ie condition is nil when machine is in the critical section or when it was stopped;
	mc.Condition.Clear()
	//------------------------------------------------------------
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	return ret
}

//------------------------------------------------------------
// machine gives up its critical section to allow another one to proceed (maybe it is itself again)
// - there is no event it must wait for
// - CAUTION: if return false -> caller must exit the machine (because stop was sent to machine)
func (s *Status) Wait4NoEvent(m *Machine, choicePointFlag bool) bool {
	return s.wait(m, NO_EVENT, 0 /* targetSystemTime */, nil /* user event */, choicePointFlag)
}

//------------------------------------------------------------
// machine waits until a time has been reached
// - targetSystemTime = absolute system time
// - CAUTION: if return false -> caller must exit the machine (because stop was sent to machine)
func (s *Status) Wait4TimeEvent(m *Machine, targetSystemTime int, choicePointFlag bool) bool {
	return s.wait(m, TIME_EVENT, targetSystemTime, nil /* user event */, choicePointFlag)
}

//------------------------------------------------------------
// machine waits for an event
// - bounded with time = absolute system time
// - CAUTION: if return false -> caller must exit the machine (because stop was sent to machine)
func (s *Status) Wait4UserEvent(m *Machine, targetSystemTime int, userEvent IEvent, choicePointFlag bool) bool {
	//------------------------------------------------------------
	// assertion
	if nil == userEvent {
		s.SystemError(fmt.Sprintf("Wait4UserEvent: empty user event struct (m=%s)", m.Key()))
	}
	//------------------------------------------------------------
	// wait
	return s.wait(m, USER_EVENT, targetSystemTime, userEvent, choicePointFlag)
}

//------------------------------------------------------------
// leave critical section
// - current machine tells status that it leaves now the critical section
// -- nb: ie it must be the machine to be in the critical section now
// -- nb: there can be at most one machine in the critical section; irrelevant which machine; just give the critical section free
// -- but for quality assurance pass m as arg to assert that it is the right mache
func (s *Status) LeaveCriticalSection(m *Machine) {
	//------------------------------------------------------------
	// debug
	// s.PrintMyGoRoutineId("LEAVE", m.Key()) // DEBUG
	//------------------------------------------------------------
	// send leave to the controller channel
	s.ControllerChannel <- NewChanSig(LEAVE, SENDER_IS_MACHINE, m.Key())
}

//------------------------------------------------------------
// start the controller:
// - trick: send controller an initial "KICK", so that it allows the first process to enter its critical section
func (s *Status) Run() {
	//------------------------------------------------------------
	// reset statistics variables for next run
	// - global ones
	s.CriticalSectionCounter = 0
	s.MachineTerminationCounter = 0
	//	// - per automaton ones
	//	// -- makes no sense
	//	s.StatusMutex.RLock() // LOCK FOR READ //
	//	for _, a := range s.Automata {
	//		a.nCriticalSections = 0
	//		a.nMachinesUsedCount = 0
	//	}
	//	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	// the trick: send KICK
	s.ControllerChannel <- NewChanSig(KICK, SENDER_IS_SYSTEM, "Run" /* msg */)
	//------------------------------------------------------------
	// start controller
	s.Controller()
}

//------------------------------------------------------------
// clean up a terminated machine
// - delete the machine control for this machine
func (s *Status) cleanUpTerminatedMachine(mKey string, mStartType MachineStartTypeEnum) {
	//------------------------------------------------------------
	// debug
	// s.PrintMyGoRoutineId("Clean", mKey) // DEBUG
	if CONTROLLER_TRACE.DoTrace() { // DEBUG
		/**/ String2TraceFile(fmt.Sprintf("CLEAN UP:  machine %s\n", mKey)) // DEBUG
	} // DEBUG
	//------------------------------------------------------------
	s.StatusMutex.Lock() // LOCK FOR WRITE //
	//------------------------------------------------------------
	// debug statistics
	// - remember machine key in done list
	s.DoneMachines = append(s.DoneMachines, mKey)
	// - how many machines were used in total
	s.MachineTerminationCounter++
	//------------------------------------------------------------
	// TBD: close channels
	// - nb: machine has receiver role on both channels...
	// -- and channel closing should be done by sender, because sending to a closed channel causes panic
	// close(s.MachineControls[mKey].CtrlChan)
	// close(s.MachineControls[mKey].MutexChan)
	//------------------------------------------------------------
	// remove machine control from MCS map
	delete(s.MachineControls, mKey)
	//------------------------------------------------------------
	s.StatusMutex.Unlock() // UNLOCK FOR WRITE //
	//------------------------------------------------------------
}

//------------------------------------------------------------
// create and add a new machine control for machine's key
// - done when the machine is started
// - set its wait event to given condition
// nb: each (sync and async) machine has its own machine control with own channels
func (s *Status) CreateAndAddMachineControl(m *Machine, cond *Event) *MachineControl {
	//------------------------------------------------------------
	// assertions
	if nil == cond {
		s.SystemError(fmt.Sprintf("can't add machine control with empty condition (m=%s)", m.Key()))
	}
	if USER_EVENT == cond.Type && nil == cond.UserEvent {
		s.SystemError(fmt.Sprintf("can't add machine control with user condition with empty user event struct (m=%s)", m.Key()))
	}
	//------------------------------------------------------------
	// create new machine control struct
	mc := NewMachineControl(m, cond)
	//------------------------------------------------------------
	// add to machine controls of current status
	s.StatusMutex.Lock() // LOCK FOR WRITE //
	s.MachineControls[m.Key()] = mc
	s.StatusMutex.Unlock() // UNLOCK FOR WRITE //
	//------------------------------------------------------------
	// return
	return mc
}

//------------------------------------------------------------
// add an automaton specification
// - returns true, if automaton was added; else false (if it existed already in the map)
func (s *Status) AddAutomaton(a *Automaton) bool {
	//------------------------------------------------------------
	// check if exists
	s.StatusMutex.RLock() // LOCK FOR READ //
	_, existsAlreadyFlag := s.Automata[a.Name]
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	// add, if not yet there
	s.StatusMutex.Lock() // LOCK FOR WRITE //
	if !existsAlreadyFlag {
		s.Automata[a.Name] = a
	}
	s.StatusMutex.Unlock() // UNLOCK FOR WRITE //
	//------------------------------------------------------------
	// return
	return !existsAlreadyFlag
}

//------------------------------------------------------------
// get ctrl channel of this machine
func (s *Status) GetCtrlChannel(m *Machine) chan ChanSig {
	//------------------------------------------------------------
	// get machine controls  of current status
	s.StatusMutex.RLock() // LOCK FOR READ //
	mc := s.MachineControls[m.Key()]
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	// assertion that not empty
	if nil == mc {
		m.SystemError("ill. machine control")
	}
	//------------------------------------------------------------
	// return control channel
	return mc.CtrlChan
}

//------------------------------------------------------------
// does automaton exist already?
func (s *Status) CheckAutomatonExistence(automatonName string) (*Automaton, bool) {
	s.StatusMutex.RLock() // LOCK FOR READ //
	a, flag := s.Automata[automatonName]
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	return a, flag
}

//------------------------------------------------------------
// select next machine for execution depending on execution mode
// - returns machine key of selected machine, or "" if no machine could be selected
// - private fu
func (s *Status) selectNextMachine() string {
	//------------------------------------------------------------
	// debug
	thisFuNm := "selectNextMachine"
	//------------------------------------------------------------
	// init ret val
	nextMachineKey := ""
	//------------------------------------------------------------
	// depending on execution mode: select machine
	switch EXECUTION_MODE {
	case MIN_ISSUE_TIME:
		//------------------------------------------------------------
		// MIN_ISSUE_TIME:
		// - select a machine whose issue time (= condition time) is smallest... very simple...
		// - nb: wait4 condition is not checked (contradics the choice point creation for model checking mode)
		// - TBD: is this mode useful? maybe better (= RANDOM and unfair) take any machine whose condition is fulfilled
		// - nb: limited randomness is given in that map access via range is not deterministic
		//------------------------------------------------------------
		// - init with not yet reached event time
		curMinIssueTime := EVENT_CLOCK + 1
		//------------------------------------------------------------
		s.StatusMutex.RLock() // LOCK FOR READ //
		for key, mc := range s.MachineControls {
			//------------------------------------------------------------
			// assertion
			if mc.Condition == nil {
				s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (a)", key))
			}
			//------------------------------------------------------------
			// debug
			if CONTROLLER_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: key = %s, mc.Condition.ConditionTime = %d\n", thisFuNm, key, mc.Condition.IssueEventTime)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// check issue time to be actual minimum; remember machine key if so
			if mc.Condition.IssueEventTime < curMinIssueTime {
				nextMachineKey = key
				curMinIssueTime = mc.Condition.IssueEventTime
			}
		}
		s.StatusMutex.RUnlock() // UNLOCK FOR READ //

	case FAIRNESS:
		//------------------------------------------------------------
		// FAIRNESS:
		// - select the one whose condition is fulfilled and whose last try to execute is furthest behind
		//------------------------------------------------------------
		// - init with not yet reached event time :-)
		minLastTryTime := CLOCK + 1
		//------------------------------------------------------------
		// caution: use lock, otherwise fatal error: concurrent map iteration and map write might sometimes occur here
		s.StatusMutex.RLock() // LOCK FOR READ //
		for key, mc := range s.MachineControls {
			//------------------------------------------------------------
			// assertion
			if mc.Condition == nil {
				s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (b)", key))
			}
			//------------------------------------------------------------
			// debug
			if CONTROLLER_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: key = %s, condition time = %d\n", thisFuNm, key, mc.Condition.IssueEventTime)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// check the condition and continue with next machine if not fulfilled
			// - nb: for model checking mode: user events were reset when recovering from a choice point
			if !s.ConditionIsFulfilled(mc) {
				continue
			}
			//------------------------------------------------------------
			// debug
			// fmt.Println(fmt.Sprintf("@@@ %s: SIMULATION: ConditionIsFulfilled is fulfilled for %s", thisFuNm, key)) // DEBUG
			//------------------------------------------------------------
			// check if machine's last try to execute is older than actual minimum one
			if mc.LastExecutionTime <= minLastTryTime {
				nextMachineKey = key
				minLastTryTime = mc.LastExecutionTime
			}
		}
		s.StatusMutex.RUnlock()
		// fmt.Println(fmt.Sprintf("%s: SIMULATION: nextMachineKey = %s (last try t = %d, t = %d)", thisFuNm, nextMachineKey, minLastTryTime, CLOCK)) // DEBUG

	case MIN_ISSUE_TIME_AND_CONDITION_FULFILLED:
		//------------------------------------------------------------
		// MIN_ISSUE_TIME_AND_CONDITION_FULFILLED:
		//------------------------------------------------------------
		// select the one whose issue time (= condition time) is smallest and condition fulfilled
		// - nb: limited randomness is given in that map access via range is not deterministic
		//------------------------------------------------------------
		// - init with not yet reached event time :-)
		curMinIssueTime := EVENT_CLOCK + 1
		//------------------------------------------------------------
		// caution: use lock, otherwise fatal error: concurrent map iteration and map write might sometimes occur here
		s.StatusMutex.RLock() // LOCK FOR READ //
		for key, mc := range s.MachineControls {
			//------------------------------------------------------------
			// assertion
			if mc.Condition == nil {
				s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (c)", key))
			}
			//------------------------------------------------------------
			// debug
			if CONTROLLER_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: key = %s, condition time = %d\n", thisFuNm, key, mc.Condition.IssueEventTime)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// check check condition and issue time to be actual minimum; remember machine key if so
			if mc.Condition.IssueEventTime < curMinIssueTime {
				fulfilledFlag := s.ConditionIsFulfilled(mc)
				if fulfilledFlag {
					nextMachineKey = key
					curMinIssueTime = mc.Condition.IssueEventTime
				}
			}
		}
		s.StatusMutex.RUnlock() // UNLOCK FOR READ //
		// fmt.Println(fmt.Sprintf("%s: FAIRNESS: nextMachineKey = %s", thisFuNm, nextMachineKey)) // DEBUG

	default:
		//------------------------------------------------------------
		// DEFAULT:
		//------------------------------------------------------------
		Panic("ill. EXECUTION_MODE")
	}
	//------------------------------------------------------------
	// return
	return nextMachineKey
}

//------------------------------------------------------------
// controller
// - controls next run
// - handles all signals on its control channel in an 'endless' loop and
// -- treate them depending on verficiation and execution modi of the system
// - when it finishes, all machine threads are stopped (except of the controller thread)
func (s *Status) Controller() {
	//------------------------------------------------------------
	// local vars:
	var stopFlag bool
	var stopMsg string
	var nextMachineKey string
	curCpDepth := 0
	// - only for MC
	prevSpaceUpdates := SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT
	//------------------------------------------------------------
	// debug:
	thisFuNm := "Controller"    // DEBUG
	if REPLAY_TRACE.DoTrace() { // DEBUG
		/**/ String2ReplayTraceFile(fmt.Sprintf("\nRUN = %d\n", RUN_COUNT)) // DEBUG
	} // DEBUG

	//////////////////////////////////////////////////////////////
	// controller loop start:
	//////////////////////////////////////////////////////////////
controllerLoop:
	for {
		//------------------------------------------------------------
		// debug:
		if CONTROLLER_TRACE.DoTrace() { // DEBUG
			/**/ String2TraceFile(fmt.Sprintf("%s: Start Next Controler Loop\n", thisFuNm)) // DEBUG
		} // DEBUG
		//------------------------------------------------------------
		// reset key for the next machine that will get the permission to ENTER:
		nextMachineKey = ""
		//------------------------------------------------------------
		// wait for control message from any machine:
		ctrlMsg := <-s.ControllerChannel
		//------------------------------------------------------------
		// debug:
		if CONTROLLER_TRACE.DoTrace() { // DEBUG
			/********/ String2TraceFile(fmt.Sprintf("%s: ctrlMsg = %s, sender=%s\n", thisFuNm, ctrlMsg, ctrlMsg.SenderMachineKey)) // DEBUG
		} // DEBUG
		//------------------------------------------------------------
		// treat received signal:
		switch ctrlMsg.Sig {
		case TERMINATED:
			//============================================================
			// TERMINATED:
			//============================================================
			//------------------------------------------------------------
			// debug
			// - s.SystemInfo(fmt.Sprintf("%s: TERMINATED received from %s", thisFuNm, ctrlMsg.SenderMachineKey)) // DEBUG
			//------------------------------------------------------------
			// info of a machine that it has terminated received
			// - nothing to be done here
			// - machine has already cleaned up its MC
			// - continue with for loop
			continue

		case STOP:
			//============================================================
			// STOP:
			//============================================================
			//------------------------------------------------------------
			// stop all machines and end the controller loop
			stopFlag = true
			stopMsg = "USER"
			break controllerLoop

		case KICK:
			//============================================================
			// KICK:
			//============================================================
			// info from system/controller that critical section is free and a
			// - next machine can be let in (same case as leave)
			//------------------------------------------------------------
			// only for MC
			prevSpaceUpdates = SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT
			//------------------------------------------------------------
			fallthrough

		case LEAVE:
			//============================================================
			// LEAVE:
			//============================================================
			// a machine signals the LEAVE of its critical section
			// - select a machine as next one to ENTER the critical section
			// -- depending on verification and execution mode
			//------------------------------------------------------------
			// check if sig is leave, because kick signal is also treated here
			if ctrlMsg.Sig == LEAVE {
				//------------------------------------------------------------
				// assert that it is sent from a machine
				if SENDER_IS_SYSTEM == ctrlMsg.WhoIsSender {
					/**/ s.SystemError(fmt.Sprintf("system sent leave instead of machine"))
				}
				// //------------------------------------------------------------
				// // TBD: assert that it is the right machine that is in the critical section
				// // - caution: will not work if also sync machines may call wait4 fus...
				// // -- because the sub machine executes in same CS as the caller
				// if s.CurMachineKey != ctrlMsg.SenderMachineKey {
				//    s.SystemError(fmt.Sprintf("ill. machine leaves critical section: '%s' instead of '%s'", ctrlMsg.SenderMachineKey, s.CurMachineKey))
				// }
				//------------------------------------------------------------
				// clear info who is in critical section
				s.CurMachineKey = ""
				//------------------------------------------------------------
				// debug:
				// - print all containers if there was a space change made by the last machine execution that sent the leave
				if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() && SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT > prevSpaceUpdates { // DEBUG
					String2TraceFile(fmt.Sprintf("SPACE CHANGED:    t=%d, et=%d\n", CLOCK, EVENT_CLOCK))
					if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
						s.MetaContext.SpacePrint(TRACE0, IND, false /* printAlsoEmptyContainersFlag */) // DEBUG
					} // DEBUG
				} // DEBUG
			}
			//------------------------------------------------------------
			// SELECT A NEXT MACHINE
			//------------------------------------------------------------
			if VERIFICATION_MODE != MODEL_CHECKING {
				//------------------------------------------------------------
				// VERIFICATION_MODE != MODEL_CHECKING
				// - select machine from all existing machines
				//------------------------------------------------------------
				nextMachineKey = s.selectNextMachine()
				//------------------------------------------------------------
			} else {
				//------------------------------------------------------------
				// VERIFICATION_MODE == MODEL_CHECKING
				//------------------------------------------------------------
				if MC_VARS.UseCurChoiceAsNextMachineFlag {
					//------------------------------------------------------------
					// NEXT RUN IS TO BE STARTED WITH A SELECTED CHOICE FROM A RECOVERED CHOICE POINT (DONE BY RUNTIME)
					// - passed via global vars by runtime
					//------------------------------------------------------------
					// assertions
					if MC_VARS.CurChoicePoint == nil {
						s.SystemError(fmt.Sprintln("can't recover from empty choice point"))
					}
					if MC_VARS.CurChoice == "" {
						s.SystemError(fmt.Sprintln("can't recover from empty choice"))
					}
					if RUN_COUNT == 1 {
						s.SystemError(fmt.Sprintln("first run must not recover from choice point"))
					}
					//------------------------------------------------------------
					// set next machine to current choice
					nextMachineKey = MC_VARS.CurChoice
					//------------------------------------------------------------
					// set cur depth of path of the current choice
					curCpDepth = MC_VARS.CurChoicePoint.Depth
					//------------------------------------------------------------
					// caution: do not create a new choice point here!!
					// - this was done by runtime which updated the CP (ie removed the current choice)
					//------------------------------------------------------------
					// debug:
					// - TBD: redundant
					if MODEL_CHECKING_DETAILS3_TRACE.DoTrace() { // DEBUG
						String2TraceFile(fmt.Sprintf("RECOVER FROM CHOICE POINT CP[%d]: t=%d, et=%d, %d CPs, %d GIDs\n",
							MC_VARS.CurChoicePoint.Id, CLOCK, EVENT_CLOCK, len(MC_VARS.ChoicePoints), runtime.NumGoroutine())) // DEBUG
					}
					//------------------------------------------------------------
				} else {
					//------------------------------------------------------------
					// FIRST RUN AT ALL, OR N-TH (N > 1) CHOICE OF AN ALREADY EXECUTING RUN
					// - select machine like above from all existing machines
					nextMachineKey = s.selectNextMachine()
					//------------------------------------------------------------
					// if state has changed since last CP
					// - construct a choice point with all possible alternative candidates for this status
					// -- and add them to the global choice points list
					//------------------------------------------------------------
					if SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT > prevSpaceUpdates {
						//------------------------------------------------------------
						// tentatively alloc a new choice point
						newCp := NewChoicePoint()
						//------------------------------------------------------------
						s.StatusMutex.RLock() // LOCK FOR READ //
						//------------------------------------------------------------
						// debug for MODEL_CHECKING_TRACE.DoTrace()
						noCandidateMs := "" // DEBUG
						//------------------------------------------------------------
						// check *all* machines to be a candidate:
						// - and remember all candidates (except of the selected one) in the potentially new choice point
						for key, mc := range s.MachineControls {
							//------------------------------------------------------------
							// assertion
							if mc.Condition == nil {
								s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (d)", key))
							}
							//------------------------------------------------------------
							// debug
							if CONTROLLER_TRACE.DoTrace() { // DEBUG
								String2TraceFile(fmt.Sprintf("\n%s: %s: CLOCK=%d\n", thisFuNm, key, CLOCK)) // DEBUG
								mc.Condition.Println(TAB)                                                   // DEBUG
							} // DEBUG
							//------------------------------------------------------------
							// check definition of being a candidate:
							// String2TraceFile(fmt.Sprintf("?? CHECK: %s %s, cpflag=%t, t=%d\n", key, mc.Condition.String(), mc.Condition.ChoicePointFlag, CLOCK)) // @@@DEBUG
							candidateFlag := false
							//------------------------------------------------------------
							// (a) not the selected key?
							if key != nextMachineKey {
								//------------------------------------------------------------
								// (b) if it is a condition caused by a wait4 event that requires a choice point, and
								if mc.Condition.GenerateChoiceFlag {
									//------------------------------------------------------------
									// (c) if condition is not resetted user event (which must not trigger a choice!)
									if !mc.UserConditionResettedFlag {
										//------------------------------------------------------------
										// (d) if condition is fulfilled to enter the critical section
										fulfilledFlag := s.ConditionIsFulfilled(mc)
										if fulfilledFlag {
											//------------------------------------------------------------
											// debug
											if CONTROLLER_TRACE.DoTrace() { // DEBUG
												String2TraceFile("  fulfilled\n") // DEBUG
											} // DEBUG
											// - set flag
											candidateFlag = true
											//------------------------------------------------------------
											// add machine key to new choice point
											// - nb: val is unused
											// - nb: do not break from for loop, as we must compute *all* candidate choices that exist at this point
											newCp.Choices[key] = 0
										}
									}
								}
							} // candidate determination
							//------------------------------------------------------------
							// debug: print the non candidates
							if key != nextMachineKey && !candidateFlag && MODEL_CHECKING_DETAILS1_TRACE.DoTrace() { // DEBUG
								noCandidateMs = fmt.Sprintf("%s  %s\n", noCandidateMs, s.MachineControls[key].KeyInfo()) // DEBUG
							}
							//------------------------------------------------------------
						} // for (candidate determination)
						if noCandidateMs != "" && MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
							String2TraceFile(fmt.Sprintf("\nNO CHOICE POINT CHOICE CANDIDATE(S): \n%s", noCandidateMs)) // DEBUG
						}
						//------------------------------------------------------------
						s.StatusMutex.RUnlock() // UNLOCK FOR READ //
						//------------------------------------------------------------
						// add choice point to MC_CHOICE_POINTS, if there are there more than 1 choices open
						// - for backtracking -- done by runtime
						// nb: here we just execute 1 path in the depth
						// nb: create CP only if there was any state change since the last CPs
						if 1 < len(newCp.Choices) {
							//------------------------------------------------------------
							// clone status for a new run
							newCp.S = s.Clone4NewRun(SYSTEM_TTL)
							//------------------------------------------------------------
							// set
							// - event clokck
							newCp.EventClock = EVENT_CLOCK
							// - clock
							newCp.Clock = CLOCK
							// - depth:
							newCp.Depth = curCpDepth
							curCpDepth++
							// -- debug
							MC_VARS.CurPathDepth = curCpDepth
							// - debug: counting is 1, 2, 3, ...
							newCp.Id = MC_VARS.ChoicePointUuid
							MC_VARS.ChoicePointUuid++
							//------------------------------------------------------------
							// add a new choice point
							MC_VARS.ChoicePoints = append(MC_VARS.ChoicePoints, newCp)
							//------------------------------------------------------------
							// debug
							//............................................................
							if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
								//............................................................
								// only for debug
								s.StatusMutex.RLock() // LOCK FOR READ //
								//............................................................
								String2TraceFile(fmt.Sprintf("ADD CHOICE POINT: CP[%d] for %d/%d machines: depth=%d, %d state updates since last CP, ancestor=CP[%d]\n",
									newCp.Id, len(newCp.Choices), len(s.MachineControls), newCp.Depth, SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT, MC_VARS.CurChoicePoint.Id)) // DEBUG
								//............................................................
								if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
									for key, mc := range s.MachineControls { // DEBUG
										if newCp.ContainsKey(key) { // DEBUG
											String2TraceFile(fmt.Sprintf("  %s\n", mc.KeyInfo())) // DEBUG
										} else if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
											String2TraceFile(fmt.Sprintf("  (%s, no choice)\n", mc.KeyInfo())) // DEBUG
										} // DEBUG
									} // DEBUG
								} // DEBUG
								//............................................................
								// only for debug
								s.StatusMutex.RUnlock() // UNLOCK FOR READ //
								//............................................................
							} // DEBUG
							//------------------------------------------------------------
						} else { // if/else: are there choices, ie was a cp created?
							//------------------------------------------------------------
							// debug
							if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() && nextMachineKey != "" { // DEBUG
								String2TraceFile(fmt.Sprintf("NO NEW CP:     depth=%d\n", MC_VARS.CurPathDepth)) // DEBUG
							}
						} // else, debug
					} // if state has changed since last CP
				} // if key was not recovered from choice point
				//------------------------------------------------------------
			} // if model checking mode

		default:
			//============================================================
			// DEFAULT
			//============================================================
			Panic(fmt.Sprintf("%s: ill. ctrl msg received = %s", thisFuNm, ctrlMsg.Sig))

		} // end of switch ctrlMsg

		//============================================================
		// SELECTION OF NEXT MACHINE DONE:
		// - NB: there is no machine in the critical section at this place
		// - NB: the possible new candidate is stored in nextMachineKey
		//============================================================

		//------------------------------------------------------------
		// only for MC
		prevSpaceUpdates = SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT
		//------------------------------------------------------------
		// process all ripe scheduler slots & remove them from scheduler:
		// - eg for PM work needs to be done, if ettl expired: namely entry hunting
		// - CAUTION: it must be guaranteed that only controller is active when executing the scheduler !!!
		// -- TBD: assert that this is the case; nb: machines might still execute some "termination/clean up code" after having left the CS,
		// --- but this code does not access the scheduler...
		var nextRipeSlot *Slot
		for {
			//------------------------------------------------------------
			// debug: for assertion below
			schedulerLenBefore := len(s.Scheduler)
			//------------------------------------------------------------
			// get and remove next ripe slot, if any
			nextRipeSlot, s.Scheduler, stopFlag = s.Scheduler.GetAndRemoveNextRipeSlot()
			//------------------------------------------------------------
			// debug
			// - fmt.Println(fmt.Sprintf("after GetAndRemoveNextRipeSlot: len scheduler = %d", len(s.Scheduler))) // DEBUG
			//------------------------------------------------------------
			// shall we stop now?
			// - ie sys ttl has been reached
			if stopFlag {
				//------------------------------------------------------------
				// set stop msg for sys info
				stopMsg = fmt.Sprintf("SYSTEM TTL %d exceeded", SYSTEM_TTL)
				//------------------------------------------------------------
				// break from controller loop
				break controllerLoop
			}
			//------------------------------------------------------------
			// if a scheduler slot was found and removed -> execute it now
			if nil != nextRipeSlot {
				//------------------------------------------------------------
				// assertion that it is a user slot
				if nextRipeSlot.Type != USER_SLOT {
					s.SystemError(fmt.Sprintf("ill. slot type"))
				}
				//------------------------------------------------------------
				// assertion that scheduler has shrinked by 1
				if (schedulerLenBefore - 1) != len(s.Scheduler) {
					s.SystemError(fmt.Sprintf("ill. scheduler access: len before = %d, len after = %d", schedulerLenBefore, len(s.Scheduler)))
				}
				//------------------------------------------------------------
				// execute the slot
				// - in the user model
				s.MetaContext.ProcessRipeUserSlot(nextRipeSlot.UserSlot, &s.Scheduler)
				//------------------------------------------------------------
				// debug:
				// - fmt.Printf("call ProcessRipeUserSlot, clock=%d, slot time=%d, len(scheduler)=%d\n", CLOCK, nextRipeSlot.Time, len(s.Scheduler)) // DEBUG
			} else {
				//------------------------------------------------------------
				// no more ripe slot -> break from *this* for-loop
				break
			}
		}
		//------------------------------------------------------------
		// A NEXT MACHINE WAS SELECTD ABOVE FOR EXECUTION:
		// - resume it, ie allow it to enter the critical section
		if "" != nextMachineKey {
			//------------------------------------------------------------
			s.StatusMutex.RLock() // LOCK FOR READ //
			nextMc := s.MachineControls[nextMachineKey]
			//------------------------------------------------------------
			// assertions
			if nextMc == nil {
				s.SystemError(fmt.Sprintf("next selected machine '%s' not in MCS", nextMachineKey))
			}
			if nextMc.M == nil {
				s.SystemError("M of next selected machine in MCS is empty")
			}
			//------------------------------------------------------------
			if VERIFICATION_MODE == MODEL_CHECKING {
				//------------------------------------------------------------
				if MODEL_CHECKING_DETAILS3_TRACE.DoTrace() || CONTROLLER_TRACE.DoTrace() { // DEBUG
					String2TraceFile(fmt.Sprintf("\nENTER %s, time=%d, state=%s\n", nextMachineKey, CLOCK, nextMc.M.CurrentState)) // DEBUG
				} // DEBUG
				//------------------------------------------------------------
				// debug
				if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
					String2TraceFile(fmt.Sprintf("\n")) // DEBUG
				} // DEBUG
				if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
					String2TraceFile(fmt.Sprintf("EXECUTE:          %s\n", nextMc.KeyInfo())) // DEBUG
				} // DEBUG
			}
			//------------------------------------------------------------
			// debug
			if REPLAY_TRACE.DoTrace() { // DEBUG
				String2ReplayTraceFile(fmt.Sprintf("%s\n", nextMachineKey)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			s.StatusMutex.RUnlock() // UNLOCK FOR READ //
			//------------------------------------------------------------
			// RESUME
			//------------------------------------------------------------
			// resume the selected machine now !!!!!!!!!!!!!!!!!!!!!!!!!!!
			// - caution: locks status
			// - nb: does not block
			s.resume(nextMachineKey)
			//------------------------------------------------------------
			// complicated: for model checking:
			// - clear flag that machine's user event condition was resetted to no event;
			// - ie try it without user event condition and do so as if the event was there;
			// - nb: after that the next user wait4 event incl. its event times is used;
			// CAUTION: do it after resume -- machine must get 1 chance with no event
			if VERIFICATION_MODE == MODEL_CHECKING && nextMc.UserConditionResettedFlag {
				nextMc.UserConditionResettedFlag = false
			}
		}
		//------------------------------------------------------------
		// if no machine could be selected and time must be advanced -> kick the controller :-)
		if "" == nextMachineKey {
			//------------------------------------------------------------
			// advance clock to next interesting time
			if len(s.Scheduler) > 0 {
				CLOCK = s.Scheduler[0].Time
				if MODEL_CHECKING_DETAILS1_TRACE.DoTrace() { // DEBUG
					String2TraceFile(fmt.Sprintf("\nNO MACHINE SELECTED\nADVANCE CLOCK TO %d of scheduler slot=%s", CLOCK, s.Scheduler[0].ToString(0))) // DEBUG
				} // DEBUG
			}
			//------------------------------------------------------------
			// debug
			if CONTROLLER_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: no machine's condition fulfilled\n", thisFuNm)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// trick: re-send controller the "KICK" so that it does not hang
			// - and hopefully lets a next machine enter the CS...
			s.ControllerChannel <- NewChanSig(KICK, SENDER_IS_SYSTEM, "Controller" /* msg */)
		}
		//------------------------------------------------------------
		// increment the system time aka CLOCK
		// - "stepper motor"
		CLOCK++
		//------------------------------------------------------------
		// debug
		if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
			String2TraceFile(fmt.Sprintf("\nTICK: t=%d, et=%d, %d CPs, %d GIDs \n", CLOCK, EVENT_CLOCK, len(MC_VARS.ChoicePoints), runtime.NumGoroutine())) // DEBUG
		} // DEBUG
		//------------------------------------------------------------
		// signal that time has changed
		// - nb: for debug only
		TimeChangeEvent()
		//------------------------------------------------------------
		// for model checking mode: reset the flag, because next selected machine is already allowed to create a CP
		if VERIFICATION_MODE == MODEL_CHECKING {
			MC_VARS.UseCurChoiceAsNextMachineFlag = false
		}
		//------------------------------------------------------------
	} // for

	//////////////////////////////////////////////////////////////
	// end of controller loop
	//////////////////////////////////////////////////////////////

	//------------------------------------------------------------
	// stop flag?
	// - nb: check here, because the execution of slots could also set the stop flag
	if stopFlag {
		//------------------------------------------------------------
		// stop all machines
		//------------------------------------------------------------
		// debug:
		// s.PrintStatistics() // DEBUG
		String2TraceFile("\n")                                            // DEBUG
		s.SystemInfo(fmt.Sprintf("STOP MACHINES because of %s", stopMsg)) // DEBUG
		//------------------------------------------------------------
		s.StatusMutex.RLock() // LOCK FOR READ //
		//------------------------------------------------------------
		// EXPLICITLY TERMINATE EVERY MACHINE IN MCS
		// - nb: could also be a sync machine that is running in the thread of a (async) machine
		nMachinesStopped := 0
		for key, mc := range s.MachineControls {
			//------------------------------------------------------------
			// debug: print which machine is stopped
			if CONTROLLER_TRACE.DoTrace() { // DEBUG
				// mc.M.SystemInfo(fmt.Sprintf("%s: STOP %s, GID=%d", thisFuNm, key, mc.Gid)) // SYS INFO
				mc.M.SystemInfo(fmt.Sprintf("STOP %s", key)) // SYS INFO
			} //DEBUG
			//------------------------------------------------------------
			// 1.) send "stop" to machine's ctrl channel
			mc.CtrlChan <- NewChanSig(STOP, SENDER_IS_SYSTEM, "Controller" /* msg */)
			// - count number of stopped machines
			nMachinesStopped++
		}
		//------------------------------------------------------------
		s.StatusMutex.RUnlock() // UNLOCK FOR READ //
		//------------------------------------------------------------
		// debug
		// - s.SystemInfo(fmt.Sprintf("%s: waiting for %d machines to send terminated signal", thisFuNm, nMachinesStopped)) // DEBUG
		//		//------------------------------------------------------------
		//		// 2.) wait until all machines, that were told to stop, send either leave or terminated
		//		// - nb: each sub (sync) machine of an (async) machine has its own MC with its own channels
		//		// -- the stop is therefore sent to sub machines as well as to their callers
		//		// - bit complicated logic... a submachine terminates on its own and will also send either leave or terminate
		//		i := 0
		//		for {
		//			//------------------------------------------------------------
		//			// wait for next msg
		//			msg := <-s.ControllerChannel
		//			//------------------------------------------------------------
		//			// count number of terminated or leave signals received
		//			// - nb: machine will send either one... bit complicated logic...
		//			// caution: do not use wait group to check that all machines have really terminated their execution
		//			// - because if the wait for waitgroup call comes after all machines have already stopped, it might panic with
		//			// -- "fatal error: all goroutines are asleep - deadlock", if the decrement from 1 to 0 machine has happened before
		//			// --- the wait call; this is an ill. semantics of waitgroup...
		//			if msg.Sig == TERMINATED || msg.Sig == LEAVE {
		//				//------------------------------------------------------------
		//				// debug:
		//				// - s.SystemInfo(fmt.Sprintf("%s: %s received from stopped %s; chan = %v", thisFuNm, msg.Sig, msg.SenderMachineKey, s.ControllerChannel)) // DEBUG
		//				i++
		//				//------------------------------------------------------------
		//				// break if enough msgs were received
		//				if i >= nMachinesStopped {
		//					break
		//				}
		//			}
		//			// else: ignore any other msg
		//		} // for
	} // if (stop flag)
	//------------------------------------------------------------
	// TBD: close controller channel
	// - nb: controller has receiver role on both channels...
	// -- and channel closing should be done by sender, because sending to a closed channel causes panic
	// close(s.ControllerChannel)
	//------------------------------------------------------------
	// debug:
	s.SystemInfo(fmt.Sprintf("Controller: EXIT (modi: %s and %s; %d go routines left)", VERIFICATION_MODE, EXECUTION_MODE, runtime.NumGoroutine())) // DEBUG
	// - fmt.Println(fmt.Sprintf("Scheduler = %s\n", s.Scheduler.ToString(0)))                                      // DEBUG                                                                                    // DEBUG
}

//------------------------------------------------------------
// resume a machine
// - send it the enter signal
// - does not block
// - nb: all start conditions were already checked by caller
// - private fu
func (s *Status) resume(machineKey string) {
	//------------------------------------------------------------
	s.StatusMutex.RLock() // LOCK FOR READ //
	//------------------------------------------------------------
	// get the mutex channel for the given machine (via its key)
	//............................................................
	// - get the machine's machine control (= mc)
	mc := s.MachineControls[machineKey]
	//............................................................
	// - assertion that mc and m exist
	if nil == mc || nil == mc.M {
		s.SystemInfo(fmt.Sprintf("cant' resume machine %s", machineKey)) // DEBUG
		s.StatusMutex.RUnlock()                                          // UNLOCK FOR READ //
		s.Panic()
	}
	//............................................................
	// - get its mutex channel
	mutexChan := mc.MutexChan
	//------------------------------------------------------------
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	// send ENTER signal to machine's mutex channel
	mutexChan <- NewChanSig(ENTER, SENDER_IS_SYSTEM, "Resume" /* msg */)
	//------------------------------------------------------------
	// for debug only: increment the machine's nCriticalSections counter (in its mc)
	// - CAUTION: use locking, otherwise "fatal error: concurrent map iteration and map write" might be reported here
	//------------------------------------------------------------
	s.StatusMutex.Lock() // LOCK FOR WRITE //
	//------------------------------------------------------------
	// statistics:
	s.MachineControls[machineKey].NCriticalSections++
	s.MachineControls[machineKey].LastExecutionTime = CLOCK
	//------------------------------------------------------------
	s.StatusMutex.Unlock() // UNLOCK FOR WRITE //
}

//============================================================
// system messages:
// - machine parameter can be nil: if so -> no machine info printed
//============================================================

//------------------------------------------------------------
// debug
// - print the id of my go routine (i = whoAmI)
func (s *Status) PrintMyGoRoutineId(msg string, whoAmI string) {
	s.SystemInfo(fmt.Sprintf("%s GID=%d: %s, ngids=%d, t=%d", msg, GetGoRoutineID(), whoAmI, runtime.NumGoroutine(), CLOCK)) // DEBUG
}

//------------------------------------------------------------
// debug: print space status & statistics
// - caution: MCS are empty as all machines are gone
// - nb: for the peer model also the system peers IOP and STOP are included with their machines
func (s *Status) PrintStatistics() {
	//------------------------------------------------------------
	// print all containers:
	s.MetaContext.SpacePrint(TRACE0, IND, false /* printAlsoEmptyContainersFlag */)
	//------------------------------------------------------------
	if STATISTICS_TRACE.DoTrace() {
		//------------------------------------------------------------
		// s.Scheduler.Println(0)
		//------------------------------------------------------------
		s.SystemInfo("STATISTICS:")
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- event time et=%d", EVENT_CLOCK))
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- %d critical sections (CS)", s.CriticalSectionCounter))
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- %d machines (M)", s.MachineTerminationCounter))
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- execution mode: %s", EXECUTION_MODE))
		//------------------------------------------------------------
		helpS := fmt.Sprintf("- verification mode: %s", VERIFICATION_MODE)
		switch VERIFICATION_MODE {
		case SIMULATION:
			s.SystemInfo(fmt.Sprintf("%s (%d RUNS)", helpS, RUN_COUNT))
		case MODEL_CHECKING:
			ncps := 0
			for _, cp := range MC_VARS.ChoicePoints {
				ncps += len(cp.Choices)
			}
			s.SystemInfo(fmt.Sprintf("%s (%d runs, %d CPs, depth=%d, still %d CPs and %d choices)", helpS, RUN_COUNT, MC_VARS.ChoicePointUuid-1, MC_VARS.CurPathDepth, len(MC_VARS.ChoicePoints), ncps))
		default:
			s.SystemInfo(fmt.Sprintf("%s", helpS))
		}
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- %s", s.AutomataToString()))
		//------------------------------------------------------------
		s.SystemInfo(fmt.Sprintf("- %d go routines running", runtime.NumGoroutine()))
		//------------------------------------------------------------
		String2TraceFile("\n")
	}
}

//------------------------------------------------------------
// print a system message that starts with 3 stars
// does the padding...
// - @@@assumes the following consts for tab stops: 55 + 125
// - @@@cf the analogous function in machine
func (s *Status) PrintlnStarMessage(msgType MachineMessageTypeEnum, msgText string) {

	msg0 := fmt.Sprintf("*** %s ", msgType)
	msg0 = Padding(msg0, 4+11+1+3, "-")

	msg1 := fmt.Sprintf(" %s ", msgText)
	msg1 = Padding(msg1, 55, "-") // @@@

	msg0_1 := Padding(fmt.Sprintf("%s%s", msg0, msg1), 125, "-") // @@@

	msg3 := fmt.Sprintf("-- CLOCK=%d", CLOCK)

	msg := fmt.Sprintf("%s%s\n", msg0_1, msg3)
	/**/ String2TraceFile(msg)
}

//------------------------------------------------------------
func (s *Status) Panic() {
	// print status ie all containers - for debug only
	s.MetaContext.SpacePrint(TRACE0, 0, true /* printAlsoEmptyContainersFlag */)

	// stop everything
	// - makes no sense, because there is only one machine in the critical section
	// -- and probably also the controller has stopped so machines cannot execute any more their stop signal
	// s.StopAllMachines("Panic")

	// print status - debug
	// s.SpacePrint(0, false /* printAlsoEmptyContainersFlag */)

	// exit
	syscall.Exit(-1)
}

//------------------------------------------------------------
// should not occur, e.g. wrong model
func (s *Status) UserError(err string) {
	/**/ s.PrintlnStarMessage(USER_ERROR, err)
	s.Panic()
}

//------------------------------------------------------------
// should not occur, system failure
func (s *Status) SystemError(err string) {
	/**/ s.PrintlnStarMessage(SYSTEM_ERROR, err)
	s.Panic()
}

//------------------------------------------------------------
func (s *Status) UserWarning(w string) {
	/**/ s.PrintlnStarMessage(USER_WARNING, w)
}

//------------------------------------------------------------
func (s *Status) SystemWarning(w string) {
	/**/ s.PrintlnStarMessage(SYSTEM_WARNING, w)
}

//------------------------------------------------------------
func (s *Status) SystemInfo(info string) {
	// /**/ String2TraceFile("\n")
	/**/
	s.PrintlnStarMessage(SYSTEM_INFO, info)
}

//------------------------------------------------------------
// signal time change event
// - nothing needs to be done
func TimeChangeEvent() {
	//------------------------------------------------------------
	// debug
	if EVENT_CONDITION_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("TimeChangeEvent:et=%d, t=%d \n", EVENT_CLOCK, CLOCK))
	}
}

//------------------------------------------------------------
// check if condition of the machine control is fulfilled
// - ie either user event condition or time condition
// - or no event ... anyhow fulfilled
// CAUTION: special treatment for model checking after a user condition was resetted (complicated)
func (s *Status) ConditionIsFulfilled(mc *MachineControl) bool {
	//------------------------------------------------------------
	// debug
	// - fmt.Println(fmt.Sprintf("ConditionIsFulfilled: key=%s, userConditionResettedFlag=%t", mc.M.Key(), mc.UserConditionResettedFlag))
	//------------------------------------------------------------
	condition := mc.Condition
	//------------------------------------------------------------
	// assertion
	if condition == nil {
		s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (e)", mc.M.Key()))
	}
	//------------------------------------------------------------
	// complicated: for model checking:
	// - if user condition was resetted, simulate that its event is fulfilled
	// nb: is only done for the first execution of machine after choice point recovery
	if VERIFICATION_MODE == MODEL_CHECKING && mc.UserConditionResettedFlag {
		return true
	}
	//------------------------------------------------------------
	fulfilledFlag := false
	switch condition.Type {
	//------------------------------------------------------------
	case EMPTY_CONDITION:
		fulfilledFlag = true
	//------------------------------------------------------------
	case USER_EVENT:
		// assertion
		if nil == condition.UserEvent {
			s.SystemError(fmt.Sprintf("ConditionIsFulfilled: user event struct is empty in USER EVENT for m=%s", mc.M.Key()))
		}
		// check condition
		fulfilledFlag = s.MetaContext.ConditionIsFulfilled(condition)
		// !!! fallthrough because every event has also a time condition !!!
		// !!! nb: it sufficies if *one* of both conditions is fulfilled !!!
		fallthrough
	//------------------------------------------------------------
	case TIME_EVENT:
		if condition.Wait4Time <= CLOCK {
			fulfilledFlag = true
		}
	//------------------------------------------------------------
	case NO_EVENT:
		fulfilledFlag = true
	//------------------------------------------------------------
	default:
		Panic(fmt.Sprintf("ill. condition type='%s'", condition.Type))
	}
	return fulfilledFlag
}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// only automaton names
func (s *Status) AutomataToString() string {
	//------------------------------------------------------------
	// init local vars
	sep := ""
	res1 := ""
	nM := 0
	//------------------------------------------------------------
	s.StatusMutex.RLock() // LOCK FOR READ //
	//------------------------------------------------------------
	// sort automata names for nice output
	var sortedAutomataNames Strings
	for key, _ := range s.Automata {
		sortedAutomataNames = sortedAutomataNames.SortedInsertString(key)
	}
	//------------------------------------------------------------
	// generate output
	for i := 0; i < len(sortedAutomataNames); i++ {
		aName := sortedAutomataNames[i]
		a := s.Automata[aName]
		res1 = fmt.Sprintf("%s%s%s (%dM", res1, sep, aName, a.nMachinesUsedCount)
		nM += a.nMachinesUsedCount
		if a.nCriticalSections > 0 {
			// - eg sync machines have no critical section entered
			res1 = fmt.Sprintf("%s,%dCS", res1, a.nCriticalSections)
		}
		res1 = fmt.Sprintf("%s)", res1)
		sep = ", "
	}
	res2 := fmt.Sprintf("%d automata used: %s", len(s.Automata), res1)
	//------------------------------------------------------------
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	//------------------------------------------------------------
	return res2
}

//------------------------------------------------------------
// only amount and machine keys
func (s *Status) MCSToString() string {
	s.StatusMutex.RLock() // LOCK FOR READ //
	sep := ""
	res := fmt.Sprintf("%d MCs: ", len(s.MachineControls))
	for key, _ := range s.MachineControls {
		res = fmt.Sprintf("%s%s%s", res, sep, key)
		sep = " / "
	}
	s.StatusMutex.RUnlock() // UNLOCK FOR READ //
	return res
}

//------------------------------------------------------------
// print
func (s *Status) MCSPrintln(msg string) {
	String2TraceFile(fmt.Sprintf("%s: %s\n", msg, s.MCSToString()))
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
