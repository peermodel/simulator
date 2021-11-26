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
	. "github.com/peermodel/simulator/debug"
	"fmt"
)

type Locks struct {
	// Transaction IDs
	RLocks map[string]int
	WLocks map[string]int
	DLocks map[string]int
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
func NewLocks() Locks {
	return Locks{RLocks: map[string]int{}, WLocks: map[string]int{}, DLocks: map[string]int{}}
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (locks Locks) Copy() Locks {
	//------------------------------------------------------------
	// alloc
	newLocks := NewLocks()
	//------------------------------------------------------------
	// copy all fields:
	// - RLocks:
	for l, i := range locks.RLocks {
		newLocks.RLocks[l] = i
	}
	// - WLocks:
	for l, i := range locks.WLocks {
		newLocks.WLocks[l] = i
	}
	// - DLocks:
	for l, i := range locks.DLocks {
		newLocks.DLocks[l] = i
	}
	//------------------------------------------------------------
	// return
	return newLocks
}

// ----------------------------------------
func (locks *Locks) AddLock(spaceOp SpaceOpTypeEnum, txid string) {
	switch spaceOp {
	case READ:
		// check if lock exists; if not -> k = 0:
		k := locks.RLocks[txid]
		// increment lock counter:
		locks.RLocks[txid] = k + 1
	case WRITE:
		// check if lock exists; if not -> k = 0:
		k := locks.WLocks[txid]
		// increment lock counter:
		locks.WLocks[txid] = k + 1
	case DELETE:
		// check if lock exists; if not -> k = 0:
		k := locks.DLocks[txid]
		// increment lock counter:
		locks.DLocks[txid] = k + 1
	default:
		Panic(fmt.Sprintf("AddLock: unsupported Lock Type (= space op) %s", spaceOp))
	}
}

// ----------------------------------------
// returns true if there is
// - no WRITE lock (of another tx than me) and
// - no DELETE lock (of any tx)
func (locks *Locks) WriteLockedByOtherTxOrDeleteLocked(mytxid string) bool {
	for txid, k1 := range locks.WLocks {
		if k1 > 0 && txid != mytxid {
			return true
		}
	}
	for _, k2 := range locks.DLocks {
		if k2 > 0 {
			return true
		}
	}
	return false
}

// ----------------------------------------
// returns true if there is no READ lock of another tx than me:
func (locks *Locks) ReadLockedByOtherTx(myTxId string) bool {
	for txid, k := range locks.RLocks {
		if txid != myTxId && k > 0 {
			return true
		}
	}
	return false
}

// ----------------------------------------
// returns true if there is *any* lock set by ** transaction:
// @@@ must one really check k or is isn't it sufficient to check whether a lock is in the list?!
func (locks *Locks) Locked() bool {
	for _, k := range locks.RLocks {
		if k > 0 {
			return true
		}
	}
	for _, k := range locks.WLocks {
		if k > 0 {
			return true
		}
	}
	for _, k := range locks.DLocks {
		if k > 0 {
			return true
		}
	}
	return false
}

// ----------------------------------------
/*
	Remove a specific lock in the lock list of the given type.

	The lock type (lType) can be READ, WRITE, DELETE.
*/
func (locks *Locks) RemoveLock(spaceOp SpaceOpTypeEnum, txid string) {
	switch spaceOp {
	case READ:
		k := locks.RLocks[txid]
		if k > 1 {
			locks.RLocks[txid] = k - 1
		} else {
			delete(locks.RLocks, txid)
		}
	case WRITE:
		k := locks.WLocks[txid]
		if k > 1 {
			locks.WLocks[txid] = k - 1
		} else {
			delete(locks.WLocks, txid)
		}
	case DELETE:
		k := locks.DLocks[txid]
		if k > 1 {
			locks.DLocks[txid] = k - 1
		} else {
			delete(locks.DLocks, txid)
		}
	default:
		Panic(fmt.Sprintf("ill. lock type (= space op) = %s", spaceOp))
	}
}

// ----------------------------------------
/*
	This method removes all locks of the given lock type for the given txid.

	The lock type (lType) can be READ, WRITE, DELETE.
*/
func (locks *Locks) removeAllLocksOfGivenType(spaceOp SpaceOpTypeEnum, txid string) error {
	switch spaceOp {
	case READ:
		if 0 < locks.RLocks[txid] {
			delete(locks.RLocks, txid)
		}
	case WRITE:
		if 0 < locks.WLocks[txid] {
			delete(locks.WLocks, txid)
		}
	case DELETE:
		if 0 < locks.DLocks[txid] {
			delete(locks.DLocks, txid)
		}
	default:
		Panic(fmt.Sprintf("ill. lock type (= space op) = %s", spaceOp))
	}
	return nil
}

// ----------------------------------------
/*
	Clears all locks (e.g. on an entry) for given transaction id (txid).
*/
func (locks *Locks) RemoveAllLocks(txid string) {
	locks.removeAllLocksOfGivenType(READ, txid)
	locks.removeAllLocksOfGivenType(WRITE, txid)
	locks.removeAllLocksOfGivenType(DELETE, txid)
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (locks *Locks) IsEmpty() bool {
	if nil == locks ||
		(len(locks.WLocks) == 0 && len(locks.RLocks) == 0 && len(locks.DLocks) == 0) {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (locks *Locks) ToString(tab int) string {
	s := NBlanksToString("", tab)

	// s = fmt.Sprintf("%s{", s)

	sep := ""
	// @@@ sollte eleganter gehen:
	if 0 < len(locks.RLocks) {
		s = fmt.Sprintf("%s%sRlocks={", s, sep)
		sep = ""
		for lockOp, count := range locks.RLocks {
			s = fmt.Sprintf("%s%s%s:%d", s, sep, lockOp, count)
			sep = ", "
		}
		s = fmt.Sprintf("%s}", s)
	}
	if 0 < len(locks.DLocks) {
		s = fmt.Sprintf("%s%sDlocks={", s, sep)
		sep = ""
		for lockOp, count := range locks.DLocks {
			s = fmt.Sprintf("%s%s%s:%d", s, sep, lockOp, count)
			sep = ", "
		}
		s = fmt.Sprintf("%s}", s)
	}
	if 0 < len(locks.WLocks) {
		s = fmt.Sprintf("%s%sWLocks={", s, sep)
		sep = ""
		for lockOp, count := range locks.WLocks {
			s = fmt.Sprintf("%s%s%s:%d", s, sep, lockOp, count)
		}
		sep = ", "
		s = fmt.Sprintf("%s}", s)
	}
	// s = fmt.Sprintf("%s}", s, )
	return s
}

// ----------------------------------------
func (locks *Locks) Print(tab int) {
	/**/ String2TraceFile(locks.ToString(tab))
}

// ----------------------------------------
func (locks *Locks) Println(tab int) {
	/**/ locks.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
