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
////////////////////////////////////////

package debug

import (
	"os"
)

// general trace file, e.g. "trace.log" or "stdout"
const TRACE_FILE_NAME string = "" // <<<<<<<<<<<<<<<<<<<<
// default = stdout
var TRACE_FILE *os.File = os.Stdout

// model checking trace file, e.g. "model_checking.log" or "stdout"
const MODEL_CHECKING_TRACE_FILE_NAME string = "model_checking.log" // <<<<<<<<<<<<<<<<<<<<
// default = stdout
var MODEL_CHECKING_TRACE_FILE *os.File = os.Stdout

// replay trace file, e.g. "replay.log" or "stdout"
const REPLAY_TRACE_FILE_NAME string = "replay.log" // <<<<<<<<<<<<<<<<<<<<
// default = stdout
var REPLAY_TRACE_FILE *os.File = os.Stdout

// @@@ gehört nicht hierher !! gehört ins framework !!
// latex file for output of meta model of use case, e.g.  "use_case.tex"or "stdout"
const LATEX_FILE_NAME string = "use_case.tex" // <<<<<<<<<<<<<<<<<<<<
// default = stdout
var LATEX_FILE *os.File = os.Stdout

////////////////////////////////////////
// ENUMs
////////////////////////////////////////

////////////////////////////////////////
// trace level
// set them (as many as you like) in TRACES map!
// @@@ create new package "debug" for trace issues!
////////////////////////////////////////

type TraceLevelEnum int

const (
	TRACE0 TraceLevelEnum = iota
	TRACE1
	TRACE2
	TRACE3
	TRACE4
	TRACE5
	ARGS_EVAL_TRACE
	CONTROLLER_TRACE
	EVENT_CONDITION_TRACE
	INIT_TRACE
	MACHINE_START_TRACE
	MODEL_CHECKING_TRACE
	MODEL_CHECKING_DETAILS1_TRACE // adds info to MODEL_CHECKING_TRACE
	MODEL_CHECKING_DETAILS2_TRACE // adds info to MODEL_CHECKING_DETAILS1_TRACE
	MODEL_CHECKING_DETAILS3_TRACE // adds info to MODEL_CHECKING_DETAILS2_TRACE
	MODEL_CHECKING_DETAILS4_TRACE // adds info to MODEL_CHECKING_DETAILS4_TRACE
	QUERY_TRACE
	REPLAY_TRACE
	RUN_TRACE
	SCHEDULER_DETAILS_TRACE
	SCHEDULER_TRACE
	SERVICE_TRACE
	SIMULATION_TRACE
	STATISTICS_TRACE
)

func (t TraceLevelEnum) String() string {
	switch t {
	case TRACE0:
		return "TRACE0"
	case TRACE1:
		return "TRACE1"
	case TRACE2:
		return "TRACE2"
	case TRACE3:
		return "TRACE3"
	case TRACE4:
		return "TRACE4"
	case TRACE5:
		return "TRACE5"
	case ARGS_EVAL_TRACE:
		return "ARGS_EVAL_TRACE"
	case CONTROLLER_TRACE:
		return "CONTROLER_TRACE"
	case EVENT_CONDITION_TRACE:
		return "EVENT_CONDITION_TRACE"
	case INIT_TRACE:
		return "INIT_TRACE"
	case MACHINE_START_TRACE:
		return "MACHINE_START_TRACE"
	case MODEL_CHECKING_TRACE:
		return "MODEL_CHECKING_TRACE"
	case MODEL_CHECKING_DETAILS1_TRACE:
		return "MODEL_CHECKING_DETAILS1_TRACE"
	case MODEL_CHECKING_DETAILS2_TRACE:
		return "MODEL_CHECKING_DETAILS2_TRACE"
	case MODEL_CHECKING_DETAILS3_TRACE:
		return "MODEL_CHECKING_DETAILS3_TRACE"
	case MODEL_CHECKING_DETAILS4_TRACE:
		return "MODEL_CHECKING_DETAILS4_TRACE"
	case QUERY_TRACE:
		return "QUERY_TRACE"
	case REPLAY_TRACE:
		return "REPLAY_TRACE"
	case RUN_TRACE:
		return "RUN_TRACE"
	case SCHEDULER_DETAILS_TRACE:
		return "SCHEDULER_DETAILS_TRACE"
	case SCHEDULER_TRACE:
		return "SCHEDULER_TRACE"
	case SERVICE_TRACE:
		return "SERVICE_TRACE"
	case SIMULATION_TRACE:
		return "SIMULATION_TRACE"
	case STATISTICS_TRACE:
		return "STATISTICS_TRACE"
	default:
		return "ill. log type"
	}
}

//--------------------------------------
// shall i print trace for this trace level?
// @@@ eigenes file
func (tl TraceLevelEnum) DoTrace() bool {
	if true == TRACES[tl] {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
