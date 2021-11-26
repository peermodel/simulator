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
	"fmt"
)

// @@@ should better be a map! with key = eid
type EntryPtrs []*Entry

////////////////////////////////////////
// methods
////////////////////////////////////////

// ----------------------------------------
func (inEs EntryPtrs) Copy() EntryPtrs {
	outEs := EntryPtrs{}

	for _, inE := range inEs {
		// clone outE from inE:
		outE := inE.Copy()
		// add to outE outEs:
		outEs = append(outEs, outE)
	}

	return outEs
}

// ----------------------------------------
// return first element, or nil if the es is empty
func (es EntryPtrs) GetFirstEntry() *Entry {
	if (es != nil) && (cap(es) > 0) && (len(es) > 0) {
		return es[0]
	}
	return nil
}

// ----------------------------------------
func (inEs EntryPtrs) CopyAndStripLocks() EntryPtrs {
	outEs := EntryPtrs{}

	for _, inE := range inEs {
		// clone outE from inE:
		outE := inE.Copy()
		// clear all locks in outE:
		outE.Locks.RLocks = make(map[string]int)
		outE.Locks.WLocks = make(map[string]int)
		outE.Locks.DLocks = make(map[string]int)
		// add to outE outEs:
		outEs = append(outEs, outE)
	}

	return outEs
}

// ----------------------------------------
// remove entry in a given set of entry pointers; entry is denoted by its unique entry id
// caller must assure that entry exists in entry collection
// nb: removes unchanged entry set if entry is not found
func (es EntryPtrs) RemoveEntry(eid string) EntryPtrs {
	entryIndex := -1
	for i, e := range es {
		if e.Id == eid {
			entryIndex = i
			break
		}
	}
	if -1 == entryIndex {
		return es
	}
	// remove entry
	i := entryIndex
	es = append(es[:i], es[i+1:]...)

	return es
}

// ----------------------------------------
// replace alle eid by a new one for all entries
func (es EntryPtrs) ExchangeEidByNewOne() {
	for _, e := range es {
		e.Id = Uuid("e")
	}
}

// ----------------------------------------
// apply entry props of link to all entries:
// first resolve each eprop
// overwrite!
func (es EntryPtrs) EvalAndApply(vars Vars, eprops EProps) {
	if len(eprops) > 0 {
		for _, tmpE := range es {
			for label, eprop := range eprops {

				if !eprop.Eval(vars, tmpE) {
					if ARGS_EVAL_TRACE.DoTrace() {
						/**/ PRINT_ARGS_NAME = "eprops"
						/**/ PRINT_ARG_DETAILS_FLAG = false
						/**/ PRINT_ARG_TYPE_FLAG = true
						/**/ eprops.Println(TAB)
					}
					Panic(fmt.Sprintf("can't eval eprop with label=%s", label))
				}

				switch label {
				case TTS:
					tmpE.SetIntVal(TTS, eprop.IntVal)
				case TTL:
					tmpE.SetIntVal(TTL, eprop.IntVal)
				case DEST:
					tmpE.SetStringVal(DEST, eprop.StringVal)
				case FID:
					tmpE.SetStringVal(FID, eprop.StringVal)
				default:
					// alloc if not yet
					if nil == tmpE.EProps {
						tmpE.EProps = EProps{}
					}
					switch eprop.Type {
					case STRING:
						if ARGS_EVAL_TRACE.DoTrace() {
							/**/ String2TraceFile(fmt.Sprintf("label=%s, val=%s\n", label, eprop.StringVal))
						}
						tmpE.SetStringVal(label, eprop.StringVal)
					case INT:
						if ARGS_EVAL_TRACE.DoTrace() {
							/**/ String2TraceFile(fmt.Sprintf("label=%s, val=%d\n", label, eprop.IntVal))
						}
						tmpE.SetIntVal(label, eprop.IntVal)
					case BOOL:
						if ARGS_EVAL_TRACE.DoTrace() {
							/**/ String2TraceFile(fmt.Sprintf("label=%s, val=%t\n", label, eprop.BoolVal))
						}
						tmpE.SetBoolVal(label, eprop.BoolVal)
					default:
						Panic(fmt.Sprintf("ill. eprop type = %s", eprop.Type))
					}
				}
			}
		}
	}
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (es EntryPtrs) IsEmpty() bool {
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
func (es EntryPtrs) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%s{", s)
	sep := ""
	nextTab := 0
	for _, e := range es {
		s = fmt.Sprintf("%s%s", s, sep)
		s = fmt.Sprintf("%s%s", s, e.ToString(nextTab))
		sep = ", \n"
		nextTab = tab + TAB
	}
	s = fmt.Sprintf("%s}", s)
	return s
}

// ----------------------------------------
func (es EntryPtrs) Print(tab int) {
	/**/ String2TraceFile(es.ToString(tab))
}

// ----------------------------------------
func (es EntryPtrs) Println(tab int) {
	/**/ es.Print(tab)
	/**/ String2TraceFile("\n")
}

// ----------------------------------------
func (es EntryPtrs) ToStringInOneRow() string {
	s := "{"
	sep := ""
	for _, e := range es {
		s = fmt.Sprintf("%s%s", s, sep)
		s = fmt.Sprintf("%s%s", s, e.ToString(0))
		sep = ", "
	}
	s = fmt.Sprintf("%s}", s)
	return s
}

// ----------------------------------------
func (es EntryPtrs) PrintInOneRow() {
	/**/ String2TraceFile(es.ToStringInOneRow())
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
