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
// runtime with different execution modi
// - the actual execution mode must be configured
// - start the system with "Run"
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package runtime

import (
	. "cca/config"
	. "cca/debug"
	. "cca/framework"
	. "cca/latex"
	. "cca/scheduler"
	"fmt"
	"runtime"
)

//------------------------------------------------------------
// init and manage the test case run depending on execution and verification mode
// - nb: if Run finishes, all machines have stopped (by user, or by system ttl);
// - caution: initially, all machines must be created by the caller and reflected in the machine controls of status;
// -- ie: call NewTestCase (via testPreparation that generates the status) before running it,
// -- otherwise SIMULATION_COUNT, VERIFICATION_MODE etc. are not initialized;
func Run(s *Status, testCaseName string, testCaseLatexConfig LatexConfig, systemTtl int) {
	//------------------------------------------------------------
	// init debugging
	DebugInit()
	//------------------------------------------------------------
	// debug: print the entire status/space (only for the first run)
	// - s.ModelPrint(TRACE0, IND, testCaseId)
	// - s.SpacePrint(TRACE0, IND, false /* printAlsoEmptyContainersFlag */)
	//------------------------------------------------------------
	// generate latex docu (only for the first run)
	s.MetaContext.MetaModel2Latex(testCaseName, &testCaseLatexConfig)
	// ----------------------------------------
	// generate event-b xml (only for the first run)
	// - nb: Peer Model has not yet been implemented in Event-B
	// -- s.MetaContext.MetaModel2EventB(testCaseName, &testCaseLatexConfig)
	//------------------------------------------------------------
	// set clocks
	CLOCK = 0
	EVENT_CLOCK = 0
	//------------------------------------------------------------
	// execute the test case depending on execution and verification mode
	switch VERIFICATION_MODE {
	//------------------------------------------------------------
	case ONE_RUN:
		//============================================================
		// ONE_RUN
		//============================================================
		// run the test case exactly once
		//------------------------------------------------------------
		// debug:
		//   s.PrintMyGoRoutineId("RUN", "Runtime") // DEBUG
		//------------------------------------------------------------
		// run the test case:
		s.Run()
		//------------------------------------------------------------
		// debug:
		s.PrintStatistics() // DEBUG

	case SIMULATION:
		//============================================================
		// SIMULATION
		//============================================================
		// a major difference to model checking is that each simulation run starts at the beginning and
		// - not at a choice point, ie it is completely fresh initialized
		//------------------------------------------------------------
		nextS := s
		//------------------------------------------------------------
		for RUN_COUNT = 1; RUN_COUNT <= SIMULATION_COUNT; RUN_COUNT++ {
			//------------------------------------------------------------
			// debug:
			//............................................................
			if SIMULATION_TRACE.DoTrace() { // DEBUG
				Banner2TraceFile("\n\n", fmt.Sprintf("%d. MODEL SIMULATION RUN:\n", RUN_COUNT), "\n") // DEBUG
			} // DEBUG
			//............................................................
			// print my go routine id:
			//   s.PrintMyGoRoutineId("RUN", "Runtime") // DEBUG
			//............................................................
			// print machine controls:
			//   for _, mc := range nextS.MachineControls { // DEBUG
			//   	 fmt.Println(fmt.Sprintf("machine control = %s", mc.M.Key())) // DEBUG
			//   } // DEBUG
			//------------------------------------------------------------
			// run the test case
			nextS.Run()
			//------------------------------------------------------------
			// debug
			nextS.PrintStatistics() // DEBUG
			//------------------------------------------------------------
			// TBD: close all channels that still exist
			//------------------------------------------------------------
			// prepare everything for the next run:
			// - generate a completely fresh new status (as a copy of the orig metacontext copied from s);
			// - init the test case and start all machines (async) which also generates the machine controls;
			nextS = s.InitAppUseCaseFu()
			//------------------------------------------------------------
			// reset clocks
			CLOCK = 0
			EVENT_CLOCK = 0
		}

	case MODEL_CHECKING:
		//============================================================
		// MODEL CHECKING
		//============================================================
		//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
		//------------------------------------------------------------
		// INIT:
		//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
		//------------------------------------------------------------
		// init & clear global model checking vars shared with Controller
		InitModelCheckingVars()
		//------------------------------------------------------------
		// set next status to initial one
		nextS := s
		//------------------------------------------------------------
		// debug: local vars
		traceS := ""
		//------------------------------------------------------------
		// try out all possibilities (= paths)
		// - which are determined by wait points of async machines;
		// - however bounded by the configurable MC_BOUND;
		// - nb: a run (except of the first one) starts at a choice point which has a certain depth in the overall possibilities tree;
		for RUN_COUNT = 1; RUN_COUNT <= MC_BOUND; RUN_COUNT++ {
			//------------------------------------------------------------
			// reset global var that counts the number of space updates made by this run
			SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT = 0
			//------------------------------------------------------------
			// debug
			//............................................................
			traceS = fmt.Sprintf("%d. MODEL CHECKING PATH RUN:\n", RUN_COUNT)
			//............................................................
			if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
				if MODEL_CHECKING_DETAILS1_TRACE.DoTrace() { // DEBUG
					Banner2TraceFile("\n\n", traceS, "\n") // DEBUG
				} else { // DEBUG
					ThinBanner2TraceFile("\n\n", traceS, "\n") // DEBUG
				} //  DEBUG
			} // DEBUG
			//............................................................
			// print my go routine id:
			//   s.PrintMyGoRoutineId("RUN", "Runtime") // DEBUG
			//............................................................
			// print entire space, ie all containers, before the next run starts
			if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
				String2TraceFile(fmt.Sprintf("BACKTRACKING:     TO CP[%d] (depth=%d, t=%d, et=%d)\n",
					MC_VARS.CurChoicePoint.Id, MC_VARS.CurChoicePoint.Depth, MC_VARS.CurChoicePoint.Clock, MC_VARS.CurChoicePoint.EventClock)) // DEBUG
			} // DEBUG
			if MODEL_CHECKING_DETAILS1_TRACE.DoTrace() { // DEBUG
				nextS.MetaContext.SpacePrint(TRACE0, IND, false /* printAlsoEmptyContainersFlag */) // DEBUG
			} // DEBUG
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// RUN:
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// try the next path (represented by next status) until stopped (by user, or by system ttl)
			// - nb: the run may create new choice points
			nextS.Run()
			//------------------------------------------------------------
			// debug:
			// - info about run completion
			if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%d. RUN completed\n", RUN_COUNT)) // DEBUG
			} // DEBUG
			// - print entire status (incl. user app space data) and statistics after the run
			s.PrintStatistics() // DEBUG
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// DONE?
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// 1) have all choice points been tried?
			if 0 == len(MC_VARS.ChoicePoints) {
				//------------------------------------------------------------
				// debug
				if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
					/**/ String2TraceFile(fmt.Sprintf("MODEL CHECKING COMPLETED: all paths have been tried, RUN_COUNT=%d\n", RUN_COUNT)) // DEBUG
				} // DEBUG
				//------------------------------------------------------------
				// finished -> all paths have been tried
				break
			}
			//------------------------------------------------------------
			// 2) has model checking bound been reached?
			if RUN_COUNT >= MC_BOUND {
				//------------------------------------------------------------
				// debug
				if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
					/**/ String2TraceFile(fmt.Sprintf("MODEL CHECKING END: MC_BOUND=%d reached\n", RUN_COUNT)) // DEBUG
				} // DEBUG
				//------------------------------------------------------------
				// finished -> mc bound reached
				break
			}
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// 3) do BACKTRACKING: there are still choices
			//_._._._._._._._._._._._._._._._._._._._._._._._._._._._._._.
			//------------------------------------------------------------
			// get next choice (= machine key) and its updated CP (ie without the selected choice) from CP list:
			// - possibly also removing the CP from the global CPs collection, if it was its last choice;
			// - nb: there must exist a choice (because there is at least one choice point open -- see check 1) above);
			// - nb: the chosen choice (and possibly also its CP) is removed from CP list;
			MC_VARS.CurChoice, MC_VARS.CurChoicePoint = MC_VARS.ChoicePoints.EasyFetchChoice()
			//------------------------------------------------------------
			// set flag that next run starts with a recovered choice
			// - caution: flag must be set every new run, because Controller resets the variable!
			MC_VARS.UseCurChoiceAsNextMachineFlag = true
			//------------------------------------------------------------
			// reset clocks:
			CLOCK = MC_VARS.CurChoicePoint.Clock
			EVENT_CLOCK = MC_VARS.CurChoicePoint.EventClock
			//------------------------------------------------------------
			// debug
			MC_VARS.CurPathDepth = MC_VARS.CurChoicePoint.Depth
			//------------------------------------------------------------
			// debug:
			if MODEL_CHECKING_TRACE.DoTrace() { // DEBUG
				//............................................................
				// recovery of a choice from a CP:
				// - TBD: trace is redundant
				traceS = fmt.Sprintf("RECOVER CHOICE:   FROM CP[%d] (depth=%d, t=%d, et=%d):",
					MC_VARS.CurChoicePoint.Id, MC_VARS.CurChoicePoint.Depth, MC_VARS.CurChoicePoint.Clock, MC_VARS.CurChoicePoint.EventClock) // DEBUG
				if MODEL_CHECKING_DETAILS3_TRACE.DoTrace() { // DEBUG
					String2TraceFile(fmt.Sprintf("\n%s\n  %s\n", traceS, MC_VARS.CurChoice)) // DEBUG
				}
				//............................................................
				if len(MC_VARS.CurChoicePoint.Choices) == 0 { // DEBUG
					//............................................................
					// CP removed:
					String2TraceFile(fmt.Sprintf("REMOVE CP:        CP[%d]\n", MC_VARS.CurChoicePoint.Id)) // DEBUG
				} else { // DEBUG
					//............................................................
					// CP updated:
					if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
						String2TraceFile(fmt.Sprintf("UPDATED CP[%d]: %s\n", MC_VARS.CurChoicePoint.Id, MC_VARS.CurChoicePoint.ToString(0))) // DEBUG
					} // DEBUG
				} // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// set next status to cur CP's status:
			// - optimization if it was the last choice of this CP, ie CP was removed
			if len(MC_VARS.CurChoicePoint.Choices) == 0 {
				//------------------------------------------------------------
				// reuse the status of the CP
				nextS = MC_VARS.CurChoicePoint.S
			} else {
				//------------------------------------------------------------
				// the CP still has open choices
				// - copy/share/renew status for next run (starting with CP's status);
				// - nb: copies also machine controls (incl. wait conditions and machines with their LVS);
				nextS = MC_VARS.CurChoicePoint.S.Clone4NewRun(SYSTEM_TTL)
			}
			//------------------------------------------------------------
			nextS.StatusMutex.RLock() // LOCK FOR READ //
			//------------------------------------------------------------
			// for the copied LVSs in the copied machines of the copied status:
			// - the aliases need to be recomputed
			// - the function for doing this is found in the shared automaton A
			for _, mc := range nextS.MachineControls {
				mc.M.LocalVariables = mc.M.A.CompleteLocalVariablesAliasFunction(nextS, mc.M, mc.M.LocalVariables)
			}
			//------------------------------------------------------------
			// start *all* (async) machines maintained by machine controls:
			// - nb: start machine under old machine number; we have completely new channels and the old
			// -- go routines will fade after they got their stop signal; even if they continue
			// --- for a while it is not a problem, because we operate on deeply copied status, lvss and new channels
			//------------------------------------------------------------
			// debug
			if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("START %d MACHINES:\n", len(nextS.MachineControls))) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// start machines
			for key, mc := range nextS.MachineControls {
				//------------------------------------------------------------
				// assertion
				if mc.Condition == nil {
					s.SystemError(fmt.Sprintf("machine %s is in machine controls but has empty condition (f)", key))
				}
				//------------------------------------------------------------
				// debug
				if MODEL_CHECKING_DETAILS2_TRACE.DoTrace() { // DEBUG
					/**/ String2TraceFile(fmt.Sprintf("  %s (state %s)\n", mc.M.Key(), mc.M.CurrentState)) // DEBUG
				} // DEBUG
				//------------------------------------------------------------
				// complicated:
				// - if the machine is waiting, set flag that the machine is interrupted by a choice point;
				// - this causes that the original wait4 event of the machine is used to re-enter execution,
				// - ie the machine goes again into the same wait state; the flag indicates that here the
				// - machine must skip the wait4 condition; this is crucial, because re-issuing the wait4
				// - condition would otherwise be done with a too high event issue time and then the mmachine
				// - cannot be waked up by an event that happened before machine resumption after the choice point!
				// caution: check that all wait4 events are mentioned here
				if mc.Condition.Type == NO_EVENT || mc.Condition.Type == USER_EVENT || mc.Condition.Type == TIME_EVENT {
					//------------------------------------------------------------
					// !!! do not generate a new choice for this machine -- that was done already !!!
					mc.Condition.GenerateChoiceFlag = false
					//------------------------------------------------------------
					// !!! signal that this machine shall skip the LEAVE/ENTER sections in wait fu after being allowed to execute
					// - due to the original wait4 event
					mc.Condition.WaitInterruptedByCP_Flag = true
				}
				//------------------------------------------------------------
				// start async execution of machine
				if ASYNC == mc.M.StartType {
					go mc.M.Execute(nextS, mc)
				} else {
					//------------------------------------------------------------
					// assertion:
					// - sync machine: this should not happen
					s.SystemError(fmt.Sprintf("sync machine (for automaton = %s) found in machine controls", mc.M.A.Name))
				}
			}
			//------------------------------------------------------------
			nextS.StatusMutex.RUnlock() // UNLOCK FOR READ //
			//------------------------------------------------------------
			// debug:
			// - info that all machines were started
			if MODEL_CHECKING_DETAILS3_TRACE.DoTrace() { // DEBUG
				String2TraceFile(fmt.Sprintf("-> all machines started\n")) // DEBUG
			} // DEBUG
			// // TBD: redundant
			// // - print space on which the next MC run starts
			// if MODEL_CHECKING_DETAILS1_TRACE.DoTrace() { // DEBUG
			// 	nextS.MetaContext.SpacePrint(TRACE0, TAB, false) // DEBUG
			// } // DEBUG
			//------------------------------------------------------------
			// TBD: do garbage collection
			runtime.GC()
		}

	default:
		//============================================================
		// DEFAULT
		//============================================================
		Panic("ill. execution mode")
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
