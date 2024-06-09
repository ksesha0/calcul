package inputloop

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"TestProject1/utils"
)

// бесконечный цикл для ввода
func RunInputLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите арифметическое выражение через пробел (или 'exit' для выхода): ")
		infix, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		infix = strings.TrimSpace(infix)

		if infix == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		rpn := utils.ConvertToRPN(infix)
		result, err := utils.EvaluateRPN(rpn)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Результат выражения '%s' равен: %.2f\n", infix, result)
		}
	}
}
