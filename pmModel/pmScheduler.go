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
	. "github.com/peermodel/simulator/config"
	. "github.com/peermodel/simulator/debug"
	. "github.com/peermodel/simulator/scheduler"
	"fmt"
)

// ----------------------------------------
// scheduler is informed that an entry was written (and committed) into a pic or poc
// -> i.e. it must update/insert tts/ttl slots for the entry
// returns the updated scheduler
func SetEttsAndEttlSlot(scheduler Scheduler, eid string, etts int, ettl int) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		// /**/ String2TraceFile(fmt.Sprintf("SetEttsAndEttlSlot for entry=%s, entry type=%s, t=%d\n", eid, e.GetType(), CLOCK))
		/**/
		String2TraceFile(fmt.Sprintf("SetEttsAndEttlSlot for entry=%s, t=%d\n", eid, CLOCK))
	}
	// -------------------
	// get entry times: convert infinite ttl to system ttl:
	// nb: tts cant be infinite
	tts := CheckInfiniteTts(etts)
	ttl := ConvertInfiniteTtl(ettl)
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("  tts=%d, ttl=%d\n", tts, ttl))
	}
	// -------------------
	// search for entry in the list:
	// update if tt* >= current time AND tt* <= SYSTEM_TTL, else remove the slot
	// do it in extra loops, because the list might be changed
	// tts:
	ttsFoundFlag := false
	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && ETTS == slot.UserSlot.(*PMSlot).Type && eid == slot.UserSlot.(*PMSlot).Eid {
			// found:
			ttsFoundFlag = true
			// update or remove the slot:
			if CLOCK <= tts && SYSTEM_TTL >= tts {
				slot.Time = tts
			} else {
				// remove
				schedulerTmp := append(scheduler[:i], scheduler[i+1:]...)
				scheduler = schedulerTmp
			}
			break
		}
	}
	// ttl:
	ttlFoundFlag := false
	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && ETTL == slot.UserSlot.(*PMSlot).Type && eid == slot.UserSlot.(*PMSlot).Eid {
			// found:
			ttlFoundFlag = true
			// update the slot:
			if CLOCK <= ttl && SYSTEM_TTL >= ttl {
				slot.Time = ttl
			} else {
				// remove
				scheduler = append(scheduler[:i], scheduler[i+1:]...)
			}
			break
		}
	}
	// -------------------
	// insert slots if not found & treated above and if tt* >= current time AND tt* <= SYSTEM_TTL:
	// tts:
	if !ttsFoundFlag && CLOCK <= tts && SYSTEM_TTL >= tts {
		scheduler = scheduler.SortedInsert(NewUserSlot(tts, NewEttsSlot(eid)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new tts slot inserted\n"))
		}
	}
	// ttl:
	if !ttlFoundFlag && CLOCK <= ttl && SYSTEM_TTL >= ttl {
		scheduler = scheduler.SortedInsert(NewUserSlot(tts, NewEttlSlot(eid)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new ttl slot inserted\n"))
		}
	}
	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed about new WTTS
// -> i.e. it must insert tts slot for the wiring
// returns the updated scheduler
func SetWttsSlot(scheduler Scheduler, time int, wid string) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("SetWttsSlot for wid=%s, time=%d, t=%d\n", wid, time, CLOCK))
	}
	// -------------------
	// insert slot if time >= current time AND time <= SYSTEM_TTL:
	if CLOCK <= time && SYSTEM_TTL >= time {
		scheduler = scheduler.SortedInsert(NewUserSlot(time, NewWttsSlot(wid)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new wtts slot with time=%d inserted\n", time))
		}
	}
	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed about new WTTL
// -> i.e. it must insert ttl slot for the wiring
// returns the updated scheduler
func SetWttlSlot(scheduler Scheduler, time int, wid string) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("SetWttlSlot for wid=%s, time=%d, t=%d\n", wid, time, CLOCK))
	}
	// -------------------
	// insert slot if time >= current time AND time <= SYSTEM_TTL:
	if CLOCK <= time && SYSTEM_TTL >= time {
		scheduler = scheduler.SortedInsert(NewUserSlot(time, NewWttlSlot(wid)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new wttl slot with time=%d inserted\n", time))
		}
	}
	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed about new LTTS
// -> i.e. it must insert tts slot for the link
// returns the updated scheduler
// nb: wid servers only for docu/trace
func SetLttsSlot(scheduler Scheduler, time int, wid string, wiid string, linkNo int) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("SetLttsSlot for wid=%d, wiid=%s, linkNo=%d, time=%d, t=%d\n", wid, wiid, linkNo, time, CLOCK))
	}
	// -------------------
	// insert slot if time >= current time AND time <= SYSTEM_TTL:
	if CLOCK <= time && SYSTEM_TTL >= time {
		scheduler = scheduler.SortedInsert(NewUserSlot(time, NewLttsSlot(wid, wiid, linkNo)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new ltts slot with time=%d inserted\n", time))
		}
	}
	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed about new LTTL
// -> i.e. it must insert ttl slot for the link
// returns the updated scheduler
// nb: wid servers only for docu/trace
func SetLttlSlot(scheduler Scheduler, time int, wid string, wiid string, linkNo int) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("SetLttlSlot for wid=%d, wiid=%s, linkNo=%d, time=%d, t=%d\n", wid, wiid, linkNo, time, CLOCK))
	}
	// -------------------
	// insert slot if time >= current time AND time <= SYSTEM_TTL:
	if CLOCK <= time && SYSTEM_TTL >= time {
		scheduler = scheduler.SortedInsert(NewUserSlot(time, NewLttlSlot(wid, wiid, linkNo)))
		if SCHEDULER_DETAILS_TRACE.DoTrace() {
			/**/ String2TraceFile(fmt.Sprintf("  new lttl slot with time=%d inserted\n", time))
		}
	}
	// -------------------
	// return changed scheduler
	return scheduler
}

