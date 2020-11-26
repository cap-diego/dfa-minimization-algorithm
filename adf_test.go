package adf

import (
	testing "testing"
	"fmt"
)


func TestPrimero(t *testing.T) {
	fmt.Println("Primer test")
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

	//fmt.Println("Antes de asignar delta")
	delta[q0] = map[int]State{0: q1}
	delta[q0] = map[int]State{1: q0}
	//
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

	fmt.Println("Antes de llamar a min")
	M := Adf{States: states, InitialState: q0, FinalStates: fs, Delta: delta, Alphabet: alphabet}
	MinADF(&M)
}