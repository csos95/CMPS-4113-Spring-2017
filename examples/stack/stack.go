//This example implements a basic stack datastructure and tests it
package stack

import (
	"fmt"
)

//Stack is a simple LIFO datastructure
type Stack struct {
	Items []interface{}
	Top   int
}

//String returns a string representation of the stack
func (s Stack) String() string {
	var output = "Stack: "
	for _, item := range s.Items {
		output += fmt.Sprintf(" %v", item)
	}
	return output
}

//NewStack returns a new stack of the specified size
func NewStack(size int) Stack {
	return Stack{make([]interface{}, size), 0}
}

//Push an item onto the stack if there is space
func (s *Stack) Push(item interface{}) error {
	if s.Top < len(s.Items) {
		s.Items[s.Top] = item
		s.Top++
		return nil
	}
	return fmt.Errorf("Error: the stack is full\n")
}

//Pop an item off of the stack if there are any
func (s *Stack) Pop() (interface{}, error) {
	var item interface{}
	if s.Top > 0 {
		s.Top--
		item = s.Items[s.Top]
		return item, nil
	}
	return nil, fmt.Errorf("Error: the stack is empty\n")
}

//Reverse the stack
func (s *Stack) Reverse() {
	var top, bottom int = s.Top - 1, 0
	for top > bottom {
		s.Items[top], s.Items[bottom] = s.Items[bottom], s.Items[top]
		top--
		bottom++
	}
}

//Clear the stack
func (s *Stack) Clear() {
	s.Items = make([]interface{}, len(s.Items))
	s.Top = 0
}

//IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Top == 0
}
