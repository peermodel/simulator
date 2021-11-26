//////////////////////////////////////////////////////////////
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
// Author: eva Kühn
// Date: 2015
////////////////////////////////////////

package pmModel

import (
	. "github.com/peermodel/simulator/debug"
	"fmt"
	"strings"
)

// key = label or variable name
type Args map[string]Arg

// aliases of type Args
type LProps Args
type EProps Args
type WProps Args
type Vars Args

////////////////////////////////////////
// interface
////////////////////////////////////////

type IArgs interface {
	// @@@ fkt so nicht
	// Copy() Args
	// set:
	SetIntVal(label string, val int)
	SetStringVal(label string, val string)
	SetStringEtype(label string, varName string)
	SetStringUrl(label string, varName string)
	SetBoolVal(label string, val bool)
	SetIntVar(label string, varName string)
	SetStringVar(label string, varName string)
	SetBoolVar(label string, varName string)
	SetIntLabel(label string, labelName string)
	SetStringLabel(label string, labelName string)
	SetBoolLabel(label string, labelName string)
	// get:
	GetIntVal(label string) int
	GetStringVal(label string) string
	GetBoolVal(label string) bool
	IsEmpty() bool
	// Debug
	ToString(tab int) string
	ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool) string
	//Print(tab int, text string, detailsFlag bool, printTypeFlag bool)
	Print(tab int)
	//Println(tab int, text string, detailsFlag bool, printTypeFlag bool)
	Println(tab int)
}

////////////////////////////////////////
// methods
////////////////////////////////////////

//------------------------------------------------------------
// deep copy
func (args Args) Copy() Args {
	//------------------------------------------------------------
	// alloc
	newArgs := map[string]Arg{}
	//------------------------------------------------------------
	// copy props
	for label, value := range args {
		newArgs[label] = value
	}
	//------------------------------------------------------------
	// return
	return newArgs
}

// ----------------------------------------
// set functions:
// - add a new property
// - if exists already: overwrite the old property
// ----------------------------------------

// ----------------------------------------
func (args *Args) SetIntVal(label string, val int) {
	(*args)[label] = IVal(val)
}
func (args *Args) SetStringVal(label string, val string) {
	(*args)[label] = SVal(val)
}
func (args *Args) SetStringEtype(label string, val string) {
	(*args)[label] = SEtype(val)
}
func (args *Args) SetStringUrl(label string, val string) {
	(*args)[label] = SUrl(val)
}
func (args *Args) SetBoolVal(label string, val bool) {
	(*args)[label] = BVal(val)
}
func (args *Args) SetIntVar(label string, varName string) {
	(*args)[label] = IVar(varName)
}
func (args *Args) SetStringVar(label string, varName string) {
	(*args)[label] = SVar(varName)
}
func (args *Args) SetBoolVar(label string, varName string) {
	(*args)[label] = BVar(varName)
}
func (args *Args) SetIntLabel(label string, labelName string) {
	(*args)[label] = ILabel(labelName)
}
func (args *Args) SetStringLabel(label string, labelName string) {
	(*args)[label] = SLabel(labelName)
}
func (args *Args) SetBoolLabel(label string, labelName string) {
	(*args)[label] = BLabel(labelName)
}

// ----------------------------------------
// get functions:
// - they return default values for the data type if arg does not exist
// ----------------------------------------

// ----------------------------------------
func (args Args) GetIntVal(label string) int {
	a := args[label]
	if "" != a.Kind {
		return a.IntVal
	} else {
		return 0 // default value - if not set
	}
}

// ----------------------------------------
func (args Args) GetStringVal(label string) string {
	a := args[label]
	if "" != a.Kind {
		return a.StringVal
	} else {
		return ""
	}
}

// ----------------------------------------
func (args Args) GetBoolVal(label string) bool {
	a := args[label]
	if "" != a.Kind {
		return a.BoolVal
	} else {
		return false
	}
}

