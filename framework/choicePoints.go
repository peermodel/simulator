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
/// System: Peer Model State Machine
// Author: eva Kühn
// Date: 2015
//------------------------------------------------------------
// for model checking mode:
// - choice points: data type and logic
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package framework

import (
	//	. "cca/config"
	. "cca/debug"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// slice of choice points
type ChoicePoints []*ChoicePoint

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create new choice points
func NewChoicePoints() *ChoicePoints {
	//------------------------------------------------------------
	// alloc
	cps := new(ChoicePoints)
	//------------------------------------------------------------
	// return
	return cps
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// fetch a choice and remove it from the corresponding choice points,
// - and possibly also remove its choice point if it becomes empty;
// returns:
// - key  ... the selected machine key
// - cp   ... the selected & updated choice point (ie the selected choice is removed);
//            is needed by caller; caution might be removed from choice points by this function (and empty)
// CAUTION: caller must assure that there are still choices when calling this fu
func (cpsPtr *ChoicePoints) EasyFetchChoice() (string, *ChoicePoint) {
	//------------------------------------------------------------
	// local vars
	theCpNo := -1
	//------------------------------------------------------------
	// ret vars
	var retCp *ChoicePoint
	retKey := ""
	//------------------------------------------------------------
	// go through all choice points:
	// - get
	cps := *cpsPtr
	// - try to fetch a choice
	for i := 0; i < len(cps); i++ {
		// nb: fetch updates the CP, i.e. removes the choice and possibly also its key in the CP's map, if found:
		retKey = cps[i].EasyFetchChoice()
		if retKey != "" {
			// set return vals
			theCpNo = i
			retCp = cps[i]
			break
		}
	}
	if retKey == "" {
		Panic("corrupted choice point list: no open choice found")
	}
	//------------------------------------------------------------
	// if CP's choices is empty -> remove CP from choice points:
	if 0 == len(retCp.Choices) {
		cps = append(cps[:theCpNo], cps[theCpNo+1:]...)
	}
	//------------------------------------------------------------
	// set output parameter:
	*cpsPtr = cps
	//------------------------------------------------------------
	// return fetched choice et al.
	return retKey, retCp
}

////------------------------------------------------------------
//// fetch a choice and remove it from the corresponding choice points,
//// - and possibly also remove its choice point if it becomes empty;
//// consider the criteria:
//// - keyCriterion .... FIRST_KEY, RANDOM_KEY, ...
//// - timeCriterion ... FIRST_TIME, RANDOM_TIME, ...
//// returns:
//// - key ... the selected machine key
//// - int ... the selected time
//// - cp  ... the choice point. is needed by caller. might be removed from choice points by this function.
//// - cpNo... for docu only. the ex-number of the choice point in the list.
//// - cpWasRemovedFlag... info whether the CP was removed (i.e. has no more open choices)
//// nb: caller must take care that there are still choices
//func (cpsPtr *ChoicePoints) FetchChoice(keyCriterion ChoiceSelectionCriterionTypeEnum, timeCriterion ChoiceSelectionCriterionTypeEnum) (string, int, *ChoicePoint, int, bool) {
//	//------------------------------------------------------------
//	// local vars
//	foundFlag := false
//	//------------------------------------------------------------
//	// ret vars
//	var retCp *ChoicePoint
//	retKey := ""
//	retTime := -1
//	retCpWasRemovedFlag := false
//	retI := -1

//	//------------------------------------------------------------
//	// go through all choice points:
//	cps := *cpsPtr
//	for i := 0; i < len(cps); i++ {
//		// nb: FetchChoice updates the CP, i.e. removes the choice and possibly also its key in the CP's map, if found:
//		retKey, retTime, foundFlag = cps[i].FetchChoice(keyCriterion, timeCriterion)
//		if foundFlag {
//			retI = i
//			break
//		}
//	}
//	if !foundFlag {
//		Panic("corrupted choice point list: no open choice found")
//	}
//	//------------------------------------------------------------
//	// set return CP:
//	//------------------------------------------------------------
//	retCp = cps[retI]
//	// if CP's choices is empty -> remove CP from choice points:
//	if 0 == len(retCp.Choices) {
//		cps = append(cps[:retI], cps[retI+1:]...)
//		retCpWasRemovedFlag = true
//	}
//	//------------------------------------------------------------
//	// set output parameter:
//	*cpsPtr = cps
//	//------------------------------------------------------------
//	// return fetched choice et al.
//	return retKey, retTime, retCp, retI, retCpWasRemovedFlag
//}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
func (cps ChoicePoints) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%sChoice Points:\n", s)
	for _, cp := range cps {
		s = fmt.Sprintf("%s%s\n", s, cp.ToString(tab+TAB))
	}
	return s
}

//------------------------------------------------------------
func (cps ChoicePoints) Print(tab int) {
	/**/ String2TraceFile(cps.ToString(tab))
}

//------------------------------------------------------------
// the same as Print
func (cps ChoicePoints) Println(tab int) {
	/**/ cps.Print(tab)
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
