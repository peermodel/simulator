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

////////////////////////////////////////
// DOCU:
// - convention: all functions must use Panic in case of error
// -- because they must not use m./s. user error -> cyclic packages...
////////////////////////////////////////

package pmModel

import (
	. "cca/debug"
	. "cca/scheduler"
	"fmt"
	"strings"
)

// TBD: caution: cannot be exchanged...
const ArrayChar byte = '#'

// nb: variable and entry property use arg in map where the key is the label or variable to be set
// @@@ use union
type Arg struct {
	// Kind is obligatory field that must be set to one of:
	// - VAL
	// - LABEL
	// - VAR
	// - EXPR
	// - FU
	// - DYN_ARRAY_REF (uses StringVal and arrayArg)
	// - TYPED_ARRAY_LABEL (uses Type and arrayArg)
	// - TYPED_ARRAY_VAL (uses Type and arrayArg)
	// TBD: @@@ KindTypeEnum ... problem: what is the null value in order to find out that an arg is empty, ie not contained in an args list????
	Kind string

	// INT, STRING, BOOL; (set for all kinds except for expr);
	// for expr it is temporarily overwritten during eval:
	Type DataTypeEnum
	// NORMAL, URL, ENTRY_TYPE
	StringSubType StringSubTypeEnum

	// name label or variable name:
	Name string

	// function name
	FuName SystemFunctionEnum

	// value (temporarily overwritten during eval for non val kinds):
	// nb: entry type and url use string val!
	// - basic value
	IntVal    int
	StringVal string
	BoolVal   bool

	// expression
	// - must be pointer; otherwise Go reports recursion problem:
	// also used as array arg (for array access)
	// - nb: array arg uses only left arg of expression!
	ExprVal *Expr
}

////////////////////////////////////////
// constructors
// - the following are all Arg-aliases of different kinds
////////////////////////////////////////

// ----------------------------------------
func XVal(left Arg, op OpTypeEnum, right Arg) Arg {
	return Arg{Kind: EXPR, ExprVal: &Expr{Left: left, Op: op, Right: right}}
}

// ----------------------------------------
func XValP(left Arg, op OpTypeEnum, right Arg) *Arg {
	xv := XVal(left, op, right)
	return &xv
}

// ----------------------------------------
func IVal(val int) Arg {
	return Arg{Kind: VAL, Type: INT, IntVal: val}
}

// ----------------------------------------
func IFu(fuName SystemFunctionEnum) Arg {
	return Arg{Kind: FU, Type: INT, FuName: fuName}
}

// ----------------------------------------
func ILabel(name string) Arg {
	return Arg{Kind: LABEL, Type: INT, Name: name}
}

// ----------------------------------------
func IVar(name string) Arg {
	return Arg{Kind: VAR, Type: INT, Name: name}
}

// ----------------------------------------
func SVal(val string) Arg {
	return Arg{Kind: VAL, Type: STRING, StringVal: val}
}

// ----------------------------------------
func SFu(fuName SystemFunctionEnum) Arg {
	return Arg{Kind: FU, Type: STRING, FuName: fuName}
}

// ----------------------------------------
func SEtype(val string) Arg {
	return Arg{Kind: VAL, Type: STRING, StringSubType: ENTRY_TYPE, StringVal: val}
}

// ----------------------------------------
// TBD? fkt nicht auf rechter seite von dest zb
func SUrl(val string) Arg {
	return Arg{Kind: VAL, Type: STRING, StringSubType: URL, StringVal: val}
}

// ----------------------------------------
func SLabel(name string) Arg {
	return Arg{Kind: LABEL, Type: STRING, Name: name}
}

// ----------------------------------------
func SVar(name string) Arg {
	return Arg{Kind: VAR, Type: STRING, Name: name}
}

// ----------------------------------------
func BVal(val bool) Arg {
	return Arg{Kind: VAL, Type: BOOL, BoolVal: val}
}

