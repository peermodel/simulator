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
	. "github.com/peermodel/simulator/debug"
	"fmt"
)

type Xxx struct {
	Yyy int
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
func AllocXxx() *Xxx {
	x := new(Xxx)

	// x.Yyy = Yyy{}

	return x
}

// ----------------------------------------
func NewXxx() *Xxx {
	x := AllocXxx()

	// x.Yyy = Yyy{}

	return x
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (x *Xxx) IsEmpty() bool {
	if nil == x { // || len/cap(x) == 0
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// methods
////////////////////////////////////////

// ----------------------------------------
// deep copy
func (x *Xxx) Copy() *Xxx {
	newX := AllocXxx()

	// copy all fields:
	// - Yyy:
	newX.Yyy = x.Yyy

	return newX
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// --------------------------------------------
func (x *Xxx) Print(ind int) {
	/**/ NBlanks2TraceFile(ind)
	/**/ String2TraceFile(fmt.Sprintf("%d", x.Yyy))
}

// --------------------------------------------
func (x *Xxx) Println(ind int) {
	/**/ x.Print(ind)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
