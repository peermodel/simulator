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
	"fmt"
)

type Peer struct {
	Id            string
	Pic           string
	Poc           string
	Wirings       map[string]*Wiring
	IsSysPeerFlag bool
	// for debug only
	WiringWids Strings
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// create peer and resolve the names for its PIC and POC:
// - Pic := <peerId> <Sep> <PIC>
// - Poc := <peerId> <Sep> <POC>
func NewPeer(id string) *Peer {
	p := new(Peer)

	p.Id = id
	p.Pic = fmt.Sprintf("%s%s%s", id, SEP, PIC)
	p.Poc = fmt.Sprintf("%s%s%s", id, SEP, POC)

	p.Wirings = make(map[string]*Wiring)
	p.IsSysPeerFlag = false
	p.WiringWids = Strings{}

	return p
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (p *Peer) Copy() *Peer {
	//------------------------------------------------------------
	// alloc
	newP := NewPeer(p.Id)
	//------------------------------------------------------------
	// copy all fields:
	// - Id:
	newP.Id = p.Id
	// - Pic:
	newP.Pic = p.Pic
	// - Poc:
	newP.Poc = p.Poc
	// - Wirings:
	for wid, w := range p.Wirings {
		newP.Wirings[wid] = w.Copy()
	}
	// - WiringWids:
	newP.WiringWids = p.WiringWids.Copy()
	// - IsSysPeerFlag
	newP.IsSysPeerFlag = p.IsSysPeerFlag
	//------------------------------------------------------------
	// return
	return newP
}

// ----------------------------------------
// add wiring to peer and thereby resolve all names:
// caution: changes also name of wiring (as w is passed by reference!)
// @@@check: treatment of IOP  using the subpid field....
func (p *Peer) AddWiring(w *Wiring) {

	// resolve wiring name:
	newWId := fmt.Sprintf("%s%s%s", p.Id, SEP, w.Id)
	w.Id = newWId

	// resolve wiring container name:
	w.WCId = fmt.Sprintf("%s%s%s", w.Id, SEP, WC)

	// resolve container names used in links:
	for _, l := range w.Links {
		switch l.Type {

		// guard:
		// - C1 := <peerId>_<modeledC> where peer is either this peer or a subpeer
		// - C2 := resolved wiring container name
		case GUARD:
			if IOP_PEER == l.SubPid {
				l.C1 = fmt.Sprintf("%s%s%s", IOP_PEER, SEP, l.modelC)
			} else if "" == l.SubPid {
				l.C1 = fmt.Sprintf("%s%s%s", p.Id, SEP, l.modelC)
			} else {
				l.C1 = fmt.Sprintf("%s%s%s%s%s", p.Id, SEP, l.SubPid, SEP, l.modelC)
			}
			l.C2 = w.WCId

		// service in:
		// - C1 := resolved wiring container name
		// - C2 := <resolvedWId> <SEP> <serviceId> <SEP> <SIC>
		case SERVICE_IN:
			l.C1 = w.WCId
			l.C2 = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, l.Sid, SEP, SIC)

		// service call:
		// - C1 := <resolvedWId> <SEP> <serviceId> <SEP> <SIC>
		// - C2 := <resolvedWId> <SEP> <serviceId> <SEP> <SOC>
		case SERVICE:
			l.C1 = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, l.Sid, SEP, SIC)
			l.C2 = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, l.Sid, SEP, SOC)

		// service out:
		// - C1 := <resolvedWId> <SEP> <serviceId> <SEP> <SOC>
		// - C2 := resolved wiring container name
		case SERVICE_OUT:
			l.C1 = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, l.Sid, SEP, SOC)
			l.C2 = w.WCId

		// action:
		// - C1 := resolved wiring container name
		// - C2 := <peerId> <SEP> <modeledC>, where peer is either this peer or a subpeer
		case ACTION:
			l.C1 = w.WCId
			if IOP_PEER == l.SubPid {
				l.C2 = fmt.Sprintf("%s%s%s", IOP_PEER, SEP, l.modelC)
			} else if "" == l.SubPid {
				l.C2 = fmt.Sprintf("%s%s%s", p.Id, SEP, l.modelC)
			} else {
				l.C2 = fmt.Sprintf("%s%s%s%s%s", p.Id, SEP, l.SubPid, SEP, l.modelC)
			}

		default:
			Panic(fmt.Sprintf("AddWiring: resolve meta model: ill. link type"))
		}
	}

	// resolve container names used in service wrappers:
	// see above!
	for sid, sw := range w.ServiceWrappers {
		sw.InCid = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, sid, SEP, SIC)
		sw.OutCid = fmt.Sprintf("%s%s%s%s%s", w.Id, SEP, sid, SEP, SOC)
	}

	// add wiring to peer under the new Id
	if nil == p.Wirings {
		p.Wirings = map[string]*Wiring{}
	}
	p.Wirings[newWId] = w

	// for debug - sorted printing
	p.WiringWids = p.WiringWids.SortedInsertString(newWId)
}

// ----------------------------------------
func (p *Peer) RemoveWiring(w *Wiring) {
	delete(p.Wirings, w.Id)
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (p *Peer) IsEmpty() bool {
	if nil == p || p.Id == "" {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (p *Peer) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%sPeer %s:", s, p.Id)
	for _, w := range p.Wirings {
		/**/ w.Println(tab + TAB)
		s = fmt.Sprintf("%s%s:", s, w.ToString(tab+TAB))
	}
	return s
}

// ----------------------------------------
func (p *Peer) Print(tab int) {
	/**/ String2TraceFile(p.ToString(tab))
}

// ----------------------------------------
func (p *Peer) Println(tab int) {
	p.Print(tab)
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
