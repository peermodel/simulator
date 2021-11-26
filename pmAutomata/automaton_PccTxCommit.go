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
// PccTxCommit AUTOMATON:
// --------------------------------------
func NewAutomaton_PccTxCommit(automatonName string, createAutomatonFlag bool, a *Automaton) (*Automaton, *Machine) {
    // --------------------------------------
    // declare local variables (LVS) interface struct:
    // --------------------------------------
    type localVariables struct {
        // --------------------------------
        // alias variables: point into meta model -> do not deep copy but recompute!
        wtx *Tx
        c *Container
        // --------------------------------
        // ordinary variables:
        cid string
        k int
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
        newLvs.c = GetContainerAlias(theM, s)
        // --------------------------------
        // cast <-:
        return (interface{})(newLvs)
    }


    // --------------------------------------
    // init: INIT STATE
    //   - GVars: [Wtxid, RetErr]
    // --------------------------------------
    a.AddState("init", "PccTxCommit(ctx.Wtxid, ctx.RetErr): PccTxCommit automaton", func(s *Status, m *Machine) StateRetEnum {
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
        
        m.CurrentState = "7"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnY(TRACE0, TAB, "= RetErr", ctx.RetErr)
        
        return OK
        })

    // --------------------------------------
    // 1: ACTION STATE
    //   - LVars: [k, cid]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("1", "does a container exist on which wtx has a lock?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- k", lvs.k)
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.k = len(lvs.wtx.Pcc.LockedCids)
        lvs.cid = ""
        
        m.CurrentState = "8"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= k", lvs.k)
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 2: ACTION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [cid]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("2", "remove all entries in the container that are DELETE-locked by wtx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
        // assert that nil != lvs.c
        tmpEs := Entries{}
        	for _, tmpE := range lvs.c.Entries {
        		if tmpE.DLocks[ctx.Wtxid] > 0 {
        			// found -> do not copy to tmpEs
        			// inform scheduler
        			s.Scheduler = ClearEttsAndEttlSlot(s.Scheduler, tmpE.Id)
        		} else {
        			// not found -> keep entry ie copy it to tmpEs
        			tmpEs = append(tmpEs, tmpE)
        		}
        	}
        lvs.c.Entries = tmpEs
        
        m.CurrentState = "3"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 3: ACTION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [cid]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("3", "on all entries in the container: remove all READ-locks of wtx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
        // assert nil != lvs.c
        tmpEs := Entries{}
        for _, tmpE := range lvs.c.Entries {
        	if tmpE.RLocks[ctx.Wtxid] > 0 {
        		delete(tmpE.RLocks, ctx.Wtxid)
        		tmpEs = append(tmpEs, tmpE)
        	}
        }
        // check if entry"s ttl has expired and wrap it into exception if so, 
        // write the exception into the right poc
        // do it in extra loop, because lvs.c is changed by check ettl
        for _, tmpE := range tmpEs {
        	s.MetaContext.(*MetaContext).PeerSpace.CheckEttl(& tmpE, lvs.c, & s.Scheduler)
        }
        
        m.CurrentState = "4"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 4: ACTION STATE
    //   - GVars: [Wtxid]
    //   - LVars: [cid]
    //   - Aliases:[c]
    // --------------------------------------
    a.AddState("4", "on all entries in the container: remove WRITE-lock of tx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- c", lvs.c)
        
        lvs.c = s.MetaContext.(*MetaContext).PeerSpace.Containers[lvs.cid]
        // assert nil != lvs.c
        for _, tmpE := range lvs.c.Entries {
        	// remove write lock on tmpE: 
        	delete(tmpE.WLocks, ctx.Wtxid)
        	// inform scheduler about new entry - 
         	// scheduler must insert GetAbsTts(ctx) and GetAbsTtl(ctx) slots for it: 
        	s.Scheduler = SetEttsAndEttlSlot(s.Scheduler, tmpE.Id, tmpE.GetTts(), tmpE.GetTtl())
        }
        
        m.CurrentState = "5"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= c", lvs.c)
        
        return OK
        })

    // --------------------------------------
    // 5: ACTION STATE
    //   - LVars: [cid]
    // --------------------------------------
    a.AddState("5", "signal container change event", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        
        s.MetaContext.(*MetaContext).PeerSpace.ContainerChangeEvent(lvs.cid)
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        
        return OK
        })

    // --------------------------------------
    // 6: EXIT STATE
    //   - GVars: [RetErr]
    // --------------------------------------
    a.AddState("6", "exit", func(s *Status, m *Machine) StateRetEnum {
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

    // --------------------------------------
    // 7: ACTION STATE
    //   - GVars: [Wtxid]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("7", "init variables", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        ctx := m.Context.(*Context)
        
        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "- Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.wtx = s.MetaContext.(*MetaContext).Transactions[ctx.Wtxid]
        
        m.CurrentState = "1"

        // debug: 
        /**/ m.PrintlnS(TRACE0, TAB, "= Wtxid", ctx.Wtxid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })

    // --------------------------------------
    // 8: CONDITION STATE
    //   - LVars: [k]
    // --------------------------------------
    a.AddState("8", "are there any containers on which wtx possesses a lock?", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- k", lvs.k)
        
        if (lvs.k > 0) { m.CurrentState = "9" } else { m.CurrentState = "6" }

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= k", lvs.k)
        
        return OK
        })

    // --------------------------------------
    // 9: ACTION STATE
    //   - LVars: [k, cid]
    //   - Aliases:[wtx]
    // --------------------------------------
    a.AddState("9", "yes: remove cid from locked cids of wtx", func(s *Status, m *Machine) StateRetEnum {
        lvs := m.LocalVariables.(*localVariables)
        
        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "- k", lvs.k)
        /**/ m.PrintlnS(TRACE0, TAB, "- cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "- wtx", lvs.wtx)
        
        lvs.cid = lvs.wtx.Pcc.LockedCids[0]
        // shrink the slice
        lvs.wtx.Pcc.LockedCids = lvs.wtx.Pcc.LockedCids[1:lvs.k]
        
        m.CurrentState = "2"

        // debug: 
        /**/ m.PrintlnI(TRACE0, TAB, "= k", lvs.k)
        /**/ m.PrintlnS(TRACE0, TAB, "= cid", lvs.cid)
        /**/ m.PrintlnX(TRACE0, TAB, "= wtx", lvs.wtx)
        
        return OK
        })
    }

    return a, m
}

////////////////////////////////////////
// EOF of PccTxCommit AUTOMATON
////////////////////////////////////////
