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

type Expr struct {
	Left  Arg
	Op    OpTypeEnum
	Right Arg
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (expr *Expr) Copy() Expr {
	//------------------------------------------------------------
	// alloc
	newExpr := new(Expr)
	//------------------------------------------------------------
	// copy all fields:
	// - Left:
	newExpr.Left = expr.Left.Copy()
	// - Op:
	newExpr.Op = expr.Op
	// - Right:
	newExpr.Right = expr.Right.Copy()
	//------------------------------------------------------------
	// return
	return *newExpr
}

// ----------------------------------------
func (expr Expr) String() string {
	s := fmt.Sprintf("%s", expr.Left)
	s = fmt.Sprintf("%s {$ %s $} ", s, expr.Op)
	s = fmt.Sprintf("%s%s", s, expr.Right)
	return s
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (expr *Expr) IsEmpty() bool {
	if nil == expr {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (expr *Expr) ToString(tab int) string {
	s := NBlanksToString("", tab)
	// unary operator?
	if NOT == expr.Op {
		s = fmt.Sprintf("%s %s (%s)", s, expr.Op.String(), expr.Left.ToString(0))
	} else {
		s = fmt.Sprintf("%s(%s%s%s)", s, expr.Left.ToString(0), expr.Op.String(), expr.Right.ToString(0))
	}
	return s
}

// ----------------------------------------
func (expr *Expr) Print(tab int) {
	/**/ String2TraceFile(expr.ToString(tab))
}

// ----------------------------------------
func (expr *Expr) Println(tab int) {
	/**/ expr.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
