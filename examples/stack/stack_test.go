//This is a test file.
//A test file is signified by the suffix _test on a go file.
//There are two ways of doing testing in go: use go examples
//or the go testing library.
//This files uses go examples.

package stack

import (
	"fmt"
)

//This is a example function.
//A example function is signified by the prefix example
//The comments in the function specify what the acceptable output is.
//If the output does not match, the example is a failure.
func ExampleStack_Push() {
	var stack = NewStack(2)
	fmt.Println(stack.Push(1))
	fmt.Println(stack.Push(2))
	fmt.Println(stack.Push(3))
	//Output:
	// <nil>
	// <nil>
	// Error: the stack is full
}

func ExampleStack_Pop() {
	var stack = NewStack(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	//Output:
	// 3 <nil>
	// 2 <nil>
	// 1 <nil>
}

func ExampleStack_Reverse() {
	var stack = NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Reverse()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	//Output:
	// 1 <nil>
	// 2 <nil>
	// 3 <nil>
}

func ExampleStack_Clear() {
	var stack = NewStack(10)
	stack.Push(1)
	stack.Clear()
	fmt.Println(stack.Pop())
	//Output:
	// <nil> Error: the stack is empty
}

func ExampleStack_IsEmpty() {
	var stack = NewStack(10)
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	stack.Clear()
	fmt.Println(stack.IsEmpty())
	//Output:
	// false
	// true
}
