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

// @@@ IPrint interface implementation is missing

package pmModel

import (
	. "cca/config"
	. "cca/debug"
	. "cca/eventInterface"
	. "cca/helpers"
	. "cca/latex"
	. "cca/scheduler"
	. "cca/slotInterface"
	"fmt"
	"os"
	"strings"
)

//------------------------------------------------------------
type PeerSpace struct {
	//------------------------------------------------------------
	// key = pid
	Peers map[string]*Peer
	//------------------------------------------------------------
	// key = cid
	Containers map[string]*Container
	//------------------------------------------------------------
	// for debug only: is needed to be able to print containers always in the same order...
	PeerPids      Strings
	ContainerCids Strings
}

////////////////////////////////////////
// functions
////////////////////////////////////////

func NewPeerSpace() *PeerSpace {
	ps := new(PeerSpace)
	ps.Peers = make(map[string]*Peer)
	ps.Containers = make(map[string]*Container)
	return ps
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (ps *PeerSpace) Copy() *PeerSpace {
	//------------------------------------------------------------
	// alloc
	newPS := NewPeerSpace()
	//------------------------------------------------------------
	// copy dynamic fields:
	// - Peers:
	for pid, p := range ps.Peers {
		newPS.Peers[pid] = p.Copy()
	}
	//------------------------------------------------------------
	// - Containers:
	for cid, c := range ps.Containers {
		newPS.Containers[cid] = c.Copy()
	}
	//------------------------------------------------------------
	// - PeerPids:
	newPS.PeerPids = ps.PeerPids.Copy()
	//------------------------------------------------------------
	// - ContainerCids:
	newPS.ContainerCids = ps.ContainerCids.Copy()
	//------------------------------------------------------------
	// return
	return newPS
}

// ----------------------------------------
func (ps *PeerSpace) AddPeer(p *Peer) {
	ps.Peers[p.Id] = p
	// for debug only:
	ps.PeerPids = ps.PeerPids.SortedInsertString(p.Id)
}

// ----------------------------------------
func (ps *PeerSpace) AddContainer(c *Container) {
	ps.Containers[c.Id] = c
	// for debug only:
	ps.ContainerCids = ps.ContainerCids.SortedInsertString(c.Id)
}

// =========================================================
// user space API
// =========================================================

// ----------------------------------------
// read one entry:
func (ps *PeerSpace) Read(cid string, typ string, sel *Arg, vars Vars) *Entry {
	c := ps.Containers[cid]
	if nil == c {
		UserError(fmt.Sprintf("Read: ill. cid=%s\n", cid))
	}
	// read:
	return c.SelectEntry(vars, typ, sel)
}

// ----------------------------------------
// take one entry:
func (ps *PeerSpace) Take(cid string, typ string, sel *Arg, vars Vars) *Entry {
	c := ps.Containers[cid]
	if nil == c {
		UserError(fmt.Sprintf("Take: ill. cid=%s\n", cid))
	}
	// read:
	e := c.SelectEntry(vars, typ, sel)
	if nil != e {
		// remove:
		ps.Containers[cid].RemoveEntry(e.Id)
	}
	return e
}

// ----------------------------------------
// internal helper method for Write and Emit
// "committed" write of one entry into a container
// eval entry propoerties
// raise container change event
// trigger a new tts/ttl slot in the scheduler for emitted entry
// @@@ TBD: create new eid for each entry that is written!?
func (ps *PeerSpace) doWrite(cid string, e *Entry, vars Vars, raiseContainerChangeEventFlag bool, informSchedulerFlag bool, scheduler *Scheduler) {
	// ----------
	// find container:
	c := ps.Containers[cid]
	if nil == c || "" == c.Id {
		UserError(fmt.Sprintf("DoWrite: ill. cid=%s\n", cid))
	}
	// ----------
	// eval entry properties that are not yet basic values against vars, now:
	for label, eprop := range e.EProps {
		if VAL != eprop.Kind {
			// nb: eval changes eprop:
			if !eprop.Eval(vars, nil /* entry */) {
				UserError(fmt.Sprintf("DoWrite: can't eval entry property with label=%s", label))
			}
			// nb: if ok evaluated, its value must be set (albeit its kind is kept)
			// overwrite e's property with the evaluated value
			switch eprop.Type {
			case INT:
				e.SetIntVal(label, eprop.IntVal)
			case STRING:
				e.SetStringVal(label, eprop.StringVal)
			case BOOL:
				e.SetBoolVal(label, eprop.BoolVal)
			default:
				SystemError("ill. eprop type")
			}
		}
	}
	// ----------
	// write:
	c.Entries = append(c.Entries, *e)
	// ----------
	// raise container change event:
	if raiseContainerChangeEventFlag {
		ContainerPtrChangeEvent(c)
	}
	// ----------
	// inform scheduler about (new) entry - scheduler must insert TTS and TTL slots for it:
	// @@@ e.GetId()
	if informSchedulerFlag {
		*scheduler = SetEttsAndEttlSlot(*scheduler, e.Id, e.GetTts(), e.GetTtl())
	}
}

// ----------------------------------------
// committed write one entry into a pic or poc
// eval entry propoerties
// raise container change event
// trigger a new tts/ttl slot in the scheduler for emitted entry
func (ps *PeerSpace) Write(cid string, e *Entry, vars Vars, scheduler *Scheduler) {
	ps.doWrite(cid, e, vars, true /* raiseContainerChangeEventFlag */, true /* informSchedulerFlag */, scheduler)
}

// ----------------------------------------
// emit one entry: used to write into SOUT from a service
// eval entry proporties
// nb: does not raise cpontainer change event
//     does not trigger a new tts/ttl slot in the scheduler for emitted entry
func (ps *PeerSpace) Emit(cid string, e *Entry, vars Vars, scheduler *Scheduler) {
	ps.doWrite(cid, e, vars, false /* raiseContainerChangeEventFlag */, false /* informSchedulerFlag */, scheduler)
}

////////////////////////////////////////
// methods: add meta models for built-in peer types
////////////////////////////////////////

////////////////////////////////////////
// IOP
////////////////////////////////////////

func (ps *PeerSpace) AddMetaModel_IOP_PEER(wprops WProps) {
	var lprops LProps

	// --------------------------------------------------
	// link properties:
	arg := wprops[TTL]
	if "" != arg.Kind {
		lprops = LProps{TTL: IVal(wprops[TTL].IntVal)}
	} else {
		lprops = LProps{}
	}

	// --------------------------------------------------
	// service wrappers:
	swSendService := NewServiceWrapper(SendService, "SendService")

	// --------------------------------------------------
	// Peer IOP:
	p := NewPeer("IOP")

	// is a system peer
	p.IsSysPeerFlag = true

	// Wiring W1:
	p_w1 := NewWiring("W1")

	p_w1.AddServiceWrapper("S1", swSendService)

	p_w1.AddGuard("", PIC, TAKE, Query{Typ: SEtype(WILDCARD), Count: IVal(1)}, lprops, EProps{}, Vars{})
	p_w1.AddSin(TAKE, Query{Typ: SEtype(WILDCARD), Count: IVal(1)}, "S1", lprops, EProps{}, Vars{})
	p_w1.AddScall("S1", LProps{TTL: wprops[TTL], COMMIT: BVal(true)}, EProps{}, Vars{})
	p_w1.WProps = wprops

	// add wirings to peer & resolve names:
	p.AddWiring(p_w1)

	// --------------------------------------------------
	// add peers to peer space:
	ps.AddPeer(p)
}

////////////////////////////////////////
// Stop
////////////////////////////////////////

func (ps *PeerSpace) AddMetaModel_STOP_PEER(wprops WProps) {
	// --------------------------------------------------
	// service wrappers:
	sw := NewServiceWrapper(StopService, "StopService")

	// --------------------------------------------------
	// Stop Peer:
	p := NewPeer(STOP_PEER)

	// is a system peer
	p.IsSysPeerFlag = true

	// Wiring W1:
	// ----------
	p_w1 := NewWiring("W1")

	p_w1.AddServiceWrapper("S1", sw)

	// - G1: do not commit, becaus otherwies if SIN link times out the stop entry is lost and repetition does not find it any more!
	p_w1.AddGuard("", PIC, DELETE, Query{Typ: SEtype("STOP"), Count: IVal(1)}, LProps{TTL: IVal(SYSTEM_TTL)}, EProps{}, Vars{})
	// - SIN1:
	p_w1.AddScall("S1", LProps{TTL: IVal(SYSTEM_TTL), COMMIT: BVal(true)}, EProps{}, Vars{})

	// wprops
	p_w1.WProps = wprops

	// add wirings to peer & resolve names:
	p.AddWiring(p_w1)

	// --------------------------------------------------
	// add peers to peer space:
	ps.AddPeer(p)
}

// =========================================================
// search for entry with eid in all peer's PICs or POCs in order to check whether it has expired:
// if its ttl has expired AND if it is not locked:
//  - remove it from container
//  - find out the corresponding poc container (could be the same as c)
//  - wrap entry into exception and write the exception wrap to this poc
// if entry is found but locked -> the wtx that currently locks it must care for entry's ttl treatment
// @@@ not very efficient
// caution: system ttl could already be overridden -> so check also whether ttl is infinite
// TBD: if entry is locked... später nochmal probieren!!!
func (ps *PeerSpace) EntryHunter(eid string, scheduler *Scheduler) {
	for _, p := range ps.Peers {
		picC := ps.Containers[p.Pic]
		pocC := ps.Containers[p.Poc]
		// assertion:
		if nil == picC || nil == pocC {
			Panic("ill. pic or poc")
		}
		// search in pic:
		for _, e := range picC.Entries {
			if eid == e.Id && e.GetTtl() != INFINITE && CLOCK > ConvertInfiniteTtl(e.GetTtl()) && !e.Locked() {
				// entry found and not locked
				RaiseExceptionForOutdatedEntry(ps, eid, picC, pocC.Id, scheduler)
				return
			}
		}
		// search in poc:
		for _, e := range pocC.Entries {
			if eid == e.Id && e.GetTtl() != INFINITE && CLOCK > ConvertInfiniteTtl(e.GetTtl()) && !e.Locked() {
				// entry found and not locked
				RaiseExceptionForOutdatedEntry(ps, eid, pocC, pocC.Id, scheduler)
				return
			}
		}
	}
}

//// =========================================================
//// @@@ not needed
//// check all entries of given peer whether they are outdated
//// if the entry's ttl has expired AND if it is not locked:
////  - remove it from container
////  - find out the corresponding poc container (could be the same as c)
////  - wrap entry into exception and write the exception wrap to this poc
//// if entry is found but locked -> the wtx that currently locks it must care for entry's ttl treatment
//// @@@ not very efficient
//// rewind the entries hunter using the same time interval // @@@ improve: add extra interval
//// if peer does not exist -> stop the hunting
//func (ps *PeerSpace) PeerEntriesHunter(pid string, time int, scheduler *Scheduler) {
//	/**/ SystemInfo(fmt.Sprintf("@@@DEBUG: PeerEntriesHunter called for pid = %s", pid))

//	p := ps.Peers[pid]
//	if nil == p {
//		return
//	}
//	picC := ps.Containers[p.Pic]
//	pocC := ps.Containers[p.Poc]
//	// assertion:
//	if nil == picC || nil == pocC {
//		Panic("ill. pic or poc")
//	}

//	// check all entries of pic in pic:
//	for _, e := range picC.Entries {
//		if CLOCK > ConvertInfiniteTtl(e.GetTtl()) && !e.Locked() {
//			// an outdated entry found that is not locked
//			RaiseExceptionForOutdatedEntry(ps, e.Id, picC, pocC.Id, scheduler)
//		}
//	}
//	// search in poc:
//	for _, e := range pocC.Entries {
//		if CLOCK > ConvertInfiniteTtl(e.GetTtl()) && !e.Locked() {
//			// an outdated entry found and not locked
//			RaiseExceptionForOutdatedEntry(ps, e.Id, pocC, pocC.Id, scheduler)
//		}
//	}

//	// rewind the slot
//	// @@@ ? does not work?!
//	*scheduler = SetPeerEntriesHuntSlot(*scheduler, pid, time)
//}

// =========================================================
// @@@ not needed
// if found: raise the entry exception for e:
// remove it from foundC and write the exception into pocCid
func RaiseExceptionForOutdatedEntry(ps *PeerSpace, eid string, foundC *Container, pocCid string, scheduler *Scheduler) {
	// remove entry from container
	e := foundC.RemoveEntry(eid)

	// /**/ SystemInfo(fmt.Sprintf("@@@DEBUG: raise exception for outdated entry = %s", e.ToString(0)))
	if SCHEDULER_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("  raise exception for outdated entry = %s\n", e.ToString(0)))
	}

	// wrap entry into exception
	excE := e.ExceptionWrap("@@@DUMP-STELLE-3")

	// write exception to peer's poc
	ps.Write(pocCid, excE, nil /* @@@ no vars */, scheduler)
	if SCHEDULER_TRACE.DoTrace() {
		/**/ ps.Containers[pocCid].Println(TAB)
		/**/ scheduler.Println(0)
	}
	return
}

