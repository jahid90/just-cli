package lib

import (
	"errors"
	"fmt"
)

// RuneStack A stack of runes
type RuneStack struct {
	current  int    // index where the next element will be inserted
	length   int    // the length of the stack
	elements []rune // holder for the elemnts of the stack
}

// IsEmpty Returns true when the stack is empty, false otherwise
func (s *RuneStack) IsEmpty() bool {
	return s.current <= 0
}

// Size Returns the size of the stack
func (s *RuneStack) Size() int {
	return s.length
}

// Top Returns the top element from the stack, without modifying the stack
func (s *RuneStack) Top() (rune, error) {

	if s.current <= 0 {
		return -1, errors.New("error: stack is empty")
	}

	return s.elements[s.current-1], nil
}

// Pop Pops the top element from the stack and returns it
// On pop, just the indices are updated, no space is released. On subsequent pushes, this space is reused
func (s *RuneStack) Pop() (rune, error) {

	elem, err := s.Top()
	if err != nil {
		return -1, err
	}

	s.current--
	s.length--

	return elem, nil
}

// Push Pushes an element to the top of the stack
// On push, space is allocated if needed, and element is inserted at index current
func (s *RuneStack) Push(e rune) error {

	if s.current < len(s.elements) {
		s.elements[s.current] = e
	} else {
		s.elements = append(s.elements, e)
	}

	s.length++
	s.current++

	return nil
}

// AsString Converts the stack of runes into a string
func (s *RuneStack) AsString() string {
	result := ""

	for i := 0; i < s.length; i++ {
		result += string(s.elements[i])
	}

	return result
}

// NewRuneStack Creates an instance of a stack of runes
func NewRuneStack() *RuneStack {
	return &RuneStack{
		elements: make([]rune, 0),
		length:   0,
		current:  0,
	}
}

// Print Prints the stack state
func (s *RuneStack) Print() {
	fmt.Println("Current: " + fmt.Sprint(s.current) + ", Size: " + fmt.Sprint(s.Size()))
	fmt.Println(s.AsString())
}