//// ----------------------------------------
//// add a hunting slot to find outdated entries for one wiring to the scheduler to be executed at the given time
//// - unused
//func SetPeerEntriesHuntSlot(scheduler Scheduler, pid string, time int) Scheduler {
//	if SCHEDULER_DETAILS_TRACE.DoTrace() {
//		/**/ String2TraceFile(fmt.Sprintf("SetMegaHuntSlot time=%d, t=%d\n", time, CLOCK))
//	}
//	// -------------------
//	// insert slot if time >= current time AND time <= SYSTEM_TTL:
//	if CLOCK <= time && SYSTEM_TTL >= time {
//		scheduler = scheduler.SortedInsert(NewUserSlot(time, NewPeerEntriesHuntSlot(pid, time)))
//		if SCHEDULER_DETAILS_TRACE.DoTrace() {
//			/**/ String2TraceFile(fmt.Sprintf("  new wiring entries hunt slot with time=%d inserted\n", time))
//		}
//	}
//	return scheduler
//}

// ----------------------------------------
// scheduler is informed that an entry was deleted (and committed) in a pic or poc
// -> i.e. it must delete tts/ttl slots for the entry
// returns the updated scheduler
func ClearEttsAndEttlSlot(scheduler Scheduler, eid string) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		// /**/ String2TraceFile(fmt.Sprintf("ClearEttsAndEttlSlot for entry=%s, entry type=%s\n", e.Id, e.GetType()))
		/**/
		String2TraceFile(fmt.Sprintf("ClearEttsAndEttlSlot for entry=%s\n", eid))
	}

	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && ETTS == slot.UserSlot.(*PMSlot).Type && eid == slot.UserSlot.(*PMSlot).Eid {
			// found: remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && ETTL == slot.UserSlot.(*PMSlot).Type && eid == slot.UserSlot.(*PMSlot).Eid {
			// found: remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ scheduler.Println(TAB * 2)
	}

	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed that a wiring terminated
// -> i.e. it must delete tts/ttl slots for the wiring
// returns the updated scheduler
func ClearWttsAndWttlSlot(scheduler Scheduler, wid string) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("ClearWttsAndWttlSlot for wid=%s\n", wid))
	}

	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && WTTS == slot.UserSlot.(*PMSlot).Type && wid == slot.UserSlot.(*PMSlot).Wid {
			// found: remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && WTTL == slot.UserSlot.(*PMSlot).Type && wid == slot.UserSlot.(*PMSlot).Wid {
			// found:  remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ scheduler.Println(TAB * 2)
	}

	// -------------------
	// return changed scheduler
	return scheduler
}

// ----------------------------------------
// scheduler is informed that a link terminated
// -> i.e. it must delete tts/ttl slots for the link
// returns the updated scheduler
func ClearLttsAndLttlSlot(scheduler Scheduler, wiid string, linkNo int) Scheduler {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("ClearLttsAndLttlSlot for wiid=%s, linkNo=%d\n", wiid, linkNo))
	}

	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && LTTS == slot.UserSlot.(*PMSlot).Type && wiid == slot.UserSlot.(*PMSlot).Wiid && linkNo == slot.UserSlot.(*PMSlot).LinkNo {
			// found:  remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	for i, slot := range scheduler {
		if USER_SLOT == slot.Type && WTTL == slot.UserSlot.(*PMSlot).Type && wiid == slot.UserSlot.(*PMSlot).Wiid && linkNo == slot.UserSlot.(*PMSlot).LinkNo {
			// found:  remove:
			scheduler = append(scheduler[:i], scheduler[i+1:]...)
			break
		}
	}
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ scheduler.Println(TAB * 2)
	}

	// -------------------
	// return changed scheduler
	return scheduler
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
