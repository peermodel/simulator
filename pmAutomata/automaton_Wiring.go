////////////////////////////////////////
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
// Copyright: eva Kuehn
////////////////////////////////////////

package pmAutomata

import (
    "errors"
    "fmt"
    . "github.com/peermodel/simulator/contextInterface"
    . "github.com/peermodel/simulator/debug"
    . "github.com/peermodel/simulator/helpers"
    . "github.com/peermodel/simulator/scheduler"
    . "github.com/peermodel/simulator/pmModel"
    . "github.com/peermodel/simulator/framework"
)

// --------------------------------------
// Wiring AUTOMATON:
// --------------------------------------
func NewAutomaton_Wiring(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        l *Link
        p *Peer
        w *Wiring
        // --------------------------------
        // ordinary variables:
        exc error
        repeatCount int
        readEs EntryPtrs
        res bool
        iopEs EntryPtrs
        dwMachineNumber int
        dwid string
        wTtl int
        es2 EntryPtrs
        wTts int
        wait4Cid string
        lTtl int
        lTts int
        startTime int
        e *Entry
        writeEs EntryPtrs
        nLinks int
    }

    // --------------------------------------
    // create new automaton:
    // --------------------------------------
    if createAutomatonFlag {
        a = NewAutomaton(automatonName)
    }

    // --------------------------------------
    // create new machine:
    // --------------------------------------
    m := NewMachine(a)

    // --------------------------------------
    // alloc LVS:
    // --------------------------------------
    m.LocalVariables = new(localVariables)

    if createAutomatonFlag {
    // --------------------------------------
    // define LVS copy function:
    // --------------------------------------
    a.LocalVariablesCopyFunction = func(theM *Machine, lvs interface{}) interface{} {
        // --------------------------------
        // cast ->:
        tmpOrigLvs := lvs.(*localVariables)
        // --------------------------------
        // alloc LVS:
        tmpNewLvs := new(localVariables)
        // --------------------------------
        // copy static fields:
        *tmpNewLvs = *tmpOrigLvs
        // --------------------------------
        // copy dynamic fields:
        if ! (IsPointer(tmpOrigLvs.readEs) && nil == tmpOrigLvs.readEs) {
            tmpNewLvs.readEs = tmpOrigLvs.readEs.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.iopEs) && nil == tmpOrigLvs.iopEs) {
            tmpNewLvs.iopEs = tmpOrigLvs.iopEs.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.es2) && nil == tmpOrigLvs.es2) {
            tmpNewLvs.es2 = tmpOrigLvs.es2.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.e) && nil == tmpOrigLvs.e) {
            tmpNewLvs.e = tmpOrigLvs.e.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.writeEs) && nil == tmpOrigLvs.writeEs) {
            tmpNewLvs.writeEs = tmpOrigLvs.writeEs.Copy()
        }
        // --------------------------------
        // cast <-:
        return (interface{})(tmpNewLvs)
    }

    // --------------------------------------
    // define LVS alias function:
    // --------------------------------------
    a.CompleteLocalVariablesAliasFunction = func(s *Status, theM *Machine, lvs interface{}) interface{} {
        // --------------------------------
        // cast ->:
        newLvs := lvs.(*localVariables)
        // --------------------------------
        newLvs.l = GetLinkAlias(theM, s)
        newLvs.p = GetPeerAlias(theM, s)
        newLvs.w = GetWiringAlias(theM, s)
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [Wid, Pid]
    // --------------------------------------
    a.AddState("init", "Wiring(ctx.Pid, ctx.Wid, ...): Wiring automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        
        // dummy code: helps that all imports are needed by every automaton
        s.DummyString = fmt.Sprintf("dummy")
        DummyHelpersFu()
        DummySchedulerFu()
        if ctx.DummyIContextFu().(IContext) == nil {}
        // reset all error vars
        ctx.RetErr = errors.New("")
        ctx.RetErr = nil
        
        m.CurrentState = "61"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - GVars: [Wfid, Wid, Wiid, Pid]
    //   - LVars: [wTtl, wTts]
    //   - Aliases:[w]
    // --------------------------------------
    a.AddState("1", "wiring instance start: set Wiid; compute wiring tts and ttl; inform scheduler; reset wfid; reset system vars;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wiid", ctx.Wiid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "- wTtl", lvs.wTtl)
        /**/ m.PrintlnI(TRACE0, TAB, "- wTts", lvs.wTts)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        
        ctx.Wiid = Uuid("wiid")
        lvs.wTts = lvs.w.GetAbsTts(ctx)
        lvs.wTtl = lvs.w.GetAbsTtl(ctx)
        ctx.Wfid = ""
        ctx.Vars.SetStringVal("$$PID", ctx.Pid)
        ctx.Vars.SetStringVal("$$WID", ctx.Wid)
        ctx.Vars.SetIntVal("$$WTTS", lvs.wTts)
        ctx.Vars.SetIntVal("$$WTTL", lvs.wTtl)
        ctx.Vars.SetStringVal("$$TXCC", lvs.w.GetTxcc(ctx))
        ctx.Vars.SetIntVal("$$CLOCK", CLOCK)
        ctx.Vars.SetIntVal("$$CNT", 0)
        ctx.Vars.SetStringVal("$$FID", "")
        ctx.Vars.SetIntVal("$$MAX_THREADS",   lvs.w.GetMaxThreads(ctx))
        ctx.Vars.SetIntVal("$$REPEAT_COUNT",   lvs.w.GetRepeatCount(ctx))
        s.Scheduler = SetWttsSlot(  s.Scheduler, lvs.wTts, ctx.Wid)
        s.Scheduler = SetWttlSlot(  s.Scheduler, lvs.wTtl, ctx.Wid)
        
        m.CurrentState = "60"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wiid", ctx.Wiid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "= wTtl", lvs.wTtl)
        /**/ m.PrintlnI(TRACE0, TAB, "= wTts", lvs.wTts)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        
        return OK
        })

    // --------------------------------------
    // 2: WAIT STATE
    //   - LVars: [wTts]
    // --------------------------------------
    a.AddState("2", "wait until wiring tts is reached", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- wTts", lvs.wTts)
        
        if ! s.Wait4TimeEvent(m, lvs.wTts, NO_CP) {
            m.CurrentState = "stopped" // for docu
            return STOPPED
        } else {
        m.CurrentState = "4"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= wTts", lvs.wTts)
        } 
    
        return OK
        })

    // --------------------------------------
    // 3: ACTION STATE
    //   - LVars: [repeatCount]
    // --------------------------------------
    a.AddState("3", "increment repeat counter", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- repeatCount", lvs.repeatCount)
        
        lvs.repeatCount = lvs.repeatCount + 1
        
        m.CurrentState = "52"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= repeatCount", lvs.repeatCount)
        
        return OK
        })

    // --------------------------------------
    // 4: CONDITION STATE
    //   - LVars: [wTts]
    // --------------------------------------
    a.AddState("4", "is wTts fulfilled?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- wTts", lvs.wTts)
        
        if (lvs.wTts <= CLOCK) { m.CurrentState = "5" } else { m.CurrentState = "2" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= wTts", lvs.wTts)
        
        return OK
        })

    // --------------------------------------
    // 5: ACTION STATE
    //   - GVars: [LinkNo, EvalEs, Wid, Query, Wiid, RetErr]
    //   - LVars: [exc, readEs, lTtl, wTtl, iopEs, lTts, writeEs]
    //   - Aliases:[w, l]
    // --------------------------------------
    a.AddState("5", "link start: get lTts and lTtl as absolute times; nb: lTtl is bounded by wTtl; inform scheduler; init vars", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnX(TRACE0, TAB, "- EvalEs", ctx.EvalEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wiid", ctx.Wiid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        /**/ m.PrintlnX(TRACE0, TAB, "- readEs", lvs.readEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- lTtl", lvs.lTtl)
        /**/ m.PrintlnI(TRACE0, TAB, "- wTtl", lvs.wTtl)
        /**/ m.PrintlnX(TRACE0, TAB, "- iopEs", lvs.iopEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- lTts", lvs.lTts)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.exc = nil
        lvs.l = lvs.w.Links[ctx.LinkNo]
        ctx.Query = lvs.l.Q
        lvs.lTts = lvs.l.GetAbsTts(ctx)
        lvs.lTtl = Min(lvs.l.GetAbsTtl(ctx), lvs.wTtl)
        ctx.Vars.SetIntVal("$$CNT", 0)
        ctx.Vars.SetIntVal("$$CLOCK", CLOCK)
        ctx.RetErr = nil
        ctx.EvalEs = EntryPtrs{}
        lvs.readEs = EntryPtrs{}
        lvs.writeEs = EntryPtrs{}
        lvs.iopEs = EntryPtrs{}
        s.Scheduler = SetLttsSlot(  s.Scheduler, lvs.lTts, ctx.Wid, ctx.Wiid, ctx.LinkNo)
        s.Scheduler = SetLttlSlot(  s.Scheduler, lvs.lTtl, ctx.Wid, ctx.Wiid, ctx.LinkNo)
        
        m.CurrentState = "6"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnX(TRACE0, TAB, "= EvalEs", ctx.EvalEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wiid", ctx.Wiid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        /**/ m.PrintlnX(TRACE0, TAB, "= readEs", lvs.readEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= lTtl", lvs.lTtl)
        /**/ m.PrintlnI(TRACE0, TAB, "= wTtl", lvs.wTtl)
        /**/ m.PrintlnX(TRACE0, TAB, "= iopEs", lvs.iopEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= lTts", lvs.lTts)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 6: CONDITION STATE
    //   - LVars: [lTts]
    // --------------------------------------
    a.AddState("6", "is lTts fulfilled?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- lTts", lvs.lTts)
        
        if (lvs.lTts <= CLOCK) { m.CurrentState = "9" } else { m.CurrentState = "7" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= lTts", lvs.lTts)
        
        return OK
        })

    // --------------------------------------
    // 7: WAIT STATE
    //   - LVars: [lTts]
    // --------------------------------------
    a.AddState("7", "wait until link tts is reached", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- lTts", lvs.lTts)
        
        if ! s.Wait4TimeEvent(m, lvs.lTts, NO_CP) {
            m.CurrentState = "stopped" // for docu
            return STOPPED
        } else {
        m.CurrentState = "6"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= lTts", lvs.lTts)
        } 
    
        return OK
        })

    // --------------------------------------
    // 8: CALL STATE
    //   - GVars: [LinkNo, Wid, RetErr, Wtxid, Pid]
    // --------------------------------------
    a.AddState("8", "undo wtx", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceUndo") 
        theNewAutomaton, m1 := NewAutomaton_SpaceUndo("SpaceUndo", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "62"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        
        return OK
        })

    // --------------------------------------
    // 9: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("9", "source property set on link?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetSource(ctx) != "") { m.CurrentState = "10" } else { m.CurrentState = "12" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 10: ACTION STATE
    //   - LVars: [dwid]
    //   - Aliases:[p, w, l]
    // --------------------------------------
    a.AddState("10", "treat SOURCE: create dynamic wiring", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- dwid", lvs.dwid)
        /**/ m.PrintlnX(TRACE0, TAB, "- p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.dwid = CreateDynamicWiring(m, s, lvs.p, lvs.w, lvs.l)
        
        m.CurrentState = "12"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= dwid", lvs.dwid)
        /**/ m.PrintlnX(TRACE0, TAB, "= p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 11: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("11", "check mandatory", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetMandatory(ctx)) { m.CurrentState = "8" } else { m.CurrentState = "32" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 12: CONDITION STATE
    //   - LVars: [lTtl]
    // --------------------------------------
    a.AddState("12", "lTtl expired?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- lTtl", lvs.lTtl)
        
        if (lvs.lTtl >= CLOCK) { m.CurrentState = "15" } else { m.CurrentState = "58" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= lTtl", lvs.lTtl)
        
        return OK
        })

    // --------------------------------------
    // 13: WAIT STATE
    //   - GVars: [Query]
    //   - LVars: [lTtl, wait4Cid]
    // --------------------------------------
    a.AddState("13", "wait until C1 has changed; @@@check retval of all Wait4 functions", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnI(TRACE0, TAB, "- lTtl", lvs.lTtl)
        /**/ m.PrintlnS(TRACE0, TAB, "- wait4Cid", lvs.wait4Cid)
        
        if !s.Wait4UserEvent(m, lvs.lTtl, NewPMUserEvent(CONTAINER_CHANGE_EVENT, lvs.wait4Cid, ctx.Query.Typ.StringVal), NO_CP) {
            m.CurrentState = "stopped" // for docu
            return STOPPED
        } else {
        m.CurrentState = "12"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnI(TRACE0, TAB, "= lTtl", lvs.lTtl)
        /**/ m.PrintlnS(TRACE0, TAB, "= wait4Cid", lvs.wait4Cid)
        } 
    
        return OK
        })

    // --------------------------------------
    // 14: ACTION STATE
    //   - LVars: [writeEs]
    // --------------------------------------
    a.AddState("14", "create entries according to query", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.writeEs = CreateEntries(m)
        
        m.CurrentState = "21"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })

    // --------------------------------------
    // 15: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("15", "link op is CALL?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Op == CALL) { m.CurrentState = "37" } else { m.CurrentState = "63" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 16: CALL STATE
    //   - GVars: [LinkNo, Wfid, Wid, Query, Vars, RetErr, Wtxid, Pid, RetEs, Cid]
    // --------------------------------------
    a.AddState("16", "call space read", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceRead") 
        theNewAutomaton, m1 := NewAutomaton_SpaceRead("SpaceRead", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        ctx.RetEs = ctx2.(*Context).RetEs
        ctx.Wfid = ctx2.(*Context).Wfid
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "17"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 17: CONDITION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("17", "was space read ok?", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        if (ctx.RetErr == nil) { m.CurrentState = "65" } else { m.CurrentState = "66" }

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 18: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("18", "link source property set?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetSource(ctx) != "") { m.CurrentState = "19" } else { m.CurrentState = "21" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 19: ACTION STATE
    //   - LVars: [readEs, dwid, writeEs, dwMachineNumber]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("19", "treat source property: unwrap entries; cleanup dynamic wiring;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- readEs", lvs.readEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- dwid", lvs.dwid)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- dwMachineNumber", lvs.dwMachineNumber)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.writeEs  = DestUnWrap(m, lvs.readEs)
        CleanUpDynamicWiring(m, s, lvs.l, lvs.dwid,   lvs.dwMachineNumber)
        
        m.CurrentState = "21"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= readEs", lvs.readEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= dwid", lvs.dwid)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= dwMachineNumber", lvs.dwMachineNumber)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 20: ACTION STATE
    //   - LVars: [writeEs]
    // --------------------------------------
    a.AddState("20", "no entries to be written", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.writeEs = EntryPtrs{}
        
        m.CurrentState = "32"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })

    // --------------------------------------
    // 21: ACTION STATE
    //   - GVars: [Vars]
    //   - LVars: [writeEs]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("21", "resolve and set vars (ii): do again - because new entry propos might have been set", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        ctx.Vars = lvs.l.ResolveLinkArgs(ctx.Vars, lvs.writeEs)
        
        m.CurrentState = "48"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 22: ACTION STATE
    //   - LVars: [iopEs, e]
    // --------------------------------------
    a.AddState("22", "add e to write set for IOP", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- iopEs", lvs.iopEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        lvs.iopEs = append(lvs.iopEs, lvs.e)
        
        m.CurrentState = "74"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= iopEs", lvs.iopEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 23: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("23", "no destination set for link", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetDest(ctx) == "") { m.CurrentState = "24" } else { m.CurrentState = "25" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 24: ACTION STATE
    //   - LVars: [es2, writeEs]
    // --------------------------------------
    a.AddState("24", "check dest property of all selected entries to be written and distribute entries accordingly;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.es2 = lvs.writeEs
        lvs.writeEs = EntryPtrs{}
        
        m.CurrentState = "74"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })

    // --------------------------------------
    // 25: ACTION STATE
    //   - LVars: [iopEs, e, writeEs]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("25", "link dest property set: wrap entries", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- iopEs", lvs.iopEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.e = DestWrap(m, lvs.l, lvs.writeEs)
        lvs.iopEs = append(lvs.iopEs, lvs.e)
        lvs.writeEs = EntryPtrs{}
        
        m.CurrentState = "26"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= iopEs", lvs.iopEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 26: ACTION STATE
    //   - GVars: [Es, WMNo, Cid]
    //   - LVars: [writeEs]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("26", "writeEs: apply eprops; ???TBD;set fot ctx of next machine call: - Es to entries to be written;- Cid to C2;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnI(TRACE0, TAB, "- WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        cidptr := lvs.l.ConvertC2toM(ctx.WMNo)
        ctx.Cid = *cidptr
        ctx.Es = lvs.writeEs
        
        m.CurrentState = "67"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnI(TRACE0, TAB, "= WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 27: CONDITION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("27", "check if space write to link&apos;s C2 was OK", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        if (ctx.RetErr == nil) { m.CurrentState = "29" } else { m.CurrentState = "28" }

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 28: ACTION STATE
    //   - GVars: [Es, Cid]
    //   - LVars: [readEs, wait4Cid]
    // --------------------------------------
    a.AddState("28", "set Cid (to wait4Cid, i.e. C1) and Es (to readEs) for undo of read into C1", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- readEs", lvs.readEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- wait4Cid", lvs.wait4Cid)
        
        ctx.Cid = lvs.wait4Cid 
        ctx.Es = lvs.readEs
        
        m.CurrentState = "54"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= readEs", lvs.readEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= wait4Cid", lvs.wait4Cid)
        
        return OK
        })

    // --------------------------------------
    // 29: ACTION STATE
    //   - GVars: [Es, Cid]
    //   - LVars: [iopEs]
    // --------------------------------------
    a.AddState("29", "set ct vars Cid and Es for next machine call;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- iopEs", lvs.iopEs)
        
        ctx.Cid = IOP_PIC
        ctx.Es = lvs.iopEs
        
        m.CurrentState = "68"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= iopEs", lvs.iopEs)
        
        return OK
        })

    // --------------------------------------
    // 30: CONDITION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("30", "check if space write to IOP&apos;s PIC was OK", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        if (ctx.RetErr == nil) { m.CurrentState = "32" } else { m.CurrentState = "64" }

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 31: CALL STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("31", "call space undo write", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceUndoWrite") 
        theNewAutomaton, m1 := NewAutomaton_SpaceUndoWrite("SpaceUndoWrite", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "28"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 32: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("32", "check commit", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetCommit(ctx)) { m.CurrentState = "33" } else { m.CurrentState = "34" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 33: CALL STATE
    //   - GVars: [RetErr, Wtxid]
    // --------------------------------------
    a.AddState("33", "call space tx commit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceTxCommit") 
        theNewAutomaton, m1 := NewAutomaton_SpaceTxCommit("SpaceTxCommit", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "57"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 34: ACTION STATE
    //   - GVars: [LinkNo, Wiid]
    // --------------------------------------
    a.AddState("34", "link done; inform scheduler about link termination;", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wiid", ctx.Wiid)
        
        s.Scheduler = ClearLttsAndLttlSlot(     s.Scheduler, ctx.Wiid, ctx.LinkNo)
        
        m.CurrentState = "56"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wiid", ctx.Wiid)
        
        return OK
        })

    // --------------------------------------
    // 35: ACTION STATE
    //   - GVars: [LinkNo]
    // --------------------------------------
    a.AddState("35", "increment current link number", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        
        ctx.LinkNo = ctx.LinkNo + 1
        
        m.CurrentState = "5"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        
        return OK
        })

    // --------------------------------------
    // 36: CONDITION STATE
    //   - LVars: [repeatCount]
    //   - Aliases:[w]
    // --------------------------------------
    a.AddState("36", "repeat?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- repeatCount", lvs.repeatCount)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        
        if (lvs.repeatCount < lvs.w.GetRepeatCount(ctx)) { m.CurrentState = "3" } else { m.CurrentState = "43" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= repeatCount", lvs.repeatCount)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        
        return OK
        })

    // --------------------------------------
    // 37: ACTION STATE
    //   - Aliases:[w, l]
    // --------------------------------------
    a.AddState("37", "call service;TBD: return check of service call;why continue with state 30?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        CallService(m, s, lvs.w, lvs.l)
        
        m.CurrentState = "30"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 38: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("38", "is link a GUARD?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Type == GUARD) { m.CurrentState = "18" } else { m.CurrentState = "21" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 39: ACTION STATE
    //   - GVars: [Vars]
    //   - LVars: [writeEs]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("39", "CREATE, READ, TAKE: @@@check it; apply eprops to writeEs", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if 0 < len(lvs.writeEs) {
        			lvs.writeEs.EvalAndApply(ctx.Vars, lvs.l.EProps)
        }
        
        m.CurrentState = "42"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 40: CONDITION STATE
    //   - GVars: [LinkNo]
    //   - LVars: [nLinks]
    // --------------------------------------
    a.AddState("40", "was is not yet the last link?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnI(TRACE0, TAB, "- nLinks", lvs.nLinks)
        
        if (lvs.nLinks > (ctx.LinkNo + 1)) { m.CurrentState = "35" } else { m.CurrentState = "72" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnI(TRACE0, TAB, "= nLinks", lvs.nLinks)
        
        return OK
        })

    // --------------------------------------
    // 41: CALL STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("41", "space undo automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceUndo") 
        theNewAutomaton, m1 := NewAutomaton_SpaceUndo("SpaceUndo", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "53"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 42: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("42", "if ACTION -> treat DEST property", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Type == ACTION) { m.CurrentState = "23" } else { m.CurrentState = "26" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 43: CALL STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("43", "space undo automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceUndo") 
        theNewAutomaton, m1 := NewAutomaton_SpaceUndo("SpaceUndo", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "51"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 44: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("44", "read not ok; check mandatory", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetMandatory(ctx)) { m.CurrentState = "13" } else { m.CurrentState = "32" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 45: ACTION STATE
    //   - GVars: [Query, Vars]
    //   - LVars: [res]
    // --------------------------------------
    a.AddState("45", "apply/compute query selector (without entries); caution: query selector it may use only variables and basic values;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnB(TRACE0, TAB, "- res", lvs.res)
        
        lvs.res = ctx.Query.Sel.Apply(ctx.Vars, nil /* no entry! */)
        
        m.CurrentState = "46"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnB(TRACE0, TAB, "= res", lvs.res)
        
        return OK
        })

    // --------------------------------------
    // 46: CONDITION STATE
    //   - LVars: [res]
    // --------------------------------------
    a.AddState("46", "check res of selector application", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnB(TRACE0, TAB, "- res", lvs.res)
        
        if (lvs.res == true) { m.CurrentState = "47" } else { m.CurrentState = "11" }

        // debug: 
        /**/ m.PrintlnB(TRACE0, TAB, "= res", lvs.res)
        
        return OK
        })

    // --------------------------------------
    // 47: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("47", "is operation CREATE?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Op == CREATE) { m.CurrentState = "14" } else { m.CurrentState = "21" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 48: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("48", "no entries to be written?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if ((lvs.l.Op == DELETE) ||  (lvs.l.Op == NOOP) ||  (lvs.l.Op == TEST)) { m.CurrentState = "20" } else { m.CurrentState = "78" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 49: EXIT STATE
    //   - GVars: [LinkNo, Wid, Pid, WMNo]
    // --------------------------------------
    a.AddState("49", "exit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "- WMNo", ctx.WMNo)
        
        // Exit(ctx.Wid, ctx.Pid, Q, ctx.WMNo,  ctx.LinkNo)
        
        m.CurrentState = "exit" // docu
        return EXIT

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "= WMNo", ctx.WMNo)
        
        return OK
        })

    // --------------------------------------
    // 50: ACTION STATE
    //   - GVars: [Wid]
    // --------------------------------------
    a.AddState("50", "inform scheduler -> to clean up wtts and wttl slots", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        
        s.Scheduler = ClearWttsAndWttlSlot(s.Scheduler, ctx.Wid)
        
        m.CurrentState = "49"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        
        return OK
        })

    // --------------------------------------
    // 51: ACTION STATE
    //   - LVars: [exc]
    // --------------------------------------
    a.AddState("51", "raise wiring repeat exception", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        
        lvs.exc = errors.New("WIRING-REPEAT-EXCEPTION")
        
        m.CurrentState = "50"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        
        return OK
        })

    // --------------------------------------
    // 52: WAIT STATE
    // --------------------------------------
    a.AddState("52", "optionally: give up critical section", func(s *Status, m *Machine) StateRetEnum {
        
        // debug: 
        
        if !s.Wait4NoEvent(m, CP) {
            m.CurrentState = "stopped" // for docu
            return STOPPED
        } else {
        m.CurrentState = "1"

        // debug: 
        } 
    
        return OK
        })

    // --------------------------------------
    // 53: ACTION STATE
    //   - LVars: [exc]
    // --------------------------------------
    a.AddState("53", "raise wiring ttl exception", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        
        lvs.exc = errors.New("WIRING-GetAbsTtl(ctx)-EXCEPTION")
        
        m.CurrentState = "75"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        
        return OK
        })

    // --------------------------------------
    // 54: CALL STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("54", "call space undo read", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceUndoRead") 
        theNewAutomaton, m1 := NewAutomaton_SpaceUndoRead("SpaceUndoRead", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "13"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 55: INFO STATE
    //   - LVars: [exc]
    // --------------------------------------
    a.AddState("55", "print exception", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        
        s.SystemInfo(fmt.Sprintf(lvs.exc.Error()))

        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        
        return OK
        })

    // --------------------------------------
    // 56: CONDITION STATE
    //   - LVars: [wTtl]
    // --------------------------------------
    a.AddState("56", "check wiring ttl", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- wTtl", lvs.wTtl)
        
        if (lvs.wTtl < CLOCK) { m.CurrentState = "41" } else { m.CurrentState = "40" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= wTtl", lvs.wTtl)
        
        return OK
        })

    // --------------------------------------
    // 57: CALL STATE
    //   - GVars: [RetErr, Wtxid]
    // --------------------------------------
    a.AddState("57", "call space create tx automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceCreateTx") 
        theNewAutomaton, m1 := NewAutomaton_SpaceCreateTx("SpaceCreateTx", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        ctx.Wtxid = ctx2.(*Context).Wtxid
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "34"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 58: ACTION STATE
    //   - GVars: [LinkNo, Wiid]
    // --------------------------------------
    a.AddState("58", "lTtl has expired; inform scheduler about link termination", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wiid", ctx.Wiid)
        
        s.Scheduler = ClearLttsAndLttlSlot(  s.Scheduler, ctx.Wiid, ctx.LinkNo)
        
        m.CurrentState = "11"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wiid", ctx.Wiid)
        
        return OK
        })

    // --------------------------------------
    // 59: ACTION STATE
    //   - GVars: [LinkNo]
    //   - Aliases:[w]
    // --------------------------------------
    a.AddState("59", "reset current link number;clear all entry collections of the wiring: its WC and all SICs and SOUTCs;", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        
        ctx.LinkNo = 0
        s.MetaContext.(*MetaContext).PeerSpace.ClearEntryCollections(lvs.w, m.Number)
        
        m.CurrentState = "4"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        
        return OK
        })

    // --------------------------------------
    // 60: CALL STATE
    //   - GVars: [Wid, RetErr, Wtxid, Pid]
    // --------------------------------------
    a.AddState("60", "call space create tx", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceCreateTx") 
        theNewAutomaton, m1 := NewAutomaton_SpaceCreateTx("SpaceCreateTx", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        ctx.Wtxid = ctx2.(*Context).Wtxid
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "59"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        
        return OK
        })

    // --------------------------------------
    // 61: ACTION STATE
    //   - GVars: [Wid, Pid]
    //   - LVars: [startTime, nLinks, repeatCount]
    //   - Aliases:[p, w]
    // --------------------------------------
    a.AddState("61", "init variables and create scheduler slot that repeatedly hunts for outdated entries", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "- startTime", lvs.startTime)
        /**/ m.PrintlnI(TRACE0, TAB, "- nLinks", lvs.nLinks)
        /**/ m.PrintlnI(TRACE0, TAB, "- repeatCount", lvs.repeatCount)
        /**/ m.PrintlnX(TRACE0, TAB, "- p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        
        lvs.p = s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid]
        lvs.w = lvs.p.Wirings[ctx.Wid]
        lvs.nLinks = len(lvs.w.Links)
        lvs.repeatCount = 0
        lvs.startTime = CLOCK
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnI(TRACE0, TAB, "= startTime", lvs.startTime)
        /**/ m.PrintlnI(TRACE0, TAB, "= nLinks", lvs.nLinks)
        /**/ m.PrintlnI(TRACE0, TAB, "= repeatCount", lvs.repeatCount)
        /**/ m.PrintlnX(TRACE0, TAB, "= p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        
        return OK
        })

    // --------------------------------------
    // 62: ACTION STATE
    //   - LVars: [exc]
    // --------------------------------------
    a.AddState("62", "raise link ttl exception", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        
        lvs.exc = errors.New(  "LINK-GetAbsTtl(ctx)-EXCEPTION")
        
        m.CurrentState = "75"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        
        return OK
        })

    // --------------------------------------
    // 63: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("63", "link op is CREATE or NOOP?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if ((lvs.l.Op == CREATE)  || (lvs.l.Op == NOOP)) { m.CurrentState = "45" } else { m.CurrentState = "73" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 64: ACTION STATE
    //   - GVars: [Es, WMNo, Cid]
    //   - LVars: [writeEs]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("64", "set Cid (to C2) and Es (to writeEs) for undo of write into C2", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnI(TRACE0, TAB, "- WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        cidptr := lvs.l.ConvertC2toM(ctx.WMNo)
        ctx.Cid = *cidptr
        ctx.Es = lvs.writeEs
        
        m.CurrentState = "31"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnI(TRACE0, TAB, "= WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 65: ACTION STATE
    //   - GVars: [EvalEs, Wfid, RetEs]
    //   - LVars: [readEs, writeEs]
    // --------------------------------------
    a.AddState("65", "set readEs and EvalEs (needed to eval args) to RetEs and copy it also to writeEs; set $$CNT to number of read entries", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- EvalEs", ctx.EvalEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- readEs", lvs.readEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.readEs = ctx.RetEs
        ctx.EvalEs = ctx.RetEs
        lvs.writeEs = ctx.RetEs.Copy()
        ctx.Vars.SetIntVal("$$CNT", len(ctx.RetEs))
        ctx.Vars.SetStringVal("$$FID", ctx.Wfid)
        
        m.CurrentState = "38"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= EvalEs", ctx.EvalEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= readEs", lvs.readEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })

    // --------------------------------------
    // 66: ACTION STATE
    //   - GVars: [LinkNo, Wiid]
    // --------------------------------------
    a.AddState("66", "inform scheduler about link termination", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wiid", ctx.Wiid)
        
        s.Scheduler = ClearLttsAndLttlSlot(s.Scheduler, ctx.Wiid, ctx.LinkNo)
        
        m.CurrentState = "44"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wiid", ctx.Wiid)
        
        return OK
        })

    // --------------------------------------
    // 67: CALL STATE
    //   - GVars: [LinkNo, Wid, RetErr, Wtxid, Pid, Es, Cid]
    // --------------------------------------
    a.AddState("67", "call space write into C2", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceWrite") 
        theNewAutomaton, m1 := NewAutomaton_SpaceWrite("SpaceWrite", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "27"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 68: CALL STATE
    //   - GVars: [LinkNo, Wid, RetErr, Wtxid, Pid, Es, Cid]
    // --------------------------------------
    a.AddState("68", "call space write into IOP PIC", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceWrite") 
        theNewAutomaton, m1 := NewAutomaton_SpaceWrite("SpaceWrite", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "30"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 69: ACTION STATE
    //   - LVars: [e, es2]
    // --------------------------------------
    a.AddState("69", "get and remove first entry", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        
        lvs.e = lvs.es2[0]
        lvs.es2 = append(lvs.es2[:0], lvs.es2[1:]...)
        
        m.CurrentState = "70"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        
        return OK
        })

    // --------------------------------------
    // 70: CONDITION STATE
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("70", "is dest property not set on entry?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e.GetDest() == "") { m.CurrentState = "71" } else { m.CurrentState = "22" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 71: ACTION STATE
    //   - LVars: [e, writeEs]
    // --------------------------------------
    a.AddState("71", "add e to writeEs", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.writeEs = append(lvs.writeEs, lvs.e)
        
        m.CurrentState = "74"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })

    // --------------------------------------
    // 72: CONDITION STATE
    //   - LVars: [exc]
    // --------------------------------------
    a.AddState("72", "wiring instance finished;was there any exception?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- exc", lvs.exc)
        
        if (lvs.exc != nil) { m.CurrentState = "55" } else { m.CurrentState = "36" }

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= exc", lvs.exc)
        
        return OK
        })

    // --------------------------------------
    // 73: ACTION STATE
    //   - GVars: [WMNo, Cid]
    //   - LVars: [wait4Cid]
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("73", "link op is READ, TAKE, TEST or DELETE: set Cid and remember it in wait4Cid", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnS(TRACE0, TAB, "- wait4Cid", lvs.wait4Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        cidptr := lvs.l.ConvertC1toM(ctx.WMNo)
        ctx.Cid = *cidptr
        lvs.wait4Cid = ctx.Cid
        
        m.CurrentState = "16"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= WMNo", ctx.WMNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnS(TRACE0, TAB, "= wait4Cid", lvs.wait4Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 74: CONDITION STATE
    //   - LVars: [es2]
    // --------------------------------------
    a.AddState("74", "still one entry in es2?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        
        if (len(lvs.es2) > 0) { m.CurrentState = "69" } else { m.CurrentState = "26" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        
        return OK
        })

    // --------------------------------------
    // 75: CONDITION STATE
    //   - Aliases:[w]
    // --------------------------------------
    a.AddState("75", "WTX abort -> is on abort action set?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        
        if (lvs.w.GetOnAbort(ctx) == true) { m.CurrentState = "76" } else { m.CurrentState = "36" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        
        return OK
        })

    // --------------------------------------
    // 76: ACTION STATE
    //   - GVars: [Es, Cid]
    //   - Aliases:[p]
    // --------------------------------------
    a.AddState("76", "get PIC; create 1 on abort entry (as 'list');", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- p", lvs.p)
        
        ctx.Cid = lvs.p.Pic
        ctx.Es = CreateOnAbortEntry(m)
        
        m.CurrentState = "77"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= p", lvs.p)
        
        return OK
        })

    // --------------------------------------
    // 77: CALL STATE
    //   - GVars: [LinkNo, Wid, RetErr, Wtxid, Pid, Es, Cid]
    // --------------------------------------
    a.AddState("77", "write exception into PIC", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("SpaceWrite") 
        theNewAutomaton, m1 := NewAutomaton_SpaceWrite("SpaceWrite", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "36"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 78: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("78", "read op?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Op == READ) { m.CurrentState = "79" } else { m.CurrentState = "39" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 79: ACTION STATE
    //   - LVars: [writeEs]
    // --------------------------------------
    a.AddState("79", "exchange all eids by new ones", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- writeEs", lvs.writeEs)
        
        lvs.writeEs.ExchangeEidByNewOne()
        
        m.CurrentState = "39"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= writeEs", lvs.writeEs)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of Wiring AUTOMATON
////////////////////////////////////////
