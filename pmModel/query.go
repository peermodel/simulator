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

type Query struct {
	// entry type; nb: must be Arg so that vars can be used for entry type!
	Typ Arg
	// selector: must be pointer in order to figure out whether it is set or not!
	// nb: result must be boolean;
	Sel *Arg
	// count: user interface; Cnt=k is shortcut for Min=k and Max=k
	Count Arg
	// Min must evaluate to number >= 0
	// Max must evaluate to number >= 0, ALL or NONE
	// nb: both must not contain entry labels - only expressions with vars and basic values
	Min Arg
	Max Arg
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

func NewQuery() *Query {
	return new(Query)
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy;
// CAUTION: keep up to date with Query struct
func (q Query) Copy() Query {
	//------------------------------------------------------------
	// alloc
	newQ := NewQuery()
	//------------------------------------------------------------
	// copy all fields:
	// - Typ:
	newQ.Typ = q.Typ
	// - Sel:
	if nil != q.Sel {
		s := q.Sel.Copy()
		newQ.Sel = &s
	}
	// - Count:
	newQ.Count = q.Count.Copy()
	// - Min:
	newQ.Min = q.Min.Copy()
	// - Max:
	newQ.Max = q.Max.Copy()
	//------------------------------------------------------------
	// return
	return *newQ
}

////////////////////////////////////////
// get (and eval) Count, Min, Max, Typ -- without given entry
// nb: vars are considered
////////////////////////////////////////

// --------------------------------------------
// @@@ unused?
func (q *Query) GetCount(vars Vars) int {
	if q.Count.Eval(vars, nil /* no entry */) && INT == q.Count.Type {
		return (q.Count.IntVal)
	} else {
		Panic(fmt.Sprintf("Query: ill. query count specification: q = %s", q.ToString(0)))
		return 0
	}
}

// --------------------------------------------
func (q *Query) GetMin(vars Vars) int {
	if q.Min.Eval(vars, nil /* no entry */) && INT == q.Min.Type {
		return (q.Min.IntVal)
	} else {
		Panic(fmt.Sprintf("Query: ill. query min specification: q = %s", q.ToString(0)))
		return 0
	}
}

// --------------------------------------------
func (q *Query) GetMax(vars Vars) int {
	if q.Max.Eval(vars, nil /* no entry */) && INT == q.Max.Type {
		return (q.Max.IntVal)
	} else {
		Panic(fmt.Sprintf("Query: ill. query max specification: q = %s", q.ToString(0)))
		return 0
	}
}

// --------------------------------------------
func (q *Query) GetTyp(vars Vars) string {
	if q.Typ.Eval(vars, nil /* no entry */) && STRING == q.Typ.Type {
		return (q.Typ.StringVal)
	} else {
		Panic(fmt.Sprintf("Query: ill. query typ specification: q = %s", q.ToString(0)))
		return ""
	}
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (q Query) IsEmpty() bool {
	return false
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// --------------------------------------------
func (q *Query) ToString(ind int) string {
	s := NBlanksToString("", ind)
	if "" != q.Typ.Kind {
		s = fmt.Sprintf(" %s%s[%s <= cnt <= %s]", s, q.Typ.ToString(0), q.Min.ToString(0), q.Max.ToString(0))
		if nil != q.Sel {
			s = fmt.Sprintf("%s [[", s)
			s = fmt.Sprintf("%s%s", s, q.Sel.ToString(0))
			s = fmt.Sprintf("%s]]", s)
		}
	}
	return s
}

// --------------------------------------------
func (q Query) Print(ind int) {
	/**/ String2TraceFile(q.ToString(ind))
}

// --------------------------------------------
func (q Query) Println(ind int) {
	/**/ q.Print(ind)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
