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
// Date: 2015, 2016
//------------------------------------------------------------
// controller channel type
// - transports signals for the controller
//------------------------------------------------------------
// Code Review: 2021 Apr, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package controller

//------------------------------------------------------------
// controller channel type
type ControllerChannel chan ChanSig

//------------------------------------------------------------
// // stop the system
// // - unused
//func (controllerChan ControllerChannel) SendStop() {
//	// does not need any entry: just stop the system
//	controllerChan <- STOP
//}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
