// +build estructures all

package DFA

import testing "testing"


func TestPartitionNotEmpty(t *testing.T) {
	q0 := State{IsFinal: false, Value: 0}

	partition := Partition{ q0 }

	if partition.IsEmpty() {
		t.Error()
	}
}
func TestPartitionEmpty(t *testing.T) {
	partition := Partition{}

	if !partition.IsEmpty() {
		t.Error()
	}
	q0 := State{IsFinal: false, Value: 0}

	partition.Add(q0)

	if partition.IsEmpty() {
		t.Error()
	}
}

func TestPartitionExtractElem(t *testing.T) {
	q0 := State{IsFinal: false, Value: 0}
	partition := Partition{ q0 }

	partition.ExtractElem(q0)

	if !partition.IsEmpty() {
		t.Error()
	}
	q1 := State{IsFinal: true, Value: 1}
	partition.Add(q0)
	partition.Add(q1)
	partition.ExtractElem(q1)
	if !partition.Includes(q0) {
		t.Error()
	}
	if partition.Includes(q1) {
		t.Error()
	}
}

func TestPartitionExtractPartitionWithNoIntersection(t *testing.T) {
	q1 := State{IsFinal: true, Value: 1}
	q2 := State{IsFinal: false, Value: 2}
	q3 := State{IsFinal: false, Value: 3}
	q4 := State{IsFinal: false, Value: 4}

	p1 := Partition{ q1, q2 }
	p2 := Partition{ q3, q4 }
	oldSizeSum := p1.Size() + p2.Size()

	p1.Extract(p2)
	if oldSizeSum != p1.Size() + p2.Size() {
		t.Error()
	}
}
func TestPartitionExtractPartitionWithIntersection(t *testing.T) {
	q1 := State{IsFinal: true, Value: 1}
	q2 := State{IsFinal: false, Value: 2}
	q3 := State{IsFinal: false, Value: 3}
	q4 := State{IsFinal: false, Value: 4}

	p1 := Partition{ q1, q2, q4 }
	p2 := Partition{ q3, q4 }
	oldSizeP1 := p1.Size()
	p1.Extract(p2)
	if oldSizeP1 != p1.Size() + 1 {
		t.Error()
	}
}

func TestPartitionExtractPartitionBiggerBecomesEmpty(t *testing.T) {
	q1 := State{IsFinal: true, Value: 1}
	q2 := State{IsFinal: false, Value: 2}
	q3 := State{IsFinal: false, Value: 3}
	q4 := State{IsFinal: false, Value: 4}

	p1 := Partition{ q1, q2, q3, q4 }
	p2 := Partition{ q3, q4 }
	p2.Extract(p1)
	if !p2.IsEmpty() {
		t.Error()
	}
}

