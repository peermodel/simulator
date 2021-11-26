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
// Copyright: eva Kuehn
// 2016
//////////////////////////////////////////////////////////////

package pmAutomata

import (
	. "cca/debug"
	. "cca/framework"
	. "cca/helpers"
	. "cca/pmModel"
	. "cca/scheduler"
	"fmt"
)

//////////////////////////////////////////////////////////////
// data type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// peer model automata generator data struct
type PeerModelAutomataGenerator struct {
	// automaton Id
	///////////	aId AutomatonID
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// create a peer model automata generator
func NewPeerModelAutomataGenerator() *PeerModelAutomataGenerator {
	return new(PeerModelAutomataGenerator)
}

//////////////////////////////////////////////////////////////
// methods
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// init the runtime model and start its machines:
// - create all needed containers
// - start all wiring machines
func (a PeerModelAutomataGenerator) InitRuntimeModelAndStartMachines(s *Status) {
	// - create containers
	a.CreateContainers4RuntimeModel(s)
	// - start wiring machines
	a.StartWiringMachines4RuntimeModel(s)
}

//------------------------------------------------------------
// init the runtime model (part a):
// - create all needed containers
func (a PeerModelAutomataGenerator) CreateContainers4RuntimeModel(s *Status) {
	//------------------------------------------------------------
	// debug
	if RUN_TRACE.DoTrace() { // DEBUG
		/**/ String2MCTraceFile("create containers & start wirings\n") // DEBUG
	} // DEBUG
	//------------------------------------------------------------
	// for all peers:
	// - create their pic and poc containers
	// - start their wiring machines
	for _, p := range s.MetaContext.(*MetaContext).PeerSpace.Peers {
		//------------------------------------------------------------
		// debug
		if RUN_TRACE.DoTrace() { // DEBUG
			/**/ String2MCTraceFile(fmt.Sprintf("next peer %s\n", p.Id)) // DEBUG
			/**/ String2MCTraceFile(fmt.Sprintf("create and add pic and poc containers; pic = %s, poc = %s\n", p.Pic, p.Poc))
		}
		//------------------------------------------------------------
		// create pic and poc containers for peer
		pic := *NewContainer(p.Pic)
		poc := *NewContainer(p.Poc)
		//------------------------------------------------------------
		// debug
		if RUN_TRACE.DoTrace() { // DEBUG
			/**/ String2MCTraceFile(fmt.Sprintf("add pic and poc containers to peer space: pic = %s, poc = %s\n", p.Pic, p.Poc)) // DEBUG
		} // DEBUG
		//------------------------------------------------------------
		// add pic and poc to peer space, which is contained in the meta context
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&pic)
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&poc)
	}
}

//------------------------------------------------------------
// init the runtime model (part b):
// - create wiring automata and machines and start them
func (a PeerModelAutomataGenerator) StartWiringMachines4RuntimeModel(s *Status) {
	//------------------------------------------------------------
	// debug
	if RUN_TRACE.DoTrace() { // DEBUG
		/**/ String2MCTraceFile("create containers & start wirings\n") // DEBUG
	} // DEBUG
	//------------------------------------------------------------
	// for all peers
	// - create and start their wiring machines
	for _, p := range s.MetaContext.(*MetaContext).PeerSpace.Peers {
		//------------------------------------------------------------
		// debug
		if RUN_TRACE.DoTrace() { // DEBUG
			/**/ String2TraceFile("start all wirings of peer:}n") // DEBUG
		} // DEBUG
		//------------------------------------------------------------
		// start given number of wiring instance(s) (= wiring machine(s)) for each wiring
		// - incl. its WIC (wiring internal container)
		for _, w := range p.Wirings {
			//------------------------------------------------------------
			// get max-threads property of the wiring
			// - TBD: no eval required?!
			maxthreads := w.GetMaxThreads(nil /* ctx */)
			//------------------------------------------------------------
			// debug
			if RUN_TRACE.DoTrace() { // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: Status.Run: starting wid %s, MaxThreads=%d\n", w.Id, maxthreads)) // DEBUG
				/**/ String2TraceFile(fmt.Sprintf("%s: wid=%s, MaxThreads=%d\n", w.Id, maxthreads)) // DEBUG
			} // DEBUG
			//------------------------------------------------------------
			// async start wiring
			// - in as many threads as specified by max thread count
			for i := 0; i < maxthreads; i++ {
				// creates automaton with the code, if not yet exists, and create and starts the wiring machine
				asyncStartWiring(s, w, p)
			}
		}
		//------------------------------------------------------------
		// TBD:
		// - scheduler shall continuously hunt for outdated entries contained in POC or POC of this wiring in the given time interval;
		// - s.Scheduler = SetPeerEntriesHuntSlot(s.Scheduler, p.Id, PEER_ENTRIES_HUNT_REPEAT_INTERVAL)
	}
}

