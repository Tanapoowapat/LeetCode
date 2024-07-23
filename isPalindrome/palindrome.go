package ispalindrome

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}
func (s *Stack) IsEmpty() bool { return len(*s) == 0 }

func (s *Stack) Len() int { return len(*s) }

func (s *Stack) Pop() (string, int) {
	if len(*s) == 0 {
		return "", -1
	}

	lastindex := len(*s) - 1
	pop := (*s)[lastindex]
	*s = (*s)[:lastindex]
	return pop, lastindex
}

func InitStack(str string) Stack {
	stack := Stack{}

	for _, c := range str {
		stack.Push(string(c))
	}

	return stack
}

func isPalindrome(x string) (string, bool) {
	s := InitStack(x)
	newX := ""
	for _, c := range s {

		//fmt.Printf("Stack Val : %v\n", s)
		temp, _ := s.Pop()

		if c == "#" && temp == "#" {
			c, temp = "z", "z"
		}

		if c != temp {
			if temp == "#" {
				temp = c
			} else if c == "#" {
				c = temp
			} else {
				return "", false
			}
		}
		newX += c
	}
	return newX, true
}