// =========================================================
// check plausibility of the meta model:
// - there must be at least one guard link
// -- tbd: subconditions @@@ zb nur lesen alleine reicht nicht, wenn zb kein
// --      take aus dem selben container erfolgt etc.
// --      oder: 1. guard must ein take sein (tbd)
// - last link must have a commit
// - ttl on optional link must be 0
// - ttl on noop link must be 0
// testCaseName serves only for docu
// raises user error
// @@@ TBD:
// - wiring tts must not be infinite
func (ps *PeerSpace) MetaModelCheck(testCaseName string) {
	// prefix for user error:
	preString := fmt.Sprintf("%s: Meta Model Check: ", testCaseName)
	// check all peers:
	for _, p := range ps.Peers {
		// check all wirings:
		for wid, w := range p.Wirings {
			// is there at least one guard link?
			// @@@ tbd: subconditions
			firstGuardIndex := -1
			for i, l := range w.Links {
				if -1 == firstGuardIndex && GUARD == l.Type {
					firstGuardIndex = i
				}

				// check if ttl on optional link is 0:
				if !l.GetMandatory(nil /* ctx */) && 0 != l.GetTtl(nil /* ctx */) {
					Panic(fmt.Sprintf("%s%s, LinkNo=%d: ttl on non mandatory link must be 0", preString, wid, i))
				}
				// check if ttl on noop link is 0:
				if NOOP == l.Op && 0 != l.GetTtl(nil /* ctx */) {
					Panic(fmt.Sprintf("%s%s, LinkNo=%d: ttl on NOOP link must be 0", preString, wid, i))
				}
			}
			// check if there is at least one guard link:
			if -1 == firstGuardIndex {
				Panic(fmt.Sprintf("%s%s: no guard link", preString, wid))
			}
			// check if last link has commit set:
			if false == w.Links[len(w.Links)-1].GetCommit(nil /* ctx */) {
			}
		}
	}
}

