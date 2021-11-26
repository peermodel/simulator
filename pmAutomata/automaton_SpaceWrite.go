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
// SpaceWrite AUTOMATON:
// --------------------------------------
func NewAutomaton_SpaceWrite(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        wtx *Tx
        l *Link
        // --------------------------------
        // ordinary variables:
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
        newLvs.l = GetLinkAlias(theM, s)
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [LinkNo, Wid, Wtxid, Pid, Es, Cid]
    // --------------------------------------
    a.AddState("init", "SpaceWrite(ctx.Pid, ctx.Wid, ctx.LinkNo, ctx.Wtxid, ctx.Cid, ctx.Es): SpaceWrite automaton: write Es to Cid; link type determines whether wtx shall be used;", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // dummy code: helps that all imports are needed by every automaton
        s.DummyString = fmt.Sprintf("dummy")
        DummyHelpersFu()
        DummySchedulerFu()
        if ctx.DummyIContextFu().(IContext) == nil {}
        // reset all error vars
        ctx.RetErr = errors.New("")
        ctx.RetErr = nil
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - GVars: [LinkNo, Wid, Wtxid, Pid]
    //   - Aliases:[wtx, l]
    // --------------------------------------
    a.AddState("1", "init vars", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        lvs.l = s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid].  Wirings[ctx.Wid].Links[ctx.LinkNo]
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        
        m.CurrentState = "2"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= LinkNo", ctx.LinkNo)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 2: CONDITION STATE
    //   - Aliases:[l]
    // --------------------------------------
    a.AddState("2", "if link type is ACTION do transactional write; otherwise non-transactional (for GUARD &amp; service)", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- l", lvs.l)
        
        if (lvs.l.Type == ACTION) { m.CurrentState = "4" } else { m.CurrentState = "3" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= l", lvs.l)
        
        return OK
        })

    // --------------------------------------
    // 3: CALL STATE
    //   - GVars: [RetErr, Es, Cid]
    // --------------------------------------
    a.AddState("3", "call non-transactional write", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("Write") 
        theNewAutomaton, m1 := NewAutomaton_Write("Write", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "7"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 4: CONDITION STATE
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("4", "switch txcc", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        if (lvs.wtx.Txcc == PCC) { m.CurrentState = "6" } else { m.CurrentState = "5" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 5: ERROR STATE
    // --------------------------------------
    a.AddState("5", "OCC not yet implemented", func(s *Status, m *Machine) StateRetEnum {
        
        // debug: 
        
        SystemError(fmt.Sprintf("ill. txcc"))

        m.CurrentState = "7"

        // debug: 
        
        return OK
        })

    // --------------------------------------
    // 6: CALL STATE
    //   - GVars: [RetErr, Wtxid, Es, Cid]
    // --------------------------------------
    a.AddState("6", "call transactional write with pessimistic concurrency control", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("PccWrite") 
        theNewAutomaton, m1 := NewAutomaton_PccWrite("PccWrite", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "7"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 7: EXIT STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("7", "exit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // Exit(ctx.RetErr)
        
        m.CurrentState = "exit" // docu
        return EXIT

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of SpaceWrite AUTOMATON
////////////////////////////////////////
