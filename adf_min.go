package adf

import "fmt"

func MinADF(M* Adf) {
	//lastEquivalenece := 0
	equivalenceClasses := make([]int, 100)

	table := make([][]int, 100)
	for i := range table {
		table[i] = make([]int, 100)
	}
	fmt.Println("0")
	for i, q := range M.States {
		if q.IsFinal {
			equivalenceClasses[i] = 1
		}else {
			equivalenceClasses[i] = 0
		}
	}

	fmt.Println("1")
	for i, q := range M.States {
		for j, symb := range M.Alphabet {
			table[i][j] = equivalenceClasses[M.idxOfState(M.Delta[q][symb])]
		}
	}
	fmt.Println("2")


	lastEquivalenceClass := 0
	equivalenceClasses[0] = lastEquivalenceClass
	for i, _ := range M.States {
		fmt.Printf("i: %d, table[i]]:%v \n eq: %d", i, table[i], equivalenceClasses[i])
		row := append(table[i], equivalenceClasses[i])
		if rowEquivalent := eqClassExists(&row, &table, i); rowEquivalent != -1 {
			fmt.Printf("Row eq: %d",rowEquivalent)
			equivalenceClasses[i] = equivalenceClasses[rowEquivalent]
		} else {
			fmt.Printf("ACA")
			equivalenceClasses[i] = lastEquivalenceClass + 1
		}
	}

}
func (M* Adf) idxOfState(q State) int {
	for i, v := range M.States {
		if v == q {
			return i
		}
	}
	panic("Error")
}

func eqClassExists(rowCmp* []int, table* [][]int, untilRow int) int {
	fmt.Printf("Entrando a eqClassExists, untl:%d", untilRow)
	for j, row := range *table{
		fmt.Printf("retorna j:%d\n",j)
		if j > untilRow {
			fmt.Print("retorna -1")
			return -1
		}
		for i, _ := range *rowCmp {
			fmt.Printf("retorna, i:%d vs %d\n", i, len(*rowCmp))
			count := 0
			if (*rowCmp)[i] == row[i]{
				count++
			}
			if count == len(*rowCmp) {
				fmt.Print("retorna j")
				return j
			}
		}
	}
	fmt.Print("retorna -1")
	return -1
}