//------------------------------------------------------------
// process a ripe pm slot
// - currently only for entry ttl expired an actions is performed
func (ps *PeerSpace) ProcessRipePMSlot(userSlot ISlot, scheduler *Scheduler) {
	pmSlot := userSlot.(*PMSlot)
	switch pmSlot.Type {
	case ETTL:
		//------------------------------------------------------------
		// hunt the entry and wrap it into exception, if found and not locked
		ps.EntryHunter(pmSlot.Eid, scheduler)
		//------------------------------------------------------------
	// for the following cases: nothing needs to be done
	case ETTS:
		fallthrough
	case LTTS:
		fallthrough
	case LTTL:
		fallthrough
	case WTTS:
		fallthrough
	case WTTL:
		if SCHEDULER_TRACE.DoTrace() {
			/**/ scheduler.Println(0)
		}
		break
	default:
		Panic(fmt.Sprintf("ill. pm user slot = %s", pmSlot.Type))
	}
}

// =========================================================
// latex export of meta model:
// @@@ TBD: CREATE, DELETE!!!!
// @@@ PIC / POC abfrage ist nicht optimal/nicht ganz richtig
// @@@ TBD sub peer treatment / IOP treatment
// @@@ aution: latex bug: poc guards must be *before* services -> so we output all services first
func (ps *PeerSpace) MetaModel2Latex(testCaseName string, testCaseLatexConfig *LatexConfig) {
	file := LATEX_FILE

	// for report:
	file.WriteString("%%======================================================================= \n")
	file.WriteString("%%======================================================================= \n")
	file.WriteString(fmt.Sprintf("\\section{%s} \n\n", ConvertString2LatexString(testCaseName)))

	// print all peers:
	if nil != ps {
		for pIndex := 0; pIndex < len(ps.PeerPids); pIndex++ {
			pid := ps.PeerPids[pIndex]
			p := ps.Peers[pid]

			file.WriteString("%%======================================================================= \n")
			file.WriteString(fmt.Sprintf("\\subsection{%s} \n\n", ConvertString2LatexString(p.Id)))

			// print all wirings: sorted:
			for wIndex := 0; wIndex < len(p.WiringWids); wIndex++ {
				wid := p.WiringWids[wIndex]
				w := p.Wirings[wid]

				file.WriteString(fmt.Sprintf("%%----------------------------------------------------------------------- \n"))
				// copy for temporarty manipulations
				tmpLatexConfig := testCaseLatexConfig

				file.WriteString(fmt.Sprintf("\\begin{flushleft}\n"))
				file.WriteString(fmt.Sprintf("\\scalebox{%f}{ \n", tmpLatexConfig.Scalebox))
				// artificial alignment to the left to gain more space
				file.WriteString(fmt.Sprintf("\\hspace*{-0.5cm}"))
				file.WriteString(fmt.Sprintf("\\begin{peerless}"))
				// artificial alignment to adapt the line spacing in the wiring
				file.WriteString(fmt.Sprintf("\\setstretch{0.60}"))

				// are there any pic links? or poc links; the test is not really good... @@@
				// it tests whether a PIC or POC is involved from my own peer (ie not from a sup peer) in any link
				anyPicLinks := false
				anyPocLinks := false
				// the concrete number is needed for graphical heuristics below
				nPicLinks := 0
				nPocLinks := 0
				for _, l := range w.Links {
					// if "" == l.SubPid && (-1 != strings.Index(l.C1, PIC) || -1 != strings.Index(l.C2, PIC)) {
					if "" == l.SubPid && (CidIsPic(l.C1) || CidIsPic(l.C2)) {
						anyPicLinks = true
						nPicLinks++
					}
					// if "" == l.SubPid && (-1 != strings.Index(l.C1, POC) || -1 != strings.Index(l.C2, POC)) {
					if "" == l.SubPid && (CidIsPoc(l.C1) || CidIsPoc(l.C2)) {
						anyPocLinks = true
						nPocLinks++
					}
				}
				if !anyPicLinks {
					file.WriteString("\\noPIC]")
					//					// optimization: adjust arrow length
					//					tmpLatexConfig.WiringArrowRightWidth = testCaseLatexConfig.WiringArrowLeftWidth + testCaseLatexConfig.WiringArrowRightWidth
				}
				if !anyPocLinks {
					file.WriteString("[\\noPOC]")
					//					// optimization: adjust arrow length
					//					tmpLatexConfig.WiringArrowLeftWidth = testCaseLatexConfig.WiringArrowLeftWidth + testCaseLatexConfig.WiringArrowRightWidth
				}
				file.WriteString("\n\n")

				// --------------------------------------
				// optimization for graphical output - a bit experimental...:

				// if there are more than 5 service related links, shrink a bit the service space and
				// if more than 10 -> shrink also the right and left link lenght
				k := 0
				for i := 0; i < len(w.Links); i++ {
					l := w.Links[i]
					if SERVICE_IN == l.Type || SERVICE == l.Type || SERVICE_OUT == l.Type {
						k++
					}
				}
				if k >= 5 {
					tmpLatexConfig.WiringArrowServiceSpace = tmpLatexConfig.WiringArrowServiceSpace * 0.6
					tmpLatexConfig.WiringArrowLeftWidth = tmpLatexConfig.WiringArrowLeftWidth * 0.7
					tmpLatexConfig.WiringArrowRightWidth = tmpLatexConfig.WiringArrowRightWidth * 0.7
				}
				if k >= 10 {
					tmpLatexConfig.WiringArrowServiceSpace = tmpLatexConfig.WiringArrowServiceSpace * 0.3
					tmpLatexConfig.WiringArrowLeftWidth = tmpLatexConfig.WiringArrowLeftWidth * 0.3
					tmpLatexConfig.WiringArrowRightWidth = tmpLatexConfig.WiringArrowRightWidth * 0.3
				}

				// if there is <= 1 pic and <= 1 poc link -> increase the box height
				// so that the docu fits into the wiring box
				//				if 1 >= nPicLinks && 1 >= nPocLinks {
				//					tmpLatexConfig.SlotHeight = tmpLatexConfig.SlotHeight * 2
				//				}

				// --------------------------------------
				// config:
				file.WriteString(fmt.Sprintf("  \\setSlotHeight[%f] \n", tmpLatexConfig.SlotHeight))                             // hight of guard and action boxes
				file.WriteString(fmt.Sprintf("  \\setSlotWidth[%f] \n", tmpLatexConfig.SlotWidth))                               // width of guard and action boxes
				file.WriteString(fmt.Sprintf("  \\setWiringArrowLeftWidth[%f] \n", tmpLatexConfig.WiringArrowLeftWidth))         // guard arrows length
				file.WriteString(fmt.Sprintf("  \\setWiringArrowRightWidth[%f] \n", tmpLatexConfig.WiringArrowRightWidth))       // action arrows length
				file.WriteString(fmt.Sprintf("  \\setWiringArrowServiceHeight[%f] \n", tmpLatexConfig.WiringArrowServiceHeight)) // height of in and out service arrows
				file.WriteString(fmt.Sprintf("  \\setWiringArrowServiceSpace[%f] \n", tmpLatexConfig.WiringArrowServiceSpace))   // horizontal distance between service arrows
				file.WriteString(fmt.Sprintf("  \\setContainerWidth[%f] \n", tmpLatexConfig.ContainerWidth))                     // width of PIC and POC containers
				file.WriteString(fmt.Sprintf("  \\setMinWiringWidth[%f] \n", tmpLatexConfig.MinWiringWidth))                     // minimal width of wiring box
				file.WriteString(fmt.Sprintf("\n"))

				file.WriteString(fmt.Sprintf("  \\BeginWiring{ \n"))
				file.WriteString(fmt.Sprintf("    \\wiringDefinition{}{%s} \n", ConvertString2LatexString(w.Id)))
				file.WriteString(fmt.Sprintf("    \\wiringDocuSpace \n"))
				// wiring docu: WProps:
				// --------------
				wpropsstring1 := "\\wiringDocuLine{" + w.WProps.ToString(0) + "}\n"
				// @@@ maximal 100 times
				wpropsstring2 := strings.Replace(wpropsstring1, " ", "}\n    \\wiringDocuLine{", 100)
				file.WriteString(fmt.Sprintf("    %s", ConvertString2LatexString(wpropsstring2)))
				// --------------
				//				// max threads:
				//				file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=%d} \n", ConvertString2LatexString(MAX_THREADS), w.GetMaxThreads(nil /* ctx */)))
				//				// tts:
				//				tts := w.GetTts(nil /* no ctx */)
				//				if INFINITE == tts {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=INFINITE} \n", ConvertString2LatexString(TTS)))
				//				} else {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=%d} \n", ConvertString2LatexString(TTS), tts))
				//				}
				//				// ttl:
				//				ttl := w.GetTtl(nil /* no ctx */)
				//				if INFINITE == ttl {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=INFINITE} \n", ConvertString2LatexString(TTL)))
				//				} else {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=%d} \n", ConvertString2LatexString(TTL), ttl))
				//				}
				//				// txcc:
				//				file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=%s} \n", ConvertString2LatexString(TXCC), ConvertString2LatexString(w.GetTxcc(nil /* no ctx */))))
				//				// repeat count:
				//				if INFINITE == w.GetRepeatCount(nil /* no ctx */) {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=INFINITE} \n", ConvertString2LatexString(REPEAT_COUNT)))
				//				} else {
				//					file.WriteString(fmt.Sprintf("    \\wiringDocuLine{%s=%d} \n", ConvertString2LatexString(REPEAT_COUNT), w.GetRepeatCount(nil /* no ctx */)))
				//				}
				//				// --------------
				file.WriteString(fmt.Sprintf("    \\wiringDocuSpace \n"))
				file.WriteString(fmt.Sprintf("    \\wiringDocuEnd \n"))
				// end of wiring
				file.WriteString(fmt.Sprintf("  } \n"))

				nextService := false

				// --------------------------------------
				// output all links:

				// --------------------------------------
				// output services first (in the right order):
				for i := 0; i < len(w.Links); i++ {
					l := w.Links[i]

					switch l.Type {
					case SERVICE_IN:
						// print start service if not yet:
						if !nextService {
							nextService = true
							if nil == w.ServiceWrappers[l.Sid] {
								Panic(fmt.Sprintf("MetaModel2Latex: ill. meta model: service %s does not exist", l.Sid))
							}
							file.WriteString(fmt.Sprintf("  \\BeginService{%s:%s} \n", ConvertString2LatexString(l.Sid), ConvertString2LatexString(w.ServiceWrappers[l.Sid].Name)))
						}
						switch l.Op {
						case READ:
							file.WriteString("    \\inCopyServiceArrow")
						case TAKE:
							file.WriteString("    \\inMoveServiceArrow")
						case TEST:
							file.WriteString("    \\inTestServiceArrow")
						case DELETE:
							file.WriteString("    \\inDeleteServiceArrow")
						case CREATE:
							file.WriteString("    \\inCreateServiceArrow")
						case NOOP:
							file.WriteString("    \\inNoopServiceArrow")
						}
						typeCntQueryEPropsVarsLProps2Latex(file, l)
					case SERVICE:
						// print start service if not yet:
						if !nextService {
							nextService = true
							file.WriteString(fmt.Sprintf("  \\BeginService{%s:%s} \n", ConvertString2LatexString(l.Sid), ConvertString2LatexString(w.ServiceWrappers[l.Sid].Name)))
						}
						// file.WriteString("    \\callServiceArrow{}{}{}{}{}{} \n")
						file.WriteString("    \\callServiceArrow")
						typeCntQueryEPropsVarsLProps2Latex(file, l)
					case SERVICE_OUT:
						// only move exists for out
						file.WriteString("    \\outServiceArrow")
						typeCntQueryEPropsVarsLProps2Latex(file, l)
					default:
						// treated below
						break
					}
				}
				// --------------------------------------
				// print end service for last service (if there was at least one service)
				if nextService {
					nextService = false
					file.WriteString(fmt.Sprintf("    \\EndService \n"))
				}

				// --------------------------------------
				// then output all other links (in the right order):
				for i := 0; i < len(w.Links); i++ {
					l := w.Links[i]

					// which container (or sub peer):
					cId := ""
					ltype := ""
					if GUARD == l.Type {
						cId = l.C1
						ltype = "Guard"
					} else if ACTION == l.Type {
						cId = l.C2
						ltype = "Action"
					} else {
						// service
						continue
					}

					file.WriteString("  \\")

					if "" != l.SubPid {
						file.WriteString("subPeer")
					} else {
						if -1 != strings.Index(cId, PIC) {
							file.WriteString("pic")
						}
						if -1 != strings.Index(cId, POC) {
							file.WriteString("poc")
						}
					}

					// Copy, Move, Test, Delete or Create:
					op := ""
					switch l.Op {
					case READ:
						op = "Copy"
					case TAKE:
						op = "Move"
					case DELETE:
						op = "Delete"
					case TEST:
						op = "Test"
					case CREATE:
						op = "Create"
					case NOOP:
						op = "Noop"
					}
					file.WriteString(fmt.Sprintf("%s", op))

					// ltype:
					file.WriteString(fmt.Sprintf("%s", ltype))

					if "" != l.SubPid {
						file.WriteString(fmt.Sprintf("{%s}", l.SubPid))
					}
					typeCntQueryEPropsVarsLProps2Latex(file, l)
				}

				file.WriteString(fmt.Sprintf("  \\EndWiring \n"))
				file.WriteString(fmt.Sprintf("\\end{peerless} \n"))
				file.WriteString(fmt.Sprintf("} \n"))
				file.WriteString(fmt.Sprintf("\\end{flushleft}\n\n"))
			}
		}
	} else {
		file.WriteString(fmt.Sprintf("meta model not yet initialized \n"))
	}
	file.WriteString(fmt.Sprintf("%%======================================================================= \n"))
	file.WriteString(fmt.Sprintf("%% EOF \n"))
	file.WriteString(fmt.Sprintf("%%======================================================================= \n"))

	// file.Close()
	file.Sync()
}

