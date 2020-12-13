package DFA

func HopcroftDFAMin(A DFA) DFA{
	/* Minimise finite automaton */

	var pi []Partition

	pi = append(pi, A.FinalStates)
	pi = append(pi, A.States)
	pi[1].Extract(pi[0])

	var worklist []Splitter

	for _, symb := range A.Alphabet { // Step 4
		var a0 = pi[0].StatesWithIncomingTransitionWith(symb,  &A)
		var a1 = pi[1].StatesWithIncomingTransitionWith(symb, &A)
		if a0.Size() < a1.Size() {
			// Use partition with finals states
			worklist = append(worklist, Splitter{Partition: pi[0], Symbol: symb})
		} else {
			// Use partition with non finals states
			worklist = append(worklist, Splitter{Partition: pi[1], Symbol: symb})
		}
	}
	for len(worklist) > 0 {
		// Pick a splitter <Partition P, Symbol a> from the splitter set (and delete it)
		currentSplitter := PickOneAndRemove(&worklist)
		for i, R := range pi { //Step 7
			R1, R2, splitted := R.SplitBy(&currentSplitter, &A)
			if splitted { // Hay refinamiento
				// R == Bj
				// R1 == B'j  ==>  Bj
				// R2 == B''j ==>  Bk
				pi = append(pi, R2)
				for _, c := range A.Alphabet {
					spR := Splitter{ Partition: R, Symbol: c}
					spR1 := Splitter{ Partition: R1, Symbol: c}
					spR2 := Splitter{ Partition: R2, Symbol: c}
					if !ReplaceInWorklistIfSplitterExists(&worklist, spR, spR1, spR2) { 
							ar1 := R1.StatesWithIncomingTransitionWith(c, &A)
							ar2 := R2.StatesWithIncomingTransitionWith(c, &A)
							if ( (ar1.Size() < ar2.Size())) { 
								worklist = append(worklist, Splitter{Partition: R1, Symbol: c})
							}else {
								worklist = append(worklist, Splitter{Partition: R2, Symbol: c})
							}
					}
				}
				pi = append(pi, R1)
				pi[i] = pi[len(pi)-1]
				pi = pi[:len(pi)-1]
			}
		}
	}
	var finalStatesMinim Partition
	var statesMinim Partition
	var initialStateMinim State 
	for i, part := range pi {
			statesMinim.Add(i)
			tieneEstadosFinales := false
			tieneEstadoInicial := false
			for _, s := range part {
				if A.FinalStates.Includes(s) {
					tieneEstadosFinales = true
					continue;
				}else if(A.InitialState == s) {
					tieneEstadoInicial = true
				}
				if tieneEstadosFinales {
					finalStatesMinim.Add(s)	
				}
				if tieneEstadoInicial {
					initialStateMinim = s
				}	
			}
	}
	Ar := DFA{States: statesMinim, Alphabet: A.Alphabet, InitialState: initialStateMinim, FinalStates: finalStatesMinim, Delta: A.Delta}
	return Ar
}

func ReplaceInWorklistIfSplitterExists(worklist *[]Splitter, spR Splitter, spR1 Splitter, spR2 Splitter) bool {
	newWorklist := *worklist
	worklistSize := len(*worklist)
	for i, w := range (*worklist) {
		if w.Partition.Equals(spR.Partition) && w.Symbol == spR.Symbol {
			newWorklist[i] = (*worklist)[worklistSize-1]
			newWorklist = append(newWorklist, spR1)
			newWorklist = append(newWorklist, spR2)
			*worklist = newWorklist
			return true
		}
	}
	return false
}

func PickOneAndRemove(worklist *[]Splitter) Splitter {
	worklistSize := len(*worklist)
	newWorklist := *worklist
	sp := (*worklist)[worklistSize - 1]
	newWorklist = (*worklist)[:worklistSize-1]
	*worklist = newWorklist
	return sp
}
