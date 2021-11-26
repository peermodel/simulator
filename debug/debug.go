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
////////////////////////////////////////

// debug traces
// turn them on or off
// overall -> for everything
// per machine (if set to true it overrides overall trace flag if set to false)
// per state (if set to true it overrides machine trace flag if set to false)
package debug

import (
	"fmt"
	"os"
)

////////////////////////////////////////
// configuration vars and consts of printing of Arg and Args
// @@@ model dependent
// can be set before a trace is issued
////////////////////////////////////////

var PRINT_ARGS_NAME string = ""
var PRINT_ARG_DETAILS_FLAG bool = false
var PRINT_ARG_TYPE_FLAG bool = false
var PRINT_OMIT_DEFAULTS_FLAG = false // @@@ überall verwendet?!

////////////////////////////////////////
// functions
////////////////////////////////////////

// ----------------------------------------
// keep consistent also with trace setting in machine!!
func DebugInit() {
	var err error

	// -----------------------------
	// MAX_MACHINE_NAME_CHARS
	for nm, _ := range MACHINE_TRACE_FLAGS {
		if len(nm) > MAX_MACHINE_NAME_CHARS {
			MAX_MACHINE_NAME_CHARS = len(nm)
		}
	}
	// IND
	IND = 1 + 1 + 5 + 1 + MAX_MACHINE_NAME_CHARS
	// init strings for blanks:
	IND_BLANKS = ""
	for i := 0; i < IND; i++ {
		IND_BLANKS += " "
	}
	TAB_BLANKS = ""
	for j := 0; j < TAB; j++ {
		TAB_BLANKS += " "
	}
	SERVICE_BLANKS = fmt.Sprintf("%s%s=== %s", TAB_BLANKS, IND_BLANKS, TAB_BLANKS)
	SERVICE_START_INFO = fmt.Sprintf("%s=== SERVICE", TAB_BLANKS)
	SERVICE_END_INFO = fmt.Sprintf("%s=== %s END", TAB_BLANKS, TAB_BLANKS)

	// -----------------------------
	// open general trace file
	// trick: if trace file name = "stdout" -> use stdout which is set by default
	if "stdout" != TRACE_FILE_NAME && "" != TRACE_FILE_NAME {
		TRACE_FILE, err = os.Create(TRACE_FILE_NAME)
		if err != nil {
			/**/ panic(err)
			return
		}
	}
	// -----------------------------
	// open model checking trace file
	// trick: if trace file name = "stdout" -> use stdout which is set by default
	if "stdout" != MODEL_CHECKING_TRACE_FILE_NAME && "" != MODEL_CHECKING_TRACE_FILE_NAME {
		MODEL_CHECKING_TRACE_FILE, err = os.Create(MODEL_CHECKING_TRACE_FILE_NAME)
		if err != nil {
			/**/ panic(err)
			return
		}
	}
	// -----------------------------
	// open replay trace file
	// trick: if trace file name = "stdout" -> use stdout which is set by default
	if "stdout" != REPLAY_TRACE_FILE_NAME && "" != REPLAY_TRACE_FILE_NAME {
		REPLAY_TRACE_FILE, err = os.Create(REPLAY_TRACE_FILE_NAME)
		if err != nil {
			/**/ panic(err)
			return
		}
	}
	// -----------------------------
	// open latex file
	// @@@ gehört nicht hierher !! gehört ins framework !!
	LATEX_FILE, err = os.Create(LATEX_FILE_NAME)
	if err != nil {
		/**/ panic(err)
		return
	}
}

// ----------------------------------------
// add n blanks at the end of a string and return the new string
func NBlanksToString(inString string, nBlanks int) string {
	blankString := ""
	for i := 0; i < nBlanks; i++ {
		blankString = fmt.Sprintf("%s ", blankString)
	}
	return fmt.Sprintf("%s%s", inString, blankString)
}

// ----------------------------------------
// write n blanks to trace file
// @@@ eliminate this function...
func NBlanks2TraceFile(nBlanks int) {
	s := ""
	for i := 0; i < nBlanks; i++ {
		s = fmt.Sprintf("%s ", s)
	}
	/**/ String2TraceFile(s)
}

// ----------------------------------------
// write a string to general trace file
func String2TraceFile(s string) {
	/**/ TRACE_FILE.WriteString(s)
}

// ----------------------------------------
// write a string to model checking trace file
func String2MCTraceFile(s string) {
	/**/ MODEL_CHECKING_TRACE_FILE.WriteString(s)
}

// ----------------------------------------
// write a string to replay trace file
func String2ReplayTraceFile(s string) {
	/**/ REPLAY_TRACE_FILE.WriteString(s)
}

// ----------------------------------------
// write a string surrounded by banner to general trace file
func Banner2TraceFile(sBefore string, s string, sAfter string) {
	// ----------------------------------------
	String2TraceFile(sBefore)
	// ----------------------------------------
	SlashBorderLine2TraceFile()
	SlashBorderLine2TraceFile()
	// ----------------------------------------
	String2TraceFile(s)
	// ----------------------------------------
	SlashBorderLine2TraceFile()
	SlashBorderLine2TraceFile()
	// ----------------------------------------
	String2TraceFile(sAfter)
}

// ----------------------------------------
// write a string surrounded by banner to general trace file
func ThinBanner2TraceFile(sBefore string, s string, sAfter string) {
	// ----------------------------------------
	String2TraceFile(sBefore)
	// ----------------------------------------
	SlashBorderLine2TraceFile()
	// ----------------------------------------
	String2TraceFile(s)
	// ----------------------------------------
	SlashBorderLine2TraceFile()
	// ----------------------------------------
	String2TraceFile(sAfter)
}

// ----------------------------------------
// write a star border line
func StarBorderLine2TraceFile() {
	String2TraceFile(fmt.Sprintf("**********************************************************************\n"))
}

// ----------------------------------------
// write a slash border line
func SlashBorderLine2TraceFile() {
	String2TraceFile(fmt.Sprintf("//////////////////////////////////////////////////////////////////////\n"))
}

// ----------------------------------------
// write an equal border line
func EqualBorderLine2TraceFile() {
	String2TraceFile(fmt.Sprintf("======================================================================\n"))
}

// ----------------------------------------
// write a dash border line
func DashBorderLine2TraceFile() {
	String2TraceFile(fmt.Sprintf("----------------------------------------------------------------------\n"))
}

// ----------------------------------------
// write a dot border line
func DotBorderLine2TraceFile() {
	String2TraceFile(fmt.Sprintf("......................................................................\n"))
}

// ----------------------------------------
// write a border line with lineLen chars;
// unused
func BorderLine2TraceFile(char byte, lineLen int) {
	for i := 0; i < lineLen; i++ {
		String2TraceFile(fmt.Sprintf("%s", char))
	}
	String2TraceFile("\n")
}

////////////////////////////////////////
// Panic
////////////////////////////////////////

// ----------------------------------------
func Panic(s string) {
	// ----------------------------------------
	String2TraceFile("\n\n\n")
	StarBorderLine2TraceFile()
	StarBorderLine2TraceFile()
	// ----------------------------------------
	String2TraceFile(fmt.Sprintf("PANIC: %s\n", s))
	// ----------------------------------------
	StarBorderLine2TraceFile()
	StarBorderLine2TraceFile()
	String2TraceFile("\n\n\n")
	// ----------------------------------------
	panic("-- THE END --")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
