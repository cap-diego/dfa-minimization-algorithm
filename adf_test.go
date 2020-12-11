// +build all minimization

package DFA

import testing "testing"

func TestTrivial(t *testing.T) {
	q0 := State{IsFinal: false, Value: 0}
	q1 := State{IsFinal: true, Value: 1}
	states := []State{ q0, q1 }
	alphabet := []int{0, 1}
	delta := make(map[State]map[int]State)
	delta[q0] = map[int]State{0: q0}
	delta[q0] = map[int]State{1: q1}
	delta[q1] = map[int]State{0: q0}
	delta[q1] = map[int]State{1: q1}
	M := DFA{States: states, InitialState: q0, FinalStates: []State{q1}, Delta: delta, Alphabet: alphabet}
	Min := HopcroftDFAMin(M)
	if Min.Size() != 2 {
		t.Error("size of automata not right")
	}
}

func aTestPrimero(t *testing.T) {
	var states []State

	q0 := State{IsFinal: false, Value: 0}
	q1 := State{IsFinal: false, Value: 1}
	q2 := State{IsFinal: false, Value: 2}
	q3 := State{IsFinal: false, Value: 3}
	q4 := State{IsFinal: false, Value: 4}
	q5 := State{IsFinal: false, Value: 5}
	states = append(states, q0, q1, q2, q3, q4, q5)
	fs := []State{q3, q4, q5}
	alphabet := []int{0, 1}

	delta := make(map[State]map[int]State)

	delta[q0] = map[int]State{0: q1}
	delta[q0] = map[int]State{1: q0}

	delta[q0] = map[int]State{0: q2}
	delta[q0] = map[int]State{1: q0}

	delta[q2] = map[int]State{0: q3}
	delta[q2] = map[int]State{1: q0}

	delta[q3] = map[int]State{0: q3}
	delta[q3] = map[int]State{1: q4}

	delta[q4] = map[int]State{0: q5}
	delta[q4] = map[int]State{1: q4}

	delta[q5] = map[int]State{0: q3}
	delta[q5] = map[int]State{1: q4}

	M := DFA{States: states, InitialState: q0, FinalStates: fs, Delta: delta, Alphabet: alphabet}

	HopcroftDFAMin(M)	
}