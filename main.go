package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Вывод ошибки,", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	result, error := calculator(input)
	if error != nil {
		fmt.Println("Вывод ошибки,", err)
		os.Exit(1)
	}
	fmt.Printf(result)
}

func calculator(input string) (string, error) {
	arg1, operator, arg2, isRoman, err := parseExpression(input)
	if err != nil {
		return fmt.Println("Вывод ошибки,", err)
	}

	if arg1 < 1 || arg1 > 10 || arg2 < 1 || arg2 > 10 {
		fmt.Println("Вывод ошибки, оба числа должны быть от 1 до 10 включительно")
	}

	result := performOperation(arg1, operator, arg2)

	if isRoman {
		romanResult, romanErr := arabicToRoman(result)
		if romanErr != nil {
			fmt.Println("Вывод ошибки,", romanErr)
		}
		return romanResult, nil
	} else {
		return fmt.Sprint(result), nil
	}
}

func parseExpression(expr string) (int, string, int, bool, error) {
	parts := strings.Fields(expr)

	if len(parts) == 1 || len(parts) == 2 {
		return 0, "", 0, false, fmt.Errorf("так как строка не является математической операцией.")
	}

	if len(parts) != 3 {
		return 0, "", 0, false, fmt.Errorf("так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	arg1, err1 := parseOperand(parts[0])
	if err1 != nil {
		return 0, "", 0, false, err1
	}

	operator := parts[1]

	arg2, err2 := parseOperand(parts[2])
	if err2 != nil {
		return 0, "", 0, false, err2
	}

	isRoman := isRomanNumeral(parts[0]) && isRomanNumeral(parts[2])

	if (isRoman && !isRomanNumeral(parts[2])) || (!isRoman && isRomanNumeral(parts[2])) || (isRoman && !isRomanNumeral(parts[0])) || (!isRoman && isRomanNumeral(parts[0])) {
		return 0, "", 0, false, fmt.Errorf("используйте либо арабские, либо римские числа")
	}

	return arg1, operator, arg2, isRoman, nil
}

func parseOperand(operand string) (int, error) {
	if isRomanNumeral(operand) {
		return romanToInt(operand)
	} else {
		num, err := strconv.Atoi(operand)
		if err != nil {
			return 0, fmt.Errorf("ошибка преобразования в число: %s", err)
		}
		return num, nil
	}
}

func isRomanNumeral(s string) bool {
	if _, ok := romanNumerals[s]; !ok {
		return false
	}
	return true
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

func romanToInt(s string) (int, error) {

	if _, ok := romanNumerals[s]; ok {
		return romanNumerals[s], nil
	} else {
		return 0, fmt.Errorf("ошибка преобразования в число: %s", s)
	}

}

func arabicToRoman(num int) (string, error) {

	if num < 1 {
		return "", fmt.Errorf("так как в римской системе нет отрицательных чисел. - %d", num)
	}
	romanSymbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	romanValues := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}

	result := ""

	for i := len(romanValues) - 1; i >= 0; i-- {
		for num >= romanValues[i] {
			result += romanSymbols[i]
			num -= romanValues[i]
		}
	}
	return result, nil
}
