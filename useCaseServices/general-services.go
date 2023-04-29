//************************************************************
//************************************************************
//** AUTO GENERATED FILE START -- DO NOT CHANGE !!!
//************************************************************
//** Copyright: Eva Maria Kuehn (C) 2021
//** File:      ../examples/_GO-AUTOMATON/src/useCases/Amigoo/Amigoo_One/use-case/use-case_Amigoo_One.go
//** Generated: 2023-04-29 18:01:56
//************************************************************
//************************************************************
// Peer Model Tool Chain
// Copyright (C) 2021 Eva Maria Kuehn
// -----------------------------------------------------------
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
//************************************************************
//************************************************************

//////////////////////////////////////////////////////////////
// System: PMMM Use Case Go Code for Peer Model State Machine
// Author: Eva Maria Kuehn
// Date:   2015; 2021
//////////////////////////////////////////////////////////////

package pmUseCases 

import ( 
	. "github.com/peermodel/simulator/controller" 
	. "github.com/peermodel/simulator/debug" 
	. "github.com/peermodel/simulator/pmModel" 
	. "github.com/peermodel/simulator/scheduler" 
	. "github.com/peermodel/simulator/useCaseServices" 
	"fmt" 
) 

//////////////////////////////////////////////////////////////
// type
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// config: empty for auto generated code, but needed by go automaton
type UseCaseAmigoo_One struct { 
}

//////////////////////////////////////////////////////////////
// constructor
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// alloc new use case
// if use case is manually coded, vars could be shared here by the use case peers 
func NewUseCaseAmigoo_One() *UseCaseAmigoo_One { // <<<<<<
	uc := new(UseCaseAmigoo_One)  // <<<<<<
	return uc
}

///////////////////////////////////////////////////////////
// add meta model for Amigoo_One
///////////////////////////////////////////////////////////

