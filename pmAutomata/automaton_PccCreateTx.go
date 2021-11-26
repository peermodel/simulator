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
// PccCreateTx AUTOMATON:
// --------------------------------------
func NewAutomaton_PccCreateTx(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
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
    //   - GVars: [Wtxid]
    // --------------------------------------
    a.AddState("init", "PccCreateTx(ctx.Wtxid): PccCreateTx automaton", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        
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
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - GVars: [Wtxid]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("1", "init pcc; check wtx not empty", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        lvs.wtx.Pcc.LockedCids = Strings{}
        
        m.CurrentState = "2"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 2: EXIT STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("2", "exit", func(s *Status, m *Machine) StateRetEnum {
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
// EOF of PccCreateTx AUTOMATON
////////////////////////////////////////
