package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Вывод ошибки,", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)

	arg1, operator, arg2, err := parseExpression(input)
	if err != nil {
		fmt.Println("Вывод ошибки,", err)
		os.Exit(1)
	}

	if arg1 < 1 || arg1 > 10 || arg2 < 1 || arg2 > 10 {
		fmt.Println("Вывод ошибки, оба числа должны быть от 1 до 10 включительно")
		os.Exit(1)
	}

	result := performOperation(arg1, operator, arg2)

	fmt.Println(result)

}
func parseExpression(expr string) (int, string, int, error) {
	parts := strings.Fields(expr)

	if len(parts) == 1 || len(parts) == 2 {
		return 0, "", 0, fmt.Errorf("так как строка не является математической операцией.")
	}

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	arg1, err1 := parseOperand(parts[0])
	if err1 != nil {
		return 0, "", 0, err1
	}

	operator := parts[1]

	arg2, err2 := parseOperand(parts[2])
	if err2 != nil {
		return 0, "", 0, err2
	}

	return arg1, operator, arg2, nil
}

func parseOperand(operand string) (int, error) {
	num, err := strconv.Atoi(operand)
	if err != nil {
		return 0, fmt.Errorf("ошибка преобразования в число: %s", err)
	}
	return num, nil
}

func performOperation(arg1 int, operator string, arg2 int) int {
	switch operator {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":
		return arg1 / arg2
	default:
		fmt.Println("Неподдерживаемый оператор:", operator)
		os.Exit(1)
		return 0
	}
}