//////////////////////////////////////////////////////////////
// start wiring
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// asynchronous wiring start
// - create automaton that holds the program, if not yet exists
// - create and start a new machine for the wiring based on that automaton in a new thread
// - create all containers of wiring for this machine: WC, SICs and SOCs
// - return its machine number; needed that a dynamic wiring can be cleaned up
// private fu
func asyncStartWiring(s *Status, w *Wiring, p *Peer) int {
	//------------------------------------------------------------
	// create new automaton (of not yet) and new wiring machine
	// - exists?
	foundAutomaton, foundFlag := s.CheckAutomatonExistence("Wiring")
	// - create if not yet
	a, wm := NewAutomaton_Wiring("Wiring", !foundFlag /* createAutomatonFlag */, foundAutomaton /* Automaton */)
	// - add automaton to status
	s.AddAutomaton(a)
	//------------------------------------------------------------
	// create wiring container
	// - nb:  machine number is needed for the WIC name
	wcNamePtr := ConvertCtoM(w.WCId, wm.Number)
	wic := *NewContainer(*wcNamePtr)
	// /**/ m.PrintlnA(TRACE0, TAB, "wcName", *wcNamePtr)
	//------------------------------------------------------------
	// add WIC to peer space
	s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&wic)
	//------------------------------------------------------------
	// create and init context for the machine:
	// - alloc
	ctx := NewContext()
	// - set pid
	ctx.Pid = p.Id
	// - set wid
	ctx.Wid = w.Id
	// - set wmno
	ctx.WMNo = wm.Number
	//------------------------------------------------------------
	// debug
	// m.PrintlnA(TRACE0, TAB, "Pid", ctx.Pid) // DEBUG
	// m.PrintlnA(TRACE0, TAB, "Wid", ctx.Wid) // DEBUG
	// m.PrintlnA(TRACE0, TAB, "WMNo", ctx.WMNo) // DEBUG
	wm.PrintlnI(TRACE0, TAB, "start asynchronous wiring machine WMNo", wm.Number) // DEBUG
	//------------------------------------------------------------
	// start wiring machine in asynchronous thread with the shared status and its individual context;
	// - nb: the status uses mutual exclusion to control all machines of the system
	// - start with context
	wm.StartAsync(s, ctx)
	//------------------------------------------------------------
	// create all sin and sout containers for the wiring's service wrappers
	// - caution: do this only after wiring machine has been started, because its number is needed below:
	for _, sw := range w.ServiceWrappers {
		//------------------------------------------------------------
		// service InCid:
		// - convert container cid to machine
		inCidNamePtr := ConvertCtoM(sw.InCid, wm.Number)
		inC := *NewContainer(*inCidNamePtr)
		//------------------------------------------------------------
		// add service inCid to peer space:
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&inC)
		//------------------------------------------------------------
		// service OutCid:
		// - convert container cid to machine
		outCidNamePtr := ConvertCtoM(sw.OutCid, wm.Number)
		outC := *NewContainer(*outCidNamePtr)
		//------------------------------------------------------------
		// add service outCid to peer space:
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&outC)
		//------------------------------------------------------------
		// debug
		// m.Println2(TRACE0, TAB, "Service InCid", *incidNamePtr, "Service OutCid", *outcidNamePtr)
	}
	//------------------------------------------------------------
	// retun number of the new wiring machine
	return wm.Number
}

