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
	. "cca/config"
	. "cca/debug"
	. "cca/scheduler"
	"fmt"
	"strings"
)

type Entry struct {
	// system properties that can only be system defined:
	Id string
	// entry properties (system or user) that can be defined by the user:
	// system: TYPE, TTS, TTL, DEST, FID, ORIGINATOR
	// exception: FROMPEER, FROMLINK, FROMWIRING, ERRMSG, ERRTYPE, TIME,
	//            ERRTREATMENTPEER (default = ORIGINATOR), TTL
	// nb: exception entry wraps entire entry if it raised the exception (via data)
	// nb: they contain also the entry type
	EProps
	// data (modeled as an entry collection):
	Data EntryPtrs
	// locks:
	Locks
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
func AllocEntry() *Entry {
	e := new(Entry)

	// @@@MCs: wozu?
	e.EProps = EProps{}
	e.Data = EntryPtrs{}
	e.Locks = NewLocks()

	return e
}

// ----------------------------------------
func NewEntry(eType string) *Entry {
	e := AllocEntry()

	e.Id = Uuid("e")
	e.SetStringVal(TYPE, eType)

	// for debug
	// e.SetIntVal("created", Clock())

	return e
}

////////////////////////////////////////
// methods
////////////////////////////////////////

// ----------------------------------------
// wrap entry e into exception entry
// it inherits e's flow id @@@ what else?
// @@@ workaround: because the access of data. ... is complicated to model ...
// - keep all e props of e and just change type to exception and etype is set to e's original type;
//   also: set its TTL to INFINITE and TTS to 0
// - richtig wäre: e becomes data of the exception entry
func (e *Entry) ExceptionWrap(dump string) *Entry {
	// assert that e's ttl is really expired
	if CLOCK < ConvertInfiniteTtl(e.GetTtl()) {
		Panic(fmt.Sprintf("Exception raised for non-expired entry: entry ttl = %d, CLOCK = %d, SYSTEM_TTL = %d, INFINITE = %d, e = %s, dump = %s", e.GetTtl(), CLOCK, SYSTEM_TTL, INFINITE, e.ToString(0), dump))
	}
	if e.GetTtl() == INFINITE {
		Panic(fmt.Sprintf("Exception raised for entry with INFINITE ttl: entry ttl = %d, CLOCK = %d, SYSTEM_TTL = %d, INFINITE = %d, e = %s, dump = %s", e.GetTtl(), CLOCK, SYSTEM_TTL, INFINITE, e.ToString(0), dump))
	}
	if e.GetTtl() >= SYSTEM_TTL {
		Panic(fmt.Sprintf("Exception raised for entry with ttl = %d, CLOCK = %d, SYSTEM_TTL = %d, INFINITE = %d, e = %s, dump = %s", e.GetTtl(), CLOCK, SYSTEM_TTL, INFINITE, e.ToString(0), dump))
	}

	excE := NewEntry(EXCEPTION_WRAP)

	// workaround: caution: overwrites also the type propoerty -> set it again below
	excE.EProps = e.EProps.Copy()
	excE.SetIntVal("ettl", e.GetTtl())
	excE.SetIntVal(TTL, INFINITE)
	excE.SetIntVal(TTS, 0)

	// set etype
	excE.SetStringEtype("etype", e.GetStringVal("type"))

	// set type
	excE.SetStringEtype("type", EXCEPTION_WRAP)

	// that would be correct:
	// excE.Data = append(excE.Data, e)
	// excE.SetStringVal(FID, e.GetStringVal(FID))

	// for debug only:
	excE.SetIntVal("exc_time", CLOCK)

	return excE
}

// --------------------------------------------
// get properties: return default values if property is not set:
// ----------------------------------------

// ----------------------------------------
func (e *Entry) GetType() string {
	arg := e.EProps[TYPE]
	if "" == arg.Kind {
		return "" // default @@@ tbd
	} else {
		return arg.StringVal
	}
}

// ----------------------------------------
func (e *Entry) GetTts() int {
	arg := e.EProps[TTS]
	if "" == arg.Kind {
		return 0 // default
	} else {
		return arg.IntVal
	}
}

// ----------------------------------------
func (e *Entry) GetTtl() int {
	arg := e.EProps[TTL]
	if "" == arg.Kind {
		return INFINITE // default
	} else {
		return arg.IntVal
	}
}

// ----------------------------------------
func (e *Entry) GetDest() string {
	arg := e.EProps[DEST]
	if "" == arg.Kind {
		return "" // default
	} else {
		return arg.StringVal
	}
}

// ----------------------------------------
func (e *Entry) GetFid() string {
	arg := e.EProps[FID]
	if "" == arg.Kind {
		return "" // default
	} else {
		return arg.StringVal
	}
}

//------------------------------------------------------------
// deep copy
func (e Entry) Copy() *Entry {
	//------------------------------------------------------------
	// alloc
	newE := AllocEntry()
	//------------------------------------------------------------
	// copy all fields:
	// - Id:
	// -- nb: id remains the same, but the enry is a copied one
	newE.Id = e.Id
	// - Args:
	for label, arg := range e.EProps {
		newArg := arg.Copy()
		newE.EProps[label] = newArg
	}
	// - Data:
	for _, subE := range e.Data {
		subEPtr := subE.Copy()
		subE1 := *subEPtr
		newE.Data = append(newE.Data, &subE1)
	}
	// - Locks:
	newE.Locks = e.Locks.Copy()
	//------------------------------------------------------------
	// return
	return newE
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (e *Entry) IsEmpty() bool {
	if nil == e || e.Id == "" {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
// print only fields that are set resp. that are not default...
func (e *Entry) ToString(tab int) string {
	sep := ", "
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%s<", s)
	s = fmt.Sprintf("%sId=%s", s, e.Id)
	// type: print it first - and not with args
	s = fmt.Sprintf("%s%s%s=%s", s, sep, TYPE, e.GetType())
	// print all other args (type was printed above): nb: 1...for type which must exist!
	if 1 < len(e.EProps) {
		s = fmt.Sprintf("%s%s, %s", s, sep, e.EProps.ToStringWithDetails(0, "", false /* detailsFlag */, false /* printTypeFlag */, false /* omitDefaultsFlag */))
	}
	if 0 < len(e.Data) {
		// @@@ bessere ausgabe...
		s = fmt.Sprintf("%s%sData=%s", s, sep, e.Data.ToStringInOneRow())
	}
	if 0 < len(e.Locks.RLocks) || 0 < len(e.Locks.DLocks) || 0 < len(e.Locks.WLocks) {
		s = fmt.Sprintf("%s%s", s, sep)
		s = fmt.Sprintf("%s%s", s, e.Locks.ToString(0))
	}
	s = fmt.Sprintf("%s>", s)
	// @@@ workaround @@@ hack @@@ igitt: replace ", , " by ", "; once
	return strings.Replace(s, ", , ", ", ", 1)
}

// ----------------------------------------
// print only fields that are set resp. that are not default...
func (e *Entry) Print(tab int) {
	/**/ String2TraceFile(e.ToString(tab))
}

// ----------------------------------------
func (e *Entry) Println(tab int) {
	/**/ e.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
