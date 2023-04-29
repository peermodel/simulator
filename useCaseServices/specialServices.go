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

package useCaseServices

import (
	"fmt"
	"strings"

	. "github.com/peermodel/simulator/controller"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/pmModel"
	. "github.com/peermodel/simulator/scheduler"
)

//////////////////////////////////////////////////////////////
// Built-in Service "SelectAdvice"
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// User: select one advice out of N > 1
// Mockup: RANDOM selection (using map in go lang)
func SelectAdvice(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, inCid string, outCid string, controllerChannel ControllerChannel) {
	//------------------------------------------------------------
	// print service name and cid of container
	String2TraceFile("\n")
	String2TraceFile(fmt.Sprintf("%s: SELECT ADVICE (time = %d):\n", inCid, CLOCK))
	//------------------------------------------------------------
	// take all entries from the wiring container, print them and remember one of them (... the first one as mock up)
	selectedEntry := AllocEntry()
	// done := false
	var Es EntryPtrs
	// use map, because access is random
	entriesMap := make(map[int]*Entry)
	i := 0
	for {
		//------------------------------------------------------------
		// take entry
		entry := ps.Take(inCid, "*", nil /* no selector */, vars)
		//------------------------------------------------------------
		// done?
		if nil == entry {
			break
		}
		//------------------------------------------------------------
		if 0 == strings.Compare("event", entry.GetType()) {
			// trace
			String2TraceFile("    ")
			/**/ entry.Println(0 /* ind */)
			// append to map
			Es = append(Es, entry)
			entriesMap[i] = entry
			// caution: key must be dense... ?!
			i++
		}
	}
	//------------------------------------------------------------
	// "randomly" pick one entry
	j := 0
	for _, val := range entriesMap {
		j++
		// TBD :-) take the "middle" one
		if j >= i/2 {
			selectedEntry = val.Copy()
			// done = true
		}
	}
	//------------------------------------------------------------
	// emit the selected entry back to wiring container
	ps.Emit(outCid, selectedEntry, vars, scheduler)
	// trace
	String2TraceFile("    ==> selected entry: ")
	/**/ selectedEntry.Println(0 /* ind */)
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
