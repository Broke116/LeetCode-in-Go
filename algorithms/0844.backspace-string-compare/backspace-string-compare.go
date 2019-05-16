package problem0844

import "fmt"

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// BackspaceCompare is the solution of the problem 0844
func BackspaceCompare(S string, T string) bool {
	stack1 := make(stack, 0)
	stack2 := make(stack, 0)

	fmt.Println("stack1 length ", len(stack1))
	fmt.Println("stack2 length ", len(stack2))

	stack1 = appendStack(stack1, S)
	stack2 = appendStack(stack2, T)

	fmt.Println("stack1 length ", len(stack1))
	fmt.Println("stack2 length ", len(stack2))

	if len(stack1) != len(stack2) {
		return false
	}

	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}

func appendStack(stack stack, V string) []string { // #a#c
	for i := range V {
		stack = stack.Push(string(V[i]))
		if V[i] == 35 { // #
			stack, _ = stack.Pop()
			if len(stack) > 0 {
				stack, _ = stack.Pop()
			}
		}
	}
	return stack
}
