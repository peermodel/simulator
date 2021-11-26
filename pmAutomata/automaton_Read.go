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
// Read AUTOMATON:
// --------------------------------------
func NewAutomaton_Read(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
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
        es2 EntryPtrs
        min int
        e *Entry
        max int
        cnt int
        qTyp string
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
        if ! (IsPointer(tmpOrigLvs.es2) && nil == tmpOrigLvs.es2) {
            tmpNewLvs.es2 = tmpOrigLvs.es2.Copy()
        }
        if ! (IsPointer(tmpOrigLvs.e) && nil == tmpOrigLvs.e) {
            tmpNewLvs.e = tmpOrigLvs.e.Copy()
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
    //   - GVars: [Wid, Vars, Pid, Cid]
    // --------------------------------------
    a.AddState("init", "Read(ctx.Cid, ctx.Pid, ctx.Vars, ctx.Wid): Read automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
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
        
        m.CurrentState = "13"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 1: CONDITION STATE
    //   - LVars: [max, cnt]
    // --------------------------------------
    a.AddState("1", "max entries found?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        
        if (lvs.max == lvs.cnt) { m.CurrentState = "2" } else { m.CurrentState = "4" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        
        return OK
        })

    // --------------------------------------
    // 2: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("2", "is entry deleted?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if ((lvs.l.Op == TAKE)  ||  (lvs.l.Op == DELETE)) { m.CurrentState = "6" } else { m.CurrentState = "3" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 3: ACTION STATE
    //   - LVars: [e, es2]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("3", "restore selected entries in container", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        for _, lvs.e = range lvs.es2 {
          lvs.c.Entries = append(lvs.c.Entries, *lvs.e)
        }
        
        m.CurrentState = "6"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 4: ACTION STATE
    //   - GVars: [Vars]
    //   - LVars: [e, qTyp]
    //   - Aliases:[c, l]
    // --------------------------------------
    a.AddState("4", "get next entry that fulfills query", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnS(TRACE0, TAB, "- qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.e = lvs.c.SelectEntry(ctx.Vars, lvs.qTyp, lvs.l.Q.Sel)
        
        m.CurrentState = "7"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnS(TRACE0, TAB, "= qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 5: CONDITION STATE
    //   - LVars: [min, cnt]
    // --------------------------------------
    a.AddState("5", "select ALL", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        
        if (lvs.min <= lvs.cnt) { m.CurrentState = "2" } else { m.CurrentState = "11" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        
        return OK
        })

    // --------------------------------------
    // 6: ACTION STATE
    //   - GVars: [RetEs]
    //   - LVars: [es2]
    // --------------------------------------
    a.AddState("6", "read was successful: set return entries", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        
        ctx.RetEs = lvs.es2
        
        m.CurrentState = "14"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        
        return OK
        })

    // --------------------------------------
    // 7: CONDITION STATE
    //   - LVars: [e]
    // --------------------------------------
    a.AddState("7", "entry found?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        
        if (lvs.e != nil) { m.CurrentState = "9" } else { m.CurrentState = "8" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        
        return OK
        })

    // --------------------------------------
    // 8: CONDITION STATE
    //   - LVars: [max]
    // --------------------------------------
    a.AddState("8", "is max == NONE?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        
        if (lvs.max == NONE) { m.CurrentState = "10" } else { m.CurrentState = "5" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        
        return OK
        })

    // --------------------------------------
    // 9: ACTION STATE
    //   - LVars: [es2, e, cnt]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("9", "add e to es; remove e (temporarily) from c; cnt+", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.es2 = append(lvs.es2, lvs.e)
        lvs.c.RemoveEntry(lvs.e.Id)
        lvs.cnt = lvs.cnt + 1
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 10: CONDITION STATE
    //   - LVars: [min]
    // --------------------------------------
    a.AddState("10", "is min > cnt?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        
        if (lvs.min > lvs.min) { m.CurrentState = "12" } 
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        
        return OK
        })

    // --------------------------------------
    // 11: ACTION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("11", "user errror: read failed", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        ctx.RetErr = errors.New("USER: read failed: not enough entries there")
        
        m.CurrentState = "12"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 12: ACTION STATE
    //   - LVars: [e, es2]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("12", "restore container", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        for _, lvs.e = range lvs.es2 {
          lvs.c.Entries = append(lvs.c.Entries, *lvs.e)
        }
        
        m.CurrentState = "14"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 13: ACTION STATE
    //   - GVars: [LinkNo, Wid, Query, Vars, Wtxid, Pid, Cid]
    //   - LVars: [min, es2, max, cnt, qTyp]
    //   - Aliases:[c, wtx, l]
    // --------------------------------------
    a.AddState("13", "init vars", func(s *Status, m *Machine) StateRetEnum {
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
        /**/ m.PrintlnI(TRACE0, TAB, "- min", lvs.min)
        /**/ m.PrintlnX(TRACE0, TAB, "- es2", lvs.es2)
        /**/ m.PrintlnI(TRACE0, TAB, "- max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "- cnt", lvs.cnt)
        /**/ m.PrintlnS(TRACE0, TAB, "- qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.l = s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid].  Wirings[ctx.Wid].Links[ctx.LinkNo]
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[ctx.Cid]
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        lvs.cnt = 0
        lvs.es2 = EntryPtrs{}
        lvs.min = ctx.Query.GetMin(ctx.Vars)
        lvs.max = ctx.Query.GetMax(ctx.Vars)
        lvs.qTyp = ctx.Query.GetTyp(ctx.Vars)
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Query", ctx.Query)
        /**/ m.PrintlnX(TRACE0, TAB, "= Vars", ctx.Vars)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnI(TRACE0, TAB, "= min", lvs.min)
        /**/ m.PrintlnX(TRACE0, TAB, "= es2", lvs.es2)
        /**/ m.PrintlnI(TRACE0, TAB, "= max", lvs.max)
        /**/ m.PrintlnI(TRACE0, TAB, "= cnt", lvs.cnt)
        /**/ m.PrintlnS(TRACE0, TAB, "= qTyp", lvs.qTyp)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 14: EXIT STATE
    //   - GVars: [RetErr, RetEs]
    // --------------------------------------
    a.AddState("14", "exit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "- RetEs", ctx.RetEs)
        
        // Exit(ctx.RetErr, ctx.RetEs)
        
        m.CurrentState = "exit" // docu
        return EXIT

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "= RetEs", ctx.RetEs)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of Read AUTOMATON
////////////////////////////////////////