// =========================================================
// get aliases into the status for a machine: using the machine context variables;
// they raise system error upon failure
// @@@tbd: system or user error? or selectable
// @@@ tbd: in welches package/zu welchem type gehören diese functions?
// =========================================================

// =========================================================
// get peer alias: via m.Context's Pid
func GetPeerAlias(m *Machine, s *Status) *Peer {
	ctx := m.Context.(*Context)

	p := s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid]

	// detection if found:
	if nil == p || "" == p.Id {
		m.SystemError(fmt.Sprintf("ill. pid=%s", ctx.Pid))
	}
	return p
}

// =========================================================
// get wiring alias: via m.Context's Pid, Wid
func GetWiringAlias(m *Machine, s *Status) *Wiring {
	ctx := m.Context.(*Context)

	p := GetPeerAlias(m, s)
	w := p.Wirings[ctx.Wid]
	// detection if found:
	if nil == w || "" == w.Id {
		m.SystemError(fmt.Sprintf("ill. wid=%s", ctx.Wid))
	}
	return w
}

// =========================================================
// get dynamic wiring alias: via given peer and dwid
func GetDynWiringAlias(m *Machine, s *Status, pid string, dwid string) *Wiring {
	ctx := m.Context.(*Context)

	p := s.MetaContext.(*MetaContext).PeerSpace.Peers[pid]
	// detection if found:
	if nil == p || "" == p.Id {
		m.SystemError(fmt.Sprintf("ill. pid=%s", ctx.Pid))
	}
	w := p.Wirings[dwid]
	// detection if found:
	if nil == w || "" == w.Id {
		m.SystemError(fmt.Sprintf("ill. wid=%s", dwid))
	}
	return w
}

// =========================================================
// get link alias: via m.Context's Pid, Wid, LinkNo
func GetLinkAlias(m *Machine, s *Status) *Link {
	ctx := m.Context.(*Context)

	w := GetWiringAlias(m, s)
	l := w.Links[ctx.LinkNo]
	// detection if found:
	if nil == l {
		m.SystemError(fmt.Sprintf("ill. link no=%d", ctx.LinkNo))
	}
	return l
}

// =========================================================
// get wtx alias: via m.Context's Wtxid
func GetWtxAlias(m *Machine, s *Status) *Tx {
	ctx := m.Context.(*Context)

	wtx := s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
	// detection if found:
	if nil == wtx {
		m.SystemError(fmt.Sprintf("ill. wtx=%s", ctx.Wtxid))
	}
	return wtx
}

// =========================================================
// get container alias: via m.Context's Cid
func GetContainerAlias(m *Machine, s *Status) *Container {
	ctx := m.Context.(*Context)

	c := s.MetaContext.(*MetaContext).PeerSpace.Containers[ctx.Cid]
	// detection if found:
	if nil == c || "" == c.Id {
		m.SystemError(fmt.Sprintf("ill. cid=%s", ctx.Cid))
	}
	return c
}