// =========================================================
// {type}{cnt}{sel}{eprops}{vars}{lprops}:
func typeCntQueryEPropsVarsLProps2Latex(file *os.File, l *Link) {
	var typeString, cntString, selString, epropsString, varsString, lpropsString string
	// ----------------------------------------
	// type:
	if SERVICE != l.Type {
		if NOOP != l.Op {
			typeString = ConvertString2LatexString(l.Q.Typ.ToString(0))
			// // strip the latex quotes from the string that represents the entry type -- hack; "\ttdqt " has blank afterwards...
			// typeString = strings.Replace(typeString, "\\ttdqt ", "", 2 /* twice */)
		}
	} else {
		typeString = "{\\bf call}"
	}

	// ----------------------------------------
	// cnt:
	// @@@ vereinfachen!!!!
	// @@@ direct usage of ALL and NONE
	if SERVICE != l.Type {
		if NOOP != l.Op {
			// if min == max and both are ints -> print only int count:
			if VAL == l.Q.Min.Kind && INT == l.Q.Min.Type &&
				VAL == l.Q.Max.Kind && INT == l.Q.Max.Type && l.Q.Min.IntVal == l.Q.Max.IntVal {
				cntString = fmt.Sprintf("%d", l.Q.Max.IntVal)
			} else {
				// if count == var -> print variable name:
				if VAR == l.Q.Count.Kind {
					cntString = fmt.Sprintf("%s", l.Q.Count.Name)
				} else {
					// if max == ALL and min == 0 -> print only ALL:
					if VAL == l.Q.Max.Kind && INT == l.Q.Max.Type && ALL == l.Q.Max.IntVal &&
						VAL == l.Q.Min.Kind && INT == l.Q.Min.Type && 0 == l.Q.Min.IntVal {
						cntString = fmt.Sprintf("\\bf{%s}", "ALL")
					} else {
						// if max == NONE and min == 1 -> print only NONE:
						if VAL == l.Q.Max.Kind && INT == l.Q.Max.Type && NONE == l.Q.Max.IntVal &&
							VAL == l.Q.Min.Kind && INT == l.Q.Min.Type && 1 == l.Q.Min.IntVal {
							cntString = fmt.Sprintf("\\bf{%s}", "NONE")
						} else {

							// if min == ALL -> print ALL:
							if VAL == l.Q.Min.Kind && INT == l.Q.Min.Type && ALL == l.Q.Min.IntVal {
								cntString = fmt.Sprintf("\\bf{%s} ", "ALL")
							} else {
								// if min == NONE -> print NONE:
								if VAL == l.Q.Min.Kind && INT == l.Q.Min.Type && NONE == l.Q.Min.IntVal {
									cntString = fmt.Sprintf("\\bf{%s} ", "NONE")
								} else {
									cntString = fmt.Sprintf("%s", l.Q.Min.ToString(0))
								}
							}

							// if max == ALL -> print ALL:
							if VAL == l.Q.Max.Kind && INT == l.Q.Max.Type && ALL == l.Q.Max.IntVal {
								cntString = fmt.Sprintf("%s; \\bf{%s}", cntString, "ALL")
							} else {
								// if max == NONE -> print NONE:
								if VAL == l.Q.Max.Kind && INT == l.Q.Max.Type && NONE == l.Q.Max.IntVal {
									cntString = fmt.Sprintf("%s; \\bf{%s}", cntString, "NONE")
								} else {
									cntString = fmt.Sprintf("%s; %s", cntString, l.Q.Max.ToString(0))
								}
							}
						}
					}
				}
			}
		}
	}

	// ----------------------------------------
	// sel:
	if SERVICE != l.Type && nil != l.Q.Sel {
		selString = l.Q.Sel.ToString(0)
	}

	// ----------------------------------------
	// eprops:
	if SERVICE != l.Type {
		epropsString = l.EProps.ToString(0)
	}

	// ----------------------------------------
	// vars:

	if SERVICE != l.Type {
		varsString = l.LVars.ToString(0)
	}

	// ----------------------------------------
	// lprops:
	lpropsString = l.LProps.ToString(0)

	// ----------------------------------------
	file.WriteString(ConvertString2LatexString(fmt.Sprintf("{%s}{%s}{%s}{%s}{%s}{%s}\n",
		typeString, cntString, selString, epropsString, varsString, lpropsString)))
}

