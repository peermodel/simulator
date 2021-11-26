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
	"fmt"
)

var UUID int

// specialized peers and containers:
const IOP_PEER string = "IOP"
const IOP_PIC string = "IOP_PIC"
const IOP_POC string = "IOP_POC"

const STOP_PEER string = "Stop"

// count or max values: must be < 0:

const ALL int = -100
const NONE int = -200

// arg kind:
const VAR string = "VAR"
const VAL string = "VAL"
const LABEL string = "LABEL"
const EXPR string = "EXPR"
const FU string = "FU"
const DYN_ARRAY_REF string = "DYN_ARRAY_REF"
const TYPED_ARRAY_LABEL string = "TYPED_ARRAY_LABEL"
const TYPED_ARRAY_VAL string = "TYPED_ARRAY_VAL"

// prop type: system defined property labels
const COMMIT string = "commit"
const DEST string = "dest"
const FID string = "fid"
const FLOW string = "flow"
const MANDATORY string = "mandatory"
const MAX_THREADS string = "max_threads"
const ON_ABORT string = "on_abort"
const REPEAT_COUNT string = "repeat_count"
const TTS string = "tts"
const TTL string = "ttl"
const TXCC string = "txcc"
const TYPE string = "type"
const SOURCE string = "source"

// reserved entry types:
const DEST_WRAP string = "DEST_WRAP"
const SOURCE_WRAP string = "SOURCE_WRAP"
const EXCEPTION_WRAP string = "EXCEPTION_WRAP"
const EXCEPTION_ON_ABORT string = "EXCEPTION_ON_ABORT"
const WILDCARD string = "*"

// container name & construction type:
const PIC string = "PIC"
const POC string = "POC"
const WC string = "WC"
const SIC string = "SIC"
const SOC string = "SOC"
const SEP string = "_"

// txcc type:
const PCC string = "pcc"
const OCC string = "occ"

// repeat interval for entry hunting for each peer:
const PEER_ENTRIES_HUNT_REPEAT_INTERVAL int = 1

////////////////////////////////////////
// ENUMs
////////////////////////////////////////

////////////////////////////////////////
// slot type
////////////////////////////////////////

type PMSlotTypeEnum int

// tts and ttl for entry, link and wiring
// WIRING_ENTRY_HUNT:
//  - "mega hunting" for all possibly out dated entries of one wiring
//  - very expensive; but problem is that entries might not be found when
//    they are ripe because they are locked at that moment or in a WIC
//  - so do the exhaustive search from time to time (controlled via a scheduler slot)

const (
	ETTS PMSlotTypeEnum = iota
	ETTL
	LTTS
	LTTL
	WIRING_ENTRIES_HUNT
	WTTS
	WTTL
)

func (t PMSlotTypeEnum) String() string {
	switch t {
	case ETTS:
		return "ETTS"
	case ETTL:
		return "ETTL"
	case LTTS:
		return "LTTS"
	case LTTL:
		return "LTTL"
	case WIRING_ENTRIES_HUNT:
		return "WIRING_ENTRIES_HUNT"
	case WTTS:
		return "WTTS"
	case WTTL:
		return "WTTL"
	default:
		return fmt.Sprintf("ill. pm slot type = %s", t)
	}
}

////////////////////////////////////////
// event type
////////////////////////////////////////

////////////////////////////////////////
// link type
////////////////////////////////////////

type LinkTypeEnum int

const (
	GUARD LinkTypeEnum = iota
	ACTION
	SERVICE_IN
	SERVICE_OUT
	SERVICE
)

func (t LinkTypeEnum) String() string {
	switch t {
	case GUARD:
		return "guard"
	case ACTION:
		return "action"
	case SERVICE_IN:
		return "sin"
	case SERVICE_OUT:
		return "sout"
	case SERVICE:
		return "service"
	default:
		return "ill. link type"
	}
}

////////////////////////////////////////
// exception types
////////////////////////////////////////

type ExceptionTypeEnum int

