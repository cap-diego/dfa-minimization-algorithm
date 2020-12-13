package DFA

import testing "testing"

func TestTrivial(t *testing.T) {
	q0 := 0
	q1 := 1
	states := []State{ q0, q1 }
	alphabet := []int{0, 1}
	delta := make(map[State]map[int]State)
	delta[q0] = map[int]State{0: q0, 1: q1}
	delta[q1] = map[int]State{0: q0, 1: q1}
	M := DFA{States: states, InitialState: q0, FinalStates: []State{q1}, Delta: delta, Alphabet: alphabet}
	Min := HopcroftDFAMin(M)
	if Min.Size() != 2 {
		t.Error("size of automata not right")
	}
}

func TestDFAMinimo(t *testing.T) {
	var states []State

	q1 := 1
	q2 := 2
	q3 := 3
	q4 := 4
	q5 := 5
	q6 := 6
	
	states = append(states ,q1, q2, q3, q4, q5, q6)
	
	fs := []State{ q6 }

	alphabet := []int{0, 1}

	delta := make(map[State]map[int]State)

	delta[q1] = map[int]State{0: q2, 1: q1}

	delta[q2] = map[int]State{0: q3, 1: q2}

	delta[q3] = map[int]State{0: q4, 1: q3}

	delta[q4] = map[int]State{0: q5, 1: q4}

	delta[q5] = map[int]State{0: q6, 1: q5}

	delta[q6] = map[int]State{0: q6, 1: q6}

	M := DFA{States: states, InitialState: q1, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	Min := HopcroftDFAMin(M)

	if (Min.States.Size() != M.States.Size()) {
		t.Errorf("the minimized dfa should have the same number of states, expected 6 have %d", Min.States.Size())
	} 
}

func TestDFANoMinimo(t *testing.T) {
	var states []State
	q1 := 1
	q2 := 2
	q3 := 3
	q4 := 4
	q5 := 5
	q6 := 6
	
	states = append(states ,q1, q2, q3, q4, q5, q6)
	
	fs := []State{ q4, q5, q6 }

	alphabet := []int{0, 1}

	delta := make(map[State]map[int]State)

	delta[q1] = map[int]State{0: q2, 1: q1}

	delta[q2] = map[int]State{0: q3, 1: q1}

	delta[q3] = map[int]State{0: q4, 1: q1}

	delta[q4] = map[int]State{0: q4, 1: q5}

	delta[q5] = map[int]State{0: q6, 1: q5}

	delta[q6] = map[int]State{0: q4, 1: q5}

	M := DFA{States: states, InitialState: q1, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	Min := HopcroftDFAMin(M)

	if (Min.States.Size() == M.States.Size()) {
		t.Errorf("the minimized dfa should have less states, expected 4, have %d", Min.States.Size())
	} 
}

