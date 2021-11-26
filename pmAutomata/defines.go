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

package pmAutomata

////////////////////////////////////////
// ENUMs
////////////////////////////////////////

////////////////////////////////////////
// automata ids
////////////////////////////////////////

type AutomatonID int

const (
	PCC_READ_AUTOMATON AutomatonID = iota
	PCC_TX_COMMIT_AUTOMATON
	READ_AUTOMATON
	SERVICE_AUTOMATON
	SPACE_CREATE_TX_AUTOMATON
	SPACE_READ_AUTOMATON
	SPACE_TX_COMMIT_AUTOMATON
	SPACE_UNDO_AUTOMATON
	SPACE_UNDO_READ_AUTOMATON
	SPACE_UNDO_WRITE_AUTOMATON
	SPACE_WRITE_AUTOMATON
	WIRING_AUTOMATON
)

func (id AutomatonID) String() string {
	switch id {
	case PCC_READ_AUTOMATON:
		return "PCC_READ_AUTOMATON"
	case PCC_TX_COMMIT_AUTOMATON:
		return "PCC_TX_COMMIT_AUTOMATON"
	case READ_AUTOMATON:
		return "READ_AUTOMATON"
	case SERVICE_AUTOMATON:
		return "SERVICE_AUTOMATON"
	case SPACE_CREATE_TX_AUTOMATON:
		return "SPACE_CREATE_TX_AUTOMATON"
	case SPACE_READ_AUTOMATON:
		return "SPACE_READ_AUTOMATON"
	case SPACE_TX_COMMIT_AUTOMATON:
		return "SPACE_TX_COMMIT_AUTOMATON"
	case SPACE_UNDO_AUTOMATON:
		return "SPACE_UNDO_AUTOMATON"
	case SPACE_UNDO_READ_AUTOMATON:
		return "SPACE_UNDO_READ_AUTOMATON"
	case SPACE_UNDO_WRITE_AUTOMATON:
		return "SPACE_UNDO_WRITE_AUTOMATON"
	case SPACE_WRITE_AUTOMATON:
		return "SPACE_WRITE_AUTOMATON"
	case WIRING_AUTOMATON:
		return "WIRING_AUTOMATON"
	default:
		return "ill. automaton id"
	}
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
