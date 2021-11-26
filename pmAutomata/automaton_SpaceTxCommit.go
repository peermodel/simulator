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
// SpaceTxCommit AUTOMATON:
// --------------------------------------
func NewAutomaton_SpaceTxCommit(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        wtx *Tx
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
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [Wtxid, RetErr]
    // --------------------------------------
    a.AddState("init", "SpaceTxCommit(ctx.Wtxid, ctx.RetErr): SpaceTxCommit automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        // dummy code: helps that all imports are needed by every automaton
        s.DummyString = fmt.Sprintf("dummy")
        DummyHelpersFu()
        DummySchedulerFu()
        if ctx.DummyIContextFu().(IContext) == nil {}
        // reset all error vars
        ctx.RetErr = errors.New("")
        ctx.RetErr = nil
        
        m.CurrentState = "6"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 1: CALL STATE
    //   - GVars: [RetErr, Wtxid]
    // --------------------------------------
    a.AddState("1", "call pcc tx commit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        // create and call new machine: 
        foundAutomaton, foundFlag := s.CheckAutomatonExistence("PccTxCommit") 
        theNewAutomaton, m1 := NewAutomaton_PccTxCommit("PccTxCommit", ! foundFlag, foundAutomaton) 
        if !foundFlag { 
            s.AddAutomaton(theNewAutomaton)
        } 
        ctx2 := m.Context.Copy().(IContext) 
        ctx2 = m1.StartSync(s, ctx2) 
        // copy back returned context variables: 
        ctx.RetErr = ctx2.(*Context).RetErr
        // debug: 
        /**/ m.PrintlnResume()

        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 2: ACTION STATE
    // --------------------------------------
    a.AddState("2", "error state", func(s *Status, m *Machine) StateRetEnum {
        
        // debug: 
        
        SystemError("ill. wtxid or txcc")
        
        m.CurrentState = "9"

        // debug: 
        
        return OK
        })

    // --------------------------------------
    // 3: CONDITION STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("3", "check return value of pcc tx commit", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "- RetErr", ctx.RetErr)
        
        if (ctx.RetErr == nil) { m.CurrentState = "4" } else { m.CurrentState = "5" }

        // debug: 
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 4: ACTION STATE
    //   - GVars: [Wtxid]
    // --------------------------------------
    a.AddState("4", "set tx state to committed", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid].State = COMMITTED
        
        m.CurrentState = "9"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 5: ACTION STATE
    //   - GVars: [Wtxid]
    // --------------------------------------
    a.AddState("5", "set tx state to rolledback", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
        s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid].State = ROLLEDBACK
        
        m.CurrentState = "9"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 6: ACTION STATE
    //   - GVars: [Wtxid]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("6", "init variables", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        
        m.CurrentState = "7"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 7: CONDITION STATE
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("7", "check validity of wtx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        if ((lvs.wtx  == nil) ||(lvs.wtx.Id == "")) { m.CurrentState = "2" } else { m.CurrentState = "8" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 8: CONDITION STATE
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("8", "switch on txcc", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        if (lvs.wtx.Txcc == PCC) { m.CurrentState = "1" } else { m.CurrentState = "2" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 9: EXIT STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("9", "exit", func(s *Status, m *Machine) StateRetEnum {
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
// EOF of SpaceTxCommit AUTOMATON
////////////////////////////////////////
