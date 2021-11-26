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
//////////////////////////////////////////////////////////////// System: Peer Model State Machine
// Author: eva Kühn
// Date: 2021
//------------------------------------------------------------
// signals that can be sent on the controller and mutex channel
//////////////////////////////////////////////////////////////

package controller

//////////////////////////////////////////////////////////////
// enum: who is the sender
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
type WhoIsSenderTypeEnum int

//------------------------------------------------------------
const (
	SENDER_IS_MACHINE WhoIsSenderTypeEnum = iota
	SENDER_IS_SYSTEM
)

//////////////////////////////////////////////////////////////
// type
//////////////////////////////////////////////////////////////

//============================================================
// signal that is sent on a channel
// - used on controller and mutex channels
//============================================================

//------------------------------------------------------------
type ChanSig struct {
	//------------------------------------------------------------
	// the signal
	Sig SignalTypeEnum
	//------------------------------------------------------------
	// info about sender
	// - who is the sender: system or machine
	WhoIsSender WhoIsSenderTypeEnum
	// - machine key, if sender is a machine
	SenderMachineKey string
	// - machine key, if sender is system
	MsgFromSystem string
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create a new signal to be sent on a channel
// - is sender is machine, info is the machine key
// - else info is any message from the system, eg the calling fu
func NewChanSig(sig SignalTypeEnum, whoIsSender WhoIsSenderTypeEnum, info string) ChanSig {
	//------------------------------------------------------------
	// alloc
	chanSig := new(ChanSig)
	//------------------------------------------------------------
	// init
	chanSig.Sig = sig
	chanSig.WhoIsSender = whoIsSender
	//------------------------------------------------------------
	// set local vars depending on who the sender is
	if whoIsSender == SENDER_IS_MACHINE {
		chanSig.SenderMachineKey = info
		chanSig.MsgFromSystem = ""
	} else {
		chanSig.SenderMachineKey = ""
		chanSig.MsgFromSystem = info
	}
	//------------------------------------------------------------
	// retun
	return *chanSig
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