////////////////////////////////////////
// CheckEttl
////////////////////////////////////////

// ----------------------------------------
// a lock was released on entry in c:
// check for entry whether its ttl has expired and if so and entry is not locked:
//   - remove it from cid
//   - wrap it into exception
//   - move it to poc -- a bit complicated to figure out the right poc....
func (ps *PeerSpace) CheckEttl(e *Entry, c *Container, scheduler *Scheduler) {
	if SCHEDULER_DETAILS_TRACE.DoTrace() {
		/**/ String2TraceFile(fmt.Sprintf("CheckEttl: eid=%s, cid=%s\n", e.Id, c.Id))
	}
	if CLOCK > ConvertInfiniteTtl(e.GetTtl()) && !e.Locked() {
		// --
		// remove entry from container
		c.RemoveEntry(e.Id)
		// --
		// wrap entry into exception
		excE := e.ExceptionWrap("@@@DUMP-STELLE-1")
		// --
		// find poc:
		// @@@ tricky: same name - just replace substring pic by poc:
		var pocCid string = ""
		if c.IsPic() {
			pocCid = strings.Replace(c.Id, PIC, POC, 1)
			if SCHEDULER_DETAILS_TRACE.DoTrace() {
				/**/ String2TraceFile(fmt.Sprintf("  entry %s was unlocked: its ttl has expired; write EXC to pocCid=%s\n", e.Id, pocCid))
			}
		} else {
			if c.IsPoc() {
				pocCid = c.Id
			} else {
				Panic("CheckEttl: c is neither pic nor poc")
			}
		}
		// --
		// write exception to peer's poc
		ps.Write(pocCid, excE, nil /* @@@ no vars */, scheduler)
		if SCHEDULER_TRACE.DoTrace() {
			/**/ ps.Containers[pocCid].Println(TAB)
		}
	}
}

