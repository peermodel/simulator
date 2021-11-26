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

package latex

import (
	. "github.com/peermodel/simulator/helpers"
	"strings"
)

// ----------------------------------------
// convert a string into a string that can be printed by latex:
// substitute every occurrance MAX_SUBSTITUTIONS_IN_STRING times
// @@@ add more on demand -> see peer-space-latex-includes-tex
// @@@ in welches package gehört diese func?!
func ConvertString2LatexString(s string) string {
	// order is important! because of trailing blanks! a bit tricky...
	type s1s2 struct {
		s1 string
		s2 string
	}
	conversionTable := []s1s2{
		s1s2{"==", "\\tteqs\\tteqs "},
		s1s2{"<=", "\\ttlab\\tteqs "},
		s1s2{">=", "\\ttrab\\tteqs "},
		s1s2{"\"", "\\ttdqt "},
		s1s2{"$", "\\ttdlr "},
		s1s2{"=", "\\tteqs "},
		s1s2{"<", "\\ttlab "},
		s1s2{">", "\\ttrab "},
		s1s2{"EXCEPTION_WRAP", "exception"}, // caution: must be done before "_" is replaced
		s1s2{"_", "\\ttusc "},
		// these might occur in user strings
		s1s2{"&", "\\& "},
		s1s2{"%", "\\% "},
		s1s2{"#", "\\# "},
	}
	for i := 0; i < len(conversionTable); i++ {
		s = strings.Replace(s, conversionTable[i].s1, conversionTable[i].s2, MAX_SUBSTITUTIONS_IN_STRING)
	}

	return s
}

////////////////////////////////////////
// EOF
////////////////////////////////////////
