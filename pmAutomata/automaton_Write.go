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
// Write AUTOMATON:
// --------------------------------------
func NewAutomaton_Write(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        c *Container
        // --------------------------------
        // ordinary variables:
        e *Entry
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
        newLvs.c = GetContainerAlias(theM, s)
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [Es, Cid]
    // --------------------------------------
    a.AddState("init", "Write(ctx.Cid, ctx.Es): Write automaton: write Es to Cid with given transaction", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
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
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - GVars: [Cid]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("1", "init vars; @@@assert that c is not empty", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[ctx.Cid]
        
        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Cid", ctx.Cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 2: ACTION STATE
    //   - GVars: [Es]
    //   - LVars: [e]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("2", "remove first entry from Es and add it to Cid,", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        /**/ m.PrintlnX(TRACE0, TAB, "- e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.e = ctx.Es[0]
        ctx.Es = append(ctx.Es[:0], ctx.Es[1:]...)
        lvs.c.AddEntryPtr(lvs.e)
        
        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        /**/ m.PrintlnX(TRACE0, TAB, "= e", lvs.e)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 3: CONDITION STATE
    //   - GVars: [Es]
    // --------------------------------------
    a.AddState("3", "is there any further entry to be written?", func(s *Status, m *Machine) StateRetEnum {
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "- Es", ctx.Es)
        
        if (len(ctx.Es) > 0) { m.CurrentState = "2" } else { m.CurrentState = "4" }

        // debug: 
        /**/ m.PrintlnX(TRACE0, TAB, "= Es", ctx.Es)
        
        return OK
        })

    // --------------------------------------
    // 4: EXIT STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("4", "exit", func(s *Status, m *Machine) StateRetEnum {
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
// EOF of Write AUTOMATON
////////////////////////////////////////
