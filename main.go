package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		printError()
		return
	}

	expression := os.Args[1]
	tokens := tokenize(expression)

	if len(tokens) == 0 {
		printError()
		return
	}

	stack := []int{}

	for _, token := range tokens {
		if isNumber(token) {
			num := parseInt(token)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				printError()
				return
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := applyOperation(a, b, token)
			stack = append(stack, result)
		} else {
			printError()
			return
		}
	}

	if len(stack) != 1 {
		printError()
	} else {
		printInt(stack[0])
		z01.PrintRune('\n')
	}
}

func tokenize(input string) []string {
	var tokens []string
	var currentToken []rune

	for _, r := range input {
		if r == ' ' {
			if len(currentToken) > 0 {
				tokens = append(tokens, string(currentToken))
				currentToken = nil
			}
		} else {
			currentToken = append(currentToken, r)
		}
	}

	if len(currentToken) > 0 {
		tokens = append(tokens, string(currentToken))
	}

	return tokens
}

func isNumber(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func parseInt(s string) int {
	var result int
	for _, r := range s {
		result = result*10 + int(r-'0')
	}
	return result
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/", "%":
		return true
	}
	return false
}

func applyOperation(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			printError()
			os.Exit(1)
		}
		return a / b
	case "%":
		if b == 0 {
			printError()
			os.Exit(1)
		}
		return a % b
	}
	return 0
}

func printError() {
	for _, r := range "Error\n" {
		z01.PrintRune(r)
	}
}

func printInt(num int) {
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}
	var digits []int
	if num == 0 {
		digits = append(digits, 0)
	}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}
	for _, digit := range digits {
		z01.PrintRune(rune('0' + digit))
	}
}
