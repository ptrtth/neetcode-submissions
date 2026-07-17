func execOp(op string, a int, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		number, err := strconv.Atoi(token)

		if err == nil {
			stack = append(stack, number)
			continue
		}

		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack = append(stack, execOp(token, a, b))
	}

	return stack[0]
}