// ----------------------------------------
func (args Args) String() string {
	s := ""
	sep := ""
	for label, arg := range args {
		s = fmt.Sprintf("%s%s%s:%s", s, sep, strings.Replace(label, "$", "\\$", 2), arg)
		// s = fmt.Sprintf("%s%s%s:%s", s, sep, label, arg)
		sep = "; "
	}
	return s
}

////////////////////////////////////////
// empty test
////////////////////////////////////////

// ----------------------------------------
func (args Args) IsEmpty() bool {
	if nil == args || len(args) == 0 {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////
// debug
////////////////////////////////////////

// ----------------------------------------
// sort args by their key in the map
// trick: if text is provided: then print text = { ... args ... }
//        else print only arg list without brackets starting with ", "
// omit defaults == true means that properties set to default values are skipped; for all kinds of props the same rule!
func (args Args) ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	var sep string
	// build up return string here
	s := NBlanksToString("", tab)

	if "" != text {
		sep = fmt.Sprintf("%s={", text)
	} else {
		sep = ""
	}
	i := 0
	firstPropFlag := true
	sep = ""

	// sort map and build up a sorted array with the keys
	n := len(args)
	sortedLabels := make([]string, n /* len */, n /* capacity */)

	for label, _ := range args {
		// put label on right place in up-sorted map
		for m := 0; m < n; m++ {
			// the end found -> insert here
			if len(sortedLabels[m]) == 0 {
				sortedLabels[m] = label
				break
			}
			// label is larger: right place found  -> insert at place m and shift rest to the right
			if strings.Compare(sortedLabels[m], label) > 0 {
				tmp := make([]string, n /* len */, n /* capacity */)
				copy(tmp[0:], sortedLabels[0:m])
				tmp[m] = label
				copy(tmp[m+1:], sortedLabels[m:n])
				sortedLabels = tmp
				break
			}
			// label is smaller: continue search; nb equal is not possible, because of map
		}
	}

	// print args
	for m := 0; m < n; m++ {
		// get next label to be printed
		label := sortedLabels[m]
		arg := args[label]

		// skip TYPE prop of entry
		if printTypeFlag || TYPE != label {
			tmpS := arg.ToStringWithDetails(0, "", detailsFlag)
			if omitDefaultsFlag {
				// check for default values @@@ not complete... ergänzen nach Bedarf!
				// @@@ passt nicht für eprops ausgabe. zb wenn ttl explicit auf INFINITE gesetzt wird, die vorher anders war -> stimmt nicht! ACHTUNG
				// if (0 == strings.Compare(label, TTL)) && ((0 == strings.Compare(tmpS, "INFINITE")) || (arg.IntVal >= SYSTEM_TTL)) {
				if (0 == strings.Compare(label, TTL)) && (0 == strings.Compare(tmpS, "INFINITE")) {
				} else {
					if (0 == strings.Compare(label, REPEAT_COUNT)) && (0 == strings.Compare(tmpS, "INFINITE")) {
					} else {
						if (0 == strings.Compare(label, TXCC)) && (0 == strings.Compare(tmpS, "pcc")) {
						} else {
							s = fmt.Sprintf("%s%s%s=%s", s, sep, label, tmpS)
							sep = ", "
						}
					}
				}
			} else {
				// display prop also if it is set to a default value;
				// - add it at the end of s;
				s = fmt.Sprintf("%s%s%s=%s", s, sep, label, tmpS)
				if firstPropFlag {
					// change seperator to comma
					firstPropFlag = false
					sep = ", "
				}
			}
		}
		i++
	}
	if 0 < i && "" != text {
		s = fmt.Sprintf("%s}", s)
	}
	return s
}

// ----------------------------------------
// KEEP the orig version which does not sort
// trick: if text is provided: then print text = { ... args ... }
//        else print only arg list without brackets starting with ", "
// omit defaults == true means that properties set to default values are skipped; for all kinds of props the same rule!
func (args Args) ToStringWithDetails_Unsorted__Orig(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	var sep string

	s := NBlanksToString("", tab)

	if "" != text {
		sep = fmt.Sprintf("%s={", text)
	} else {
		sep = ""
	}
	i := 0
	firstPropFlag := true
	sep = ""
	for label, arg := range args {
		// nb: TYPE means entry type property
		if printTypeFlag || TYPE != label {
			tmpS := arg.ToStringWithDetails(0, "", detailsFlag)
			if omitDefaultsFlag {
				// check for default values @@@ not complete... ergänzen nach Bedarf!
				// @@@ passt nicht für eprops ausgabe. zb wenn ttl explicit auf INFINITE gesetzt wird, die vorher anders war -> stimmt nicht! ACHTUNG
				// if (0 == strings.Compare(label, TTL)) && ((0 == strings.Compare(tmpS, "INFINITE")) || (arg.IntVal >= SYSTEM_TTL)) {
				if (0 == strings.Compare(label, TTL)) && (0 == strings.Compare(tmpS, "INFINITE")) {
				} else {
					if (0 == strings.Compare(label, REPEAT_COUNT)) && (0 == strings.Compare(tmpS, "INFINITE")) {
					} else {
						if (0 == strings.Compare(label, TXCC)) && (0 == strings.Compare(tmpS, "pcc")) {
						} else {
							s = fmt.Sprintf("%s%s%s=%s", s, sep, label, tmpS)
							sep = ", "
						}
					}
				}
			} else {
				s = fmt.Sprintf("%s%s%s=%s", s, sep, label, tmpS)
				if firstPropFlag {
					firstPropFlag = false
					sep = ", "
				}
			}
		}
		i++
	}
	if 0 < i && "" != text {
		s = fmt.Sprintf("%s}", s)
	}
	return s
}

// ----------------------------------------
// use defaults for flags
func (args Args) ToString(tab int) string {
	return args.ToStringWithDetails(tab, "", false /* detailsFlag */, true /* printTypeFlag */, true /* omitDefaultsFlag */)
}

// ----------------------------------------
func (args Args) Print(tab int) {
	/**/ String2TraceFile(args.ToStringWithDetails(tab, PRINT_ARGS_NAME, PRINT_ARG_DETAILS_FLAG, PRINT_ARG_TYPE_FLAG, PRINT_OMIT_DEFAULTS_FLAG))
}

// ----------------------------------------
func (args Args) Println(tab int) {
	/**/ args.Print(tab)
	/**/ String2TraceFile("\n")
}

////////////////////////////////////////
// LPROPS
////////////////////////////////////////
func (args LProps) Copy() LProps {
	copiedArgs := (Args)(args).Copy()
	return (LProps)(copiedArgs)
}
func (args *LProps) SetIntVal(label string, val int) {
	(*Args)(args).SetIntVal(label, val)
}
func (args *LProps) SetStringVal(label string, val string) {
	(*Args)(args).SetStringVal(label, val)
}
func (args *LProps) SetStringEtype(label string, val string) {
	(*Args)(args).SetStringEtype(label, val)
}
func (args *LProps) SetStringUrl(label string, val string) {
	(*Args)(args).SetStringUrl(label, val)
}
func (args *LProps) SetBoolVal(label string, val bool) {
	(*Args)(args).SetBoolVal(label, val)
}
func (args *LProps) SetIntVar(label string, varName string) {
	(*Args)(args).SetIntVar(label, varName)
}
func (args *LProps) SetStringVar(label string, varName string) {
	(*Args)(args).SetStringVar(label, varName)
}
func (args *LProps) SetBoolVar(label string, varName string) {
	(*Args)(args).SetBoolVar(label, varName)
}
func (args *LProps) SetIntLabel(label string, labelName string) {
	(*Args)(args).SetIntLabel(label, labelName)
}
func (args *LProps) SetStringLabel(label string, labelName string) {
	(*Args)(args).SetStringLabel(label, labelName)
}
func (args *LProps) SetBoolLabel(label string, labelName string) {
	(*Args)(args).SetBoolLabel(label, labelName)
}
func (args LProps) GetIntVal(label string) int {
	return Args(args).GetIntVal(label)
}
func (args LProps) GetStringVal(label string) string {
	return Args(args).GetStringVal(label)
}
func (args LProps) GetBoolVal(label string) bool {
	return Args(args).GetBoolVal(label)
}
func (args LProps) ToString(tab int) string {
	return Args(args).ToString(tab)
}
func (args LProps) ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	return Args(args).ToStringWithDetails(tab, text, detailsFlag, printTypeFlag, omitDefaultsFlag)
}
func (args LProps) IsEmpty() bool {
	return Args(args).IsEmpty()
}
func (args LProps) Print(tab int) {
	Args(args).Print(tab)
}
func (args LProps) Println(tab int) {
	Args(args).Println(tab)
}

