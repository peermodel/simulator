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
// Date: 2017
////////////////////////////////////////

package pmModel

import (
	. "cca/helpers"
	. "cca/scheduler"
	"strconv"
)

////////////////////////////////////////
// system function
// @@@ no args supported
////////////////////////////////////////

type SystemFunctionEnum int

const (
	CLOCK_FUNCTION SystemFunctionEnum = iota
	FID_FUNCTION
	UUID_FUNCTION
)

// output function in a readable form
func (f SystemFunctionEnum) String() string {
	switch f {
	case CLOCK_FUNCTION:
		return "clock()"
	case FID_FUNCTION:
		return "fid()"
	case UUID_FUNCTION:
		return "uuid()"
	default:
		return "ill. system function"
	}
}

////////////////////////////////////////
// methods
////////////////////////////////////////

var FID_PREFIX string = "f"
var UUID_PREFIX string = "u"
var FID_CNT int = 0
var UUID_CNT int = 0
var MAX_FID int = 2147483647  // int32
var MAX_UUID int = 2147483647 // int32

// --------------------------------------------
// new fid
func Fid() string {
	FID_CNT++
	if FID_CNT > MAX_FID {
		FID_CNT = 0
		/**/ SystemInfo("FID overflow captured")
	}
	return FID_PREFIX + strconv.Itoa(FID_CNT)
}

// --------------------------------------------
// new uuid
func UuidUserFu() string {
	UUID_CNT++
	if UUID_CNT > MAX_UUID {

		UUID_CNT = 0
		/**/ SystemInfo("UUID overflow captured")
	}
	return UUID_PREFIX + strconv.Itoa(UUID_CNT)
}

// --------------------------------------------
// get clock
func Clock() int {
	return CLOCK
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
