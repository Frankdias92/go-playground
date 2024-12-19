package competitionandresources

import "fmt"

func TypeParamenters() {
	a := 2
	b := 2

	fmt.Printf("add(a, b): %v\n", add(a, b))
	fmt.Println(pair(a, b))

	myStack := &Stack[string]{}
	myStack.Push("Go")
	myStack.Push("TypeScript")
	fmt.Printf("myStack: %v\n", myStack)

	myStack.Pop()
	fmt.Printf("myStack: %v\n", myStack)
}

func add[T int](a, b T) T {
	return a + b
}

func pair[T, U any](a T, b U) (T, U) {
	return a, b
}

type Stack[T string] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}
