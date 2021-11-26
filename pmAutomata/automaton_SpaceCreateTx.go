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
// SpaceCreateTx AUTOMATON:
// --------------------------------------
func NewAutomaton_SpaceCreateTx(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        wtx *Tx
        p *Peer
        w *Wiring
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
    a.AddState("init", "SpaceCreateTx(ctx.Pid, ctx.Wid): SpaceCreateTx automaton", func(s *Status, m *Machine) StateRetEnum {
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
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - GVars: [Wid, Wtxid, Pid]
    //   - Aliases:[p, w, wtx]
    // --------------------------------------
    a.AddState("1", "init variables; @@@assert p, w not empty", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "- p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "- w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.p = s.MetaContext.(*MetaContext).PeerSpace.Peers[ctx.Pid]
        lvs.w = lvs.p.Wirings[ctx.Wid]
        lvs.wtx = new(Tx)
        lvs.wtx.Id = Uuid(fmt.Sprintf("%s_tx", ctx.Wid))
        lvs.wtx.State = RUNNING
        lvs.wtx.Txcc = lvs.w.GetTxcc(ctx)
        s.MetaContext.(*MetaContext).Transactions[lvs.wtx.Id] = lvs.wtx
        ctx.Wtxid = lvs.wtx.Id
        
        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wid", ctx.Wid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= Pid", ctx.Pid)
        /**/ m.PrintlnX(TRACE0, TAB, "= p", lvs.p)
        /**/ m.PrintlnX(TRACE0, TAB, "= w", lvs.w)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 2: ERROR STATE
    // --------------------------------------
    a.AddState("2", "error", func(s *Status, m *Machine) StateRetEnum {
        
        // debug: 
        
        SystemError(fmt.Sprintf("ill. txcc or pid or wid"))

        m.CurrentState = "5"

        // debug: 
        
        return OK
        })

    // --------------------------------------
    // 3: CONDITION STATE
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("3", "switch txcc", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        if (lvs.wtx.Txcc == PCC) { m.CurrentState = "4" } else { m.CurrentState = "2" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 4: CALL STATE
    //   - GVars: [RetErr, Wtxid]
    // --------------------------------------
    a.AddState("4", "call pcc create tx – to init wtx", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("PccCreateTx") 
        theNewAutomaton, m1 := NewAutomaton_PccCreateTx("PccCreateTx", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "5"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 5: EXIT STATE
    //   - GVars: [RetErr, Wtxid]
    // --------------------------------------
    a.AddState("5", "exit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        // Exit(ctx.RetErr, ctx.Wtxid)
        
        m.CurrentState = "exit" // docu
        return EXIT

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of SpaceCreateTx AUTOMATON
////////////////////////////////////////
