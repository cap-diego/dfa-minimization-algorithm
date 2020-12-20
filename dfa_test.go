package dfa

import (
	testing "testing"
)

func TestTrivial(t *testing.T) {
	q0 := 0
	q1 := 1
	states := []State{q0, q1}
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

	states = append(states, q1, q2, q3, q4, q5, q6)

	fs := []State{q6}

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

	if Min.States.Size() != M.States.Size() {
		t.Errorf("the minimized dfa should have the same number of states, expected 6 got %d", Min.States.Size())
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

	states = append(states, q1, q2, q3, q4, q5, q6)

	fs := []State{q4, q5, q6}

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

	if Min.States.Size() == M.States.Size() {
		t.Errorf("the minimized dfa should have less states, expected 4, got %d", Min.States.Size())
	}
}

func Test2DFANoMinimo(t *testing.T) {
	// 02 | 012
	var states []State
	q0 := 0
	q1 := 1
	q2 := 2
	q3 := 3
	q4 := 4
	q5 := 5
	states = append(states, q0, q1, q2, q3, q4, q5)
	fs := []State{q3, q5}
	alphabet := []int{0, 1, 2}
	delta := make(map[State]map[int]State)
	delta[q0] = map[int]State{1: q1}
	delta[q1] = map[int]State{2: q2, 1: q4}
	delta[q2] = map[int]State{2: q3}
	delta[q4] = map[int]State{2: q5}
	M := DFA{States: states, InitialState: q1, FinalStates: fs, Delta: delta, Alphabet: alphabet}
	Min := HopcroftDFAMin(M)

	if Min.States.Size() != 4 {
		t.Errorf("error, expected 4 got %d states", Min.States.Size())
	}
	if Min.FinalStates.Size() != 1 {
		t.Errorf("error, expected 1 final state got: %d", Min.FinalStates.Size())
	}
	if states := Min.FinalStates.StatesWithIncomingTransitionWith(q0, &Min); states.IsEmpty() {
		t.Errorf("error, expected 0 transitions to final states, got: %d\n", states.Size())
	}
	if Min.Delta[Min.InitialState] == nil {
		t.Errorf("error, expected transition from initial got: %d\n", Min.InitialState)
	}
}

func Test3DFANoMinimo(t *testing.T) {
	// a ( b | c*)
	var states []State
	q0 := 0
	q1 := 1
	q2 := 2
	q3 := 3
	states = append(states, q0, q1, q2, q3)
	fs := []State{q1, q2, q3}
	alphabet := []int{0, 1, 2}
	delta := make(map[State]map[int]State)
	delta[q0] = map[int]State{0: q1}
	delta[q1] = map[int]State{1: q2, 2: q3}
	delta[q2] = map[int]State{1: q2, 2: q3}
	delta[q3] = map[int]State{1: q2, 2: q3}

	M := DFA{States: states, InitialState: q1, FinalStates: fs, Delta: delta, Alphabet: alphabet}
	Min := HopcroftDFAMin(M)

	if Min.States.Size() != 2 {
		t.Errorf("error, expected 2 got %d states", Min.States.Size())
	}
	if Min.FinalStates.Size() != 1 {
		t.Errorf("error, expected 1 final state got: %d", Min.FinalStates.Size())
	}
	for _, c := range Min.Alphabet {
		statesWithIncomingTransitions := Min.FinalStates.StatesWithIncomingTransitionWith(c, &Min)
		if statesWithIncomingTransitions.Size() == 0 {
			t.Errorf("error, expected transition with %d to final state %d\n", c, Min.FinalStates[0])
		}
	}
}

func Test4DFANoMinimo(t *testing.T) {
	var states []State
	A := 0
	B := 1
	C := 2
	D := 3
	E := 4
	F := 5
	G := 6
	H := 7
	states = append(states, A, B, C, D, E, F, G, H)
	fs := []State{C}
	alphabet := []int{0, 1}
	delta := make(map[State]map[int]State)
	delta[A] = map[int]State{0: B, 1: F}
	delta[B] = map[int]State{0: G, 1: C}
	delta[C] = map[int]State{0: A, 1: C}
	delta[D] = map[int]State{0: C, 1: G}
	delta[E] = map[int]State{0: H, 1: F}
	delta[F] = map[int]State{0: C, 1: G}
	delta[G] = map[int]State{0: G, 1: E}
	delta[H] = map[int]State{0: G, 1: C}
	M := DFA{States: states, InitialState: A, FinalStates: fs, Delta: delta, Alphabet: alphabet}
	Min := HopcroftDFAMin(M)
	if Min.States.Size() != 5 {
		t.Errorf("error, expected 5 states got %d", Min.States.Size())
	}
	if Min.FinalStates.Size() != 1 {
		t.Errorf("error, expected 1 final state got %d", Min.FinalStates.Size())
	}
}

func Test5DFAMinimo(t *testing.T) {
	var states []State
	A := 0
	B := 1
	C := 2
	D := 3
	E := 4
	F := 5
	states = append(states, A, B, C, D, E, F)
	fs := []State{A , B, C, D, F}
	alphabet := []int{0, 1}
	delta := make(map[State]map[int]State)
	delta[A] = map[int]State{0: B, 1: E}
	delta[B] = map[int]State{0: B, 1: C}
	delta[C] = map[int]State{0: D, 1: E}
	delta[D] = map[int]State{0: F, 1: C}
	delta[E] = map[int]State{0: E, 1: E}
	delta[F] = map[int]State{0: F, 1: E}

	M := DFA{States: states, InitialState: A, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	Min := HopcroftDFAMin(M)
	
	if Min.States.Size() > M.States.Size() {
		t.Errorf("error, minimized automata should have the same number of states, got %d", Min.States.Size())
	}

}

func Test6DFAMinimo(t *testing.T) {
	// L = (11111)+
	var states []State
	A := 0
	B := 1
	C := 2
	D := 3
	E := 4
	F := 5
	states = append(states, A, B, C, D, E, F)
	fs := []State{ F }
	alphabet := []int{0, 1}
	delta := make(map[State]map[int]State)
	delta[A] = map[int]State{1: B}
	delta[B] = map[int]State{1: C}
	delta[C] = map[int]State{1: D}
	delta[D] = map[int]State{1: E}
	delta[E] = map[int]State{1: F}
	delta[F] = map[int]State{1: A}

	M := DFA{States: states, InitialState: A, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	Min := HopcroftDFAMin(M)
	
	if Min.States.Size() != M.States.Size() {
		t.Errorf("error, minimized automata should have the same number of states, got %d", Min.States.Size())
	}

}

func Test7DFAMinimo(t *testing.T) {
	// L = fee | fie
	var states []State
	A := 0
	B := 1
	C := 2
	D := 3
	E := 4
	F := 5
	states = append(states, A, B, C, D, E, F)
	fs := []State{ D, F }
	alphabet := []int{1, 2, 3}
	delta := make(map[State]map[int]State)
	delta[A] = map[int]State{1: B}
	delta[B] = map[int]State{2: C}
	delta[B] = map[int]State{3: E}
	delta[C] = map[int]State{2: D}
	delta[E] = map[int]State{2: F}

	M := DFA{States: states, InitialState: A, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	Min := HopcroftDFAMin(M)
	
	if Min.States.Size() != 4 {
		t.Errorf("error, minimized automata should have less states, got %d expected 4", Min.States.Size())
	}

}