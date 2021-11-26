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
// Date: 2015, 2016
//------------------------------------------------------------
// choice point: data type and logic
// - abbreviation: CP ... choice point
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
// - TBD: enhance FetchChoice: specify a selection criterion
//////////////////////////////////////////////////////////////

package framework

import (
	//	. "cca/config"
	. "cca/debug"
	"fmt"
	//	"math/rand"
	//	"time"
)

//------------------------------------------------------------
//////////////////////////////////////////////////////////////
// consts
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// choice point flags
const (
	CP    bool = true
	NO_CP bool = false
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// collects all possible choices (for machine selection) for a status
type ChoicePoint struct {
	//------------------------------------------------------------
	// system time when choice point is created
	Clock int
	//------------------------------------------------------------
	// event time when choice point is created
	EventClock int
	//------------------------------------------------------------
	// map with all open choices for a status:
	// - key .... machine key
	// - value .... unused
	// -- TBD: find a better solution
	// nb: tried choices are removed -- and if for a machine all times were tried, also the machine key is removed
	Choices map[string]int
	//------------------------------------------------------------
	// copy of the current status:
	// - incl. machine ctrls and machines
	S *Status
	//------------------------------------------------------------
	// path depth
	Depth int
	//------------------------------------------------------------
	// debug: CP id
	Id int
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// alloc and init a new choice point
func NewChoicePoint() *ChoicePoint {
	//------------------------------------------------------------
	// alloc
	cp := new(ChoicePoint)
	//------------------------------------------------------------
	// init & reset
	cp.Clock = 0
	cp.EventClock = 0
	cp.Choices = make(map[string]int)
	cp.S = nil
	cp.Depth = 0
	cp.Id = 0
	//------------------------------------------------------------
	// return
	return cp
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// copy/share for tracking of ancestor CPs for model checking
func (cp *ChoicePoint) Clone4AncestorTracking() *ChoicePoint {
	//------------------------------------------------------------
	// alloc
	newCp := NewChoicePoint()
	//------------------------------------------------------------
	// copy
	newCp.Clock = cp.Clock
	newCp.EventClock = cp.EventClock
	newCp.Clock = cp.Clock
	// key must be copied!!
	// - nb: val unused...
	for key, val := range cp.Choices {
		newCp.Choices[key] = val
	}
	// shared!
	newCp.S = cp.S
	newCp.Id = cp.Id
	newCp.Clock = cp.Clock
	//------------------------------------------------------------
	// return
	return newCp
}

//------------------------------------------------------------
// check if a machine key exists in the map
func (cp *ChoicePoint) ContainsKey(key string) bool {
	// try to fetch entry with key
	_, found := cp.Choices[key]
	// return if found
	return found
}

// //------------------------------------------------------------
// // deprecated:
// // check if a machine key exists in the map
// // - and also for all ancestors recursively
// func (cp *ChoicePoint) ContainsKeyRecursive(key string) bool {
// 	// check me
// 	if cp.ContainsKey(key) {
// 		return true
// 	}
// 	// check ancestors
// 	if cp.ancestorCp != nil {
// 		return cp.ancestorCp.ContainsKeyRecursive(key)
// 	}
// 	return false
// }

//------------------------------------------------------------
// fetch a choice from any choice point and remove it; i.e.the respective  CP is changed;
// returns the selected machine key and "" if no choice was found;
func (cp *ChoicePoint) EasyFetchChoice() string {
	//------------------------------------------------------------
	// is there any CP? if so just take the first one
	for key, _ := range cp.Choices {
		//------------------------------------------------------------
		// remove the choice
		delete(cp.Choices, key)
		//------------------------------------------------------------
		// return
		return key
	}
	//------------------------------------------------------------
	// return
	return ""
}

////------------------------------------------------------------
//// fetch a choice and remove it from the choice point, i.e. CP is changed:
//// select the choice depending on:
//// -- keyCriterion .... FIRST_KEY, RANDOM_KEY, ...
//// -- timeCriterion ... FIRST_TIME, RANDOM_TIME, ...
//// remove the time for this machine-key from the choices map, and if the time was
//// - the last one of a machine, remove also the machine-key from the choices map;
//// returns:
//// - key ........ the selected machine key
//// - int ........ the selected time
//// - foundFlag... was a choice found?
//func (cp *ChoicePoint) FetchChoice(keyCriterion ChoiceSelectionCriterionTypeEnum, timeCriterion ChoiceSelectionCriterionTypeEnum) (string, int, bool) {
//	//------------------------------------------------------------
//	// local vars
//	chosenTimeIndex := -1
//	keyFoundFlag := false
//	var key string
//	var times Ints
//	//------------------------------------------------------------
//	// ret vars
//	retKey := ""
//	retTime := -1
//	retFoundFlag := false

//	//------------------------------------------------------------
//	// "seed" the random generator, if not yet
//	// - use the number of nanoseconds elapsed since January 1, 1970 UTC
//	if !RANDOM_GENERATOR_WAS_SEEDED_FLAG {
//		rand.Seed(time.Now().UTC().UnixNano())
//		RANDOM_GENERATOR_WAS_SEEDED_FLAG = true
//	}
//	//------------------------------------------------------------
//	// is there any CP?
//	if 0 < len(cp.Choices) {
//		//------------------------------------------------------------
//		// apply the key-criterion:
//		// - ie which machine shall be selected
//		switch keyCriterion {
//		case FIRST_KEY:
//			for key, times = range cp.Choices {
//				// take the first choice:
//				retKey = key
//				keyFoundFlag = true
//				break
//			}
//		case RANDOM_KEY:
//			chosenKeyNr := RANDOM_GENERATOR.Intn(len(cp.Choices))
//			n := 0
//			for key, times = range cp.Choices {
//				if n == chosenKeyNr {
//					retKey = key
//					keyFoundFlag = true
//					break
//				}
//				n++
//			}
//		default:
//			Panic("ill. choice selection criterion")
//		}
//		//------------------------------------------------------------
//		// apply the time-criterion:
//		// - ie which time of the machine shall be selected
//		if keyFoundFlag {
//			switch timeCriterion {
//			case FIRST_TIME:
//				for j := 0; j < len(times); j++ {
//					// take the first choice:
//					chosenTimeIndex = j
//					retTime = times[chosenTimeIndex]
//					retFoundFlag = true
//					break
//				}
//				if retFoundFlag {
//					break
//				}
//			case RANDOM_TIME:
//				chosenTimeIndex = RANDOM_GENERATOR.Intn(len(times))
//				retTime = times[chosenTimeIndex]
//				retFoundFlag = true
//			default:
//				Panic("ill. choice selection criterion")
//			}
//		}
//		//------------------------------------------------------------
//		// remove the choice, if found:
//		if retFoundFlag {
//			/********/ if MODEL_CHECKING_DETAILS_TRACE.DoTrace() {
//				/********/ fmt.Print("\nFetchChoice: #choices=", len(cp.Choices), ", chosen M=", retKey,
//					/********/ ", #times=", len(cp.Choices[retKey]), ", chosenTimeIndex=", chosenTimeIndex, ", chosen time=", retTime, "\n")
//			}
//			// remove time from Ints of machine of CP
//			cp.Choices[retKey] = append(cp.Choices[retKey][:chosenTimeIndex], cp.Choices[retKey][chosenTimeIndex+1:]...)
//			// if machine's Ints is empty -> remove key from CP's Choices
//			if 0 == len(cp.Choices[retKey]) {
//				delete(cp.Choices, retKey)
//			}
//		}
//	}
//	//------------------------------------------------------------
//	// return result
//	return retKey, retTime, retFoundFlag
//}

//////////////////////////////////////////////////////////////
// debug
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// nb: does not print status -- that would be too much
func (cp *ChoicePoint) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%sEventClock=%d; Choices=\n", s, cp.EventClock)
	ind := NBlanksToString("", tab+TAB)
	for key, _ := range cp.Choices {
		s = fmt.Sprintf("%s%s%s\n", s, ind, key)
	}
	return s
}

//------------------------------------------------------------
func (cp *ChoicePoint) Print(tab int) {
	/**/ String2TraceFile(cp.ToString(tab))
}

//------------------------------------------------------------
// the same as Print
func (cp *ChoicePoint) Println(tab int) {
	/**/ cp.Print(tab)
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
