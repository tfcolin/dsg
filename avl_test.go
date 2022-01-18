package dsg

import (
	"fmt"
	"testing"
)

func TestAvlTree(t *testing.T) {

	const (
		n  int = 15
		n1 int = 5
	)

	numbers := [n]int{5, 3, -1, 2, 6, 7, 9, 0, -4, -5, 4, -6, 10, 8, 11}
	rm_num := [5]int{0, -1, -5, 9, 2}

	var at *AvlTree = InitIntAvlTree()

	for i := 0; i < n; i++ {
		at.Add(numbers[i])
		fmt.Printf("After Adding %10d: \n", numbers[i])
		at.Print()
		fmt.Print("\n")
	}

	fmt.Print("\n")

	for i := 0; i < n1; i++ {
		at.Remove(rm_num[i])
		fmt.Printf("After Remove %10d: \n", rm_num[i])
		at.Print()
		fmt.Print("\n")
	}

}
