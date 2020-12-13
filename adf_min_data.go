package DFA

type Partition []State
type State = int // Alias to int

type Splitter struct {
	Partition Partition
	Symbol int
}

type DFA struct {
	States Partition
	Alphabet []int
	InitialState State
	FinalStates Partition
	Delta map[State]map[int]State // Given state and symbol returns the state
}

// Size returns the number of states of the automata
func (M *DFA) Size() int{
	return M.States.Size()
}

// SplitBy returns R1 and R2, where R1 are the states from the partition that has transitions to sp.Partition with sp.Symbol
// and R2 are the states from the partition that do not have transitions with sp.Symbol to sp.Partition
func (P *Partition) SplitBy(sp *Splitter, A *DFA) (R1 Partition, R2 Partition, splitted bool) {
	splitted = true
	a := sp.Symbol
	partitionSp := sp.Partition
	for _, t := range (*P) {
		if partitionSp.Includes((A.Delta[t][a])) {
			R1.Add(t)
		} else {
			R2.Add(t)
		}
	}
	if R1.IsEmpty() || R2.IsEmpty() {
		splitted = false
	}
	return
}

// NewPartition returns a new initialized partition
func NewPartition() *Partition {
	return &Partition{}
}

// StatesWithIncomingTransitionWith returns the states from the partition that has incoming transitions with a
func (P *Partition) StatesWithIncomingTransitionWith(a int, A *DFA) Partition {
	newPartition := NewPartition()
	for _, s := range A.States {
		if P.Includes(A.Delta[s][a]) {
			newPartition.Add(A.Delta[s][a])
		}
	}
	return *newPartition
}	

type Set interface {
	Extract(Q Partition)
	ExtractElem(q *State)
	Includes(q State) bool
	Size() int
	IsEmpty() bool
	Add(Q *Partition)
	Equals(Q Partition) bool
}


// Equals returns true if Q has the same elements that the partition
func (P *Partition) Equals(Q Partition) bool {
	if Q.Size() != P.Size() {
		return false
	}
	for _, p := range (*P) {
		Q.ExtractElem(p)
	}	
	return Q.IsEmpty()
}

// Add adds a new state to the partition if it does not exist
func (P *Partition) Add(q State) {
	for _, t := range (*P) {
		if  t == q {
			return
		}
	}
	*P = append(*P, q)
}

// IsEmpty returns true if the partition has 0 elements
func (P *Partition) IsEmpty() bool {
	return P.Size() == 0
}

// Size returns the number of elements of the partition
func (P *Partition) Size() int {
	return len((*P))
}

// Includes returns true if q is in the partition
func (P *Partition) Includes(q State) bool {
	for _, p := range (*P) {
		if q == p {
			return true
		}
	}
	return false
}

// Extract removes the elements of Q from the partition
func (P *Partition) Extract(Q Partition) {
	for _, e := range Q {
		P.ExtractElem(e)
	}
}

// ExtractElem removes q from the partition
func (P *Partition) ExtractElem(q State) {
	for i, e := range *P {
		if e == q {
			(*P)[i] = (*P)[len((*P))-1]
			(*P) = (*P)[:len((*P))-1]
			return
		}
	}
}