// ----------------------------------------
func BFu(fuName SystemFunctionEnum) Arg {
	return Arg{Kind: FU, Type: BOOL, FuName: fuName}
}

// ----------------------------------------
func BLabel(name string) Arg {
	return Arg{Kind: LABEL, Type: BOOL, Name: name}
}

// ----------------------------------------
func BVar(name string) Arg {
	return Arg{Kind: VAR, Type: BOOL, Name: name}
}

//--------------------------------------------------------------------------------
// array access; there are 2 forms that depend on whether used on left side of assignment or not:
// (1) if on left side of assignment (ie used for left side of a prop def):
//   - eg: <labelName>#<i1>#<i2>#<i3>#<i4> = <right side>
//   -- the code is: ArrayRef(ArrayRef(ArrayRef(ArrayRef(<labelName>, <i1>), <i2>), <i3>, i4>)
// (2) else (ie used for right side of a prop def, or in a query selector):
//   - eg: <left side> = <BLabelName>#<i1>#<i2>#<i3>#<i4>
//   -- the code for the array access is: BArrayLabel(DynArrayRef(DynArrayRef(DynArrayRef(<labelName>, <i1>), <i2>), <i3>)
//   - eg: <SLabelName>#<i1> AND <ILabelName>#<k1>#<k2>
//   -- the code for the array access is: SArrayLabel(DynArrayRef(<SLabelName>, <i1>)) and to IArrayLabel(DynArrayRef(DynArrayRef(<ILabelName>, <k1>), <k1>))
// NB: <T< is the type letter
// NB: in case (1) Go-automaton must be statically evaluate it, whereas in case (2) dynamic evaluation is possible
//     - in case (1) namely a hash map is used for entry type names and var names...
//--------------------------------------------------------------------------------

// ----------------------------------------
// static ArrayRef specification with '#';
// can be used in nested way;
// used on ***left*** side, ie as label in EProps or Vars specifications
// index must be statically resolved
func ArrayRef(name string, indexArg Arg) string {
	// try to statically evaluate the index
	if (indexArg.Eval(Vars{}, nil /* no entry */)) {
		// verify that result is INT
		if indexArg.Type != INT {
			Panic("ArrayLabel: index must have type INT")
		}
		// generate label string of the form: label # index
		s := fmt.Sprintf("%s%c%d", name, ArrayChar, indexArg.IntVal)
		// fmt.Println(fmt.Sprintf("array (static): %s", s))
		return s
	} else {
		Panic("ArrayLabel: index of ArrayRef must be statically evaluable to INT")
	}
	return "*** ill. ArrayRef usage ***" // TBD: why needed?
}