// =========================================================
// create entries according to a query
func CreateEntries(m *Machine) EntryPtrs {
	ctx := m.Context.(*Context)
	writeEs := EntryPtrs{}

	// ------------
	// evaluate max:
	max := 0
	if ctx.Query.Max.Eval(ctx.Vars, nil /* no entry */) && INT == ctx.Query.Max.Type {
		max = ctx.Query.Max.IntVal
		/**/ m.PrintlnI(TRACE0, 0, "max", max)
	} else {
		m.UserError(fmt.Sprintf("create entries: ill. max specification: Query = %s", ctx.Query.ToString(0)))
	}

	// ------------
	// evaluate qTyp:
	qTyp := ""
	if ctx.Query.Typ.Eval(ctx.Vars, nil /* no entry */) && VAL == ctx.Query.Typ.Kind && STRING == ctx.Query.Typ.Type {
		qTyp = ctx.Query.Typ.StringVal
		/**/ m.PrintlnS(TRACE0, 0, "qTyp", qTyp)
	} else {
		m.UserError("create entries: ill. qTyp specification")
	}

	// ------------
	/********/
	m.PrintlnIS(TRACE0, TAB, "max", max, "qTyp", qTyp)
	// create count entries and append them to writeEs:
	for i := 0; i < max; i++ {
		e := NewEntry(qTyp)
		writeEs = append(writeEs, e)
	}
	/********/ m.PrintlnX(TRACE0, TAB, "writeEs", writeEs)

	return writeEs
}

// =========================================================
// create one exception entry for on abort with wid set; nb: returns a "list" of EntryPtrs!!
func CreateOnAbortEntry(m *Machine) EntryPtrs {
	ctx := m.Context.(*Context)
	writeEs := EntryPtrs{}

	// create one exception entry and append it to writeEs; set its wid; set current time for debugging;
	e := NewEntry(EXCEPTION_ON_ABORT)
	e.SetStringVal("wid", ctx.Wid)
	e.SetIntVal("execTime", CLOCK)
	writeEs = append(writeEs, e)

	/********/
	m.PrintlnX(TRACE0, TAB, "writeEs", writeEs)

	return writeEs
}

// =========================================================
// unwrap dest wrap readEs's first entry e into return entries
func DestUnWrap(m *Machine, readEs EntryPtrs) EntryPtrs {
	var retEntries EntryPtrs

	// ------------------
	// set writeEs to data of first entry read:
	k := len(readEs)
	if 0 < k {
		retEntries = readEs[0].Data
	} else {
		// can this happen? maybe if query allows to be fulfilled for 0 entries; just do nothing
		/**/
		m.PrintlnS(TRACE0, TAB, "", "WARNING: empty readEs")
		retEntries = EntryPtrs{}
	}

	return retEntries
}

// =========================================================
// wrap entry into dest wrap entry e
// wrap all given entries into one and set its dest to the given dest
// use current wfid
func DestWrap(m *Machine, l *Link, es EntryPtrs) *Entry {
	ctx := m.Context.(*Context)

	e := NewEntry(DEST_WRAP)
	e.Data = es

	// 	/**/ m.PrintlnXX(TRACE0, TAB, DEST, l.LProps[DEST], true /* detailsFlag */)
	dest := l.GetDest(ctx)
	e.SetStringVal(DEST, dest)
	/**/ m.PrintlnS(TRACE0, TAB, "resolved dest", dest)

	e.SetStringVal(FID, ctx.Wfid)

	// @@@ ttl should be set to the maximum of all ttls of entries in the es set!
	// @@@ default = INFINITE
	// e.Ttl = ...
	return e
}

