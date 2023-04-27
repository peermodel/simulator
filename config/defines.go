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
// Date: 2015, 2016
//------------------------------------------------------------
// configuration: set defaults here!
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// - TBD: explain better the semantics of TICK_FREQUENCY
// - TBD: explain better the semantics of the execution modes
//------------------------------------------------------------
//////////////////////////////////////////////////////////////

package config

//////////////////////////////////////////////////////////////
// set (default) values of configuration consts
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// overall system time aka "system ttl (time to live)"
// - default = 100000
const SYSTEM_TTL int = 1000002

//------------------------------------------------------------
// increment time after N times; or if no condition is fulfilled
// - default = 1
const TICK_FREQUENCY int = 1

//------------------------------------------------------------
// determines number of simulation runs
const SIMULATION_COUNT int = 3

//------------------------------------------------------------
// limits the number of model checking iterations runs
const MC_BOUND int = 1000

//------------------------------------------------------------
// verification mode
//............................................................
// ONE_RUN
// - only one test run
//............................................................
// SIMULATION
// - uses simulation count (above)
//............................................................
// MODEL_CHECKING
// - still under construction
//............................................................
const VERIFICATION_MODE VerificationTypeEnum = ONE_RUN // <<<<<<<<<<<<<<<<<<<<<<<<

//------------------------------------------------------------
// execution mode
//............................................................
// MIN_ISSUE_TIME
// - select a machine whose issue time (= condition time) is smallest
// - simple and slow
//............................................................
// MIN_ISSUE_TIME_AND_CONDITION_FULFILLED
// - select the one whose issue time (= condition time) is smallest and condition fulfilled
// - quite fast
//............................................................
// FAIRNESS
// - select the one whose condition is fulfilled and whose last try to execute is furthest behind;
// - faster and fairer than the above ones -- simply the best!
//............................................................
const EXECUTION_MODE ExecutionTypeEnum = FAIRNESS // <<<<<<<<<<<<<<<<<<<<<<<<

//------------------------------------------------------------
// criterion of model checker: how to select choice point
//............................................................
// - FIRST_KEY / RANDOM_KEY
const MC_CP_KEY_SELECTION_CRITERION ChoiceSelectionCriterionTypeEnum = RANDOM_KEY

//............................................................
// - FIRST_TIME / RANDOM_TIME
const MC_CP_TIME_SELECTION_CRITERION ChoiceSelectionCriterionTypeEnum = FIRST_TIME

//------------------------------------------------------------
// number of space updates made by this run
// - TBD: create interface for this var
var SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT int = 0

//////////////////////////////////////////////////////////////
// enums
//////////////////////////////////////////////////////////////

//============================================================
// verification type
//============================================================

//------------------------------------------------------------
type VerificationTypeEnum int

//------------------------------------------------------------
const (
	ONE_RUN VerificationTypeEnum = iota
	SIMULATION
	MODEL_CHECKING
)

//------------------------------------------------------------
func (t VerificationTypeEnum) String() string {
	switch t {
	case ONE_RUN:
		return "ONE_RUN"
	case SIMULATION:
		return "SIMULATION"
	case MODEL_CHECKING:
		return "MODEL_CHECKING"
	default:
		return "ill. verification type"
	}
}

//============================================================
// execution type
//============================================================

//------------------------------------------------------------
type ExecutionTypeEnum int

//------------------------------------------------------------
const (
	MIN_ISSUE_TIME ExecutionTypeEnum = iota
	MIN_ISSUE_TIME_AND_CONDITION_FULFILLED
	FAIRNESS
)

//------------------------------------------------------------
func (t ExecutionTypeEnum) String() string {
	switch t {
	case MIN_ISSUE_TIME:
		return "MIN_ISSUE_TIME"
	case MIN_ISSUE_TIME_AND_CONDITION_FULFILLED:
		return "MIN_ISSUE_TIME_AND_CONDITION_FULFILLED"
	case FAIRNESS:
		return "FAIRNESS"
	default:
		return "ill. execution type"
	}
}

//============================================================
// choice selection criterion type
//============================================================

//------------------------------------------------------------
type ChoiceSelectionCriterionTypeEnum int

//------------------------------------------------------------
const (
	FIRST_KEY ChoiceSelectionCriterionTypeEnum = iota
	RANDOM_KEY
	FIRST_TIME
	RANDOM_TIME
)

//------------------------------------------------------------
func (t ChoiceSelectionCriterionTypeEnum) String() string {
	switch t {
	case FIRST_KEY:
		return "FIRST_KEY"
	case RANDOM_KEY:
		return "RANDOM_KEY"
	case FIRST_TIME:
		return "FIRST_TIME"
	case RANDOM_TIME:
		return "RANDOM_TIME"
	default:
		return "ill. choice selection criterion type"
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
