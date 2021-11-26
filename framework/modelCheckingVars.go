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
//////////////////////////////////////////////////////////////
// System: Peer Model State Machine
// Author: Eva Maria Kuehn
// Date:   2021
//------------------------------------------------------------
// for model checking mode:
// - global variables: shared between runtime and controller
//------------------------------------------------------------
// Code Review: 2021 Mai, Eva Maria Kuehn
//////////////////////////////////////////////////////////////

package framework

//	. "github.com/peermodel/simulator/config"

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// global variables: shared between runtime and controller
type ModelCheckingVars struct {
	//------------------------------------------------------------
	// all choice points for the model checker (= MC)
	// - new choice points are added at the end
	ChoicePoints ChoicePoints
	//------------------------------------------------------------
	// debug: depth of the cur choice point
	//  - we start with 0 at root
	CurPathDepth int
	//------------------------------------------------------------
	// current CHOICE POINT
	CurChoicePoint *ChoicePoint
	//------------------------------------------------------------
	// current CHOICE (ie a machine key)
	CurChoice string
	//------------------------------------------------------------
	// flag indicating that the next run starts at a recovered CP
	// - ie the next machine to be executed is the current choice taken from this CP
	UseCurChoiceAsNextMachineFlag bool
	//------------------------------------------------------------
	// debug:
	// - counter for CP ids
	// - nb: CPs are numbered with 1, 2, 3, ...
	ChoicePointUuid int
}

//------------------------------------------------------------
var MC_VARS ModelCheckingVars

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// init global model checking variables shared with Controller
func InitModelCheckingVars() {
	MC_VARS.ChoicePoints = *NewChoicePoints()
	MC_VARS.CurPathDepth = 1
	MC_VARS.CurChoicePoint = NewChoicePoint()
	MC_VARS.CurChoice = ""
	MC_VARS.UseCurChoiceAsNextMachineFlag = false
	MC_VARS.ChoicePointUuid = 1
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