// =========================================================
// call service
func CallService(m *Machine, s *Status, w *Wiring, l *Link) {
	ctx := m.Context.(*Context)

	cid1ptr := l.ConvertC1toM(ctx.WMNo)
	cid2ptr := l.ConvertC2toM(ctx.WMNo)

	// assertion:
	if "" == l.Sid || nil == w.ServiceWrappers[l.Sid] {
		m.SystemError(fmt.Sprintf("ServiceWrapper for %s is undefined", l.Sid))
	}
	/**/ m.Println(TRACE0)
	/**/ m.PrintlnSSS(TRACE0, 0, SERVICE_START_INFO, w.ServiceWrappers[l.Sid].Name, "SINC", *cid1ptr, "SOUTC", *cid2ptr)

	// call the service
	// w.ServiceWrappers[l.Sid].Fu(m, s, *cid1ptr, *cid2ptr)
	// @@@???wfid
	w.ServiceWrappers[l.Sid].Fu(s.MetaContext.(*MetaContext).PeerSpace, ctx.Wfid, ctx.Vars, &s.Scheduler, *cid1ptr, *cid2ptr, s.ControllerChannel)

	/**/
	m.Println(TRACE0)
	// /**/ m.PrintlnS(0, SERVICE_END_INFO, lvs.w.ServiceWrappers[lvs.l.Sid].Name)
	// /**/ m.Println()

	// @@@TBD: return check of service call
}

