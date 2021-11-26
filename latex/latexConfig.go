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

package latex

type LatexConfig struct {
	Scalebox                 float32 // size of entire picture
	SlotHeight               float32 // hight of guard and action boxes
	SlotWidth                float32 // width of guard and action boxes
	WiringArrowLeftWidth     float32 // guard arrows length
	WiringArrowRightWidth    float32 // action arrows length
	WiringArrowServiceHeight float32 // height of in and out service arrows
	WiringArrowServiceSpace  float32 // horizontal distance between service arrows
	ContainerWidth           float32 // width of PIC and POC containers
	MinWiringWidth           float32 // minimal width of wiring box
}

////////////////////////////////////////
// constructor - sets defaults
////////////////////////////////////////

func NewLatexConfig() LatexConfig {
	conf := LatexConfig{}

	conf.Scalebox = 0.80
	conf.SlotHeight = 2.40
	conf.WiringArrowLeftWidth = 11.0
	conf.WiringArrowRightWidth = 14.5
	conf.WiringArrowServiceHeight = 8.0
	conf.WiringArrowServiceSpace = 4.0
	conf.SlotWidth = 1.0
	conf.ContainerWidth = 0.8
	conf.MinWiringWidth = 7.0

	return conf
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
