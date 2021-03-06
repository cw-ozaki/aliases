package types

// Stack is the FILO data structure.
type Stack struct {
	hash  Hasher
	slice []interface{}
}

// Push adds data to last of Stack.
func (stack *Stack) Push(v interface{}) {
	stack.slice = append(stack.slice, v)
}

// Pop get data from first.
func (stack *Stack) Pop() interface{} {
	if len(stack.slice) == 0 {
		return nil
	}
	v := stack.slice[len(stack.slice)-1]
	stack.slice = stack.slice[:len(stack.slice)-1]
	return v
}

// Slice converts from Stack to Slice.
func (stack *Stack) Slice() []interface{} {
	slice := make([]interface{}, len(stack.slice))
	for i := len(stack.slice) - 1; i >= 0; i-- {
		slice[len(stack.slice)-i-1] = stack.slice[i]
	}
	return slice
}

// NewStack creates a new Stack.
func NewStack(hasher Hasher) *Stack {
	if hasher == nil {
		hasher = MD5
	}
	return &Stack{hasher, make([]interface{}, 0)}
}
