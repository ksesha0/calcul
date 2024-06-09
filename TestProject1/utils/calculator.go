package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// создается польская запись
func ConvertToRPN(infix string) string {
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	var rpn strings.Builder
	var operatorStack Stack

	tokens := strings.Fields(infix)
	for _, token := range tokens {
		if isNumber(token) {
			rpn.WriteString(token)
			rpn.WriteString(" ")
		} else if token == "(" {
			operatorStack.Push(token)
		} else if token == ")" {
			for !operatorStack.IsEmpty() && operatorStack.Top() != "(" {
				rpn.WriteString(operatorStack.Pop())
				rpn.WriteString(" ")
			}
			operatorStack.Pop() // Удаляем "(" из стека
		} else {
			for !operatorStack.IsEmpty() && precedence[operatorStack.Top()] >= precedence[token] {
				rpn.WriteString(operatorStack.Pop())
				rpn.WriteString(" ")
			}
			operatorStack.Push(token)
		}
	}

	for !operatorStack.IsEmpty() {
		rpn.WriteString(operatorStack.Pop())
		rpn.WriteString(" ")
	}

	return strings.TrimSpace(rpn.String())
}

// подсчет в польской записи
func EvaluateRPN(rpn string) (float64, error) {
	var operandStack Stack

	tokens := strings.Fields(rpn)
	for _, token := range tokens {
		if isNumber(token) {
			operandStack.Push(token)
		} else {
			operand2, _ := strconv.ParseFloat(operandStack.Pop(), 64)
			operand1, _ := strconv.ParseFloat(operandStack.Pop(), 64)
			result, err := Calculate(operand1, operand2, token)
			if err != nil {
				return 0, err
			}
			operandStack.Push(strconv.FormatFloat(result, 'f', -1, 64))
		}
	}

	result, _ := strconv.ParseFloat(operandStack.Pop(), 64)
	return result, nil
}

func Calculate(operand1, operand2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("неверный оператор")
	}
}

// число???
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
