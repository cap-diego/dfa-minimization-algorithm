package DFA

func HopcroftDFAMin(A DFA) DFA {
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
	finalStatesMinim := NewPartition()
	statesMinim := NewPartition()
	var initialStateMinim State
	DeltaMin := make(map[State]map[int]State)
	for i, part := range pi {
			statesMinim.Add(i)
			AddTransitions(&DeltaMin, &A, pi, i)
			partitionHasFinalStates := false
			partitionHasInitialState := false
			for _, s := range part {
				if A.FinalStates.Includes(s) {
					partitionHasFinalStates = true
				}
				if(A.InitialState == s) {
					partitionHasInitialState = true
					continue
				}
			}
			if partitionHasFinalStates {
				finalStatesMinim.Add(i)	
			}
			if partitionHasInitialState {
				initialStateMinim = i
			}	
	}
	Ar := DFA{States: *statesMinim, Alphabet: A.Alphabet, InitialState: initialStateMinim, FinalStates: *finalStatesMinim, Delta: DeltaMin}
	return Ar
}

func AddTransitions(DeltaMin *map[State]map[int]State, A *DFA, pi []Partition, i int) {
	for _, c := range A.Alphabet {
		particionDelEstado := particionDeEstado(&pi, A.Delta[pi[i][0]][c])
		(*DeltaMin)[i] = map[int]State{c: particionDelEstado} 
	}
}

func particionDeEstado(pi *[]Partition, q State) int {
	for i, part := range *pi {
		if part.Includes(q) {
			return i
		}
	}
	return 0
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
