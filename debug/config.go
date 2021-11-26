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
////////////////////////////////////////

package debug

// general trace flag
var OVERALL_TRACE_FLAG bool = false

// StatesTraceFlags
type STF map[string]bool

type TraceFlags struct {
	MFlag  bool
	SFlags STF
}

// max number of chars plus 1 that a machine name should possess
var MAX_MACHINE_NAME_CHARS int

// trace infos:
var IND int

const TAB int = 2

var IND_BLANKS string
var TAB_BLANKS string
var SERVICE_START_INFO string
var SERVICE_END_INFO string
var SERVICE_BLANKS string

////////////////////////////////////////
// configuration of log type
////////////////////////////////////////

var TRACES = map[TraceLevelEnum]bool{
	// the trace* levels control the machine rtaces, too:
	TRACE0:                        true,
	TRACE1:                        false,
	TRACE2:                        false,
	TRACE3:                        false,
	TRACE4:                        false,
	TRACE5:                        false,
	ARGS_EVAL_TRACE:               false,
	CONTROLLER_TRACE:              false,
	EVENT_CONDITION_TRACE:         false,
	INIT_TRACE:                    false,
	MACHINE_START_TRACE:           false,
	MODEL_CHECKING_TRACE:          false, // basic info about run and CPs
	MODEL_CHECKING_DETAILS1_TRACE: false, // extends MODEL_CHECKING_TRACE: space on which the next MC run starts etc.
	MODEL_CHECKING_DETAILS2_TRACE: false, // extends MODEL_CHECKING_DETAILS1_TRACE: m keys, state after each state change of a machine
	MODEL_CHECKING_DETAILS3_TRACE: false, // extends MODEL_CHECKING_DETAILS2_TRACE
	MODEL_CHECKING_DETAILS4_TRACE: false, // extends MODEL_CHECKING_DETAILS3_TRACE TBD: which wirings are entered
	QUERY_TRACE:                   false, // info about query and whether it was fulfilled and how many entries were read
	REPLAY_TRACE:                  false,
	RUN_TRACE:                     false,
	SCHEDULER_DETAILS_TRACE:       false,
	SCHEDULER_TRACE:               false,
	SERVICE_TRACE:                 true, // @@@ was ist der unterschied zu trace für state 37 von wiring?
	SIMULATION_TRACE:              true,
	STATISTICS_TRACE:              true,
}

////////////////////////////////////////
// configuration of trace flags
// caution: keep consistemt: ie store alle machine names here...
// example syntax for one item: "PccRead": {false, STF{"init": true, "1": true}},
// @@@ automata dependent
// @@@ istate number comments are not correct... dependent on profiles...
////////////////////////////////////////
var MACHINE_TRACE_FLAGS = map[string]TraceFlags{
	"PccRead": {false,
		STF{"init": false,
			"8":    false,
			"9":    false,
			"10":   false,
			"11":   false, // select next entry
			"12":   false,
			"13":   false,
			"17":   false,
			"23":   false,
			"25":   false, // init vars
			"exit": false,
		}},
	"PccTxCommit": {false,
		STF{"1": false,
			"2": false}},
	"PccWrite": {false,
		STF{}},
	"PeerModelInit": {false,
		STF{}},
	"Read": {false,
		STF{"init": false,
			"1": false, // check query count and print query
			"4": false, // select entry
		}},
	"Service": {false,
		STF{}},
	"SpaceCreateTx": {false,
		STF{}},
	"SpaceRead": {false,
		STF{}},
	"SpaceUndo": {false,
		STF{}},
	"SpaceUndoRead": {false,
		STF{}},
	"SpaceUndoWrite": {false,
		STF{}},
	"SpaceTxCommit": {false,
		STF{}},
	"SpaceWrite": {false,
		STF{"init": false,
			"1": false,
			"4": false,
		}},
	"TestCase": {false,
		STF{}},
	"Time": {false,
		STF{}},
	"Wiring": {false,
		STF{"init": false,
			"12":   false,
			"13":   false, // wait 4 user event / container changed
			"15":   false,
			"16":   false, // space read
			"58":   false,
			"exit": false,
		}},
	"Write": {false,
		STF{}},
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