func (args LProps) String() string {
	return Args(args).String()
}

////////////////////////////////////////
// EPROPS
////////////////////////////////////////
func (args EProps) Copy() EProps {
	copiedArgs := (Args)(args).Copy()
	return (EProps)(copiedArgs)
}
func (args *EProps) SetIntVal(label string, val int) {
	(*Args)(args).SetIntVal(label, val)
}
func (args *EProps) SetStringVal(label string, val string) {
	(*Args)(args).SetStringVal(label, val)
}
func (args *EProps) SetStringEtype(label string, val string) {
	(*Args)(args).SetStringEtype(label, val)
}
func (args *EProps) SetStringUrl(label string, val string) {
	(*Args)(args).SetStringUrl(label, val)
}
func (args *EProps) SetBoolVal(label string, val bool) {
	(*Args)(args).SetBoolVal(label, val)
}
func (args *EProps) SetIntVar(label string, varName string) {
	(*Args)(args).SetIntVar(label, varName)
}
func (args *EProps) SetStringVar(label string, varName string) {
	(*Args)(args).SetStringVar(label, varName)
}
func (args *EProps) SetBoolVar(label string, varName string) {
	(*Args)(args).SetBoolVar(label, varName)
}
func (args *EProps) SetIntLabel(label string, labelName string) {
	(*Args)(args).SetIntLabel(label, labelName)
}
func (args *EProps) SetStringLabel(label string, labelName string) {
	(*Args)(args).SetStringLabel(label, labelName)
}
func (args *EProps) SetBoolLabel(label string, labelName string) {
	(*Args)(args).SetBoolLabel(label, labelName)
}
func (args EProps) GetIntVal(label string) int {
	return Args(args).GetIntVal(label)
}
func (args EProps) GetStringVal(label string) string {
	return Args(args).GetStringVal(label)
}
func (args EProps) GetBoolVal(label string) bool {
	return Args(args).GetBoolVal(label)
}
func (args EProps) ToString(tab int) string {
	return Args(args).ToString(tab)
}
func (args EProps) ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	return Args(args).ToStringWithDetails(tab, text, detailsFlag, printTypeFlag, false /* omitDefaultsFlag */ /* immer false !! @@@ */)
}
func (args EProps) IsEmpty() bool {
	return Args(args).IsEmpty()
}
func (args EProps) Print(tab int) {
	Args(args).Print(tab)
}
func (args EProps) Println(tab int) {
	Args(args).Println(tab)
}