// ----------------------------------------
// same as ContainerPtrChangeEvent
func (ps *PeerSpace) ContainerChangeEvent(cid string) {
	c := ps.Containers[cid]
	if nil == c {
		Panic("ContainerChangeEvent: ill. container")
	}
	ContainerPtrChangeEvent(c)
}

////////////////////////////////////////
// clear wc
// @@@ should belong to model (i.e. Wiring); but impossible because it uses status and machine...
////////////////////////////////////////

// --------------------------------------------
// clear all entry collections of the wiring: its WC and all SICs and SOUTCs:
func (ps *PeerSpace) ClearEntryCollections(w *Wiring, machineNumber int) {
	// ------------
	// clear WC:
	// convert name to machine
	wcNamePtr := ConvertCtoM(w.WCId, machineNumber)
	// overwrite current container with a new one:
	ps.Containers[*wcNamePtr] = NewContainer(*wcNamePtr)
	// ------------
	// clear all sin and sout containers for wiring's service wrappers:
	for _, sw := range w.ServiceWrappers {
		// ----
		// Service InCid:
		// convert container cid to machine
		incidNamePtr := ConvertCtoM(sw.InCid, machineNumber)
		// overwrite current container with a new one:
		ps.Containers[*incidNamePtr] = NewContainer(*incidNamePtr)
		// ----
		// Service OutCid:
		// convert container cid to machine
		outcidNamePtr := ConvertCtoM(sw.OutCid, machineNumber)
		// overwrite current container with a new one:
		ps.Containers[*outcidNamePtr] = NewContainer(*outcidNamePtr)
	}
}

