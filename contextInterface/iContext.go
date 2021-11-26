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
// Copyright: eva Kuehn
// 2016
//------------------------------------------------------------
// context interface
// - used for communication between machines
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// TBD:
// - comment existed: "caution: non-pointer receivers requested!!" ... what is it for?
//////////////////////////////////////////////////////////////

package contextInterface

//============================================================
// interface: i/o parameters between machines
//============================================================

//------------------------------------------------------------
type IContext interface {
	// deep copy
	Copy() interface{}
	// machine identifier used as suffix
	MachineKeySuffix() string
	// require also the IPrint interface ...
	IsEmpty() bool
	Print(ind int)
	Println(ind int)
	// trick: just for code generator to guarantee that this package is used...
	DummyIContextFu() interface{}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
