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
// Copyright: eva Kuehn
// 2016
//------------------------------------------------------------
// interface for meta context
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// TBD:
// - evtl. aufteilen in context und meta context interface
// - comment existed: "caution: non-pointer receivers requested!!" ... what is it for?
//////////////////////////////////////////////////////////////

package contextInterface

import (
	. "cca/debug"
	. "cca/latex"
	. "cca/scheduler"
	. "cca/slotInterface"
)

//============================================================
// interface
//============================================================

//------------------------------------------------------------
type IMetaContext interface {
	// deep copy
	Copy() interface{}
	ProcessRipeUserSlot(userSlot ISlot, scheduler *Scheduler)
	MetaModel2Latex(testCaseName string, testCaseLatexConfig *LatexConfig)
	ConditionIsFulfilled(condition *Event) bool
	SpacePrint(tl TraceLevelEnum, nBlanks int, printAlsoEmptyContainersFlag bool)
	// require also the IPrint interface ...
	IsEmpty() bool
	Print(ind int)
	Println(ind int)
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