func (uc *UseCaseAmigoo_One) AddMetaModel(ps *PeerSpace) { // <<<<<< 
    
    p := new(Peer)
    w := new(Wiring)
    
    //============================================================
    // PEER User:user#1
    //============================================================
    
    p = NewPeer("user#1")
    
    //------------------------------------------------------------
    // WIRING executeEvent:
    //------------------------------------------------------------
    w = NewWiring("executeEvent")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("state"), EQUAL, SVal("act"))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$kind": SLabel("kind"), "$name": SLabel("name"), "$from": SLabel("from"), "$sMTI": ILabel("socialMediaTimeImpact"), "$wLI": ILabel("workLoadImpact"), "$sLI": ILabel("stressLevelImpact"), "$tVI": ILabel("teamVibeImpact"), "$aWI": ILabel("anyWorriesImpact"), "$pHI": ILabel("physicalHealthImpact"), "$rI": ILabel("restednessImpact"), "$tLI": ILabel("tasksLikingImpact"), "$fMI": ILabel("feelMovingImpact")})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$sMT": XVal(ILabel("socialMediaTime"), ADD, IVar("$sMTI")), "$wL": XVal(ILabel("workLoad"), ADD, IVar("$wLI")), "$sL": XVal(ILabel("stressLevel"), ADD, IVar("$sLI")), "$tV": XVal(ILabel("teamVibe"), ADD, IVar("$tVI")), "$aW": XVal(ILabel("anyWorries"), ADD, IVar("$aWI")), "$pH": XVal(ILabel("physicalHealth"), ADD, IVar("$pHI")), "$r": XVal(ILabel("restedness"), ADD, IVar("$rI")), "$tL": XVal(ILabel("tasksLiking"), ADD, IVar("$tLI")), "$fM": XVal(ILabel("feelMoving"), ADD, IVar("$fMI"))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(1000)})
    //............................................................
    // Guard 4:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(5000)})
    //............................................................
    // Guard 5:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(1000)})
    //............................................................
    // Guard 6:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(5000)})
    //............................................................
    // Guard 7:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(1000)})
    //............................................................
    // Guard 8:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 9:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(1000)})
    //............................................................
    // Guard 10:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(5000)})
    //............................................................
    // Guard 11:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(1000)})
    //............................................................
    // Guard 12:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(5000)})
    //............................................................
    // Guard 13:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(1000)})
    //............................................................
    // Guard 14:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(5000)})
    //............................................................
    // Guard 15:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(1000)})
    //............................................................
    // Guard 16:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(5000)})
    //............................................................
    // Guard 17:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tL": IVal(1000)})
    //............................................................
    // Guard 18:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 19:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(1000)})
    //............................................................
    // Guard 20:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(5000)})
    //............................................................
    // Guard 21:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value")})
    //............................................................
    // Guard 22:
    w.AddGuard("", PIC, NOOP, 
        Query{}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$newStatus": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(IVar("$sMT"), ADD, IVar("$wL")), ADD, IVar("$sL")), ADD, IVar("$tV")), ADD, IVar("$aW")), ADD, IVar("$pH")), ADD, IVar("$r")), ADD, IVar("$tL")), ADD, IVar("$fM")), DIV, IVal(9))})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVar("$sMT"), "workLoad": IVar("$wL"), "stressLevel": IVar("$sL"), "teamVibe": IVar("$tV"), "anyWorries": IVar("$aW"), "physicalHealth": IVar("$pH"), "restedness": IVar("$r"), "tasksLiking": IVar("$tL"), "feelMoving": IVar("$fM")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("incidentLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("incident"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("advice"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyIncidentImpact:
    //------------------------------------------------------------
    w = NewWiring("applyIncidentImpact")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("incident")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("BoringTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": XVal(IVal(0), SUB, IVal(1000)), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": XVal(IVal(0), SUB, IVal(1000)), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Holiday"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Illness"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": XVal(IVal(0), SUB, IVal(1000)), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("InterestingTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Party"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MoneyProblem"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Sport"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: incident impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyAdviceImpact:
    //------------------------------------------------------------
    w = NewWiring("applyAdviceImpact")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("CandlelightDinner"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Cinema"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Dancing"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("DoCrochet"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Jogging"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MakeAGift"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MeetFriends"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Origami"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 9:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("PlayChess"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 10:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Smile"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 11:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Spa"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 12:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Yoga"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 13:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: advice impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, DELETE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVal(3000), "workLoad": IVal(3000), "stressLevel": IVal(3000), "teamVibe": IVal(3000), "anyWorries": IVal(3000), "physicalHealth": IVal(3000), "restedness": IVal(3000), "tasksLiking": IVal(3000), "feelMoving": IVal(3000)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("registration"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"user": SVar("$$PID")}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVal(3000)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING computeUserStatus:
    //------------------------------------------------------------
    w = NewWiring("computeUserStatus")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(ILabel("socialMediaTime"), ADD, ILabel("workLoad")), ADD, ILabel("stressLevel")), ADD, ILabel("teamVibe")), ADD, ILabel("anyWorries")), ADD, ILabel("physicalHealth")), ADD, ILabel("restedness")), ADD, ILabel("tasksLiking")), ADD, ILabel("feelMoving")), DIV, IVal(9))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVar("$value")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING triggerAI:
    //------------------------------------------------------------
    w = NewWiring("triggerAI")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TEST, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(NONE)}, 
        LProps{"flow": BVal(false)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1), 
            Sel: XValP(ILabel("value"), LESS, IVal(3700))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value"), "$fid": SFu(FID_FUNCTION)})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("aiRequest"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"value": IVar("$value"), "time": IVar("$time"), "user": SVar("$$PID"), "fid": SVar("$fid")}, 
        Vars{"$time": IFu(CLOCK_FUNCTION)})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"time": IVar("$time"), "value": IVar("$value"), "fid": SVar("$fid")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING selectAdvice:
    //------------------------------------------------------------
    w = NewWiring("selectAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_SelectAdvice", NewServiceWrapper(SelectAdvice,  "SelectAdvice"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Min: IVal(1), Max: IVal(ALL), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("select")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"flow": BVal(true)}, 
        EProps{}, 
        Vars{"$value": SLabel("value"), "$time": ILabel("time")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING giveAdvice:
    //------------------------------------------------------------
    w = NewWiring("giveAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$fid": SLabel("fid"), "$time": ILabel("time"), "$user": SLabel("user"), "$value": ILabel("value"), "$replyTo": SLabel("replyTo"), "$name": SVal("NONE"), "$from": XVal(SVal(""), CONCAT, SVar("$$PID"))})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, READ, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1)}, 
        LProps{"flow": BVal(false), "ttl": IVal(0), "mandatory": BVal(false)}, 
        EProps{}, 
        Vars{"$name": SLabel("name"), "$from": XVal(XVal(SLabel("from"), CONCAT, SVar("$from")), CONCAT, SVal("/"))})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SVar("$replyTo"), "commit": BVal(true)}, 
        EProps{"kind": SVal("advice"), "name": SVar("$name"), "from": SVar("$from"), "fid": SVar("$fid"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
    //============================================================
    // PEER User:user#2
    //============================================================
    
    p = NewPeer("user#2")
    
    //------------------------------------------------------------
    // WIRING executeEvent:
    //------------------------------------------------------------
    w = NewWiring("executeEvent")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("state"), EQUAL, SVal("act"))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$kind": SLabel("kind"), "$name": SLabel("name"), "$from": SLabel("from"), "$sMTI": ILabel("socialMediaTimeImpact"), "$wLI": ILabel("workLoadImpact"), "$sLI": ILabel("stressLevelImpact"), "$tVI": ILabel("teamVibeImpact"), "$aWI": ILabel("anyWorriesImpact"), "$pHI": ILabel("physicalHealthImpact"), "$rI": ILabel("restednessImpact"), "$tLI": ILabel("tasksLikingImpact"), "$fMI": ILabel("feelMovingImpact")})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$sMT": XVal(ILabel("socialMediaTime"), ADD, IVar("$sMTI")), "$wL": XVal(ILabel("workLoad"), ADD, IVar("$wLI")), "$sL": XVal(ILabel("stressLevel"), ADD, IVar("$sLI")), "$tV": XVal(ILabel("teamVibe"), ADD, IVar("$tVI")), "$aW": XVal(ILabel("anyWorries"), ADD, IVar("$aWI")), "$pH": XVal(ILabel("physicalHealth"), ADD, IVar("$pHI")), "$r": XVal(ILabel("restedness"), ADD, IVar("$rI")), "$tL": XVal(ILabel("tasksLiking"), ADD, IVar("$tLI")), "$fM": XVal(ILabel("feelMoving"), ADD, IVar("$fMI"))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(1000)})
    //............................................................
    // Guard 4:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(5000)})
    //............................................................
    // Guard 5:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(1000)})
    //............................................................
    // Guard 6:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(5000)})
    //............................................................
    // Guard 7:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(1000)})
    //............................................................
    // Guard 8:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 9:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(1000)})
    //............................................................
    // Guard 10:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(5000)})
    //............................................................
    // Guard 11:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(1000)})
    //............................................................
    // Guard 12:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(5000)})
    //............................................................
    // Guard 13:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(1000)})
    //............................................................
    // Guard 14:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(5000)})
    //............................................................
    // Guard 15:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(1000)})
    //............................................................
    // Guard 16:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(5000)})
    //............................................................
    // Guard 17:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tL": IVal(1000)})
    //............................................................
    // Guard 18:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 19:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(1000)})
    //............................................................
    // Guard 20:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(5000)})
    //............................................................
    // Guard 21:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value")})
    //............................................................
    // Guard 22:
    w.AddGuard("", PIC, NOOP, 
        Query{}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$newStatus": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(IVar("$sMT"), ADD, IVar("$wL")), ADD, IVar("$sL")), ADD, IVar("$tV")), ADD, IVar("$aW")), ADD, IVar("$pH")), ADD, IVar("$r")), ADD, IVar("$tL")), ADD, IVar("$fM")), DIV, IVal(9))})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVar("$sMT"), "workLoad": IVar("$wL"), "stressLevel": IVar("$sL"), "teamVibe": IVar("$tV"), "anyWorries": IVar("$aW"), "physicalHealth": IVar("$pH"), "restedness": IVar("$r"), "tasksLiking": IVar("$tL"), "feelMoving": IVar("$fM")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("incidentLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("incident"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("advice"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyIncidentImpact:
    //------------------------------------------------------------
    w = NewWiring("applyIncidentImpact")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("incident")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("BoringTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": XVal(IVal(0), SUB, IVal(1000)), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": XVal(IVal(0), SUB, IVal(1000)), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Holiday"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Illness"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": XVal(IVal(0), SUB, IVal(1000)), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("InterestingTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Party"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MoneyProblem"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Sport"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: incident impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyAdviceImpact:
    //------------------------------------------------------------
    w = NewWiring("applyAdviceImpact")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("CandlelightDinner"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Cinema"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Dancing"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("DoCrochet"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Jogging"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MakeAGift"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MeetFriends"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Origami"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 9:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("PlayChess"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 10:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Smile"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 11:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Spa"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 12:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Yoga"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 13:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: advice impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, DELETE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVal(3000), "workLoad": IVal(3000), "stressLevel": IVal(3000), "teamVibe": IVal(3000), "anyWorries": IVal(3000), "physicalHealth": IVal(3000), "restedness": IVal(3000), "tasksLiking": IVal(3000), "feelMoving": IVal(3000)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("registration"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"user": SVar("$$PID")}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVal(3000)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING computeUserStatus:
    //------------------------------------------------------------
    w = NewWiring("computeUserStatus")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(ILabel("socialMediaTime"), ADD, ILabel("workLoad")), ADD, ILabel("stressLevel")), ADD, ILabel("teamVibe")), ADD, ILabel("anyWorries")), ADD, ILabel("physicalHealth")), ADD, ILabel("restedness")), ADD, ILabel("tasksLiking")), ADD, ILabel("feelMoving")), DIV, IVal(9))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVar("$value")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING triggerAI:
    //------------------------------------------------------------
    w = NewWiring("triggerAI")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TEST, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(NONE)}, 
        LProps{"flow": BVal(false)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1), 
            Sel: XValP(ILabel("value"), LESS, IVal(3700))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value"), "$fid": SFu(FID_FUNCTION)})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("aiRequest"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"value": IVar("$value"), "time": IVar("$time"), "user": SVar("$$PID"), "fid": SVar("$fid")}, 
        Vars{"$time": IFu(CLOCK_FUNCTION)})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"time": IVar("$time"), "value": IVar("$value"), "fid": SVar("$fid")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING selectAdvice:
    //------------------------------------------------------------
    w = NewWiring("selectAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_SelectAdvice", NewServiceWrapper(SelectAdvice,  "SelectAdvice"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Min: IVal(1), Max: IVal(ALL), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("select")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"flow": BVal(true)}, 
        EProps{}, 
        Vars{"$value": SLabel("value"), "$time": ILabel("time")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING giveAdvice:
    //------------------------------------------------------------
    w = NewWiring("giveAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$fid": SLabel("fid"), "$time": ILabel("time"), "$user": SLabel("user"), "$value": ILabel("value"), "$replyTo": SLabel("replyTo"), "$name": SVal("NONE"), "$from": XVal(SVal(""), CONCAT, SVar("$$PID"))})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, READ, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1)}, 
        LProps{"flow": BVal(false), "ttl": IVal(0), "mandatory": BVal(false)}, 
        EProps{}, 
        Vars{"$name": SLabel("name"), "$from": XVal(XVal(SLabel("from"), CONCAT, SVar("$from")), CONCAT, SVal("/"))})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SVar("$replyTo"), "commit": BVal(true)}, 
        EProps{"kind": SVal("advice"), "name": SVar("$name"), "from": SVar("$from"), "fid": SVar("$fid"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
    //============================================================
    // PEER User:user#3
    //============================================================
    
    p = NewPeer("user#3")
    
    //------------------------------------------------------------
    // WIRING executeEvent:
    //------------------------------------------------------------
    w = NewWiring("executeEvent")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("state"), EQUAL, SVal("act"))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$kind": SLabel("kind"), "$name": SLabel("name"), "$from": SLabel("from"), "$sMTI": ILabel("socialMediaTimeImpact"), "$wLI": ILabel("workLoadImpact"), "$sLI": ILabel("stressLevelImpact"), "$tVI": ILabel("teamVibeImpact"), "$aWI": ILabel("anyWorriesImpact"), "$pHI": ILabel("physicalHealthImpact"), "$rI": ILabel("restednessImpact"), "$tLI": ILabel("tasksLikingImpact"), "$fMI": ILabel("feelMovingImpact")})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$sMT": XVal(ILabel("socialMediaTime"), ADD, IVar("$sMTI")), "$wL": XVal(ILabel("workLoad"), ADD, IVar("$wLI")), "$sL": XVal(ILabel("stressLevel"), ADD, IVar("$sLI")), "$tV": XVal(ILabel("teamVibe"), ADD, IVar("$tVI")), "$aW": XVal(ILabel("anyWorries"), ADD, IVar("$aWI")), "$pH": XVal(ILabel("physicalHealth"), ADD, IVar("$pHI")), "$r": XVal(ILabel("restedness"), ADD, IVar("$rI")), "$tL": XVal(ILabel("tasksLiking"), ADD, IVar("$tLI")), "$fM": XVal(ILabel("feelMoving"), ADD, IVar("$fMI"))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(1000)})
    //............................................................
    // Guard 4:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sMT"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sMT": IVal(5000)})
    //............................................................
    // Guard 5:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(1000)})
    //............................................................
    // Guard 6:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$wL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$wL": IVal(5000)})
    //............................................................
    // Guard 7:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(1000)})
    //............................................................
    // Guard 8:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 9:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(1000)})
    //............................................................
    // Guard 10:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tV"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tV": IVal(5000)})
    //............................................................
    // Guard 11:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(1000)})
    //............................................................
    // Guard 12:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$aW"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$aW": IVal(5000)})
    //............................................................
    // Guard 13:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(1000)})
    //............................................................
    // Guard 14:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$pH"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$pH": IVal(5000)})
    //............................................................
    // Guard 15:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(1000)})
    //............................................................
    // Guard 16:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$r"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$r": IVal(5000)})
    //............................................................
    // Guard 17:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$tL"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$tL": IVal(1000)})
    //............................................................
    // Guard 18:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$sL"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$sL": IVal(5000)})
    //............................................................
    // Guard 19:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), LESS, IVal(1000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(1000)})
    //............................................................
    // Guard 20:
    w.AddGuard("", PIC, NOOP, 
        Query{            Sel: XValP(IVar("$fM"), GREATER, IVal(5000))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$fM": IVal(5000)})
    //............................................................
    // Guard 21:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value")})
    //............................................................
    // Guard 22:
    w.AddGuard("", PIC, NOOP, 
        Query{}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{"$newStatus": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(IVar("$sMT"), ADD, IVar("$wL")), ADD, IVar("$sL")), ADD, IVar("$tV")), ADD, IVar("$aW")), ADD, IVar("$pH")), ADD, IVar("$r")), ADD, IVar("$tL")), ADD, IVar("$fM")), DIV, IVal(9))})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVar("$sMT"), "workLoad": IVar("$wL"), "stressLevel": IVar("$sL"), "teamVibe": IVar("$tV"), "anyWorries": IVar("$aW"), "physicalHealth": IVar("$pH"), "restedness": IVar("$r"), "tasksLiking": IVar("$tL"), "feelMoving": IVar("$fM")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("incidentLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("incident"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1), 
            Sel: XValP(SVar("$kind"), EQUAL, SVal("advice"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"name": SVar("$name"), "time": IFu(CLOCK_FUNCTION), "from": SVar("$from"), "oldStatus": IVar("$value"), "newStatus": IVar("$newStatus")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyIncidentImpact:
    //------------------------------------------------------------
    w = NewWiring("applyIncidentImpact")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("incident")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("BoringTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": XVal(IVal(0), SUB, IVal(1000)), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": XVal(IVal(0), SUB, IVal(1000)), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Holiday"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Illness"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": XVal(IVal(0), SUB, IVal(1000)), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": XVal(IVal(0), SUB, IVal(1000)), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("InterestingTask"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": XVal(IVal(0), SUB, IVal(1000))}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Party"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": XVal(IVal(0), SUB, IVal(1000)), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MoneyProblem"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": XVal(IVal(0), SUB, IVal(1000)), "stressLevelImpact": XVal(IVal(0), SUB, IVal(1000)), "teamVibeImpact": IVal(0), "anyWorriesImpact": XVal(IVal(0), SUB, IVal(1000)), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Sport"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: incident impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING applyAdviceImpact:
    //------------------------------------------------------------
    w = NewWiring("applyAdviceImpact")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("apply")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("CandlelightDinner"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Cinema"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Dancing"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 4:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("DoCrochet"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 5:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Jogging"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 6:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MakeAGift"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(0), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 7:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("MeetFriends"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(1000), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 8:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Origami"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(1000), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 9:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("PlayChess"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(1000), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(0)}, 
        Vars{})
    //............................................................
    // Action 10:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Smile"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(0), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(1000), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 11:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Spa"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(0), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(1000), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 12:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("Yoga"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{"state": SVal("act"), "socialMediaTimeImpact": IVal(1000), "workLoadImpact": IVal(0), "stressLevelImpact": IVal(1000), "teamVibeImpact": IVal(0), "anyWorriesImpact": IVal(0), "physicalHealthImpact": IVal(1000), "restednessImpact": IVal(0), "tasksLikingImpact": IVal(0), "feelMovingImpact": IVal(1000)}, 
        Vars{})
    //............................................................
    // Action 13:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0), "commit": BVal(true)}, 
        EProps{"state": SVal("*** error: advice impact undefined")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, DELETE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"socialMediaTime": IVal(3000), "workLoad": IVal(3000), "stressLevel": IVal(3000), "teamVibe": IVal(3000), "anyWorries": IVal(3000), "physicalHealth": IVal(3000), "restedness": IVal(3000), "tasksLiking": IVal(3000), "feelMoving": IVal(3000)}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("registration"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"user": SVar("$$PID")}, 
        Vars{})
    //............................................................
    // Action 3:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVal(3000)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING computeUserStatus:
    //------------------------------------------------------------
    w = NewWiring("computeUserStatus")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("energyBoardUpdated"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("energyBoard"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(XVal(ILabel("socialMediaTime"), ADD, ILabel("workLoad")), ADD, ILabel("stressLevel")), ADD, ILabel("teamVibe")), ADD, ILabel("anyWorries")), ADD, ILabel("physicalHealth")), ADD, ILabel("restedness")), ADD, ILabel("tasksLiking")), ADD, ILabel("feelMoving")), DIV, IVal(9))})
    //............................................................
    // Guard 3:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"value": IVar("$value")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING triggerAI:
    //------------------------------------------------------------
    w = NewWiring("triggerAI")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TEST, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(NONE)}, 
        LProps{"flow": BVal(false)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("userStatus"), Count: IVal(1), 
            Sel: XValP(ILabel("value"), LESS, IVal(3700))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$value": ILabel("value"), "$fid": SFu(FID_FUNCTION)})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("aiRequest"), Count: IVal(1)}, 
        LProps{"dest": SUrl("ai")}, 
        EProps{"value": IVar("$value"), "time": IVar("$time"), "user": SVar("$$PID"), "fid": SVar("$fid")}, 
        Vars{"$time": IFu(CLOCK_FUNCTION)})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"time": IVar("$time"), "value": IVar("$value"), "fid": SVar("$fid")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING selectAdvice:
    //------------------------------------------------------------
    w = NewWiring("selectAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_SelectAdvice", NewServiceWrapper(SelectAdvice,  "SelectAdvice"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Min: IVal(1), Max: IVal(ALL), 
            Sel: XValP(XVal(SLabel("kind"), EQUAL, SVal("advice")), AND, XVal(SLabel("state"), EQUAL, SVal("select")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiStarted"), Count: IVal(1)}, 
        LProps{"flow": BVal(true)}, 
        EProps{}, 
        Vars{"$value": SLabel("value"), "$time": ILabel("time")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_SelectAdvice", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING giveAdvice:
    //------------------------------------------------------------
    w = NewWiring("giveAdvice")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$fid": SLabel("fid"), "$time": ILabel("time"), "$user": SLabel("user"), "$value": ILabel("value"), "$replyTo": SLabel("replyTo"), "$name": SVal("NONE"), "$from": XVal(SVal(""), CONCAT, SVar("$$PID"))})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, READ, 
        Query{Typ: SEtype("adviceLog"), Count: IVal(1)}, 
        LProps{"flow": BVal(false), "ttl": IVal(0), "mandatory": BVal(false)}, 
        EProps{}, 
        Vars{"$name": SLabel("name"), "$from": XVal(XVal(SLabel("from"), CONCAT, SVar("$from")), CONCAT, SVal("/"))})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SVar("$replyTo"), "commit": BVal(true)}, 
        EProps{"kind": SVal("advice"), "name": SVar("$name"), "from": SVar("$from"), "fid": SVar("$fid"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
    //============================================================
    // PEER AI:ai
    //============================================================
    
    p = NewPeer("ai")
    
    //------------------------------------------------------------
    // WIRING phase0_init:
    //------------------------------------------------------------
    w = NewWiring("phase0_init")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiRequest"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$user": SLabel("user"), "$value": ILabel("value"), "$fid": SLabel("fid"), "$time": ILabel("time")})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, READ, 
        Query{Typ: SEtype("registration"), Count: IVal(ALL)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("registration"), Count: IVal(ALL), 
            Sel: XValP(SLabel("user"), NOT_EQUAL, SVar("$user"))}, 
        LProps{}, 
        EProps{"fid": SVar("$fid")}, 
        Vars{"$cnt": IVar("$$CNT")})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"fid": SVar("$fid"), "phase": IVal(1), "user": SVar("$user"), "time": IVar("$time"), "value": IVar("$value"), "nNeighbors": IVar("$cnt"), "k": IVal(0)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING phase1_askNeighbor:
    //------------------------------------------------------------
    w = NewWiring("phase1_askNeighbor")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1), 
            Sel: XValP(XVal(ILabel("phase"), EQUAL, IVal(1)), AND, XVal(ILabel("k"), LESS, ILabel("nNeighbors")))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$sender": SLabel("user"), "$value": ILabel("value"), "$time": ILabel("time"), "$inc": IVal(0), "$dec": IVal(1)})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, TAKE, 
        Query{Typ: SEtype("registration"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$neighbor": SLabel("user")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{"dest": SVar("$neighbor")}, 
        EProps{"user": SVar("$sender"), "value": IVar("$value"), "time": IVar("$time"), "fid": SVar("$$FID"), "replyTo": SVar("$$PID")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"k": XVal(ILabel("k"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING phase2_AskAdviceGenerator:
    //------------------------------------------------------------
    w = NewWiring("phase2_AskAdviceGenerator")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("phase"), EQUAL, IVal(2))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$sender": SLabel("user"), "$value": ILabel("value"), "$time": ILabel("time"), "$fid": SLabel("fid")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{"dest": SUrl("adviceGenerator")}, 
        EProps{"user": SVar("$sender"), "value": IVar("$value"), "time": IVar("$time"), "fid": SVar("$fid"), "replyTo": SVar("$$PID")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"phase": IVal(3), "k": IVal(0)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING phase1_Complete:
    //------------------------------------------------------------
    w = NewWiring("phase1_Complete")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1), 
            Sel: XValP(XVal(ILabel("phase"), EQUAL, IVal(1)), AND, XVal(ILabel("k"), EQUAL, ILabel("nNeighbors")))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"phase": IVal(2)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING phase4_SendAdvices:
    //------------------------------------------------------------
    w = NewWiring("phase4_SendAdvices")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1), 
            Sel: XValP(XVal(ILabel("phase"), EQUAL, IVal(3)), AND, XVal(ILabel("k"), EQUAL, XVal(ILabel("nNeighbors"), ADD, IVal(3))))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$user": SLabel("user"), "$value": ILabel("value"), "$time": ILabel("time")})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(ALL)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$cnt": IVar("$$CNT")})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(ALL)}, 
        LProps{"dest": SVar("$user")}, 
        EProps{"state": SVal("select")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("aiLog"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"user": SVar("$user"), "value": IVar("$value"), "time": IVar("$time"), "nAdvices": IVar("$cnt"), "fid": SVar("$$FID")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING phase3_Wait4AllAdvices:
    //------------------------------------------------------------
    w = NewWiring("phase3_Wait4AllAdvices")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1), 
            Sel: XValP(XVal(ILabel("phase"), EQUAL, IVal(3)), AND, XVal(ILabel("k"), LESS, XVal(ILabel("nNeighbors"), ADD, IVal(3))))}, 
        LProps{}, 
        EProps{}, 
        Vars{"$user": SLabel("user"), "$value": ILabel("value"), "$time": ILabel("time"), "$fid": SLabel("fid")})
    //............................................................
    // Guard 2:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("advice"))}, 
        LProps{"flow": BVal(true)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("name"), NOT_EQUAL, SVal("NONE"))}, 
        LProps{"mandatory": BVal(false), "ttl": IVal(0)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("aiCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"k": XVal(ILabel("k"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
    //============================================================
    // PEER AdviceGenerator:adviceGenerator
    //============================================================
    
    p = NewPeer("adviceGenerator")
    
    //------------------------------------------------------------
    // WIRING distributeEvent#1:
    //------------------------------------------------------------
    w = NewWiring("distributeEvent#1")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(1))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING distributeEvent#2:
    //------------------------------------------------------------
    w = NewWiring("distributeEvent#2")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(2))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING distributeEvent#3:
    //------------------------------------------------------------
    w = NewWiring("distributeEvent#3")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(3))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Yoga:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Yoga")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Yoga"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Yoga"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_MakeAGift:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_MakeAGift")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("MakeAGift"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("MakeAGift"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Smile:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Smile")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Smile"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Smile"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Jogging:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Jogging")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Jogging"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Jogging"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_DoCrochet:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_DoCrochet")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("DoCrochet"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("DoCrochet"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_CandlelightDinner:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_CandlelightDinner")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("CandlelightDinner"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("CandlelightDinner"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_PlayChess:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_PlayChess")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("PlayChess"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("PlayChess"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Spa:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Spa")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Spa"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Spa"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Dancing:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Dancing")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Dancing"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Dancing"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Origami:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Origami")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Origami"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Origami"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_Cinema:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_Cinema")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("Cinema"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Cinema"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateAdvice_MeetFriends:
    //------------------------------------------------------------
    w = NewWiring("generateAdvice_MeetFriends")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("advice"), "from": SVal("/"), "name": SVal("MeetFriends"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("MeetFriends"), "repeat_count": IVal(0)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING cleanUpAdviceGenerationCtrl:
    //------------------------------------------------------------
    w = NewWiring("cleanUpAdviceGenerationCtrl")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), EQUAL, IVal(12))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(ALL), 
            Sel: XValP(SLabel("name"), EQUAL, SVal("NONE"))}, 
        LProps{"flow": BVal(false), "commit": BVal(true)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING sendAdvices:
    //------------------------------------------------------------
    w = NewWiring("sendAdvices")
    //............................................................
    // add service wrapper:
    w.AddServiceWrapper("SID_Watch", NewServiceWrapper(Watch,  "Watch"))
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("adviceRequest"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{"$fid": SLabel("fid"), "$user": SLabel("user"), "$time": ILabel("time"), "$value": ILabel("value"), "$replyTo": SLabel("replyTo")})
    //............................................................
    // Guard 2:
    w.AddGuard("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(3), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("advice"))}, 
        LProps{"flow": BVal(false)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // service wrapper:
    // SIN1:
    w.AddSin(TAKE, Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    // CALL SERVICE:
    w.AddScall("SID_Watch", LProps{}, EProps{}, Vars{})
    // SOUT1:
    w.AddSout(Query{Typ: SVal("*"), Count: IVal(ALL)}, "SID_Watch", LProps{}, EProps{}, Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, READ, 
        Query{Typ: SEtype("event"), Count: IVal(ALL)}, 
        LProps{"dest": SVar("$replyTo")}, 
        EProps{"state": SVal("select"), "fid": SVar("$fid")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", POC, READ, 
        Query{Typ: SEtype("event"), Count: IVal(ALL)}, 
        LProps{"commit": BVal(true)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("adviceGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": IVal(0)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
    //============================================================
    // PEER IncidentGenerator:incidentGenerator
    //============================================================
    
    p = NewPeer("incidentGenerator")
    
    //------------------------------------------------------------
    // WIRING distributeIncident#1:
    //------------------------------------------------------------
    w = NewWiring("distributeIncident#1")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(1))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING distributeIncident#2:
    //------------------------------------------------------------
    w = NewWiring("distributeIncident#2")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(2))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING distributeIncident#3:
    //------------------------------------------------------------
    w = NewWiring("distributeIncident#3")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1), 
            Sel: XValP(SLabel("kind"), EQUAL, SVal("incident"))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", POC, TAKE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{"dest": SArrayVal(DynArrayRef("user", IVal(3))), "commit": BVal(true)}, 
        EProps{"time": IFu(CLOCK_FUNCTION), "state": SVal("apply")}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"tts": IVal(1000)}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateBoringTaskIncident:
    //------------------------------------------------------------
    w = NewWiring("generateBoringTaskIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("BoringTask"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("BoringTask"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generatePartyIncident:
    //------------------------------------------------------------
    w = NewWiring("generatePartyIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("Party"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Party"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateHolidayIncident:
    //------------------------------------------------------------
    w = NewWiring("generateHolidayIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("Holiday"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Holiday"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateMoneyProblemIncident:
    //------------------------------------------------------------
    w = NewWiring("generateMoneyProblemIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("MoneyProblem"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("MoneyProblem"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateSportIncident:
    //------------------------------------------------------------
    w = NewWiring("generateSportIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("Sport"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Sport"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateIllnessIncident:
    //------------------------------------------------------------
    w = NewWiring("generateIllnessIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("Illness"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("Illness"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING generateInterestingTaskIncident:
    //------------------------------------------------------------
    w = NewWiring("generateInterestingTaskIncident")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), LESS, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("event"), Count: IVal(1)}, 
        LProps{}, 
        EProps{"kind": SVal("incident"), "name": SVal("InterestingTask"), "from": SVal("/"), "state": SVal("_")}, 
        Vars{})
    //............................................................
    // Action 2:
    w.AddAction("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": XVal(ILabel("n"), ADD, IVal(1))}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{"name": SVal("InterestingTask"), "repeat_count": XVal(IVal(100), SUB, IVal(1))}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING cleanUpIncidentGenerationCtrl:
    //------------------------------------------------------------
    w = NewWiring("cleanUpIncidentGenerationCtrl")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, TAKE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1), 
            Sel: XValP(ILabel("n"), EQUAL, XVal(IVal(100), MUL, IVal(7)))}, 
        LProps{"commit": BVal(true)}, 
        EProps{}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // WIRING init:
    //------------------------------------------------------------
    w = NewWiring("init")
    //............................................................
    // Guard 1:
    w.AddGuard("", PIC, DELETE, 
        Query{Typ: SEtype("INIT"), Count: IVal(1)}, 
        LProps{}, 
        EProps{}, 
        Vars{})
    //............................................................
    // Action 1:
    w.AddAction("", PIC, CREATE, 
        Query{Typ: SEtype("incidentGenerationCtrl"), Count: IVal(1)}, 
        LProps{"commit": BVal(true)}, 
        EProps{"n": IVal(0)}, 
        Vars{})
    //............................................................
    // set wprops:
    w.WProps = WProps{}
    //............................................................
    // add wiring to peer & resolve names:
    p.AddWiring(w)

    //------------------------------------------------------------
    // add peer to peer space:
    ps.AddPeer(p)
    //------------------------------------------------------------
    
}


//////////////////////////////////////////////////////////////
// Init use case: create and write INIT entry into all peer's PICs
////////////////////////////////////////////////////////////// 

//------------------------------------------------------------
// create & send an INIT entry to all user peers
// - nb: this is a convention introduced by the Peer Model tool-chain
func (uc *UseCaseAmigoo_One) Init(ps *PeerSpace, scheduler *Scheduler, controllerChannel ControllerChannel) {
    //------------------------------------------------------------
	// create INIT entry (without FID)
	initE := NewEntry("INIT")
    //------------------------------------------------------------
	// write it to the PIC of of all user peers
    for _, p := range ps.Peers {
       if p.Id == "IOP" || p.Id == "Stop" {
           continue
       }
       ps.Write(p.Pic, initE, nil /* no vars */, scheduler)
    }
}

//////////////////////////////////////////////////////////////
// Built-in Service "Consume" 
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// print all entries of the wiring container and remove them
func Consume(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, inCid string, outCid string, controllerChannel ControllerChannel) {
    //------------------------------------------------------------
    // print service name and cid of container
    String2TraceFile("\n")
    String2TraceFile(fmt.Sprintf("%s: CONSUME (time = %d):\n", inCid, CLOCK))
    //------------------------------------------------------------
    // take all entries from the wiring container and print them
    for {
        //------------------------------------------------------------
        // take entry
        entry := ps.Take(inCid, "*", nil /* no selector */, vars)
        //------------------------------------------------------------
        // done?
        if nil == entry {
            break
        }
        //------------------------------------------------------------
        // print entry
        /**/ String2TraceFile("  CONSUMED ENTRY = ")
        /**/ entry.Println(0 /* ind */)
    }
}

//////////////////////////////////////////////////////////////
// Built-in Service "Watch" 
//////////////////////////////////////////////////////////////

//------------------------------------------------------------
// print all entries of the wiring container
func Watch(ps *PeerSpace, wfid string, vars Vars, scheduler *Scheduler, inCid string, outCid string, controllerChannel ControllerChannel) {
    //------------------------------------------------------------
    // print service name and cid of container
    String2TraceFile("\n")
    String2TraceFile(fmt.Sprintf("%s: WATCH (time = %d):\n", inCid, CLOCK))
    //------------------------------------------------------------
    // take all entries from the wiring container, print them and emit them back
    for {
        //------------------------------------------------------------
        // take entry
        entry := ps.Take(inCid, "*", nil /* no selector */, vars)
        //------------------------------------------------------------
        // done?
        if nil == entry {
            break
        }
        //------------------------------------------------------------
        // print entry
        /**/ String2TraceFile("    ")
        /**/ entry.Println(0 /* ind */)
        //------------------------------------------------------------
        // emit entry back to wiring container
        ps.Emit(outCid, entry, vars, scheduler)
    }
}


//////////////////////////////////////////////////////////////
// END OF AUTO GENERATED FILE
//////////////////////////////////////////////////////////////


