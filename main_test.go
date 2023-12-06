package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {

	testCase := []struct {
		input, error, expected string
	}{
		{"1 + 2", "", "3"}, // 1 - вырожение, 2 - ошибка, 3 - ответ
		{"VI / III", "", "II"},
		{"I - II", "Вывод ошибки, так как в римской системе нет отрицательных чисел.", ""},
		{"I + 1", "Вывод ошибки, так как используются одновременно разные системы счисления." , ""},
		{"1", "Вывод ошибки, так как строка не является математической операцией.", ""},
		{"1 + 2 + 3", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).", ""},
	}

	for _, tc := range testCase {

		result, error := calculator(tc.input)

		if result != tc.expected {
			t.Errorf("Для %s ожидается %s, ошибка %s, получено %s", tc.input, error, tc.expected, result)
		}
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Инициализация тестов")
	exitCode := m.Run()
	fmt.Println("Выполнение тестов завершено")
	fmt.Println("Выходной код:", exitCode)
}

func SetUp() {
	fmt.Println("Настройка перед тестом")
}

func TearDown() {
	fmt.Println("Очистка после теста")
}