// ----------------------------------------
// dynamic DynArrayRef specification with '#';
// can be used in nested way;
// used on ***right*** side as array access to a label (in EProps, Vars on right side, or in Sel expression)
// index will be dynamically resolved
func DynArrayRef(name string, indexArg Arg) Arg {
	// generate expression with op = INDEX that must by dynamically evaluated
	return Arg{Kind: DYN_ARRAY_REF, StringVal: name, ExprVal: &Expr{Left: indexArg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// SArrayLabel
func SArrayLabel(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_LABEL, Type: STRING, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// IArrayLabel
func IArrayLabel(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_LABEL, Type: INT, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// BArrayLabel
func BArrayLabel(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_LABEL, Type: BOOL, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// SArrayVal
func SArrayVal(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_VAL, Type: STRING, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// IArrayVal
func IArrayVal(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_VAL, Type: INT, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

// ----------------------------------------
// BArrayVal
func BArrayVal(arg Arg) Arg {
	return Arg{Kind: TYPED_ARRAY_VAL, Type: BOOL, ExprVal: &Expr{Left: arg, Op: UNUSED, Right: Arg{}}}
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
// - caution: keep up to date with Arg struct
func (a *Arg) Copy() Arg {
	//------------------------------------------------------------
	// alloc
	newA := new(Arg)
	//------------------------------------------------------------
	// copy all fields:
	// - Kind:
	newA.Kind = a.Kind
	// - Type:
	newA.Type = a.Type
	// - Name:
	newA.Name = a.Name
	// - IntVal:
	newA.IntVal = a.IntVal
	// - StringVal:
	newA.StringVal = a.StringVal
	// - BoolVal:
	newA.BoolVal = a.BoolVal
	// - ExprVal:
	if nil != a.ExprVal {
		expr := a.ExprVal.Copy()
		newA.ExprVal = &expr
	}
	//------------------------------------------------------------
	// return
	return *newA
}

// ----------------------------------------
// apply selector (= arg) to vars and an entry
// - if selector is nil -> it is fulfilled
// - if entry == nil -> then only vars and direct values can be used (-> usage e.g.: "create" link)
// - nb: result of expression evaluation must be boolean!
// returns true if selector is fulfilled for e otherwise false
func (sel *Arg) Apply(vars Vars, e *Entry) bool {
	if ARGS_EVAL_TRACE.DoTrace() {
		/**/ String2TraceFile("apply selector:\n")
	}
	if nil == sel {
		return true
	}
	if !sel.Eval(vars, e) {
		return false
	}
	if BOOL != sel.Type {
		return false
	}
	return sel.BoolVal
}

// ----------------------------------------
// EVALUATION
// resolve all values in arg if not yet basic; changes values (and type) in arg for this eval!
// - vars are needed for eval
// - entry e is needed as its properties need to be considered in the eval of a label
// - nb: if entry is nil --> no entry label is allowed in expression, only vars and direct values!
// returns true if evaluation was ok, false otherwise (eg type incompatibility, label or var not found)
func (arg *Arg) Eval(vars Vars, entry *Entry) bool {
	if ARGS_EVAL_TRACE.DoTrace() {
		/**/ String2TraceFile("EVAL Arg:\n")
	}

	if nil == arg {
		Panic("Eval: arg is nil")
	}
	switch arg.Kind {
	// -------------------
	case VAL:
		// nothing to be done - value is already evaluated
		// debug only:
		switch arg.Type {
		case INT:
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("VAL=%d\n", arg.IntVal))
			}
		case STRING:
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("VAL=%s\n", arg.StringVal))
			}
		case BOOL:
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("VAL=%t\n", arg.BoolVal))
			}
		default:
			Panic(fmt.Sprintf("Eval: VAL %s: ill. arg.Type = %s", arg.Kind, arg.Type))
		}

	// -------------------
	// retrieve variable and temporarily set arg's type value to it
	case VAR:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("VAR=%s\n", arg.Name))
		}
		v := vars[arg.Name]
		if "" == v.Kind {
			// var does not exist
			Panic(fmt.Sprintf("VAR: var %s does not exist", arg.Name))
		} else {
			// set arg's value and type temporarily to var's value
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("type=%s\n", arg.Type))
			}
			v.Type = arg.Type

			switch v.Type {
			case INT:
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("val=%d\n", arg.IntVal))
				}
				arg.IntVal = v.IntVal
			case STRING:
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("val=%s\n", arg.StringVal))
				}
				arg.StringVal = v.StringVal
			case BOOL:
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("val=%t\n", arg.BoolVal))
				}
				arg.BoolVal = v.BoolVal
			default:
				Panic(fmt.Sprintf("Eval: VAR %s: ill. v.Type = %s", arg.Name, v.Type))
			}
		}

	// -------------------
	// evaluate function
	case FU:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("FU=%s\n", arg.FuName))
		}
		switch arg.Type {
		case INT:
			switch arg.FuName {
			case CLOCK_FUNCTION:
				arg.IntVal = Clock()
				arg.Type = INT
			default:
				Panic(fmt.Sprintf("Eval: ill. int fu = %s", arg.FuName))
			}
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("FU=%d\n", arg.IntVal))
			}
		case STRING:
			switch arg.FuName {
			case FID_FUNCTION:
				arg.StringVal = Fid()
				arg.Type = STRING
			case UUID_FUNCTION:
				arg.StringVal = UuidUserFu()
				arg.Type = STRING
			default:
				Panic(fmt.Sprintf("Eval: ill. string fu = %s", arg.FuName))
			}
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("FU=%s\n", arg.StringVal))
			}
			//		case BOOL:
			//			// @@@ no bool fus yet
			//			if ARGS_EVAL_TRACE.DoTrace() {
			//				/**/ String2TraceFile(fmt.Sprintf("FU=%t\n", ...))
			//			}
		default:
			Panic(fmt.Sprintf("Eval: ill. fu = %s", arg.FuName))
		}

	// -------------------
	// retrieve label on entry and temporarily set arg's  type and value to it
	case LABEL:
		// entry must not be nil!
		if nil == entry {
			Panic(fmt.Sprintf("Eval: can't eval entry label (%s) in expression if no entry is given", arg.Name))
		}
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("LABEL=%s\n", arg.Name))
		}
		// set arg's value temporarily to entry property's value:
		switch arg.Name {
		//	@@@raus:	// system property: caution: keep up to date with entry struct!
		//	@@@raus:	case TYPE:
		//	@@@raus:		arg.Type = STRING
		//	@@@raus:		arg.StringVal = entry.Type
		//	@@@raus:		if ARGS_EVAL_TRACE.DoTrace() {
		//	@@@raus:			/**/ String2TraceFile(fmt.Sprintf("val=%s\n", entry.Type))
		//	@@@raus:		}
		default:
			// get Arg for the label from the entry:
			entryArg := entry.EProps[arg.Name]
			if "" == entryArg.Kind {
				// entry does not have a property with this label
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile("entry does not have a property with this label\n")
				}

				Panic(fmt.Sprintf("Eval: LABEL \"%s\": undefined entry property = %s;\n  entry = %s;\n", arg.Name, arg.Name, entry.ToString(0)))
				return false
			}
			arg.Type = entryArg.Type
			switch entryArg.Type {
			case INT:
				arg.IntVal = entryArg.IntVal
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("int val=%d\n", entryArg.IntVal))
				}
			case STRING:
				arg.StringVal = entryArg.StringVal
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("string val=%s\n", entryArg.StringVal))
				}
			case BOOL:
				arg.BoolVal = entryArg.BoolVal
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("string bool=%s\n", entryArg.BoolVal))
				}
			default:
				Panic(fmt.Sprintf("Eval: LABEL %s: ill. type = %s", arg.Type))
			}
		}

	// -------------------
	case EXPR:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("EXPR\n")
		}
		// eval left:
		if !arg.ExprVal.Left.Eval(vars, entry) {
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("left eval = not ok\n")
			}
			Panic(fmt.Sprintf("Eval: EXPR: left eval = not ok"))
		}
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("left eval = ok\n")
		}
		// eval right arg, if not unary op
		// - TBD...umgekehrt: eval left...
		// - eval right:
		// - caution: only, if it is not a unary operator!
		// - @@@ quite explicit test here...
		if NOT != arg.ExprVal.Op && PLUS != arg.ExprVal.Op && MINUS != arg.ExprVal.Op {
			// eval right
			if !arg.ExprVal.Right.Eval(vars, entry) {
				Panic(fmt.Sprintf("right eval = not ok"))
			}
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("right eval = ok\n")
			}
			// check type compatibility:
			if arg.ExprVal.Left.Type != arg.ExprVal.Right.Type {
				Panic(fmt.Sprintf("EXPR: type incompatibility: left type = %s, right type = %s; full arg info = %s", arg.ExprVal.Left.Type.String(), arg.ExprVal.Right.Type.String(), arg.ToString(0)))
			}
		}
		// apply operator (depending on the types of both sides) and temporarily set arg's value and type:
		switch arg.ExprVal.Left.Type {
		case INT:
			switch arg.ExprVal.Op {
			// ----
			// arithmetic operators:
			case ADD:
				arg.IntVal = arg.ExprVal.Left.IntVal + arg.ExprVal.Right.IntVal
				arg.Type = INT
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%d\n", arg.IntVal))
				}
			case SUB:
				arg.IntVal = arg.ExprVal.Left.IntVal - arg.ExprVal.Right.IntVal
				arg.Type = INT
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%d\n", arg.IntVal))
				}
			case MUL:
				arg.IntVal = arg.ExprVal.Left.IntVal * arg.ExprVal.Right.IntVal
				arg.Type = INT
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%d\n", arg.IntVal))
				}
			case DIV:
				arg.IntVal = arg.ExprVal.Left.IntVal / arg.ExprVal.Right.IntVal
				arg.Type = INT
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%d\n", arg.IntVal))
				}
			case MOD:
				arg.IntVal = arg.ExprVal.Left.IntVal % arg.ExprVal.Right.IntVal
				arg.Type = INT
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%d\n", arg.IntVal))
				}
			// ----
			// relational operators:
			case EQUAL:
				arg.BoolVal = arg.ExprVal.Left.IntVal == arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case LESS:
				arg.BoolVal = arg.ExprVal.Left.IntVal < arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case LESS_EQUAL:
				arg.BoolVal = arg.ExprVal.Left.IntVal <= arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case GREATER:
				arg.BoolVal = arg.ExprVal.Left.IntVal > arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case GREATER_EQUAL:
				arg.BoolVal = arg.ExprVal.Left.IntVal >= arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case NOT_EQUAL:
				arg.BoolVal = arg.ExprVal.Left.IntVal != arg.ExprVal.Right.IntVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case CONCAT:
				if arg.ExprVal.Right.Type == STRING {
					arg.StringVal = fmt.Sprintf("%d%s", arg.ExprVal.Left.IntVal, arg.ExprVal.Right.StringVal)
				} else if arg.ExprVal.Right.Type == INT {
					arg.StringVal = fmt.Sprintf("%d%i", arg.ExprVal.Left.IntVal, arg.ExprVal.Right.IntVal)
				} else if arg.ExprVal.Right.Type == BOOL {
					arg.StringVal = fmt.Sprintf("%d%t", arg.ExprVal.Left.IntVal, arg.ExprVal.Right.BoolVal)
				} else {
					Panic(fmt.Sprintf("Eval: CONCAT: ill. right arg type"))
				}
				arg.Type = STRING
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%s\n", arg.StringVal))
				}
			default:
				Panic(fmt.Sprintf("Eval: EXPR: ill. int Op = %s", arg.ExprVal.Op))
			}
		case STRING:
			switch arg.ExprVal.Op {
			// relational operators:
			case EQUAL:
				arg.BoolVal = arg.ExprVal.Left.StringVal == arg.ExprVal.Right.StringVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case NOT_EQUAL:
				arg.BoolVal = arg.ExprVal.Left.StringVal != arg.ExprVal.Right.StringVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			// binary string operators:
			// TBD: support any types, even mixed, and convert them to string
			case CONCAT:
				if arg.ExprVal.Right.Type == STRING {
					arg.StringVal = fmt.Sprintf("%s%s", arg.ExprVal.Left.StringVal, arg.ExprVal.Right.StringVal)
				} else if arg.ExprVal.Right.Type == INT {
					arg.StringVal = fmt.Sprintf("%s%i", arg.ExprVal.Left.StringVal, arg.ExprVal.Right.IntVal)
				} else if arg.ExprVal.Right.Type == BOOL {
					arg.StringVal = fmt.Sprintf("%s%t", arg.ExprVal.Left.StringVal, arg.ExprVal.Right.BoolVal)
				} else {
					Panic(fmt.Sprintf("Eval: CONCAT: ill. right arg type"))
				}
				arg.Type = STRING
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%s\n", arg.StringVal))
				}
			default:
				Panic(fmt.Sprintf("Eval: EXPR: ill. string Op = %s", arg.ExprVal.Op))
			}
		case BOOL:
			switch arg.ExprVal.Op {
			// boolean operators:
			case AND:
				arg.BoolVal = arg.ExprVal.Left.BoolVal && arg.ExprVal.Right.BoolVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case OR:
				arg.BoolVal = arg.ExprVal.Left.BoolVal || arg.ExprVal.Right.BoolVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case NOT:
				arg.BoolVal = !arg.ExprVal.Left.BoolVal
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case EQUAL:
				arg.BoolVal = (arg.ExprVal.Left.BoolVal == arg.ExprVal.Right.BoolVal)
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case NOT_EQUAL:
				arg.BoolVal = (arg.ExprVal.Left.BoolVal != arg.ExprVal.Right.BoolVal)
				arg.Type = BOOL
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%t\n", arg.BoolVal))
				}
			case CONCAT:
				if arg.ExprVal.Right.Type == STRING {
					arg.StringVal = fmt.Sprintf("%t%s", arg.ExprVal.Left.BoolVal, arg.ExprVal.Right.StringVal)
				} else if arg.ExprVal.Right.Type == INT {
					arg.StringVal = fmt.Sprintf("%t%i", arg.ExprVal.Left.BoolVal, arg.ExprVal.Right.IntVal)
				} else if arg.ExprVal.Right.Type == BOOL {
					arg.StringVal = fmt.Sprintf("%t%t", arg.ExprVal.Left.BoolVal, arg.ExprVal.Right.BoolVal)
				} else {
					Panic(fmt.Sprintf("Eval: CONCAT: ill. right arg type"))
				}
				arg.Type = STRING
				if ARGS_EVAL_TRACE.DoTrace() {
					/**/ String2TraceFile(fmt.Sprintf("result=%s\n", arg.StringVal))
				}
			default:
				Panic(fmt.Sprintf("Eval: EXPR: ill. bool Op = %s", arg.ExprVal.Op))
			}
		default:
			Panic(fmt.Sprintf("Eval: EXPR: ill. Op type = %s", arg.ExprVal.Left.Type))
		}

	// -------------------
	case DYN_ARRAY_REF:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("DYN_ARRAY_REF\n")
		}
		// eval arg: this is the index and must resolve to INT; nb: only Left of Expr is used;
		if !arg.ExprVal.Left.Eval(vars, entry) {
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("Eval: DYN_ARRAY_REF: eval of arg = not ok\n")
			}
			Panic(fmt.Sprintf("Eval: DYN_ARRAY_REF: eval of arg = not ok"))
		}
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("DYN_ARRAY_REF arg eval = ok\n")
		}
		// check type to be INT:
		if arg.ExprVal.Left.Type != INT {
			Panic(fmt.Sprintf("DYN_ARRAY_REF: index must evaluate to INT, but found type = %s", arg.ExprVal.Left.Type.String()))
		}

		// construct the resolved label: <name> # <int> which is a STRING!
		// nb: arg.StringVal contains the basic label name
		arg.Kind = VAL
		arg.Type = STRING
		arg.StringVal = fmt.Sprintf("%s%c%d", arg.StringVal, ArrayChar, arg.ExprVal.Left.IntVal)

		// fmt.Println(fmt.Sprintf("DYN_ARRAY_REF: %s", arg.StringVal))

	// -------------------
	case TYPED_ARRAY_LABEL:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("TYPED_ARRAY_LABEL\n")
		}
		// eval arg: this is the label name (with '#'s) and must resolve to STRING; nb: only Left of Expr is used;
		if !arg.ExprVal.Left.Eval(vars, entry) {
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("Eval: TYPED_ARRAY_LABEL: eval of arg = not ok\n")
			}
			Panic(fmt.Sprintf("Eval: TYPED_ARRAY_LABEL: eval of arg = not ok"))
		}
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("TYPED_ARRAY_LABEL arg eval = ok\n")
		}
		// check type to be STRING:
		if arg.ExprVal.Left.Type != STRING {
			Panic(fmt.Sprintf("TYPED_ARRAY_LABEL: label name must evaluate to STRING, but found type = %s", arg.ExprVal.Left.Type.String()))
		}
		// construct the resolved typed label:
		arg.Kind = LABEL
		// arg.Type is set already
		arg.Name = arg.ExprVal.Left.StringVal

		// fmt.Println(fmt.Sprintf("    array label: %s; Type: %s", arg.Name, arg.Type.String()))

		// !!! and now evaluate the normal label arg !!!
		// eval arg: this is the label name (with '#'s) and must resolve to STRING; nb: only Left of Expr is used;
		if !arg.Eval(vars, entry) {
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("Eval: *TYPED_ARRAY_LABEL*: eval = not ok\n")
			}
			Panic(fmt.Sprintf("Eval: *TYPED_ARRAY_LABEL*: eval of arg = not ok"))
		}

	// -------------------
	case TYPED_ARRAY_VAL:
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("TYPED_ARRAY_VAL\n")
		}
		// eval arg: this is the label name (with '#'s) and must resolve to STRING; nb: only Left of Expr is used;
		if !arg.ExprVal.Left.Eval(vars, entry) {
			if ARGS_EVAL_TRACE.DoTrace() {
				/**/ String2TraceFile("Eval: TYPED_ARRAY_VAL: eval of arg = not ok\n")
			}
			Panic(fmt.Sprintf("Eval: TYPED_ARRAY_LABEL: eval of arg = not ok"))
		}
		if ARGS_EVAL_TRACE.DoTrace() {
			/**/ String2TraceFile("TYPED_ARRAY_VAL arg eval = ok\n")
		}
		// check type to be STRING:
		if arg.ExprVal.Left.Type != STRING {
			Panic(fmt.Sprintf("TYPED_ARRAY_VAL: label name must evaluate to STRING, but found type = %s", arg.ExprVal.Left.Type.String()))
		}
		// construct the resolved typed val:
		// - Kind
		arg.Kind = VAL
		// - Type is set already
		// - Val
		if arg.Type == STRING {
			arg.StringVal = arg.ExprVal.Left.StringVal
			// fmt.Println(fmt.Sprintf("    array val: Type=%s; Val=%s", arg.Type.String(), arg.StringVal))
		} else if arg.Type == INT {
			arg.IntVal = arg.ExprVal.Left.IntVal
			// fmt.Println(fmt.Sprintf("    array val: Type=%s; Val=%d", arg.Type.String(), arg.IntVal))
		} else if arg.Type == BOOL {
			arg.BoolVal = arg.ExprVal.Left.BoolVal
			// fmt.Println(fmt.Sprintf("    array val: Type=%s; Val=%t", arg.Type.String(), arg.BoolVal))
		} else {
			Panic(fmt.Sprintf("Eval: TYPED_ARRAY_LABEL: ill. arg type"))
		}

		arg.Name = arg.ExprVal.Left.StringVal

	// -------------------
	default:
		Panic(fmt.Sprintf("Eval: ill. arg Kind = %s", arg.Kind))
	}
	return true
}

