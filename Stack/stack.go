package stack

type stack []string

type IStack interface {
	Push(val string)
	Pop() (string, int) //popValue, lastindex int
	Len() int
	IsEmpty() bool
	Peek() string
}

func InitStack() IStack {
	return &stack{}
}

// IsEmpty implements IStack.
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Len implements IStack.
func (s *stack) Len() int {
	return len(*s)
}

// Peek implements IStack.
func (s *stack) Peek() string {
	lastindex := (*s).Len() - 1
	peek := (*s)[lastindex]
	return peek
}

// Pop implements IStack.
func (s *stack) Pop() (string, int) {
	if (*s).IsEmpty() {
		return "", -1
	}

	lastindex := (*s).Len() - 1
	popVal := (*s)[lastindex]
	*s = (*s)[:lastindex]
	return popVal, lastindex

}

// Push implements IStack.
func (s *stack) Push(val string) {
	*s = append((*s), val)
}
