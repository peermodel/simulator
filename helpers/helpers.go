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
// help functions like min, max, padding and error messages
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package helpers

import (
	"bytes"
	. "cca/debug"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

//============================================================
// min/max:
//============================================================

//------------------------------------------------------------
// return maximum
func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

//------------------------------------------------------------
// return minimum
func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

//============================================================
// padding:
//============================================================

//------------------------------------------------------------
// return text padded to padding len using padding char;
// if text len > padding len, just return text as is;
// caution: assumes that paddingchar is only one char....
func Padding(text string, paddingLen int, paddingChar string) string {
	for i := len(text); i < paddingLen; i++ {
		text = fmt.Sprintf("%s%s", text, paddingChar)
	}
	return text
}

//============================================================
// construct cid for given machine number:
//============================================================

//------------------------------------------------------------
// convert a cid to a cid that also contains the current wiring's machine number;
// - <cid>____M<wmno>
func ConvertCtoM(cid string, wmno int) *string {
	mcid := fmt.Sprintf("%s____M%d", cid, wmno)
	return &mcid
}

//============================================================
// type check:
//============================================================

//------------------------------------------------------------
// is given arg a pointer?
// - docu:
// -- https://golang.org/pkg/reflect/
// -- reflect.Struct ... for struct test
func IsPointer(arg interface{}) bool {
	// fmt.Println(fmt.Sprintf("%s is pointer = %v", reflect.ValueOf(arg).Type().Kind(), reflect.ValueOf(arg).Type().Kind() == reflect.Ptr)) // DEBUG
	return reflect.ValueOf(arg).Type().Kind() == reflect.Ptr
}

//------------------------------------------------------------
// debug
// - code from https://blog.sgmansfield.com/2015/12/goroutine-ids/
func GetGoRoutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//============================================================
// display error messages:
// - use the following only if no status and no machine is available:
// - @@@ on which file shall they print?
//============================================================

//------------------------------------------------------------
// should not occur, e.g. wrong model
func UserError(err string) {
	/**/ fmt.Println("USER_ERROR: ", err, "\n")
	Panic(err)
}

//------------------------------------------------------------
// should not occur, system failure
func SystemError(err string) {
	/**/ fmt.Println("SYSTEM_ERROR: ")
	Panic(err)
}

//------------------------------------------------------------
// user warning
func UserWarning(w string) {
	/**/ fmt.Println("USER_WARNING: ", w)
}

//------------------------------------------------------------
// can e.g. be used in automata for debugging...
func UserInfo(w string) {
	/**/ fmt.Println("  ", w)
}

//------------------------------------------------------------
// system warning
func SystemWarning(w string) {
	/**/ fmt.Println("SYSTEM_WARNING: ", w)
}

//------------------------------------------------------------
// system info
func SystemInfo(info string) {
	/**/ fmt.Println("SYSTEM_INFO: ", info)
}

//============================================================
// internal "trick":
//============================================================

//------------------------------------------------------------
// trick: just for code generator to guarantee that this package is included...
func DummyHelpersFu() {
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