// ----------------------------------------
func (arg Arg) String() string {
	s := ""
	switch arg.Kind {
	case VAL:
		switch arg.Type {
		case INT:
			s = fmt.Sprintf("%s%d", s, arg.IntVal)
		case STRING:
			// quotes @@@tbd if not sub string type is ENTRY_TYPE or URL
			s = fmt.Sprintf("%s%s", s, arg.StringVal)
		case BOOL:
			if arg.BoolVal {
				s = fmt.Sprintf("%strue", s)
			} else {
				s = fmt.Sprintf("%sfalse", s)
			}
		}
	case LABEL:
		s = fmt.Sprintf("%s%s", s, arg.Name)
	case VAR:
		s = fmt.Sprintf("%s%s", s, strings.Replace(arg.Name, "$", "\\$", 2))
	case EXPR:
		s = fmt.Sprintf("%s(%s)", s, arg.ExprVal)
	}
	return s
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
// @@@ geht es einfacher als sprintf? zb string append function?
// NB: converts INFINITE (= MAX_INT) value to string "INFINITE"
func (a *Arg) ToStringWithDetails(tab int, text string, detailsFlag bool) string {
	s := NBlanksToString("", tab)

	if "" != text {
		s = fmt.Sprintf("%s%s=", s, text)
	}
	switch a.Kind {
	case VAL:
		if detailsFlag {
			s = fmt.Sprintf("%s(VAL=", s)
		}
		switch a.Type {
		case INT:
			if a.IntVal == INFINITE {
				s = fmt.Sprintf("%sINFINITE", s)
			} else {
				s = fmt.Sprintf("%s%d", s, a.IntVal)
			}
		case STRING:
			if a.StringSubType == NORMAL {
				s = fmt.Sprintf("%s\"%s\"", s, a.StringVal)
			} else {
				// ENTRY_TYPE or URL
				s = fmt.Sprintf("%s%s", s, a.StringVal)
			}
		case BOOL:
			s = fmt.Sprintf("%s%t", s, a.BoolVal)
		default:
			s = fmt.Sprintf("%sill. arg type = %s", s, a.Type)
		}
		if detailsFlag {
			s = fmt.Sprintf("%s)", s)
		}
	case LABEL:
		if detailsFlag {
			s = fmt.Sprintf("%s(LABEL=", s)
		}
		s = fmt.Sprintf("%s%s", s, a.Name)
		if detailsFlag {
			s = fmt.Sprintf("%s)", s)
		}
	case VAR:
		if detailsFlag {
			s = fmt.Sprintf("%s(VAR=", s)
		}
		s = fmt.Sprintf("%s%s", s, a.Name)
		if detailsFlag {
			s = fmt.Sprintf("%s)", s)
		}
	case EXPR:
		if detailsFlag {
			s = fmt.Sprintf("%s(EXPR=", s)
		}
		s = fmt.Sprintf("%s%s", s, a.ExprVal.ToString(0))
		if detailsFlag {
			s = fmt.Sprintf("%s)", s)
		}
	case FU:
		if detailsFlag {
			s = fmt.Sprintf("%s(FU=", s)
		}
		s = fmt.Sprintf("%s%s", s, a.FuName)
		if detailsFlag {
			s = fmt.Sprintf("%s)", s)
		}
	default:
		if a.IsEmpty() {
			s = fmt.Sprintf("%s<empty>", s)
		} else {
			s = fmt.Sprintf("%sill. arg kind = \"%s\"", s, a.Kind)
		}
	}
	return s
}

// ----------------------------------------
// default assumptions for the iprint interface:
func (a *Arg) ToString(tab int) string {
	return a.ToStringWithDetails(tab, "", false /* detailsFlag */)
}

// ----------------------------------------
func (a Arg) ArgPrint(tab int, text string, detailsFlag bool) {
	/**/ String2TraceFile(a.ToStringWithDetails(tab, text, detailsFlag))
}

// ----------------------------------------
func (a Arg) ArgPrintln(tab int, text string, detailsFlag bool) {
	/**/ a.ArgPrint(tab, text, detailsFlag)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// iprint interface
// IPRINT INTERFACE
////////////////////////////////////////

// ----------------------------------------
func (a *Arg) IsEmpty() bool {
	if nil == a || a.Kind == "" {
		return true
	} else {
		return false
	}
}

// ----------------------------------------
// default assumptions for the iprint interface:
func (a *Arg) Print(tab int) {
	a.ArgPrint(tab, "", false /* detailsFlag */)
}

// ----------------------------------------
// default assumptions for the iprint interface:
func (a *Arg) Println(tab int) {
	a.ArgPrintln(tab, "", false /* detailsFlag */)
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
