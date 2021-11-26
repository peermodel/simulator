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
	. "cca/helpers"
	. "cca/scheduler"
	"fmt"
)

type Link struct {
	// sub peer reference (for C1 and C2; if set -> ally PIC and POC to sub peer; otherwise to this peer)
	SubPid string
	// resolved from/to container names:
	C1 string
	C2 string
	// operation type: CALL, DELETE, READ, TAKE, TEST:
	Op SpaceOpTypeEnum
	// link type: GUARD, ACTION, SERVICE_IN, SERVICE, SERVICE_OUT:
	Type LinkTypeEnum
	// user modeled container name: depending on the link type it is either the C1 or C2 name
	modelC string
	// query:
	Q Query
	// service id:
	// @@@ should also be a link property...
	Sid string
	// entry properties : (system or user) that can be defined by the user:
	EProps EProps
	// vars:
	LVars Vars
	// link properties:
	// TTS, TTL, DEST, SOURCE, FLOW, MANDATORY, COMMIT
	LProps
}

////////////////////////////////////////
// constructors
////////////////////////////////////////

func NewLink() *Link {
	// alloc:
	l := new(Link)
	return l
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (l *Link) Copy() *Link {
	//------------------------------------------------------------
	// alloc
	newL := NewLink()
	//------------------------------------------------------------
	// copy static fileds:
	// - SubPid:
	newL.SubPid = l.SubPid
	// - C1:
	newL.C1 = l.C1
	// - C2:
	newL.C2 = l.C2
	// - Op:
	newL.Op = l.Op
	// - Type:
	newL.Type = l.Type
	// - modelC:
	newL.modelC = l.modelC
	// - Q:
	newL.Q = l.Q.Copy()
	// - Sid:
	newL.Sid = l.Sid
	// - EProps:
	newL.EProps = l.EProps.Copy()
	// - Vars:
	newL.LVars = l.LVars.Copy()
	// - LProps:
	newL.LProps = l.LProps.Copy()
	//------------------------------------------------------------
	// return
	return newL
}

// --------------------------------------------
// get properties: they must be evaluated first against current wiring variables and a possibly given entry;
// return default values if property is not set or cannot be evaluated (in this case Eval prints warning):
// nb: if ctx == nil -> just return the unevaluated argument / or default value
// --------------------------------------------

// --------------------------------------------
// default = 0
func (l *Link) GetTts(ctx *Context) int {
	// /**/ m.PrintlnS(TRACE4, ind, TTS, "")
	arg := l.LProps[TTS]

	if nil == ctx && "" != arg.Kind {
		return arg.IntVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return 0 // default
	} else {
		return arg.IntVal
	}
}

// --------------------------------------------
// convert to absolute time
func (l *Link) GetAbsTts(ctx *Context) int {
	tts := l.GetTts(ctx)

	if INFINITE == tts {
		return INFINITE
	}
	return tts + CLOCK
}

// --------------------------------------------
// default = INFINITE
func (l *Link) GetTtl(ctx *Context) int {
	// /**/ m.PrintlnA(ind, TTL, "")
	arg := l.LProps[TTL]

	if nil == ctx && "" != arg.Kind {
		return arg.IntVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return INFINITE // default; @@@INFINITE does not work???? @@@1000 does not work
	} else {
		return arg.IntVal
	}
}

// --------------------------------------------
// convert to absolute time
func (l *Link) GetAbsTtl(ctx *Context) int {
	ttl := l.GetTtl(ctx)

	if INFINITE == ttl {
		return INFINITE
	}
	return ttl + CLOCK
}

// --------------------------------------------
// default = ""
func (l *Link) GetDest(ctx *Context) string {
	// /**/ m.PrintlnA(ind, DEST, "")
	arg := l.LProps[DEST]

	if nil == ctx && "" != arg.Kind {
		return arg.StringVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return "" // default
	} else {
		return arg.StringVal
	}
}

// --------------------------------------------
// default = ""
func (l *Link) GetSource(ctx *Context) string {
	// /**/ m.PrintlnA(ind, SOURCE, "")
	arg := l.LProps[SOURCE]

	if nil == ctx && "" != arg.Kind {
		return arg.StringVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return "" // default
	} else {
		return arg.StringVal
	}
}

// --------------------------------------------
// default = true
func (l *Link) GetFlow(ctx *Context) bool {
	// /**/ m.PrintlnS(TRACE4, ind, FLOW, "")
	arg := l.LProps[FLOW]

	if nil == ctx && "" != arg.Kind {
		return arg.BoolVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(TRACE4, ind+TAB, "", "default")
		return true // default
	} else {
		return arg.BoolVal
	}
}

// --------------------------------------------
// default = true
func (l *Link) GetMandatory(ctx *Context) bool {
	// /**/ m.PrintlnA(ind, MANDATORY, "")
	arg := l.LProps[MANDATORY]

	if nil == ctx && "" != arg.Kind {
		return arg.BoolVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return true // default
	} else {
		return arg.BoolVal
	}
}

// --------------------------------------------
// default = false
func (l *Link) GetCommit(ctx *Context) bool {
	// /**/ m.PrintlnA(ind, COMMIT, "")
	arg := l.LProps[COMMIT]

	if nil == ctx && "" != arg.Kind {
		return arg.BoolVal
	}
	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry()) {
		// /**/ m.PrintlnS(ind+TAB, "default")
		return false // default
	} else {
		return arg.BoolVal
	}
}

