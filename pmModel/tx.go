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

// tx state
var RUNNING string = "running"
var COMMITTED string = "committed"
var ROLLEDBACK string = "rolledback"

type Tx struct {
	Id string
	// RUNNING, COMMITTED, ROLLEDBACK
	State string
	// OCC or PCC
	Txcc string
	Pcc
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (tx *Tx) Copy() *Tx {
	//------------------------------------------------------------
	// alloc
	newTx := new(Tx)
	//------------------------------------------------------------
	// copy all fields:
	// - Id:
	newTx.Id = tx.Id
	// - State:
	newTx.State = tx.State
	// - Txcc:
	newTx.Txcc = tx.Txcc
	// - Pcc:
	newTx.Pcc.LockedCids = tx.Pcc.LockedCids.Copy()
	//------------------------------------------------------------
	// return
	return newTx
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (tx *Tx) IsEmpty() bool {
	if nil == tx || tx.Id == "" {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (tx *Tx) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%sId=%s, ", s, tx.Id)
	s = fmt.Sprintf("%sState=%s, ", s, tx.State)
	s = fmt.Sprintf("%sTxcc=%s, Pcc=", s, tx.Txcc)
	s = fmt.Sprintf("%s%a", s, tx.Pcc.LockedCids.ToString(0))
	return s
}

// ----------------------------------------
func (tx Tx) Print(tab int) {
	/**/ String2TraceFile(tx.ToString(tab))
}

// ----------------------------------------
func (tx Tx) Println(tab int) {
	/**/ tx.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