func (args EProps) String() string {
	return Args(args).String()
}

////////////////////////////////////////
// VARS
////////////////////////////////////////
func (args Vars) Copy() Vars {
	copiedArgs := (Args)(args).Copy()
	return (Vars)(copiedArgs)
}
func (args *Vars) SetIntVal(label string, val int) {
	(*Args)(args).SetIntVal(label, val)
}
func (args *Vars) SetStringVal(label string, val string) {
	(*Args)(args).SetStringVal(label, val)
}
func (args *Vars) SetStringEtype(label string, val string) {
	(*Args)(args).SetStringEtype(label, val)
}
func (args *Vars) SetStringUrl(label string, val string) {
	(*Args)(args).SetStringUrl(label, val)
}
func (args *Vars) SetBoolVal(label string, val bool) {
	(*Args)(args).SetBoolVal(label, val)
}
func (args *Vars) SetIntVar(label string, varName string) {
	(*Args)(args).SetIntVar(label, varName)
}
func (args *Vars) SetStringVar(label string, varName string) {
	(*Args)(args).SetStringVar(label, varName)
}
func (args *Vars) SetBoolVar(label string, varName string) {
	(*Args)(args).SetBoolVar(label, varName)
}
func (args *Vars) SetIntLabel(label string, labelName string) {
	(*Args)(args).SetIntLabel(label, labelName)
}
func (args *Vars) SetStringLabel(label string, labelName string) {
	(*Args)(args).SetStringLabel(label, labelName)
}
func (args *Vars) SetBoolLabel(label string, labelName string) {
	(*Args)(args).SetBoolLabel(label, labelName)
}
func (args Vars) GetIntVal(label string) int {
	return Args(args).GetIntVal(label)
}
func (args Vars) GetStringVal(label string) string {
	return Args(args).GetStringVal(label)
}
func (args Vars) GetBoolVal(label string) bool {
	return Args(args).GetBoolVal(label)
}
func (args Vars) ToString(tab int) string {
	return Args(args).ToString(tab)
}
func (args Vars) ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	return Args(args).ToStringWithDetails(tab, text, detailsFlag, printTypeFlag, false /* omitDefaultsFlag */ /* immer false !! @@@ */)
}
func (args Vars) IsEmpty() bool {
	return Args(args).IsEmpty()
}
func (args Vars) Print(tab int) {
	Args(args).Print(tab)
}
func (args Vars) Println(tab int) {
	Args(args).Println(tab)
}