// --------------------------------------------
// @@@ TBD: make Sid a normal link property
//func (l *Link) GetSid(ctx *Context) string {
//	// /**/ m.PrintlnS(TRACE4, ind, SID, "")
//	arg := l.LProps[SID]

//	if nil == ctx && "" != arg.Kind {
//		return arg.StringVal
//	}
//	if "" == arg.Kind || !arg.Eval(ctx.Vars, ctx.ReadEs.GetFirstEntry()) {
//		// /**/ m.PrintlnS(TRACE4, ind+TAB, "", "default")
//		return "-1" // default
//	} else {
//		return arg.StringVal
//	}
//}

// --------------------------------------------
// create links:
// CAUTION: use only these functions to create links, because they store
//   the modeled name of C1 resp. C2 which later on must be resolved!
// CAUTION: the count in queries must be converted into min and max
// --------------------------------------------

// --------------------------------------------
func NewGuard(subpid string, c1 string, op SpaceOpTypeEnum, q Query, lprops LProps, eprops EProps, vars Vars) *Link {
	l := NewLink()
	l.SubPid = subpid
	l.modelC = c1
	l.Op = op
	l.Type = GUARD
	l.Q = convertQueryCountToMinMax(q)
	l.EProps = eprops
	l.LVars = vars
	l.LProps = lprops
	return l
}

// --------------------------------------------
func NewSin(op SpaceOpTypeEnum, q Query, sid string, lprops LProps, eprops EProps, vars Vars) *Link {
	l := NewLink()
	l.Op = op
	l.Type = SERVICE_IN
	l.Q = convertQueryCountToMinMax(q)
	l.Sid = sid
	l.EProps = eprops
	l.LVars = vars
	l.LProps = lprops
	return l
}

// --------------------------------------------
func NewScall(sid string, lprops LProps, eprops EProps, vars Vars) *Link {
	l := NewLink()
	l.Op = CALL
	l.Type = SERVICE
	l.Sid = sid
	l.EProps = eprops
	l.LVars = vars
	l.LProps = lprops
	return l
}

// --------------------------------------------
func NewSout(q Query, sid string, lprops LProps, eprops EProps, vars Vars) *Link {
	l := NewLink()
	l.Op = TAKE
	l.Type = SERVICE_OUT
	l.Q = convertQueryCountToMinMax(q)
	l.Sid = sid
	l.EProps = eprops
	l.LVars = vars
	l.LProps = lprops
	return l
}

// --------------------------------------------
// NB: if IOP is used for subpid -> this is specially treated!!!!
// NB: in the model there is currently only one IOP (locally)
func NewAction(subpid string, c2 string, op SpaceOpTypeEnum, q Query, lprops LProps, eprops EProps, vars Vars) *Link {
	l := NewLink()
	l.SubPid = subpid
	l.modelC = c2
	l.Op = op
	l.Type = ACTION
	l.Q = convertQueryCountToMinMax(q)
	l.EProps = eprops
	l.LVars = vars
	l.LProps = lprops
	return l
}

