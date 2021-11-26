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

// ----------------------------------------
// CAUTION: keep Clone function consistent!!!!
// each machine has its own context;
// context is used to pass info between machines (at entering/leaving);
// used by <--> used when the machine is called (as input IN or output OUT variable) -> see visio
type Context struct {
	Pid    string
	Wid    string
	Wiid   string
	Wtxid  string
	Wfid   string
	LinkNo int
	// current and already to wiring machine converted cid for read or write operations
	Cid string
	// service id
	Sid string
	// wiring machine number: needed for conversion of modeled cids to machine cids
	WMNo int
	// Query: needed for source property treatment
	// @@@ warum nicht nur Query?
	Query Query
	// vars
	Vars
	// entries
	Es EntryPtrs
	// caution: eval entry set must always be correctly initialized/set:
	// - it(ie its first entry - if there) is used to evaluate args
	EvalEs EntryPtrs
	// return entries
	RetEs EntryPtrs
	// return value
	RetErr error
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
// CAUTION: keep up to date, if Context struct changes;
// nb: test case need not be copied; if nil -> it is a pure "system" machine that won't call a user service
func NewContext() *Context {
	ctx := new(Context)
	// clear structured fields
	ctx.Query = Query{}
	ctx.Vars = Vars{}
	ctx.Es = EntryPtrs{}
	ctx.EvalEs = EntryPtrs{}
	ctx.RetEs = EntryPtrs{}
	return ctx
}

////////////////////////////////////////
// methods
////////////////////////////////////////

////////////////////////////////////////
// IContext interface implementation:
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (ctx Context) Copy() interface{} {
	//------------------------------------------------------------
	// alloc new context
	newCtx := NewContext()
	//------------------------------------------------------------
	// copy all fields
	// - Pid:
	newCtx.Pid = ctx.Pid
	// - Wid:
	newCtx.Wid = ctx.Wid
	// - Wiid:
	newCtx.Wiid = ctx.Wiid
	// - Wtxid:
	newCtx.Wtxid = ctx.Wtxid
	// - Wfid:
	newCtx.Wfid = ctx.Wfid
	// - LinkNo:
	newCtx.LinkNo = ctx.LinkNo
	// - Cid:
	newCtx.Cid = ctx.Cid
	// - Sid:
	newCtx.Sid = ctx.Sid
	// - WMNo:
	newCtx.WMNo = ctx.WMNo
	// - Q:
	newCtx.Query = ctx.Query.Copy()
	// - Vars:
	newCtx.Vars = map[string]Arg{}
	for name, val := range ctx.Vars {
		newCtx.Vars[name] = val.Copy()
	}
	// - Es:
	newCtx.Es = ctx.Es.Copy()
	// - RetEs:
	newCtx.RetEs = ctx.RetEs.Copy()
	// - RetErr:
	newCtx.RetErr = ctx.RetErr
	//------------------------------------------------------------
	// return
	return newCtx
}

// ----------------------------------------
func (ctx Context) MachineKeySuffix() string {
	return fmt.Sprintf("%s__%s", ctx.Pid, ctx.Wid)
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (ctx Context) IsEmpty() bool {
	return false
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (ctx Context) Print(tab int) {
	/**/ NBlanks2TraceFile(tab)
	sep := "ctx={"
	/**/ String2TraceFile(sep)
	// /**/ String2TraceFile(fmt.Sprintf("CurrentState=%s, ", ctx.CurrentState, ))
	/**/
	String2TraceFile(fmt.Sprintf("Pid=%s, ", ctx.Pid))
	/**/ String2TraceFile(fmt.Sprintf("Wid=%s, ", ctx.Wid))
	/**/ String2TraceFile(fmt.Sprintf("Wiid=%s, ", ctx.Wiid))
	/**/ String2TraceFile(fmt.Sprintf("Wtxid=%s, ", ctx.Wtxid))
	/**/ if "" != ctx.Wfid {
		String2TraceFile(fmt.Sprintf("Wfid=%s, ", ctx.Wfid))
	}
	/**/ String2TraceFile(fmt.Sprintf("LinkNo=%d", ctx.LinkNo))

	sep = ", "

	if "" != ctx.Cid {
		/**/ String2TraceFile(sep)
		/**/ String2TraceFile(fmt.Sprintf("Cid=%s", ctx.Cid))
	}
	if "" != ctx.Sid {
		/**/ String2TraceFile(sep)
		/**/ String2TraceFile(fmt.Sprintf("Sid=%s", ctx.Sid))
	}
	// /**/ String2TraceFile(sep)
	// /**/ String2TraceFile(fmt.Sprintf("WMNo=%d", ctx.WMNo))
	//	if "" != ctx.Q.Typ {
	//		/**/ String2TraceFile(sep)
	//		/**/ ctx.Q.Print(0)
	//	}
	if "" != ctx.Query.Typ.Kind {
		/**/ String2TraceFile(sep)
		/**/ ctx.Query.Typ.Print(0)
	}
	if 0 < len(ctx.Vars) {
		/**/ String2TraceFile(sep)
		/**/ String2TraceFile("\n")
		/**/ PRINT_ARGS_NAME = "Vars"
		/**/ PRINT_ARG_DETAILS_FLAG = false
		/**/ PRINT_ARG_TYPE_FLAG = true
		/**/ ctx.Vars.Print(tab + 5)
	}
	// Es
	// RetEs
	// RetErr
	/**/
	String2TraceFile("}")
}

// ----------------------------------------
func (ctx Context) Println(tab int) {
	/**/ ctx.Print(tab)
	/**/ String2TraceFile("\n")
}

//--------------------------------------
// trick: just for code generator to guarantee that this package is used...
func (ctx Context) DummyIContextFu() interface{} {
	return new(Context)
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
