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
// help functions for strings
// - read about slices:
// -- blog.golang.org/slices
// -- https://stackoverflow.com/questions/42746972/golang-insert-to-a-sorted-slice
// -- https://golang.org/pkg/sort/#SearchStrings
// -- https://blog.golang.org/slices-intro
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package helpers

import (
	. "cca/debug"
	"fmt"
	"sort"
)

//////////////////////////////////////////////////////////////
// data types
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// Strings
// - slice of strings
type Strings []string

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// not deep copy
func (strings Strings) Copy() Strings {
	// create new slice of same length and capacity
	newStrings := make([]string, len(strings), cap(strings))
	// copy
	copy(newStrings, strings)
	// return
	return newStrings
}

//------------------------------------------------------------
// remove all strings whose values equals s
// - TBD: optimize: ie break from loop and copy the rest, if a string was found,
// -- assuming that there is at most one or exactly one s in strings
func (strings Strings) RemoveString(s string) Strings {
	// create new slice
	newStrings := Strings{}
	// iterate over strings
	for _, nextS := range strings {
		// copy, ie keep all strings whose value is not equal to s
		if s != nextS {
			newStrings = append(newStrings, nextS)
		}
	}
	// return
	return newStrings
}

//------------------------------------------------------------
// sorted insert
// - caution: caller must assign result to its strings var
// - cf: https://play.golang.org/p/iFzojVHSpq
func (strings Strings) SortedInsertString(s string) Strings {
	// search for index where element shall be inserted
	i := sort.SearchStrings(strings, s)
	// create one place at the end
	strings = append(strings, "")
	// shift everything to the right after the insert place
	// - copy(dest, source)
	copy(strings[i+1:], strings[i:])
	// insert element
	strings[i] = s
	// return
	return strings
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

//------------------------------------------------------------
// return true if slice is empty
func (strings Strings) IsEmpty() bool {
	// cap ... capacity of the slice
	if nil == strings || len(strings) == 0 || cap(strings) == 0 {
		return true
	} else {
		return false
	}
}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
func (strings Strings) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%s{", s)
	sep := ""
	for _, str := range strings {
		s = fmt.Sprintf("%s%s%s", s, sep, str)
		sep = ", "
	}
	s = fmt.Sprintf("%s}", s)
	return s
}

//------------------------------------------------------------
func (strings Strings) Print(tab int) {
	/**/ String2TraceFile(strings.ToString(tab))
}

//------------------------------------------------------------
func (strings Strings) Println(tab int) {
	/**/ strings.Print(tab)
	/**/ String2TraceFile("\n")
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
