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
	. "github.com/peermodel/simulator/scheduler"
	"fmt"
)

type Wiring struct {
	// system properties (system defined):
	// wiring id:
	Id string
	// wiring container id:
	WCId string
	// wiring properties: system properties that can be defined by the user:
	// - TTL, TTS, TXCC, REPEAT_COUNT, MAX_THREADS, ...
	WProps
	// services: map key = service id:
	ServiceWrappers map[string]*ServiceWrapper
	// links:
	Links []*Link
	// for debug only: denotes whether it is a dynamic wiring - to suppress traces:
	DynamicWiringFlag bool
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

func NewWiring(id string) *Wiring {
	// alloc:
	w := new(Wiring)
	w.Id = id
	w.WProps = WProps{}
	w.ServiceWrappers = map[string]*ServiceWrapper{}
	w.Links = []*Link{}
	return w
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (w *Wiring) Copy() *Wiring {
	//------------------------------------------------------------
	// alloc
	newW := NewWiring(w.Id)
	//------------------------------------------------------------
	// copy all fields:
	// - Id:
	newW.Id = w.Id
	// - WCId:
	newW.WCId = w.WCId
	// - WProps:
	newW.WProps = w.WProps.Copy()
	// - ServiceWrappers:
	for swid, sw := range w.ServiceWrappers {
		newW.ServiceWrappers[swid] = sw.Copy()
	}
	// - Links: keep the order!
	for i := 0; i < len(w.Links); i++ {
		newW.Links = append(newW.Links, w.Links[i].Copy())
	}
	// - dynamicWiringFlag:
	newW.DynamicWiringFlag = w.DynamicWiringFlag
	//------------------------------------------------------------
	// return
	return newW
}

// --------------------------------------------
// get properties: they must be evaluated first against current wiring variables
//   (found in machine context i.e. m.Vars)
//   nb: there is no entry against which the wiring props need to be evaluated -> nil is used
// return default values if property is not set or cannot be evaluated
//   (in this case Eval prints warning);
// nb: only max threads must not be evaluated, as there is no wiring machine yet...
//   (@@@ maybe later: take global system vars into consideration)
// nb: if ctx == nil -> just return the unevaluated argument / or default value
// - e.g. used when printing the meta model (eg by latex) there is no machine yet
// --------------------------------------------

// --------------------------------------------
// get MAX_THREADS
// - nb: this get function is special: it is called by system *before* wiring machine can start!
// - nb: if does not need an arg; this is only to be consistent with other get functions
func (w *Wiring) GetMaxThreads(_ *Context) int {
	// fetch the max threads property
	arg, found := w.WProps[MAX_THREADS]
	// - if arg exists, and if it is a var or an int -> return its value
	if found && (VAR != arg.Kind || INT != arg.Type) {
		return arg.IntVal
	}
	// otherwise return default value = 1
	return 1

}

// --------------------------------------------
// get TTS
func (w *Wiring) GetTts(ctx *Context) int {
	arg := w.WProps[TTS]

	if nil == ctx && "" != arg.Kind {
		return arg.IntVal
	}
	if "" == arg.Kind || (!arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry())) {
		return 0 // default
	} else {
		return arg.IntVal
	}
}

// --------------------------------------------
// get TTS converted to absolute time
func (w *Wiring) GetAbsTts(ctx *Context) int {
	tts := w.GetTts(ctx)

	if INFINITE == tts {
		return INFINITE
	}
	return tts + CLOCK
}

// --------------------------------------------
// get TTL
func (w *Wiring) GetTtl(ctx *Context) int {
	arg := w.WProps[TTL]

	if nil == ctx && "" != arg.Kind {
		return arg.IntVal
	}
	if "" == arg.Kind || (!arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry())) {
		return INFINITE // default
	} else {
		return arg.IntVal
	}
}

// --------------------------------------------
// get TTL converted to absolute time
func (w *Wiring) GetAbsTtl(ctx *Context) int {
	ttl := w.GetTtl(ctx)

	if INFINITE == ttl {
		return INFINITE
	}
	return ttl + CLOCK
}

// --------------------------------------------
// get TXCC
func (w *Wiring) GetTxcc(ctx *Context) string {
	arg := w.WProps[TXCC]

	if nil == ctx && "" != arg.Kind {
		return arg.StringVal
	}
	if "" == arg.Kind || (!arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry())) {
		return PCC // default
	} else {
		return arg.StringVal
	}
}

// --------------------------------------------
// get ON_ABORT
func (w *Wiring) GetOnAbort(ctx *Context) bool {
	arg := w.WProps[ON_ABORT]

	if nil == ctx && "" != arg.Kind {
		return arg.BoolVal
	}
	if "" == arg.Kind || (!arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry())) {
		return false // default
	} else {
		return arg.BoolVal
	}
}

// --------------------------------------------
// get REPEAT_COUNT
func (w *Wiring) GetRepeatCount(ctx *Context) int {
	arg := w.WProps[REPEAT_COUNT]

	if nil == ctx && "" != arg.Kind {
		return arg.IntVal
	}
	if "" == arg.Kind || (!arg.Eval(ctx.Vars, ctx.EvalEs.GetFirstEntry())) {
		return INFINITE // default
	} else {
		return arg.IntVal
	}
}

// --------------------------------------------
// add service wrapper:
func (w *Wiring) AddServiceWrapper(serviceId string, sw *ServiceWrapper) {
	w.ServiceWrappers[serviceId] = sw
}

// --------------------------------------------
// add links:

func (w *Wiring) AddLink(l *Link) {
	w.Links = append(w.Links, l)
}

func (w *Wiring) AddGuard(subpid string, c1 string, op SpaceOpTypeEnum, q Query, lprops LProps, eprops EProps, vars Vars) {
	w.AddLink(NewGuard(subpid, c1, op, q, lprops, eprops, vars))
}

func (w *Wiring) AddSin(op SpaceOpTypeEnum, q Query, sid string, lprops LProps, eprops EProps, vars Vars) {
	w.AddLink(NewSin(op, q, sid, lprops, eprops, vars))
}

func (w *Wiring) AddScall(sid string, lprops LProps, eprops EProps, vars Vars) {
	w.AddLink(NewScall(sid, lprops, eprops, vars))
}

func (w *Wiring) AddSout(q Query, sid string, lprops LProps, eprops EProps, vars Vars) {
	w.AddLink(NewSout(q, sid, lprops, eprops, vars))
}

func (w *Wiring) AddAction(subpid string, c2 string, op SpaceOpTypeEnum, q Query, lprops LProps, eprops EProps, vars Vars) {
	w.AddLink(NewAction(subpid, c2, op, q, lprops, eprops, vars))
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (w *Wiring) IsEmpty() bool {
	if nil == w || w.Id == "" {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// --------------------------------------------
// does a '\n' at the end
func (w *Wiring) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%sWiring %s: %s ", s, w.Id, w.WCId)
	s = fmt.Sprintf("%s%s\n", s, w.WProps.ToString(0))
	// @@@missing: print service wrappers
	for _, l := range w.Links {
		s = fmt.Sprintf("%s%s\n", s, l.ToString(tab+TAB))
	}
	return s
}

// --------------------------------------------
// does a '\n' at the end
func (w *Wiring) Print(tab int) {
	/**/ String2TraceFile(w.ToString(tab))
}

// --------------------------------------------
func (w *Wiring) Println(tab int) {
	w.Print(tab)
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