const (
	LINK_TTL_EXCEPTION ExceptionTypeEnum = iota
	WIRING_REPEAT_EXCEPTION
	WIRING_TTL_EXCEPTION
	SYSTEM_STOP
	WIRING_STOP
)

// try to keep names ca. same size (<= 13) -> is padded with that number
func (t ExceptionTypeEnum) String() string {
	switch t {
	case LINK_TTL_EXCEPTION:
		return "LINK-TTL"
	case SYSTEM_STOP:
		return "SYSTEM-STOP"
	case WIRING_REPEAT_EXCEPTION:
		return "WIRING-REPEAT"
	case WIRING_TTL_EXCEPTION:
		return "WIRING-TTL"
	case WIRING_STOP:
		return "WIRING-STOP"
	default:
		return "ill. exception type"
	}
}

////////////////////////////////////////
// arg kind
////////////////////////////////////////

//type ArgKindEnum int

//const (
//	VAL ArgKindEnum = iota
//	LABEL
//	VAR
//	EXPR
//)

//func (t ArgKindEnum) String() string {
//	switch t {
//	case VAL:
//		return "VAL"
//	case LABEL:
//		return "LABEL"
//	case VAR:
//		return "VAR"
//	case EXPR:
//		return "EXPR"
//	default:
//		return "ill. arg type"
//	}
//}

////////////////////////////////////////
// data type
////////////////////////////////////////

type DataTypeEnum int

const (
	INT DataTypeEnum = iota
	STRING
	BOOL
)

func (t DataTypeEnum) String() string {
	switch t {
	case INT:
		return "INT"
	case STRING:
		return "STRING"
	case BOOL:
		return "BOOL"
	default:
		return "ill. data type"
	}
}

////////////////////////////////////////
// what kind of string is it?
////////////////////////////////////////

type StringSubTypeEnum int

const (
	NORMAL StringSubTypeEnum = iota
	ENTRY_TYPE
	URL
)

func (t StringSubTypeEnum) String() string {
	switch t {
	case NORMAL:
		return "NORMAL"
	case ENTRY_TYPE:
		return "ENTRY_TYPE"
	case URL:
		return "URL"
	default:
		return "ill. data type"
	}
}

////////////////////////////////////////
// space op type
////////////////////////////////////////

type SpaceOpTypeEnum int

const (
	CALL SpaceOpTypeEnum = iota
	CREATE
	DELETE
	NOOP
	READ
	TAKE
	TEST
	WRITE
)

func (t SpaceOpTypeEnum) String() string {
	switch t {
	case CALL:
		return "call"
	case CREATE:
		return "create"
	case DELETE:
		return "delete"
	case NOOP:
		return "noop"
	case READ:
		return "read"
	case TAKE:
		return "take"
	case TEST:
		return "test"
	case WRITE:
		return "write"
	default:
		return "ill. space op type"
	}
}

////////////////////////////////////////
// op type
////////////////////////////////////////

type OpTypeEnum int

const (
	// binary relational operators:
	EQUAL OpTypeEnum = iota
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	NOT_EQUAL
	// binary arithmetic operators:
	ADD
	SUB
	MUL
	DIV
	MOD
	// unary artithmetic operators:
	PLUS
	MINUS
	// binaryboolean operators:
	AND
	OR
	// unary boolean operators:
	NOT
	// binary string operators:
	CONCAT
	// unused
	UNUSED
)

func (t OpTypeEnum) String() string {
	switch t {
	// binary relational operators:
	case EQUAL:
		return "=="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case NOT_EQUAL:
		return "!="
	// binary arithmetic operators:
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case MOD:
		return "%"
	// unary artithmetic operators:
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	// binary boolean operators:
	// caution: do not use &&, ||, !, because then the latex output does not work any more.... @@@ (using $...$)
	// blanks are added
	case AND:
		return " AND "
	case OR:
		return " OR "
	// unary boolean operators:
	case NOT:
		return " NOT "
	// binary string operators:
	case CONCAT:
		return " CONCAT "
	// // binary string operators:
	// case INDEX:
	// 	return " INDEX "
	// default
	default:
		return "ill. op type"
	}
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
