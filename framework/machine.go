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
// abbreviations:
// - LVS ... local variables of a machine
//------------------------------------------------------------
// TBD (in work:) Code Review: 2021 Apr, Eva Maria Kuehn
// - TBD: debug functions need to be reviewed
//////////////////////////////////////////////////////////////

package framework

import (
	. "github.com/peermodel/simulator/contextInterface"
	. "github.com/peermodel/simulator/controller"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/helpers"
	. "github.com/peermodel/simulator/scheduler"
	"fmt"
	"runtime"
)

//////////////////////////////////////////////////////////////
// vars
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// counts number of machines
var MACHINE_ID int

//////////////////////////////////////////////////////////////
// data types
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// copy first arg (= LVS) and return result
// - TBD: is there any more elegant solution for that?
type LVSCopyHandler func(*Machine, interface{}) interface{}

//------------------------------------------------------------
// complete second arg (= LVS) (compute lvs vars shared with status) and return result
// - first arg is the *new* status
// - TBD: is there any more elegant solution for that?
type LVSAliasHandler func(*Status, *Machine, interface{}) interface{}

//------------------------------------------------------------
// data of one machine
// - including an own context;
type Machine struct {
	//============================================================
	// static machine info:
	//============================================================
	//------------------------------------------------------------
	// - machine name
	Name string
	//------------------------------------------------------------
	// - machine number
	Number int
	//------------------------------------------------------------
	// - automaton specification to be executed by this machine
	A *Automaton
	//============================================================
	// dynamic machine info:
	//============================================================
	//------------------------------------------------------------
	// - id of the current state
	CurrentState string
	//============================================================
	// local variables (LVS) interface: individual per machine
	//============================================================
	//------------------------------------------------------------
	// - data struct that holds the local vars for the machine
	// - there vars are shared between the machine's states
	// - can be any struct, therefore interface is used
	LocalVariables interface{}
	//============================================================
	// context: used to pass info between machines (at entering/leaving)
	//============================================================
	//------------------------------------------------------------
	// - context: i/o variables, ie context "shared" between machines (on call and return time)
	Context IContext
	//============================================================
	// start type:
	//============================================================
	//------------------------------------------------------------
	// SYNC or ASYNC
	StartType MachineStartTypeEnum
	//============================================================
	// debug infos:
	//============================================================
	//------------------------------------------------------------
	// - shall traces be printed
	// -- for machine
	TraceFlag bool
	//------------------------------------------------------------
	// -- for the current state
	ThisStateTraceFlag bool
	//------------------------------------------------------------
	// - when has the machine been started?
	StartTime int
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new machine
// - caution: increments machine count
func NewMachine(a *Automaton) *Machine {
	//------------------------------------------------------------
	// alloc
	m := new(Machine)
	//------------------------------------------------------------
	// init:
	// - same name as automaton; TBD: redundant
	m.Name = a.Name
	m.Number = MACHINE_ID
	m.A = a
	MACHINE_ID++
	m.CurrentState = "init"
	m.TraceFlag = MACHINE_TRACE_FLAGS[m.Name].MFlag
	// - not necessary -> is set at start anyhow
	m.StartTime = CLOCK
	//------------------------------------------------------------
	// return
	return m
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// copy, share, renew everything that can/must be copied for a choice point
// - nb: machine has same number, name, current state
// - LVS are copied
// - CAUTION: caller must refresh the aliases with the current status!!!
func (m *Machine) Clone4NewRun() *Machine {
	//------------------------------------------------------------
	// alloc
	newM := NewMachine(m.A)
	//============================================================
	// static machine info:
	//============================================================
	//------------------------------------------------------------
	// Name
	newM.Name = m.Name
	//------------------------------------------------------------
	// Number (copy)
	newM.Number = m.Number
	//------------------------------------------------------------
	// Automaton (share)
	// - was set above
	//============================================================
	// dynamic machine info:
	//============================================================
	//------------------------------------------------------------
	// CurrentState (copy)
	newM.CurrentState = m.CurrentState
	//============================================================
	// LVS:
	//============================================================
	//------------------------------------------------------------
	// LocalVariables (deep copy)
	newM.LocalVariables = m.A.LocalVariablesCopyFunction(m, m.LocalVariables)
	//============================================================
	// context:
	//============================================================
	// Context (deep copy)
	newM.Context = m.Context.Copy().(IContext)
	//============================================================
	// start type:
	//============================================================
	//------------------------------------------------------------
	// StartType (copy)
	newM.StartType = m.StartType
	//============================================================
	// debug infos:
	//============================================================
	// StateComments
	// - the are already initialized (see comment above)
	//------------------------------------------------------------
	// TraceFlag (copy)
	newM.TraceFlag = m.TraceFlag
	//------------------------------------------------------------
	// ThisStateTraceFlag (copy)
	newM.ThisStateTraceFlag = m.ThisStateTraceFlag
	//------------------------------------------------------------
	// StartTime (copy)
	newM.StartTime = m.StartTime
	//============================================================
	// return
	//============================================================
	return newM
}

//============================================================
// comment for the following functions:
// - the current machine controls are found in the current status which is passed as arg
//============================================================

//------------------------------------------------------------
// generate the machine key for this machine
// - "<machine-name>__M<machine-number>__<machine-key-suffix>":
func (m *Machine) Key() string {
	return (fmt.Sprintf("%s__M%d__%s", m.Name, m.Number, m.Context.MachineKeySuffix()))
}

//------------------------------------------------------------
// execute the machine
// - if start type is SYNC the machine runs in caller's thread who must be in the critical section now
// - else caller has created a new thread for machine
// -- ie machine must enter the critical section before its execution and leave it on exit
// nb: inbetween the critical section could be left & entered again by the handler code when calling a wait4 fu
func (m *Machine) Execute(s *Status, mc *MachineControl) {
	//------------------------------------------------------------
	// assertions:
	// - automaton must exist
	if nil == m.A {
		m.SystemError(fmt.Sprintf("machine %s misses automaton", m.Key()))
	}
	//------------------------------------------------------------
	// debug:
	//............................................................
	// - init debug flag: default = no trace
	stateTraceFlag := false
	//............................................................
	// - set my go routine id
	mc.Gid = GetGoRoutineID()
	//............................................................
	// - print debug info
	// -- about go routine id(s)
	// m.PrintMyGoRoutineId("Exec ") // DEBUG
	// -- about machine start
	if MACHINE_START_TRACE.DoTrace() { // DEBUG
		/**/ String2TraceFile(fmt.Sprintf("Machine START %s M%d, m.StartTime=%d\n", m.Name, m.Number, m.StartTime)) // DEBUG
		// TBD: display also context pid:
		//  /**/ String2TraceFile(fmt.Sprintf("Machine START %s M%d, Peer=%s, m.StartTime=%d\n", m.Name, m.Number, m.Context.Pid, m.StartTime))
	} // DEBUG
	//------------------------------------------------------------
	// set flag if machine is in critical section (CS)
	var mIsInCsFlag bool
	//............................................................
	// a sync machine is in the CS
	if SYNC == m.StartType {
		mIsInCsFlag = true
	} else {
		//............................................................
		// try to enter CS if ASYNC (or fire-and-forget)
		// - nb: blocks until critical section can be entered
		// - caution: check retval, ie if enter was ok or not (ie machine was stopped)
		mIsInCsFlag = s.EnterCriticalSection(m, mc)
	}

	//------------------------------------------------------------
	// if in critical section:
	//------------------------------------------------------------
	if mIsInCsFlag {
		//============================================================
		// machine loop:
		//============================================================
		//------------------------------------------------------------
		// execute all states until STOP message is received on ctrl channel
		// - nb: while executing a state handler, no STOP message is treated
		// - nb: machine has the critical section here
	machineLoop:
		for {
			//------------------------------------------------------------
			// try to receive a message on the control channel
			//------------------------------------------------------------
			select {
			case ctrlMsg := <-mc.CtrlChan:
				//============================================================
				// ctrl message received?
				//============================================================
				// if STOP -> stop the machine (which is still in its CS)
				switch ctrlMsg.Sig {
				case STOP:
					m.SystemInfo(fmt.Sprintf("%s STOPPED", m.Key())) // DEBUG
					break machineLoop
				default:
					// - otherwise error as no other ctrl msg is allowed here
					Panic(fmt.Sprintf("machine %s: ill. ctrl msg received = %s", m.Key(), ctrlMsg.Sig))
				}

				//============================================================
				// no ctrl message received for the moment
				//============================================================
				// so we can work now, i.e. execute the next state of the machine
				// - nb: can be interrupted by a stop message when it calls a wait4 fu
			default:
				//------------------------------------------------------------
				// debug:
				// - fmt.Println(fmt.Sprintf("%s: next current state = %s", m.Key(), m.CurrentState)) // DEBUG
				//------------------------------------------------------------
				// - get & set trace flag for this state:
				stateTraceFlag = MACHINE_TRACE_FLAGS[m.Name].SFlags[m.CurrentState]
				m.ThisStateTraceFlag = stateTraceFlag
				//------------------------------------------------------------
				// get state handler
				h := m.A.StateHandlers[m.CurrentState]
				//------------------------------------------------------------
				// assertion: state handler must exist
				if nil == h {
					m.SystemError(fmt.Sprintf("no handler function for state %s found", m.CurrentState))
				}
				//------------------------------------------------------------
				// debug
				// - String2TraceFile(fmt.Sprintf("m.TraceFlag %s ThisStateTraceFlag of machine %s for state = %s = %s\n", m.TraceFlag, m.Name, m.CurrentState,stateTraceFlag)
				if m.DoTrace() { // DEBUG
					/**/ m.PrintNumberAndNameAndState() // DEBUG
					/**/ String2TraceFile(fmt.Sprintf(" -- CLOCK=%d:\n", CLOCK)) // DEBUG
					/**/ m.Context.Println(IND + TAB) // DEBUG
				} // DEBUG
				//------------------------------------------------------------
				// TBD: assert that is me or a sync submachine of mine in the critical section
				// - not so easy...
				//------------------------------------------------------------
				// !!! HERE THE CODE OF THE AUTOMATON IS EXECUTED !!!
				//------------------------------------------------------------
				// execute the state handler
				// - nb: advances current state to next state
				// - nb: may call a wait4 function, ie leave & enter the critical section
				// -- if machine gets stop signal while waiting, its returns STOPPED
				retval := h(s, m)
				//------------------------------------------------------------
				// check ret val
				switch retval {
				case OK:
					//------------------------------------------------------------
					// OK:
					// - state has successfully finished
					// - nb: machine still is in the critical section
					// - just continue
				case EXIT:
					//------------------------------------------------------------
					// EXIT:
					// - this is the "natural and successful end" of a machine
					// - nb: machine still is in the critical section
					//............................................................
					// debug
					if m.StartType == ASYNC { // DEBUG
						// m.SystemInfo(fmt.Sprintf("exit state reached by %s, GID=%d", m.Key(), GetGoRoutineID())) // DEBUG
					} // DEBUG
					//............................................................
					// break from machine loop
					break machineLoop
					//------------------------------------------------------------
				case STOPPED:
					//------------------------------------------------------------
					// STOPPED:
					// - ie machine was stopped because a STOP signal sent to it while waiting 4 an event
					// - nb:  machine is NOT any more in the critical section
					//------------------------------------------------------------
					// debug
					// m.SystemInfo(fmt.Sprintf("stopped state reached by %s, GID=%d", m.Key(), GetGoRoutineID())) // DEBUG
					//------------------------------------------------------------
					// set CS flag
					mIsInCsFlag = false
					//------------------------------------------------------------
					// break from machine loop
					break machineLoop
					//------------------------------------------------------------
				default:
					//------------------------------------------------------------
					// DEFAULT: error
					//------------------------------------------------------------
					s.SystemError(fmt.Sprintln("ill. ret of state handle = %s", retval))
				} // switch
			} // select
		} // machine loop (for-loop)
	} // if (in CS)

	//------------------------------------------------------------
	// END OF MACHINE EXECUTION
	//------------------------------------------------------------

	//------------------------------------------------------------
	// debug
	if m.DoTrace() { // DEBUG
		/**/ m.PrintNumberAndName() // DEBUG
		/**/ String2TraceFile(fmt.Sprintf("machine %s end (life time was %d-%d)\n", m.Key(), m.StartTime, CLOCK)) // DEBUG
	} // DEBUG
	//------------------------------------------------------------
	// debug: statistics
	// - how many machines were used for this automaton
	m.A.nMachinesUsedCount++
	//------------------------------------------------------------
	// leave the critical section, if still in the CS
	// - if ASYNC (or fire-and-forget)
	if SYNC != m.StartType && mIsInCsFlag {
		//............................................................
		// debug
		// m.SystemInfo(fmt.Sprintf("M %s calls leave critical section before exiting, GID=%d", m.Key(), GetGoRoutineID())) // DEBUG
		//............................................................
		// leave
		s.LeaveCriticalSection(m)
	}
	//------------------------------------------------------------
	// clean up
	// - do it for every machine
	// caution: locks machine controls in status
	s.cleanUpTerminatedMachine(m.Key(), m.StartType)
	//------------------------------------------------------------
	// if machine was stopped, ie not in the CS
	// - inform controller about TERMINATION and stop my go routine explicitly
	// TBD: only if machine is controlled by controller, ie ASYNC?!
	if !mIsInCsFlag {
		//------------------------------------------------------------
		// debug
		// m.SystemInfo(fmt.Sprintf("++++ %s: send TERMINATED to Controller", m.Key())) // DEBUG
		//------------------------------------------------------------
		// send TERMINATED to controller
		s.ControllerChannel <- NewChanSig(TERMINATED, SENDER_IS_MACHINE, m.Key())
		//------------------------------------------------------------
		// stop go routine
		// - caution: only for non sync machines, because go routine of sync machine must not be exited,
		// -- because it uses the goroutine of its caller!
		// TBD: needed?
		if SYNC != m.StartType {
			//............................................................
			// - debug
			// m.PrintMyGoRoutineId("Exit ") // DEBUG
			//............................................................
			// exit
			runtime.Goexit()
		}
	}
}

//------------------------------------------------------------
// generic preparartion of starting a machine
// - private
func (m *Machine) prepareStart(startType MachineStartTypeEnum, s *Status, ctx IContext) *MachineControl {
	//------------------------------------------------------------
	// set context
	m.Context = ctx
	//------------------------------------------------------------
	// set start type
	m.StartType = startType
	//------------------------------------------------------------
	// start time
	m.StartTime = CLOCK
	//------------------------------------------------------------
	// create and add new machine control:
	//............................................................
	// - create event without precondition for machine start
	evt := NewEvent(EMPTY_CONDITION)
	//............................................................
	// - create and add machine control struct with event
	// -- caution: *before* executing the machine
	mc := s.CreateAndAddMachineControl(m, evt)
	//------------------------------------------------------------
	// return
	return mc
}

//------------------------------------------------------------
// synchronous start of the machine
// - waits until machine has finished its execution
// - start in current thread which must be in the critical section
// - return the context (which was modified by the execution of the machine) to caller
func (m *Machine) StartSync(s *Status, ctx IContext) IContext {
	//------------------------------------------------------------
	// prepare
	mc := m.prepareStart(SYNC, s, ctx)
	//------------------------------------------------------------
	// execute machine in my thread & wait until done
	m.Execute(s, mc)
	//------------------------------------------------------------
	// return modified context to caller
	return m.Context
}

//------------------------------------------------------------
// asynchronous (i.e. parallel) start of the machine
// - start in a new thread
func (m *Machine) StartAsync(s *Status, ctx IContext) {
	//------------------------------------------------------------
	// prepare
	mc := m.prepareStart(ASYNC, s, ctx)
	//------------------------------------------------------------
	// set context
	m.Context = ctx
	//------------------------------------------------------------
	// start and execute machine in parallel
	// - nb: this will set the machine's new gid correctly
	go m.Execute(s, mc)
}

//////////////////////////////////////////////////////////////
// debug:
// - these methods print info only if the corresponding machine trace flags are set
// - and the corresponding trace levels (see TraceLevelEnum) are configured
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// debug
// - print the id of my go routine
func (m *Machine) PrintMyGoRoutineId(msg string) {
	m.SystemInfo(fmt.Sprintf("%s GID=%d: %s, ngids=%d, t=%d)", msg, GetGoRoutineID(), m.Key(), runtime.NumGoroutine(), CLOCK)) // DEBUG
}

//------------------------------------------------------------
// NB: string should be of len == 1
// @@@ inefficient,because always the trace flag must be checked by print blanks
func (m *Machine) PrintNumberAndName() {
	c := ""
	switch m.CurrentState {
	case "init":
		c = "*"
	case "exit":
		c = "-"
	default:
		c = " "
	}
	/**/ String2TraceFile(fmt.Sprintf("%s M%d", c, m.Number))
	if m.Number < 10000 {
		/**/ NBlanks2TraceFile(1)
	}
	if m.Number < 1000 {
		/**/ NBlanks2TraceFile(1)
	}
	if m.Number < 100 {
		/**/ NBlanks2TraceFile(1)
	}
	if m.Number < 10 {
		/**/ NBlanks2TraceFile(1)
	}
	/**/ String2TraceFile(fmt.Sprintf("%s", m.Name))
	l := len(m.Name)
	for i := l; i < MAX_MACHINE_NAME_CHARS; i++ {
		/**/ String2TraceFile(" ")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnNumberAndName() {
	if m.DoTrace() {
		/**/ m.PrintNumberAndName()
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintResume() {
	if m.DoTrace() {
		/**/ m.PrintNumberAndName()
		/**/ String2TraceFile(fmt.Sprintf("%s: RESUME", m.CurrentState))
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnResume() {
	if OVERALL_TRACE_FLAG || m.TraceFlag || m.ThisStateTraceFlag {
		/**/ m.PrintResume()
		//		if nil != m.Context.RetErr {
		//			/**/ String2TraceFile(fmt.Sprintf(", RetErr = %s", m.Context.RetErr))
		//		}
		/**/
		String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintNumberAndNameAndState() {
	/**/ m.PrintNumberAndName()
	/**/ String2TraceFile(fmt.Sprintf("%s: %s", m.CurrentState, m.A.StateComments[m.CurrentState]))
}

//------------------------------------------------------------
func (m *Machine) PrintlnNumberAndNameAndState() {
	/**/ m.PrintNumberAndNameAndState()
	/**/ String2TraceFile("\n")
}

//------------------------------------------------------------
// print the preamble of blanks plus the extra given amount (= ind) of blanks
// NB: this method need not be bound to m
func (m *Machine) NBlanks2TraceFile(tab int) {
	if OVERALL_TRACE_FLAG || m.TraceFlag || m.ThisStateTraceFlag {
		for i := 0; i < IND+tab; i++ {
			/**/ String2TraceFile(" ")
		}
	}
}

//------------------------------------------------------------
func (m *Machine) PrintContext(tl TraceLevelEnum, tab int) {
	if tl.DoTrace() && m.DoTrace() {
		m.Context.Print(tab)
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnContext(tl TraceLevelEnum, tab int) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.PrintContext(tl, tab)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
// shall i print trace for this machine?
func (m *Machine) DoTrace() bool {
	if OVERALL_TRACE_FLAG || m.TraceFlag || m.ThisStateTraceFlag {
		return true
	} else {
		return false
	}
}

//------------------------------------------------------------
// 0 args

//------------------------------------------------------------
func (m *Machine) PrintlnNotYet(tl TraceLevelEnum, tab int, a0 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ NBlanks2TraceFile(IND + tab)
		/**/ String2TraceFile(fmt.Sprintf("M%d: NOT YET: ", m.Number))
		/**/ String2TraceFile(a0)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
// @@@ wozu?
func (m *Machine) Panic(msg string) {
	/**/ Panic(msg)
}

//------------------------------------------------------------
// print a system message that starts with 3 stars
// does the padding...
// - @@@assumes the following consts for tab stops: 55 + 125
// caution: must not be used by model -> cyclic dependency otherwise in packages
func (m *Machine) PrintlnStarMessage(msgType MachineMessageTypeEnum, msgText string) {

	msg0 := fmt.Sprintf("*** %s ", msgType)
	msg0 = Padding(msg0, 4+11+1+3, "-")

	msg1 := fmt.Sprintf(" %s ", msgText)
	msg1 = Padding(msg1, 55, "-") // @@@

	msg2 := ""
	if nil != m {
		msg2 = fmt.Sprintf("-- %s__M%d[%s] ", m.Name, m.Number, m.CurrentState)
	}

	msg0_1_2 := Padding(fmt.Sprintf("%s%s%s", msg0, msg1, msg2), 125, "-") // @@@

	msg3 := fmt.Sprintf("-- CLOCK=%d", CLOCK)

	msg := fmt.Sprintf("%s%s\n", msg0_1_2, msg3)
	/**/ String2TraceFile(msg)
}

//------------------------------------------------------------
// should not occur, eg wrong model
func (m *Machine) UserError(err string) {
	/**/ m.PrintlnStarMessage(USER_ERROR, err)
	m.Panic(err)
}

//------------------------------------------------------------
// should not occur, system failure
func (m *Machine) SystemError(err string) {
	/**/ m.PrintlnStarMessage(SYSTEM_ERROR, err)
	m.Panic(err)
}

//------------------------------------------------------------
func (m *Machine) UserWarning(w string) {
	/**/ m.PrintlnStarMessage(USER_WARNING, w)
}

//------------------------------------------------------------
func (m *Machine) SystemWarning(w string) {
	/**/ m.PrintlnStarMessage(SYSTEM_WARNING, w)
}

//------------------------------------------------------------
func (m *Machine) SystemInfo(info string) {
	// /**/ String2TraceFile("\n")
	/**/
	m.PrintlnStarMessage(SYSTEM_INFO, info)
}

//------------------------------------------------------------
// for internal use only (help methods):

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_IndAndTab(tab int) {
	/**/ NBlanks2TraceFile(IND + tab)
}

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_LabelIsS(l string, a string) {
	if "" != l {
		/**/ String2TraceFile(fmt.Sprintf("%s = ", l))
	}
	/**/ String2TraceFile(fmt.Sprintf("%s", a))
}

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_LabelIsI(l string, a int) {
	if "" != l {
		/**/ String2TraceFile(fmt.Sprintf("%s = ", l))
	}
	/**/ String2TraceFile(fmt.Sprintf("%d", a))
}

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_LabelIsB(l string, a bool) {
	if "" != l {
		/**/ String2TraceFile(fmt.Sprintf("%s = ", l))
	}
	/**/ String2TraceFile(fmt.Sprintf("%t", a))
}

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_LabelIsY(l string, a interface{}) {
	if "" != l {
		/**/ String2TraceFile(fmt.Sprintf("%s = ", l))
	}
	// /**/ String2TraceFile(fmt.Sprintf("%s", a))
	/**/
	String2TraceFile(fmt.Sprintf("%v", a))
}

//------------------------------------------------------------
// caller must check if trace shall be printed
func (m *Machine) Print_LabelIsX(tab int, l string, x IPrint) {
	if "" != l {
		/**/ NBlanks2TraceFile(IND + tab)
		if (nil != x) && (!x.IsEmpty()) {
			/**/ String2TraceFile(fmt.Sprintf("%s = \n", l))
			/**/ x.Print(IND + tab*2)
		} else {
			/**/ String2TraceFile(fmt.Sprintf("%s = <empty>", l))
		}
	} else {
		if nil != x {
			/**/ x.Print(IND + tab)
		}
	}
}

//============================================================
//   S ... print string arg
//   I ... print int arg
//   B ... print bool arg
//   X ... print arg that implements IPrint interface
//   Y ... print arg that implements interface{} ... everything...@@@??? how does output look like?!
//         basic arg that has a String() method without any arg and can be printed in 1 line,
//         ie without passing them a tab parameter
//============================================================

//------------------------------------------------------------
func (m *Machine) Println(tl TraceLevelEnum) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
// no arg - just a string message
func (m *Machine) PrintlnString(tl TraceLevelEnum, tab int, s string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ String2TraceFile(fmt.Sprintf("%s", s))
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnS(tl TraceLevelEnum, tab int, l1 string, a1 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnY(tl TraceLevelEnum, tab int, l1 string, a1 interface{}) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsY(l1, a1)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnI(tl TraceLevelEnum, tab int, l1 string, a1 int) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsI(l1, a1)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnB(tl TraceLevelEnum, tab int, l1 string, a1 bool) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsB(l1, a1)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnX(tl TraceLevelEnum, tab int, l1 string, x1 IPrint) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_LabelIsX(tab, l1, x1)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSS(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSX(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 IPrint) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsX(tab, l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSI(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 int) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsI(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSB(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 bool) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsB(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnIS(tl TraceLevelEnum, tab int, l1 string, a1 int, l2 string, a2 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsI(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnII(tl TraceLevelEnum, tab int, l1 string, a1 int, l2 string, a2 int) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsI(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsI(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnXS(tl TraceLevelEnum, tab int, l1 string, x1 IPrint, l2 string, a2 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_LabelIsX(tab, l1, x1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnIY(tl TraceLevelEnum, tab int, l1 string, a1 int, l2 string, a2 interface{}) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_LabelIsI(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsY(l2, a2)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSSS(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 string, l3 string, a3 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l2, a2)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l3, a3)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSIS(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 int, l3 string, a3 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsI(l2, a2)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l3, a3)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnIII(tl TraceLevelEnum, tab int, l1 string, a1 int, l2 string, a2 int, l3 string, a3 int) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsI(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsI(l2, a2)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsI(l3, a3)
		/**/ String2TraceFile("\n")
	}
}

//------------------------------------------------------------
func (m *Machine) PrintlnSSSS(tl TraceLevelEnum, tab int, l1 string, a1 string, l2 string, a2 string, l3 string, a3 string, l4 string, a4 string) {
	if tl.DoTrace() && m.DoTrace() {
		/**/ m.Print_IndAndTab(tab)
		/**/ m.Print_LabelIsS(l1, a1)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l2, a2)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l3, a3)
		/**/ String2TraceFile(", ")
		/**/ m.Print_LabelIsS(l4, a4)
		/**/ String2TraceFile("\n")
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
