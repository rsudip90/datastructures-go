/*
stack implementation using array
*/

package main

import (
	"fmt"
	"strconv"
)

// Stack is an interface using which it can be implemented in different ways
// It supports the only "int" kinda data
type Stack interface {
	Push(int) bool
	Pop() (int, bool)
	IsEmpty() bool
	Peek() (int, bool)
	Print()
}

// StackArray implements stack using array fashioned way
type StackArray struct {
	top  uint
	data []int
	Cap  uint
}

// Push pushes int to data slice if it's capable
// and returns the last element pushed with flag ok
func (sa *StackArray) Push(p int) (int, bool) {
	if sa.top == sa.Cap {
		return 0, false
	}
	sa.data = append(sa.data, p)
	sa.top++
	return sa.data[len(sa.data)-1], true

}

// Pop pops the top element from the stack data slice if it's not empty
// and returns it with flag ok
func (sa *StackArray) Pop() (int, bool) {
	topItem := 0
	ok := false
	if len(sa.data) > 0 {
		topItem = sa.data[len(sa.data)-1]
		sa.data = sa.data[:len(sa.data)-1]
		sa.top--
		ok = true
	}
	return topItem, ok
}

// Peek just peeks the top int from the stack if stack is not empty with flag ok
func (sa *StackArray) Peek() (int, bool) {
	top := 0
	ok := false
	if len(sa.data) > 0 {
		top = sa.data[len(sa.data)-1]
		ok = true
	}
	return top, ok
}

// IsEmpty tells whether stack is empty or not
func (sa *StackArray) IsEmpty() bool {
	return len(sa.data) == 0
}

// Print prints the contents of a stack with the space between elements using ItemSpace
func (sa *StackArray) Print() {
	s := "["

	for i := len(sa.data) - 1; i >= 0; i-- {
		s += " " + strconv.Itoa(sa.data[i])
	}

	s += " " + "]"
	fmt.Println(s)
}

func main() {
	var i int
	var ok bool

	sa := StackArray{
		Cap: 2,
	}

	fmt.Printf("stack is empty: %t\n", sa.IsEmpty())

	// pushing first item
	i, ok = sa.Push(1)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nstack is overflow")
	}
	fmt.Print("\nStack contents: ")
	sa.Print()

	// pushing second item
	i, ok = sa.Push(2)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nstack is overflow")
	}
	fmt.Print("\nStack contents: ")
	sa.Print()

	// trying to push third item -- out of capacity for the stack
	i, ok = sa.Push(3)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nstack is overflow")
	}
	fmt.Print("\nStack contents: ")
	sa.Print()

	// first time popping item
	i, ok = sa.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty")
	}
	fmt.Print("\nStack contents: ")
	sa.Print()

	// seeking top item in stack
	i, ok = sa.Peek()
	if ok {
		fmt.Printf("\nthe TOP of stack is %d", i)
	} else {
		fmt.Println("\nno TOP found in stack, is empty")
	}

	// last item popping
	i, ok = sa.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty")
	}
	fmt.Print("\nStack contents: ")
	sa.Print()

	// after stack is emtpy, trying to pop
	i, ok = sa.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty Now!")
	}
}