// --------------------------------------------
// convert C1 to the real id used by the given machine instance (= machine number = wmno):
func (l *Link) ConvertC1toM(wmno int) *string {
	if GUARD != l.Type {
		return ConvertCtoM(l.C1, wmno)
	} else {
		return &l.C1
	}
}

// --------------------------------------------
// convert C2 to the real id used by the given machine instance (= machine number = wmno):
func (l *Link) ConvertC2toM(wmno int) *string {
	if ACTION != l.Type {
		return ConvertCtoM(l.C2, wmno)
	} else {
		return &l.C2
	}
}

// --------------------------------------------
// returns the changed vars;
// nb: vars are overwritten;
// @@@ restriction: consider only the first entry in es - if any one is there!
// @@@ order should be kept!
func (l *Link) ResolveLinkArgs(vars Vars, es EntryPtrs) Vars {
	var e *Entry
	if !es.IsEmpty() {
		e = es[0]
	} else {
		// use the empty entry
		e = new(Entry)
	}
	for label, v := range l.LVars {
		if !v.Eval(vars, e) {
			Panic(fmt.Sprintf("ResolveLinkArgs: ill. var specification for link; var name = %s", label))
		}
		a := Arg{Kind: VAL, Type: v.Type, IntVal: v.IntVal, StringVal: v.StringVal, BoolVal: v.BoolVal}
		vars[label] = a
	}
	return vars
}

////////////////////////////////////////
// functions
////////////////////////////////////////

// --------------------------------------------
// convert Count into Min Max:
func convertQueryCountToMinMax(q Query) Query {
	q1 := q.Copy()
	// if Count's Kind is VAL and if its Type == INT:
	if VAL == q1.Count.Kind && INT == q1.Count.Type {
		// set Max to Count
		q1.Max = q.Count.Copy()
		// if Count == ALL  -> set Min to 0
		if ALL == q1.Count.IntVal {
			q1.Min = IVal(0)
			return q1
		}
		// if Count == NONE -> set Min to 1
		if NONE == q1.Count.IntVal {
			q1.Min = IVal(1)
			return q1
		}
		// otherwise -> set Min to Count
		q1.Min = q.Count.Copy()
	} else {
		if VAR == q1.Count.Kind {
			q1.Min = q.Count.Copy()
			q1.Max = q.Count.Copy()
		}
		// @@@ else ???
	}
	return q1
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (l *Link) IsEmpty() bool {
	if nil == l {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// --------------------------------------------
func (l *Link) ToString(ind int) string {
	s := NBlanksToString("", ind)
	s = fmt.Sprintf("%s%s", s, l.Type)
	s = fmt.Sprintf("%s: ", s)
	if l.Type == GUARD {
		s = fmt.Sprintf("%s ", s)
	}
	s = fmt.Sprintf("%s%s --> %s: %s", s, l.C1, l.C2, l.Op)

	if l.Op != CALL {
		s = fmt.Sprintf("%s%s", s, l.Q.ToString(0))
	} else {
		s = fmt.Sprintf("%s %s", s, l.Sid)
	}
	s = fmt.Sprintf("%s%s%s%s", s,
		l.LProps.ToStringWithDetails(0, ", lprops", false /* detailsFlag */, true /* printTypeFlag */, false /* omitDefaultsFlag */),
		l.EProps.ToStringWithDetails(0, ", eprops", false /* detailsFlag */, true /* printTypeFlag */, false /* omitDefaultsFlag */),
		l.LVars.ToStringWithDetails(0, ", vars", false /* detailsFlag */, true /* printTypeFlag */, false /* omitDefaultsFlag */))
	return s
}

// --------------------------------------------
func (l *Link) Print(ind int) {
	/**/ String2TraceFile(l.ToString(ind))
}

// --------------------------------------------
func (l *Link) Println(ind int) {
	/**/ l.Print(ind)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
