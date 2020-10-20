package lexer

import (
	"errors"
	"fmt"
)

// TokenStack A stack of tokens
type TokenStack struct {
	current  int      // index where the next element will be inserted
	length   int      // the length of the stack
	elements []*Token // holder for the elemnts of the stack
}

// IsEmpty Returns true when the stack is empty, false otherwise
func (s *TokenStack) IsEmpty() bool {
	return s.current <= 0
}

// Size Returns the size of the stack
func (s *TokenStack) Size() int {
	return s.length
}

// Top Returns the top element from the stack, without modifying the stack
func (s *TokenStack) Top() (*Token, error) {

	if s.current <= 0 {
		return nil, errors.New("Stack is empty")
	}

	return s.elements[s.current-1], nil
}

// Pop Pops the top element from the stack and returns it
// On pop, just the indices are updated, no space is released. On subsequent pushes, this space is reused
func (s *TokenStack) Pop() (*Token, error) {

	elem, err := s.Top()
	if err != nil {
		return nil, err
	}

	s.current--
	s.length--

	return elem, nil
}

// Push Pushes an element to the top of the stack
// On push, space is allocated if needed, and element is inserted at index current
func (s *TokenStack) Push(e *Token) error {

	if s.current < len(s.elements) {
		s.elements[s.current] = e
	} else {
		s.elements = append(s.elements, e)
	}

	s.length++
	s.current++

	return nil
}

// Reverse Reverses the stack
func (s *TokenStack) Reverse() error {
	newStack := NewTokenStack()

	s.elements = s.elements[0:s.length]

	for {
		empty := s.IsEmpty()
		if empty {
			break
		}

		elem, err := s.Pop()
		if err != nil {
			return err
		}
		newStack.Push(elem)
	}

	*s = *newStack

	return nil
}

// NewTokenStack Creates an instance of a stack of tokens
func NewTokenStack() *TokenStack {
	return &TokenStack{
		elements: make([]*Token, 0),
		length:   0,
		current:  0,
	}
}

// Print Prints the stack state
func (s *TokenStack) Print() {
	fmt.Println("Current: " + fmt.Sprint(s.current) + ", Size: " + fmt.Sprint(s.Size()))
	fmt.Println(s.elements)
}
