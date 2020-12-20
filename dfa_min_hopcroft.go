package dfa

// HopcroftDFAMin Given A DFA returns a new DFA with minimum states
func HopcroftDFAMin(A DFA) DFA {
	completarTransiciones(&A)
	var pi []Partition
	pi = append(pi, A.FinalStates)
	pi = append(pi, A.States)
	pi[1].Extract(pi[0])

	var worklist []splitter

	for _, symb := range A.Alphabet { // Step 4
		var a0 = pi[0].StatesWithIncomingTransitionWith(symb, &A)
		var a1 = pi[1].StatesWithIncomingTransitionWith(symb, &A)
		if a0.Size() < a1.Size() && a0.Size()>0 {
			// Use partition with finals states
			worklist = append(worklist, splitter{partition: pi[0], input: symb})
		} else {
			// Use partition with non finals states
			worklist = append(worklist, splitter{partition: pi[1], input: symb})
		}
	}
	for len(worklist) > 0 {
		// Pick a splitter <Partition P, Symbol a> from the splitter set (and delete it)
		currentSplitter := pickOneAndRemove(&worklist)
		for i, R := range pi { //Step 7
			R1, R2, splitted := R.SplitBy(&currentSplitter, &A)
			if splitted { // Hay refinamiento
				// R == Bj
				// R1 == B'j  ==>  Bj
				// R2 == B''j ==>  Bk
				pi = append(pi, R2)
				for _, c := range A.Alphabet {
					spR := splitter{partition: R, input: c}
					spR1 := splitter{partition: R1, input: c}
					spR2 := splitter{partition: R2, input: c}
					if !RefinedPartitionReplacedInWorklist(&worklist, spR, spR1, spR2) {
						ar1 := R1.StatesWithIncomingTransitionWith(c, &A)
						ar2 := R2.StatesWithIncomingTransitionWith(c, &A)
						if ar1.Size() < ar2.Size() && ar1.Size()>0 {
							worklist = append(worklist, splitter{partition: R1, input: c})
						} else {
							worklist = append(worklist, splitter{partition: R2, input: c})
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
		addTransitions(&DeltaMin, &A, pi, i)
		partitionHasFinalStates := false
		partitionHasInitialState := false
		for _, s := range part {
			if A.FinalStates.Includes(s) {
				partitionHasFinalStates = true
			}
			if A.InitialState == s {
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

func addTransitions(DeltaMin *map[State]map[int]State, A *DFA, pi []Partition, i int) {
	for _, c := range A.Alphabet {
		transitionState, ok := A.Delta[pi[i][0]][c]
		if !ok {
			continue
		}
		partitionOfState := getPartitionOfState(&pi, transitionState)
		if (*DeltaMin)[i] == nil {
			(*DeltaMin)[i] = map[int]State{c: partitionOfState}
		} else {
			(*DeltaMin)[i][c] = partitionOfState
		}
	}

}

func getPartitionOfState(pi *[]Partition, q State) int {
	for i, part := range *pi {
		if part.Includes(q) {
			return i
		}
	}
	return 0
}

// RefinedPartitionReplacedInWorklist if the splitted partition was in the worklist it has to be replaced.
// If did not exist in the worklist return false
func RefinedPartitionReplacedInWorklist(worklist *[]splitter, spR splitter, spR1 splitter, spR2 splitter) bool {
	newWorklist := *worklist
	worklistSize := len(*worklist)
	for i, w := range *worklist {
		if w.partition.Equals(spR.partition) && w.input == spR.input {
			newWorklist[i] = (*worklist)[worklistSize-1]
			newWorklist = append(newWorklist, spR1)
			newWorklist = append(newWorklist, spR2)
			*worklist = newWorklist
			return true
		}
	}
	return false
}

func pickOneAndRemove(worklist *[]splitter) splitter {
	worklistSize := len(*worklist)
	newWorklist := *worklist
	sp := (*worklist)[worklistSize-1]
	newWorklist = (*worklist)[:worklistSize-1]
	*worklist = newWorklist
	return sp
}

func completarTransiciones(A *DFA) {
	qt := 9999
	for _, input := range A.Alphabet {
		if A.Delta[qt] == nil {
			A.Delta[qt] = map[int]State{input: qt}
		}else {
			A.Delta[qt][input] = qt
		}
		for _, q := range A.States {
			if A.Delta[q] == nil {
				A.Delta[q] = map[int]State{input: qt}
			}else {
				_, ok := A.Delta[q][input]
				if !ok {
					A.Delta[q][input] = qt
				}
			}
			
		}
	}
	A.Alphabet = append(A.Alphabet, qt)
}