// ----------------------------------------
func (ps *PeerSpace) ConditionIsFulfilled(condition IEvent, ConditionEventIssueTime int) bool {
	if nil == condition {
		return false
	}

	pmCondition := condition.(*PMUserEvent)

	fulfilledFlag := false

	// TBD: default is missing
	switch pmCondition.Type {
	case CONTAINER_CHANGE_EVENT:
		c := ps.Containers[pmCondition.Cid]
		if nil != c && c.LastUpdateEventTime >= ConditionEventIssueTime {
			fulfilledFlag = true
		}
	}
	return fulfilledFlag
}

// =========================================================
// debug function - print the model
func (ps *PeerSpace) ModelPrint(tl TraceLevelEnum, nBlanks int, actualTestCaseName string) {
	if tl.DoTrace() {
		/**/ NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile("============================================================\n")

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(actualTestCaseName)
		/**/ String2TraceFile("\n")

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(fmt.Sprintf("VERIFICATION_MODE=%s\n", VERIFICATION_MODE))

		if VERIFICATION_MODE == SIMULATION {
			/**/ NBlanks2TraceFile(nBlanks)
			/**/ String2TraceFile(fmt.Sprintf("SIMULATION_COUNT=%d\n", SIMULATION_COUNT))
		}

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(fmt.Sprintf("EXECUTION_MODE=%s\n", EXECUTION_MODE))

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(fmt.Sprintf("SYSTEM_TTL=%d\n", SYSTEM_TTL))

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(fmt.Sprintf("TICK_FREQUENCY=%d\n", TICK_FREQUENCY))

		/**/
		NBlanks2TraceFile(nBlanks)
		/**/ String2TraceFile(fmt.Sprintf("==================== MODEL at CLOCK = %d ====================\n", CLOCK))
		// print all peers:
		for i := 0; i < len(ps.PeerPids); i++ {
			pid := ps.PeerPids[i]
			p := ps.Peers[pid]
			/**/ NBlanks2TraceFile(nBlanks)
			/**/ String2TraceFile(fmt.Sprintf("Peer %s:\n", p.Id))
			// print wirings:
			for wIndex := 0; wIndex < len(p.WiringWids); wIndex++ {
				wid := p.WiringWids[wIndex]
				w := p.Wirings[wid]
				/**/ w.Println(nBlanks + TAB)
			}
		}
		/**/ NBlanks2TraceFile(nBlanks)
		String2TraceFile("============================================================\n")
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// =========================================================
// debug function - print the space (sorted), ie all its containers:
// - CAUTION/TBD: nBlanks is ignored...
func (ps *PeerSpace) SpacePrint(tl TraceLevelEnum, nBlanks int, printAlsoEmptyContainersFlag bool) {
	nBlanks = 4 // !!!
	if tl.DoTrace() {
		/**/ String2TraceFile("\n")
		/**/ NBlanks2TraceFile(nBlanks)
		String2TraceFile(fmt.Sprintf("-------------------- SPACE at CLOCK=%d --------------------\n", CLOCK))
		// print all containers:
		for i := 0; i < len(ps.ContainerCids); i++ {
			cid := ps.ContainerCids[i]
			c := ps.Containers[cid]
			if printAlsoEmptyContainersFlag || 0 < len(c.Entries) {
				/**/ c.Println(nBlanks)
			}
		}
		/**/ NBlanks2TraceFile(nBlanks)
		String2TraceFile("---------------------------------------------------------------\n")
	}
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