// =========================================================
// create dynamic wiring
// return dwid
func CreateDynamicWiring(m *Machine, s *Status, p *Peer, w *Wiring, l *Link) string {
	// /**/ m.PrintlnArgX(TRACE0, TAB, "Q", &m.Q)
	ctx := m.Context.(*Context)

	// /**/ m.PrintlnArgX(TRACE0, TAB, "Q", &m.Q)

	// ---------------------
	// remember original query in q; needed for guard in dynamic wiring below:
	q := ctx.Query.Copy()
	/**/ m.PrintlnX(TRACE0, TAB, "q", &q)

	// ---------------------
	// redefine Q:
	ctx.Query = Query{}
	ctx.Query.Typ = SEtype(SOURCE_WRAP)
	ctx.Query.Min = IVal(1)
	ctx.Query.Max = IVal(1)
	ctx.Query.Sel = &Arg{Kind: EXPR, ExprVal: &Expr{
		Left: Arg{Kind: EXPR, ExprVal: &Expr{
			Left:  Arg{Kind: LABEL, Type: STRING, Name: "Wiid"},
			Op:    EQUAL,
			Right: Arg{Kind: VAL, Type: STRING, StringVal: ctx.Wiid},
		}},
		Op: AND,
		Right: Arg{Kind: EXPR, ExprVal: &Expr{
			Left:  Arg{Kind: LABEL, Type: INT, Name: "LinkNo"},
			Op:    EQUAL,
			Right: Arg{Kind: VAL, Type: INT, IntVal: ctx.LinkNo},
		}},
	}}
	/**/ m.PrintlnX(TRACE0, TAB, "redefined Q", &(ctx.Query))

	// create dynamic wiring
	// trick: create it within dest peer
	// ---------------------
	// create a dynamic wiring id and wc id:
	// DW_<uuid>_...:
	dwId := fmt.Sprintf("%s", Uuid("DW"))
	// /**/ m.PrintlnS(TRACE0, TAB, "dwId", dwId)

	// default link properties:
	// eval
	defaultLinkProps := LProps{TTL: IVal(l.GetTtl(ctx))}

	// nb: GetSource function evaluates the value
	sourceAddress := l.GetSource(ctx)
	sourceP := s.MetaContext.(*MetaContext).PeerSpace.Peers[sourceAddress]
	if nil == sourceP || "" == sourceP.Id {
		helpText := fmt.Sprintf("ill. Source: '%s' does not exist (1)", sourceAddress)
		m.UserError(helpText)
	}
	// /**/ m.PrintlnS(TRACE0, TAB, "source peer", sourceP.Id)

	// ============================================
	// create the dynamic wiring:
	sw := NewServiceWrapper(SourceWrapService, "SourceWrapService")

	dw := NewWiring(dwId)

	dw.AddServiceWrapper("S1", sw)

	dw.AddGuard("", POC, l.Op, q, defaultLinkProps, EProps{}, Vars{})
	dw.AddSin(TAKE, q, "S1", defaultLinkProps, EProps{}, Vars{})
	dw.AddScall("S1", defaultLinkProps, EProps{}, Vars{})
	dw.AddSout(Query{Typ: SEtype(SOURCE_WRAP), Count: IVal(1)}, "S1", defaultLinkProps,
		EProps{"Wiid": SVal(ctx.Wiid), "LinkNo": IVal(ctx.LinkNo)},
		Vars{})

	// DEST = my own peer
	// LIMITATION of the model: everything is local; in realty we should use IOP at site of the other peer
	dw.AddAction(IOP_PEER, PIC, TAKE, Query{Typ: SEtype(SOURCE_WRAP), Count: IVal(1)},
		// nb: first evaluate all links args (= the ttl) locally, then transfer value(s) to remote
		LProps{
			TTL:       IVal(l.GetTtl(ctx)),
			MANDATORY: BVal(true),
			DEST:      SVal(p.Id),
			COMMIT:    BVal(true)},
		EProps{}, Vars{})

	// nb: first evaluate all wiring args locally, then transfer value(s) to remote
	dwProps := WProps{
		TTS:          IVal(w.GetTts(ctx)),
		TTL:          IVal(w.GetTtl(ctx)),
		TXCC:         SVal(w.GetTxcc(ctx)),
		MAX_THREADS:  IVal(1),
		REPEAT_COUNT: IVal(0)}
	dw.WProps = dwProps
	// /**/ m.PrintlnArgX(TRACE0, TAB, "dw props", &dw.Props)

	// for debug only
	dw.DynamicWiringFlag = true

	// ============================================
	// add the dynamic wiring to sourcePpeer & resolve names
	// @@@ workaround: wiring is added to sourceP; is ok; but could better be its enclosing peer?!
	// NB: dw object is changed -- it is used as ref param
	sourceP.AddWiring(dw)

	/**/
	m.PrintlnS(TRACE0, TAB, "", ">>>> dynamic wiring added:")
	/**/ m.PrintlnX(TRACE0, TAB*2, "", dw)

	// ============================================
	// start the dynamic wiring:
	// @@@@@@ check check check check
	// orig. code: lvs.dwMachineNumber = s.Start(dw, sourceP)
	// @@@ duplicated code -> see runtime start of all wirings...

	//------------------------------------------------------------
	// create context:
	ctxHelp := NewContext()

	//------------------------------------------------------------
	// create new automaton (of not yet) and new wiring machine
	// - existst?
	foundAutomaton, foundFlag := s.CheckAutomatonExistence("wiring")
	// - create if not yet there
	a, wm := NewAutomaton_Wiring("Wiring", !foundFlag /* createAutomatonFlag */, foundAutomaton /* Automaton */)
	// - add automaton to status
	s.AddAutomaton(a)

	//------------------------------------------------------------
	// create wiring container: nb: needs machine number for wiring container:
	wcNamePtr := ConvertCtoM(dw.WCId, wm.Number)
	// /**/ m.PrintlnA(TRACE0, TAB, "wcName", *wcNamePtr)
	wc := *NewContainer(*wcNamePtr)

	// add wc to peer space:
	s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&wc)

	// set context:
	ctxHelp.Pid = sourceP.Id
	// /**/ m.PrintlnA(TRACE0, TAB, "Pid", ctx.Pid)
	ctxHelp.Wid = dw.Id
	// /**/ m.PrintlnA(TRACE0, TAB, "Wid", ctx.Wid)
	ctxHelp.WMNo = wm.Number
	// /**/ m.PrintlnA(TRACE0, TAB, "WMNo", ctx.WMNo)

	// start wiring machine in asynchronous thread with the shared status and its individual context;
	// nb: the status uses mutual exclusion among all machines of the system
	/**/
	wm.PrintlnI(TRACE0, TAB, "start asynchronous wiring machine WMNo", wm.Number)
	// - start with context
	wm.StartAsync(s, ctxHelp)

	// caution: do this only after wiring machine has been started, because its number is needed below:

	// create all sin and sout containers for wiring's service wrappers
	for _, sw := range dw.ServiceWrappers {
		// Service InCid:
		// convert container cid to machine
		incidNamePtr := ConvertCtoM(sw.InCid, wm.Number)
		inc := *NewContainer(*incidNamePtr)
		// add incid to peer space:
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&inc)

		// Service OutCid:
		// convert container cid to machine
		outcidNamePtr := ConvertCtoM(sw.OutCid, wm.Number)
		outc := *NewContainer(*outcidNamePtr)
		// add outcid to peer space:
		s.MetaContext.(*MetaContext).PeerSpace.AddContainer(&outc)

		// /**/ m.Println2(TRACE0, TAB, "Service InCid", *incidNamePtr, "Service OutCid", *outcidNamePtr)
	}

	//=============================================
	// set dwid
	dwid := dw.Id

	return dwid
}

