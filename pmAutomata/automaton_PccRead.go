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
    . "cca/contextInterface"
    . "cca/debug"
    . "cca/helpers"
    . "cca/scheduler"
    . "cca/pmModel"
    . "cca/framework"
)

// --------------------------------------
// PccRead AUTOMATON:
// --------------------------------------
func NewAutomaton_PccRead(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        wtx *Tx
        c *Container
        l *Link
        // --------------------------------
        // ordinary variables:
        notokEs EntryPtrs
        okEs EntryPtrs
        min int
        e *Entry
        max int
        cnt int
        e1 *Entry
        qTyp string
        okLockedEs EntryPtrs
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
        if ! (IsPointer(tmpOrigLvs.notokEs) && nil == tmpOrigLvs.notokEs) {
            tmpNewLvs.notokEs = tmpOrigLvs.notokEs.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.okEs) && nil == tmpOrigLvs.okEs) {
            tmpNewLvs.okEs = tmpOrigLvs.okEs.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.e) && nil == tmpOrigLvs.e) {
            tmpNewLvs.e = tmpOrigLvs.e.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.e1) && nil == tmpOrigLvs.e1) {
            tmpNewLvs.e1 = tmpOrigLvs.e1.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.okLockedEs) && nil == tmpOrigLvs.okLockedEs) {
            tmpNewLvs.okLockedEs = tmpOrigLvs.okLockedEs.Copy()
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
        newLvs.wtx = GetWtxAlias(theM, s)
        newLvs.c = GetContainerAlias(theM, s)
        newLvs.l = GetLinkAlias(theM, s)
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [LinkNo, Wid, Wfid, Query, Vars, Wtxid, Pid, Cid]
    // --------------------------------------
    a.AddState("init", "PccRead(ctx.Cid, ctx.LinkNo, ctx.Query, ctx.Pid, ctx.Wid, ctx.Wfid, ctx.Wtxid, ctx.Vars): PccRead automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // dummy code: helps that all imports are needed by every automaton
        s.DummyString = fmt.Sprintf("dummy")
        DummyHelpersFu()
        DummySchedulerFu()
        if ctx.DummyIContextFu().(IContext) == nil {}
        // reset all error vars
        ctx.RetErr = errors.New("")
        ctx.RetErr = nil
        
        m.CurrentState = "25"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 1: CONDITION STATE
    //   - LVars: [max, cnt]
    // --------------------------------------
    a.AddState("1", "is count fulfilled?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        
        if (lvs.max == lvs.cnt) { m.CurrentState = "2" } else { m.CurrentState = "11" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        
        return OK
        })

    // --------------------------------------
    // 2: ACTION STATE
    //   - GVars: [RetErr, RetEs, Cid]
    //   - LVars: [okEs]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("2", "count fulfilled", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.wtx.Pcc.LockedCids =   append(lvs.wtx.Pcc.LockedCids, ctx.Cid)
        ctx.RetErr = nil
        ctx.RetEs = lvs.okEs.CopyAndStripLocks()
        
        m.CurrentState = "18"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 3: CONDITION STATE
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("3", "no entry found?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e == nil) { m.CurrentState = "4" } else { m.CurrentState = "5" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 4: CONDITION STATE
    //   - LVars: [max]
    // --------------------------------------
    a.AddState("4", "no further entry found; check if max == NONE?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        
        if (lvs.max == NONE) { m.CurrentState = "21" } else { m.CurrentState = "20" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        
        return OK
        })

    // --------------------------------------
    // 5: ACTION STATE
    //   - LVars: [e]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("5", "remove selected entry from container", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.c.RemoveEntry(lvs.e.Id)
        
        m.CurrentState = "6"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 6: CONDITION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("6", "is entry DELETE- or WRITE-locked by another tx?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e.WriteLockedByOtherTxOrDeleteLocked(ctx.Wtxid)) { m.CurrentState = "14" } else { m.CurrentState = "7" }

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 7: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("7", "is entry deleted?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if ((lvs.l.Op == TAKE)  ||  (lvs.l.Op == DELETE)) { m.CurrentState = "9" } else { m.CurrentState = "8" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 8: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("8", "check flow property of link", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.GetFlow(ctx)) { m.CurrentState = "10" } else { m.CurrentState = "17" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 9: CONDITION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("9", "delete: check if R-locked by another tx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e.ReadLockedByOtherTx(ctx.Wtxid)) { m.CurrentState = "14" } else { m.CurrentState = "8" }

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 10: CONDITION STATE
    //   - GVars: [Wfid]
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("10", "(fid of wiring) == (fid of entry)?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e.GetFid()== ctx.Wfid) { m.CurrentState = "17" } else { m.CurrentState = "12" }

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 11: ACTION STATE
    //   - GVars: [Query, Vars]
    //   - LVars: [e, qTyp]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("11", "select next entry; caution: use query from machine&apos;s context and not from link, because of source property treatment!", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnS(TRACE0, TAB, "- qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        ctx.Query.Typ.Eval(ctx.Vars, nil /* entry */)
        lvs.qTyp = ctx.Query.Typ.StringVal
        lvs.e = lvs.c.SelectEntry(ctx.Vars, lvs.qTyp, ctx.Query.Sel)
        
        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnS(TRACE0, TAB, "= qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 12: CONDITION STATE
    //   - GVars: [Wfid]
    // --------------------------------------
    a.AddState("12", "is fid of wiring empty?", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        
        if (ctx.Wfid == "") { m.CurrentState = "13" } else { m.CurrentState = "15" }

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        
        return OK
        })

    // --------------------------------------
    // 13: ACTION STATE
    //   - GVars: [Wfid]
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("13", "wiring fid is empty", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        ctx.Wfid = lvs.e.GetFid()
        
        m.CurrentState = "17"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 14: ACTION STATE
    //   - LVars: [notokEs, e]
    // --------------------------------------
    a.AddState("14", "add e to not ok entry collection", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- notokEs", lvs.notokEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        lvs.notokEs = append(lvs.notokEs, lvs.e)
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= notokEs", lvs.notokEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 15: CONDITION STATE
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("15", "wiring fid is not empty: check if entry fid is empty", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e.GetFid() == "") { m.CurrentState = "17" } else { m.CurrentState = "14" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 16: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("16", "delete?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if ((lvs.l.Op == TAKE)  ||  (lvs.l.Op == DELETE)) { m.CurrentState = "19" } else { m.CurrentState = "1" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 17: ACTION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [e, okLockedEs, cnt, e1, okEs]
    // --------------------------------------
    a.AddState("17", "READ-lock entry; okEs = okEs + entry; cnt++", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "- e1", lvs.e1)
        /**/ m.PrintlnX(TRACE0, TAB, "- okEs", lvs.okEs)
        
        lvs.okEs = append(lvs.okEs, lvs.e)
        lvs.e1 = lvs.e.Copy()
        lvs.e1.Locks.AddLock(READ, ctx.Wtxid)
        lvs.okLockedEs = append(lvs.okLockedEs, lvs.e1)
        lvs.cnt = lvs.cnt + 1
        
        m.CurrentState = "16"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "= e1", lvs.e1)
        /**/ m.PrintlnX(TRACE0, TAB, "= okEs", lvs.okEs)
        
        return OK
        })

    // --------------------------------------
    // 18: ACTION STATE
    //   - LVars: [e, okLockedEs]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("18", "the end: restore c with okLockedEs", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        if 0 < len(lvs.okLockedEs) {
          for _, lvs.e = range lvs.okLockedEs {
            lvs.c.AddEntryPtr(lvs.e)
          }}
        
        m.CurrentState = "23"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 19: ACTION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [e1]
    // --------------------------------------
    a.AddState("19", "set DELETE-lock on entry e1", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- e1", lvs.e1)
        
        lvs.e1.AddLock(DELETE, ctx.Wtxid)
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= e1", lvs.e1)
        
        return OK
        })

    // --------------------------------------
    // 20: CONDITION STATE
    //   - LVars: [min, cnt]
    // --------------------------------------
    a.AddState("20", "select ALL", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        
        if (lvs.min > lvs.cnt) { m.CurrentState = "22" } else { m.CurrentState = "2" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        
        return OK
        })

    // --------------------------------------
    // 21: CONDITION STATE
    //   - LVars: [min, cnt]
    // --------------------------------------
    a.AddState("21", "check NONE", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        
        if (lvs.min > lvs.cnt) { m.CurrentState = "24" } else { m.CurrentState = "22" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        
        return OK
        })

    // --------------------------------------
    // 22: ACTION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("22", "error: query could not be fulfilled", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        ctx.RetErr = errors.New("USER: not enough entries satisfying query")
        m.CurrentState = "24"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 23: ACTION STATE
    //   - LVars: [notokEs, e]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("23", "restore c with notokEs", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- notokEs", lvs.notokEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        if 0 < len(lvs.notokEs) {
          for _, lvs.e = range lvs.notokEs {
            lvs.c.AddEntryPtr(lvs.e)
          }
        }
        
        m.CurrentState = "26"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= notokEs", lvs.notokEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 24: ACTION STATE
    //   - LVars: [e, okEs]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("24", "restore c with okEs", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        for _, lvs.e = range lvs.okEs {
          lvs.c.AddEntryPtr(lvs.e)
        }
        
        m.CurrentState = "23"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 25: ACTION STATE
    //   - GVars: [LinkNo, Wid, Query, Vars, Wtxid, Pid, Cid]
    //   - LVars: [notokEs, min, max, okLockedEs, cnt, okEs]
    //   - Aliases:[c, wtx, l]
    // --------------------------------------
    a.AddState("25", "init vars", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- notokEs", lvs.notokEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        /**/ m.PrintlnX(TRACE0, TAB, "- okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "- okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.l = s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid].  Wirings[ctx.Wid].Links[ctx.LinkNo]
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[ctx.Cid]
        lvs.cnt = 0
        lvs.okEs = EntryPtrs{}
        lvs.okLockedEs = EntryPtrs{}
        lvs.notokEs = EntryPtrs{}
        lvs.min = ctx.Query.GetMin(ctx.Vars)
        lvs.max = ctx.Query.GetMax(ctx.Vars)
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= notokEs", lvs.notokEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        /**/ m.PrintlnX(TRACE0, TAB, "= okLockedEs", lvs.okLockedEs)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "= okEs", lvs.okEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 26: EXIT STATE
    //   - GVars: [Wfid, RetErr, RetEs]
    // --------------------------------------
    a.AddState("26", "exit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wfid", ctx.Wfid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        
        // Exit(ctx.RetErr, ctx.RetEs, ctx.Wfid)
        
        m.CurrentState = "exit" // docu
        return EXIT

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wfid", ctx.Wfid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of PccRead AUTOMATON
////////////////////////////////////////
