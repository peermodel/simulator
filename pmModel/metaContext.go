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
// Date: 2015, 2016
////////////////////////////////////////

package pmModel

import (
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/latex"
	. "github.com/peermodel/simulator/scheduler"
	. "github.com/peermodel/simulator/slotInterface"
)

type MetaContext struct {
	PeerSpace    *PeerSpace
	Transactions map[string]*Tx
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

//------------------------------------------------------------
// CAUTION: keep up to date, if MetaContext struct changes;
func NewMetaContext() *MetaContext {
	//------------------------------------------------------------
	// alloc
	metaCtx := new(MetaContext)
	//------------------------------------------------------------
	// alloc structured fields
	// - PeerSpace
	metaCtx.PeerSpace = NewPeerSpace()
	// - Transactions
	metaCtx.Transactions = make(map[string]*Tx)
	//------------------------------------------------------------
	// return
	return metaCtx
}

////////////////////////////////////////
// methods
////////////////////////////////////////

////////////////////////////////////////
// IMetaContext interface implementation:
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (metaCtx MetaContext) Copy() interface{} {
	//------------------------------------------------------------
	// alloc new context:
	newMetaCtx := NewMetaContext()
	//------------------------------------------------------------
	// copy all fields:
	// - PeerSpace:
	newMetaCtx.PeerSpace = metaCtx.PeerSpace.Copy()
	// - Transactions:
	for tid, tx := range metaCtx.Transactions {
		newMetaCtx.Transactions[tid] = tx.Copy()
	}
	//------------------------------------------------------------
	// return
	return newMetaCtx
}

// =========================================================
func (metaCtx MetaContext) EntryHunter(eid string, scheduler *Scheduler) {
	metaCtx.PeerSpace.EntryHunter(eid, scheduler)
}

// ----------------------------------------
func (metaCtx MetaContext) MetaModel2Latex(testCaseName string, testCaseLatexConfig *LatexConfig) {
	metaCtx.PeerSpace.MetaModel2Latex(testCaseName, testCaseLatexConfig)
}

// ----------------------------------------
// check if condition is fulfilled
func (metaCtx MetaContext) ConditionIsFulfilled(condition *Event) bool {
	help := metaCtx.PeerSpace.ConditionIsFulfilled(condition.UserEvent, condition.IssueEventTime)
	return help
}

// ----------------------------------------
func (metaCtx MetaContext) SpacePrint(tl TraceLevelEnum, nBlanks int, printAlsoEmptyContainersFlag bool) {
	metaCtx.PeerSpace.SpacePrint(tl, nBlanks, printAlsoEmptyContainersFlag)
}

// ----------------------------------------
// process a ripe pm slot
func (metaCtx MetaContext) ProcessRipeUserSlot(userSlot ISlot, scheduler *Scheduler) {
	metaCtx.PeerSpace.ProcessRipePMSlot(userSlot, scheduler)
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (metaCtx MetaContext) IsEmpty() bool {
	return false
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (metaCtx MetaContext) Print(tab int) {
	/**/ NBlanks2TraceFile(tab)
	/**/ String2TraceFile("metaCtx={")

	/**/
	NBlanks2TraceFile(tab)
	/**/ String2TraceFile("Peers:")
	for _, p := range metaCtx.PeerSpace.Peers {
		p.Println(tab + IND)
	}

	/**/
	NBlanks2TraceFile(tab)
	/**/ String2TraceFile("Containers:")
	for _, c := range metaCtx.PeerSpace.Containers {
		c.Println(tab + IND)
	}

	/**/
	NBlanks2TraceFile(tab)
	/**/ String2TraceFile("Transactions:")
	for _, tx := range metaCtx.Transactions {
		tx.Print(tab + IND)
	}
	/**/ String2TraceFile("}")
}

// ----------------------------------------
func (metaCtx MetaContext) Println(tab int) {
	/**/ metaCtx.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
