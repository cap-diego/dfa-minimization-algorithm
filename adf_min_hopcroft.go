package adf

type State struct {
	Value int
	IsFinal bool
}

type Splitter struct {
	Partition []State
	Symbol int
}

type Adf struct {
	States []State
	Alphabet []int
	InitialState State
	FinalStates []State
	Delta map[State]map[int]State // Given state and symbol returns the state
}



//func HopcroftAdfMin(A Adf) Adf{
//	/* Minimise finite automaton */
//
//	pi, _ := Equivalence(A.States, 0)
//	var splitterSet []Splitter
//
//	for _, symb := range A.Alphabet{
//		if len(StatesWithIncomingTransitionFromState(symb, pi[0], A)) <
//			len(StatesWithIncomingTransitionFromState(symb, pi[1], A)) {
//			// Use partition with finals states
//			splitterSet = append(splitterSet, Splitter{Partition: pi[1], Symbol: symb})
//		} else {
//			// Use partition with non finals states
//			splitterSet = append(splitterSet, Splitter{Partition: pi[0], Symbol: symb})
//		}
//	}
//	for len(splitterSet) > 0 {
//		// Pick a splitter <Partition P, Symbol a> from the splitter set (and delete it)
//		PA := splitterSet[len(splitterSet)-1]
//		splitterSet = splitterSet[:len(splitterSet)-1]
//
//		for _, R := range pi {
//			R1, R2 := A.Split(R, PA)
//			// Replace R in pi by {R1, R2}
//			for _, r2 := range R2 {
//				R = append(R1, r2)
//			}
//
//			for _, symb := range A.Alphabet {
//				Rb := Splitter{Partition: R, Symbol: symb}
//				if HasSplitter(splitterSet, Rb) {
//					// Replace <R, b> spliter  in splitterSet with <R1, b> and <R2, b>
//					R = MergeStates(R1, R2)
//					Rb = Splitter{Partition: R, Symbol: symb}
//				} else {
//					if len(StatesWithIncomingTransitionFromState(symb, R1, A)) <
//						len(StatesWithIncomingTransitionFromState(symb, R2, A)) {
//						// add <R1, b> to splitterSet
//						splitterSet = append(splitterSet, Splitter{Partition: R1, Symbol: symb})
//					} else {
//						// add <R2, b> to splitterSet
//						splitterSet = append(splitterSet, Splitter{Partition: R2, Symbol: symb})
//					}
//				}
//
//			}
//		}
//	}
//	Ar := Adf{States: statesFromMinim, Alphabet: A.Alphabet, InitialState: A.initialState, FinalStates: finalStatesFromMinim, Delta: A.Delta}
//	return Ar
//}

func MergeStates(R1 []State, R2 []State) []State {
	for _, r1 := range R1 {
		if !HasState(R2, r1) {
			R2 = append(R2, r1)
		}
	}
	return R2
}

func HasSplitter(splitterSet []Splitter, Rb Splitter) bool{
	for _, splitter := range splitterSet {
		if SplitterHasAllElems(splitter, Rb.Partition) {
			if splitter.Symbol == Rb.Symbol{
				return true
			}else {
				break
			}
		}

	}
	return false
}

func SplitterHasAllElems(splitter Splitter, Rb []State) bool {
	for _, q := range Rb {
		count := 0
		if !HasState(splitter.Partition, q) {
			break
		} else {
			count += 1
		}

		if count == len(Rb){
			return true
		}
	}
	return false
}

func (A Adf) Split(R []State, PA Splitter) (R1 []State, R2 []State){
	for _, r := range R {
		if HasState(PA.Partition, A.Delta[r][PA.Symbol]) {
			R1 = append(R1, r)
		} else {
			R2 = append(R2, r)
		}
	}
	return
}

// a(P) = {p e P : delta(q, a) = p for some p e States in A}
func StatesWithIncomingTransitionFromState(a int, P []State, A Adf) []State {
	var res []State
	for _, v := range A.States {
		if q := A.Delta[v][a]; HasState(P, q) {
			res = append(res, q)
		}
	}
	return res
}

func HasState(P []State, q State) bool {
	for _, item := range P {
		if item == q {
			return true
		}
	}
	return false
}



func Equivalence(states []State, k uint) (map[int][]State, int) {
	/* Given automaton and a positive integer returns the set of equivalence classes for k */
	var eq map[int][]State
	eq = make(map[int][]State)
	if k == 0 {
		for i := 0; i <= len(states); i++ {
			if states[i].IsFinal{
				eq[0] = append(eq[0], states[i])
			} else {
				eq[1] = append(eq[1], states[i])
			}
		}
	}
	// TODO: k > 0
	return eq, 2
}


/*

http://www.cs.cmu.edu/~cdm/resources/Knuutila01.pdf

*/