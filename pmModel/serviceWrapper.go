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
	. "cca/scheduler"
)

// nb: machine parameter is needed for m.Vars (and also for debug traces) in services
type ServiceFunc func(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, inCid, outCid string, controllerChannel ControllerChannel)

type ServiceWrapper struct {
	// service function:
	Fu ServiceFunc
	// only for docu (optional):
	Name string
	// internally used only: resolved automatically:
	InCid  string
	OutCid string
}

////////////////////////////////////////
// constructors
////////////////////////////////////////

func NewServiceWrapper(fu ServiceFunc, name string) *ServiceWrapper {
	sw := new(ServiceWrapper)
	sw.Fu = fu
	sw.Name = name
	return sw
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (sw *ServiceWrapper) Copy() *ServiceWrapper {
	//------------------------------------------------------------
	// alloc
	newSw := NewServiceWrapper(sw.Fu, sw.Name)
	//------------------------------------------------------------
	// copy all fields:
	// - Fu: copied by constructor above
	// - Name: copied by constructor above
	// - InCid:
	newSw.InCid = sw.InCid
	// - OutCid:
	newSw.OutCid = sw.OutCid
	//------------------------------------------------------------
	// return
	return newSw
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