func (args Vars) String() string {
	return Args(args).String()
}

////////////////////////////////////////
// WPROPS
////////////////////////////////////////
func (args WProps) Copy() WProps {
	copiedArgs := (Args)(args).Copy()
	return (WProps)(copiedArgs)
}
func (args *WProps) SetIntVal(label string, val int) {
	(*Args)(args).SetIntVal(label, val)
}
func (args *WProps) SetStringVal(label string, val string) {
	(*Args)(args).SetStringVal(label, val)
}
func (args *WProps) SetStringEtype(label string, val string) {
	(*Args)(args).SetStringEtype(label, val)
}
func (args *WProps) SetStringUrl(label string, val string) {
	(*Args)(args).SetStringUrl(label, val)
}
func (args *WProps) SetBoolVal(label string, val bool) {
	(*Args)(args).SetBoolVal(label, val)
}
func (args *WProps) SetIntVar(label string, varName string) {
	(*Args)(args).SetIntVar(label, varName)
}
func (args *WProps) SetStringVar(label string, varName string) {
	(*Args)(args).SetStringVar(label, varName)
}
func (args *WProps) SetBoolVar(label string, varName string) {
	(*Args)(args).SetBoolVar(label, varName)
}
func (args *WProps) SetIntLabel(label string, labelName string) {
	(*Args)(args).SetIntLabel(label, labelName)
}
func (args *WProps) SetStringLabel(label string, labelName string) {
	(*Args)(args).SetStringLabel(label, labelName)
}
func (args *WProps) SetBoolLabel(label string, labelName string) {
	(*Args)(args).SetBoolLabel(label, labelName)
}
func (args WProps) GetIntVal(label string) int {
	return Args(args).GetIntVal(label)
}
func (args WProps) GetStringVal(label string) string {
	return Args(args).GetStringVal(label)
}
func (args WProps) GetBoolVal(label string) bool {
	return Args(args).GetBoolVal(label)
}
func (args WProps) ToString(tab int) string {
	return Args(args).ToString(tab)
}
func (args WProps) ToStringWithDetails(tab int, text string, detailsFlag bool, printTypeFlag bool, omitDefaultsFlag bool) string {
	return Args(args).ToStringWithDetails(tab, text, detailsFlag, printTypeFlag, omitDefaultsFlag)
}
func (args WProps) IsEmpty() bool {
	return Args(args).IsEmpty()
}
func (args WProps) Print(tab int) {
	Args(args).Print(tab)
}
func (args WProps) Println(tab int) {
	Args(args).Println(tab)
}

func (args WProps) String() string {
	return Args(args).String()
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
