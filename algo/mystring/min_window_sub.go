package mystring

// func MinWindowSub(str, sub string) {
// 	ln := len(str)
// 	subLookup := make(map[string]int, len(sub))
// 	subFound := make(map[string]int, len(sub))
// 	for i := 0; i < len(sub); i++ {
// 		blockSymbols[string(sub[i])] = i
// 	}
// 	l, r := 0, 0
// 	for i := 0; i < ln; i++ {
// 		if len(subFound) != len(subLookup) {
// 			// finding all the elements, expanding window
// 			if val, ok := subLookup[str[r]]; ok {
// 				subFound[str[r]] = i
// 			}
// 			r++
// 		} else {
// 			//elements are found, contracting window
// 			subFound[str[r]]
// 			i++
// 		}
// 	}
// }
