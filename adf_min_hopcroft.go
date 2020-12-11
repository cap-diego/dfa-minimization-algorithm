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
		
		for _, R := range pi { //Step 7
			R1, R2, splitted := R.SplitBy(&currentSplitter, &A)
			if splitted { // Hay refinamiento
				// R == Bj
				// R1 == B'j  ==>  Bj
				// R2 == B''j ==>  Bk
				R = R1
				pi = append(pi, R2)

				for _, c := range A.Alphabet {
					for _, splitter := range worklist {
						currentPartition := splitter.Partition
						ar1 := R1.StatesWithIncomingTransitionWith(c, &A)
						ar2 := R2.StatesWithIncomingTransitionWith(c, &A)
						if (splitter.Symbol == c && R.Equals(currentPartition)) || (ar1.Size() > ar2.Size()) { 
							worklist = append(worklist, Splitter{Partition: R1, Symbol: c})
						}else {
							worklist = append(worklist, Splitter{Partition: R2, Symbol: c})
						}
					}
				}
			}
		}
	}
	// Ar := DFA{States: statesFromMinim, Alphabet: A.Alphabet, InitialState: A.initialState, FinalStates: finalStatesFromMinim, Delta: A.Delta}
	return DFA{}
}

func PickOneAndRemove(worklist *[]Splitter) Splitter {
	worklistSize := len(*worklist)
	sp := (*worklist)[worklistSize - 1]
	newWorklist := (*worklist)[:worklistSize-1]
	worklist = &newWorklist
	if len(*worklist) == worklistSize {
		panic("asd")
	}
	return sp
}