// =========================================================
// clean up dynamic wiring
func CleanUpDynamicWiring(m *Machine, s *Status, l *Link, dwid string, dwMachineNumber int) {
	ctx := m.Context.(*Context)

	ps := s.MetaContext.(*MetaContext).PeerSpace

	sourceAddress := l.GetSource(ctx)
	sourceP := ps.Peers[sourceAddress]
	dw := GetDynWiringAlias(m, s, sourceAddress /* = pid */, dwid)
	if nil == sourceP || "" == sourceP.Id {
		helpText := fmt.Sprintf("ill. Source: '%s' does not exist (2)", sourceAddress)
		m.UserError(helpText)
	}
	if nil == dw {
		m.SystemError("dw disappeared")
	}
	// - remove its WC
	cidptr := ConvertCtoM(dw.WCId, dwMachineNumber)
	delete(ps.Containers, *cidptr)
	ps.ContainerCids = ps.ContainerCids.RemoveString(*cidptr)

	// - remove all its SIC and SOUT
	for _, sw := range dw.ServiceWrappers {
		cidptr1 := ConvertCtoM(sw.InCid, dwMachineNumber)
		delete(ps.Containers, *cidptr1)
		ps.ContainerCids = ps.ContainerCids.RemoveString(*cidptr1)
		cidptr2 := ConvertCtoM(sw.OutCid, dwMachineNumber)
		delete(ps.Containers, *cidptr2)
		ps.ContainerCids = ps.ContainerCids.RemoveString(*cidptr2)
	}

	// - remove dw from "source" peer:
	//   nb: GetSource function evaluates the value
	/**/
	m.PrintlnS(TRACE0, TAB, "source peer", sourceP.Id)
	/**/ m.PrintlnSX(TRACE0, TAB, "remove DW", dwid, "from peer", sourceP)
	sourceP.RemoveWiring(dw)
}

// =========================================================
// @@@ TBD
// exception in user application: eg ttl expired
// @@@ padding len is hard coded
func Exception(m *Machine, s *Status, exc ExceptionTypeEnum, msg string) {
	//	// print some selected infos:
	//	w := s.PeerSpace.Peers[m.Pid].Wirings[m.Wid]
	//	if nil == w {
	//		/**/ m.PrintlnStarMessage(EXCEPTION, fmt.Sprintf("ill. Wid = %s", m.Wid))
	//	}

	// for debug only: only if wiring repeat: @@@
	// if exc == WIRING_REPEAT_EXCEPTION && false == w.dynamicWiringFlag {

	// @@@: if exc != WIRING_REPEAT_EXCEPTION || false == w.DynamicWiringFlag {
	if exc != WIRING_REPEAT_EXCEPTION {
		excpadded := fmt.Sprintf("%s ", exc.String())
		excpadded = Padding(excpadded, 15, "-")
		/**/ m.PrintlnStarMessage(EXCEPTION, fmt.Sprintf("%s %s", excpadded, msg))
	}
}

//////////////////////////////////////////////////////////////
// EOF
//////////////////////////////////////////////////////////////
