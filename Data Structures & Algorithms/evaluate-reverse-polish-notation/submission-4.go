func evalRPN(tokens []string) int {
	operations := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	stack := []int{}

	for _, token := range tokens {
		if number, err := strconv.Atoi(token); err == nil {
			stack = append(stack, number)
			continue
		}

		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		stack = append(stack, operations[token](a, b))
	}

	return stack[len(stack)-1]
}