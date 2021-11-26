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
	. "cca/controller"
	. "cca/debug"
	. "cca/scheduler"
	"fmt"
)

////////////////////////////////////////
// BuiltIn Services
/////////////////////////////////S///////

////////////////////////////////////////
// SendService
/////////////////////////////////S///////

// LIMITATION: in the model all peers are local!
func SendService(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, incid, outcid string, controllerChannel ControllerChannel) {
	// take entry of any type:
	// @@@ /**/ m.PrintlnS(TRACE0, TAB, "take next entry from", incid)
	e := ps.Take(incid, WILDCARD, nil /* no selector */, vars)
	if nil == e {
		// @@@ /**/ m.PrintlnS(TRACE0, TAB, "", "no entry found")
		return
	}

	// check destination:
	dest := e.GetDest()
	// @@@ /**/ m.PrintlnS(TRACE0, TAB, "dest", dest)
	if "" == dest {
		Panic(fmt.Sprintf("SendService: dest property not set on entry"))
	}
	// dest is a peer name -> add PIC:
	// @@@ normalize this computation
	resolvedDestCid := fmt.Sprintf("%s%s%s", dest, SEP, PIC)
	// @@@ /**/ m.PrintlnS(TRACE0, TAB, "resolvedDestCid", resolvedDestCid)

	if DEST_WRAP == e.GetType() {
		// dest property was set on link
		// write all entries from e's Data to picC:
		// @@@ without tx??!! this should be the wtx!? remote?!
		// @@@ /**/ m.PrintlnS(TRACE0, TAB, "", "write all entries from e's Data to IOP's PIC")
		for _, nextE := range e.Data {
			// @@@ /**/ m.PrintlnX(TRACE0, TAB*2, "", nextE)
			// add entry to picC:
			ps.Write(resolvedDestCid, nextE, vars, scheduler)
		}
	} else {
		// write entry to picC:
		// @@@ without tx??!! this should be the wtx!? remote?!
		ps.Write(resolvedDestCid, e, vars, scheduler)
	}
}

////////////////////////////////////////
// SourceWrapService
/////////////////////////////////S///////

func SourceWrapService(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, incid, outcid string, controllerChannel ControllerChannel) {
	var Es EntryPtrs
	var sel *Arg = nil
	fid := ""
	entryType := WILDCARD

	// take all entries (first might be of any type and fid);
	// they must all be the same type and belong to the same flow
	// @@@ /**/ m.PrintlnS(TRACE0, TAB, "take all entries from", incid)
	for i := 0; ; i++ {
		e := ps.Take(incid, entryType, sel, vars)
		if nil == e {
			break
		}
		// set entry type and flow for the next entry to be selected
		entryType = e.GetType()
		sel = XValP(SLabel(FID), EQUAL, SVal(fid))
		if "" == fid {
			fid = e.GetFid()
		}
		Es = append(Es, e)
	}
	if 0 == len(Es) {
		// no entry to be wrapped found
		// @@@ /**/ m.PrintlnS(TRACE0, TAB, "", "no entry found")
		return
	}

	// wrap entries into a new one:
	wrapE := NewEntry(SOURCE_WRAP)
	wrapE.SetStringVal(FID, fid)
	wrapE.Data = Es

	// add new entry to sout:
	ps.Emit(outcid, wrapE, vars, scheduler)
}

////////////////////////////////////////
// StopWrapService
/////////////////////////////////S///////

func StopService(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, incid, outcid string, controllerChannel ControllerChannel) {
	if SERVICE_TRACE.DoTrace() {
		/**/ String2TraceFile("\nStopService CALLED\n")
	}
	// does not need any entry: just stop the system
	// - TBD: sender is controller, because machine is not know here...
	controllerChannel <- NewChanSig(STOP, SENDER_IS_SYSTEM, "StopService" /* msg */)
}

////////////////////////////////////////
// EOF
/////////////////////////////////S///////
