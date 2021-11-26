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
////////////////////////////////////////
// System: Peer Model State Machine
// Author: eva Kühn
// Date: 2015
////////////////////////////////////////

package pmModel

import (
	. "cca/debug"
)

// @@@ should better be a map! with key = eid
type Entries []Entry

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (es Entries) Copy() Entries {
	//------------------------------------------------------------
	// alloc
	newEs := []Entry{}
	//------------------------------------------------------------
	// copy entrie)
	for _, e := range es {
		newEPtr := e.Copy()
		newE := *newEPtr
		newEs = append(newEs, newE)
	}
	//------------------------------------------------------------
	// return
	return newEs
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (es Entries) IsEmpty() bool {
	if nil == es || len(es) == 0 || cap(es) == 0 {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (es Entries) Print(tab int) {
	/**/ NBlanks2TraceFile(tab)
	/**/ String2TraceFile("{")
	sep := ""
	nextTab := 0
	for _, e := range es {
		/**/ String2TraceFile(sep)
		/**/ e.Print(nextTab)
		sep = ", \n"
		nextTab = tab + TAB
	}
	/**/ String2TraceFile("}")
}

// ----------------------------------------
func (es Entries) Println(tab int) {
	/**/ es.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
