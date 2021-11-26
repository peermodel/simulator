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
	. "cca/config"
	. "cca/debug"
	. "cca/scheduler"
	"fmt"
	"strings"
)

type Container struct {
	Id string
	Entries
	// for eventing: implicitly set to 0 (i.e.< start value of CLOCK)
	// caution: use event time here (EVENT_CLOCK)
	LastUpdateEventTime int
}

////////////////////////////////////////
// constructor
////////////////////////////////////////

// ----------------------------------------
func NewContainer(cid string) *Container {
	c := new(Container)

	c.Id = cid
	// set LastUpdateEventTime to time *before* the event time starts!
	c.LastUpdateEventTime = -1

	return c
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (c *Container) Copy() *Container {
	//------------------------------------------------------------
	// alloc
	newC := NewContainer(c.Id)
	//------------------------------------------------------------
	// copy all fields:
	// - Id: has been copied above already
	// - Entries:
	newC.Entries = c.Entries.Copy()
	// - LastUpdateEventTime:
	newC.LastUpdateEventTime = c.LastUpdateEventTime
	//------------------------------------------------------------
	// return
	return newC
}

// ----------------------------------------
/*
	Adds the given entry to the container. (CAPI1)

	Must only be used for non-committed write, ie the space write machines; not by user in use cases!
*/
func (c *Container) AddEntry(e Entry) {
	c.Entries = append(c.Entries, e)
}

// ----------------------------------------
func (c *Container) AddEntryPtr(e *Entry) {
	c.Entries = append(c.Entries, *e)
}

// ----------------------------------------
/*
	Returns index to the next entry that fulfills selector; in any order;
	NB: count checking is done by caller;
	Returns -1 if not found;

	if eType == WILDCARD (= "*") -> wildcard!
*/
// @@@ coordinator: not yet
func (c *Container) SelectEntryIndex(vars Vars, eType string, selector *Arg) int {
	// iterate over all entries e in the container
	for i, e := range c.Entries {
		// check entry type of e
		if WILDCARD == eType || e.GetType() == eType {
			// either empty selector, or apply selector to entry e
			if selector.Apply(vars, &e) {
				// entry fulfills selector
				return i
			}
		}
	}
	// no suitable entry found
	return -1
}

// ----------------------------------------
/*
	Returns pointer to the next entry that fulfills selector; in any order. Shared!!

	if eType == WILDCARD (= "*") -> wildcard!
*/
// @@@ coordinator: not yet
func (c *Container) GetPtrToNextEntry(vars Vars, eType string, selector *Arg) *Entry {
	entryIndex := c.SelectEntryIndex(vars, eType, selector)
	// /**/ String2TraceFile(fmt.Sprintf("etype ", eType))
	// /**/ selector.Print()
	// /**/ String2TraceFile(fmt.Sprintf("  -> entryIndex = %d", entryIndex))
	if entryIndex == -1 {
		return nil
	}
	return &c.Entries[entryIndex]
}

// ----------------------------------------
// like GetPtrToNextEntry, but returns a copy of the next entry that fulfills the query
func (c *Container) SelectEntry(vars Vars, eType string, selector *Arg) *Entry {
	e := c.GetPtrToNextEntry(vars, eType, selector)
	if nil != e {
		retEntry := e.Copy()
		return retEntry
	} else {
		return nil
	}
}

// ----------------------------------------
// caller must assure that entry exists in entry collection
// returns pointer to removed entry (if exists) - otherwise nil
// func (c *Container) RemoveEntry(eid string) error {
func (c *Container) RemoveEntry(eid string) *Entry {
	es, e := RemoveEntryFromEntries(c.Entries, eid)
	c.Entries = es
	return e
}

// ----------------------------------------
func (c *Container) RemoveAllEntryLocks(txid string) {
	for _, e := range c.Entries {
		e.RemoveAllLocks(txid)
	}
}

// ----------------------------------------
// see below functions
func (c *Container) IsPic() bool {
	return CidIsPic(c.Id)
}
func (c *Container) IsPoc() bool {
	return CidIsPoc(c.Id)
}
func (c *Container) IsPicOrPoc() bool {
	return CidIsPicOrPoc(c.Id)
}

////////////////////////////////////////
// functions
////////////////////////////////////////

// ----------------------------------------
// test if container is a PIC or a POC
// @@@ is not very good, as it forbids users to have "PIC" or "POC" as substrings in their container names
func CidIsPic(cid string) bool {
	if -1 == strings.Index(cid, PIC) {
		return false
	}
	return true
}
func CidIsPoc(cid string) bool {
	if -1 == strings.Index(cid, POC) {
		return false
	}
	return true
}
func CidIsPicOrPoc(cid string) bool {
	return CidIsPic(cid) || CidIsPoc(cid)
}

// ----------------------------------------
// same as RemoveEntryP above, but es is a set of entries
// returns possibly changed entries and (the found&removed entry or nil)
// @@@ hilfs function für container; gehoert nicht zu container sondern zu entries?!
// @@@ sollte private sein
func RemoveEntryFromEntries(es Entries, eid string) (Entries, *Entry) {
	entryIndex := -1
	var foundE *Entry = nil

	for i, e := range es {
		if e.Id == eid {
			entryIndex = i
			foundE = &e
			break
		}
	}
	if -1 == entryIndex {
		return es, foundE
	}
	// remove entry
	i := entryIndex
	es = append(es[:i], es[i+1:]...)
	return es, foundE
}

////////////////////////////////////////
// signalling
////////////////////////////////////////

// ----------------------------------------
// signal container change event
// @@@ only on real containers needed
// @@@ could be more detailed: what operation/query can be waken up? read, take? none?
func ContainerPtrChangeEvent(c *Container) {
	c.LastUpdateEventTime = EVENT_CLOCK
	if EVENT_CONDITION_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("%s ContainerChangeEvent: et=%d, t=%d \n", c.Id, EVENT_CLOCK, CLOCK))
	}
	// TBD: improve / use interface!!!!
	SPACE_UPDATE_COUNT_SINCE_LAST_CHOICE_POINT++
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (c *Container) IsEmpty() bool {
	if nil == c || c.Id == "" {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
func (c *Container) ToString(tab int) string {
	s := NBlanksToString("", tab)
	s = fmt.Sprintf("%s%s updateEvtTime=%d = {", s, c.Id, c.LastUpdateEventTime)
	sep := "\n"

	// --------
	// sorted print -- sorted by entry type:
	// @@@ a bit inefficient?!
	doneType := ""
	for {
		nextType := ""
		// get next alphabetically next type after doneType
		for _, e1 := range c.Entries {
			tryType := e1.GetType()
			// 1.) doneType is either not yet set, or tryType must be greater than doneType
			if doneType == "" || strings.Compare(tryType, doneType) > 0 {
				// AND
				// 2.) nextType is either not yet set, or tryType is less than nextType
				if nextType == "" || strings.Compare(tryType, nextType) < 0 {
					nextType = tryType
				}
			}
		}
		// finished
		if nextType == "" {
			break
		} else {
			// set done type to next type
			doneType = nextType
		}
		// nextType is the least type name -- there could be multiples of it in the slice
		// print all entries with that type
		for _, e2 := range c.Entries {
			if e2.GetType() == nextType {
				s = fmt.Sprintf("%s%s", s, sep)
				s = fmt.Sprintf("%s%s", s, e2.ToString(tab+TAB))
				sep = ", \n"
			}
		}
	}

	//	// --------
	//	// unsorted print:
	//	for _, e := range c.Entries {
	//		s = fmt.Sprintf("%s%s", s, sep)
	//		s = fmt.Sprintf("%s%s", s, e.ToString(tab+TAB))
	//		sep = ", \n"
	//	}

	// --------
	s = fmt.Sprintf("%s}", s)
	return s
}

// ----------------------------------------
func (c *Container) Print(tab int) {
	/**/ String2TraceFile(c.ToString(tab))
}

// ----------------------------------------
func (c *Container) Println(tab int) {
	/**/ c.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
