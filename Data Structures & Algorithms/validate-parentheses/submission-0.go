type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	s.items = s.items[:last]

	return value, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func isValid(s string) bool {
    op := Stack[rune]{}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			op.Push(c)
		} else {
			current, ok := op.Pop()
			if !ok {
				return false
			}
			if (current == '(' && c != ')') || (current == '[' && c != ']') || (current == '{' && c != '}') {
				return false
			}
		}
	}
	return op.Len() == 0
}
