/*
stack implementation using single linked list
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

// StackLinkedList implements stack using array fashioned way
type StackLinkedList struct {
	Head *Frame
	size uint
}

type Frame struct {
	data int
	Next *Frame
}

// Push pushes int to data slice if it's capable
// and returns the last element pushed with flag ok
func (sll *StackLinkedList) Push(p int) (int, bool) {
	// overflow is not possible in linked-list implementation unless memory is exhausted
	newNode := &Frame{
		data: p,
		Next: sll.Head,
	}
	sll.Head = newNode
	sll.size++
	return sll.Head.data, true

}

// Pop pops the top element from the stack data slice if it's not empty
// and returns it with flag ok
func (sll *StackLinkedList) Pop() (int, bool) {
	topItem := 0
	ok := false
	if (*sll.Head != Frame{}) {
		topItem = sll.Head.data
		sll.Head = sll.Head.Next
		sll.size--
		ok = true
	}
	return topItem, ok
}

// Peek just peeks the top int from the stack if stack is not empty with flag ok
func (sll *StackLinkedList) Peek() (int, bool) {
	top := 0
	ok := false
	if (*sll.Head != Frame{}) {
		top = sll.Head.data
		ok = true
	}
	return top, ok
}

// IsEmpty tells whether stack is empty or not
func (sll *StackLinkedList) IsEmpty() bool {
	return (*sll.Head == Frame{})
}

// Print prints the contents of a stack with the space between elements using ItemSpace
func (sll *StackLinkedList) Print() {
	s := "["

	topItem := sll.Head
	for {
		if (*topItem != Frame{}) {
			s += " " + strconv.Itoa(topItem.data)
			topItem = topItem.Next
		} else {
			break
		}
	}

	s += " " + "]"
	fmt.Println(s)
}

func main() {
	var i int
	var ok bool

	sll := StackLinkedList{
		Head: &Frame{},
	}

	fmt.Printf("stack is empty: %t\n", sll.IsEmpty())

	// pushing first node
	i, ok = sll.Push(1)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nstack is overflow")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// pushing second node
	i, ok = sll.Push(2)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nproblem with PUSH op")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// pushing third node
	i, ok = sll.Push(3)
	if ok {
		fmt.Printf("\npushed item: %d", i)
	} else {
		fmt.Println("\nproblem with PUSH op")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// first time popping
	i, ok = sll.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// seeking head node of the stack
	i, ok = sll.Peek()
	if ok {
		fmt.Printf("\nthe TOP of stack is %d\n", i)
	} else {
		fmt.Println("\nno TOP found in stack, is empty")
	}

	fmt.Print("stack contents: ")
	sll.Print()

	// second last node popping
	i, ok = sll.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// last node popping
	i, ok = sll.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty")
	}
	fmt.Print("\nstack contents: ")
	sll.Print()

	// after stack is empty
	i, ok = sll.Pop()
	if ok {
		fmt.Printf("\npopped item: %d", i)
	} else {
		fmt.Println("\nstack is empty Now!")
	}
}
