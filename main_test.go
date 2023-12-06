package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {

	testCase := []struct {
		input, expected string
	}{
		{"1 + 2", "3"},
		{"VI / III", "II"},
		{"I - II", "Вывод ошибки, так как в римской системе нет отрицательных чисел."},
		{"I + 1", "Вывод ошибки, так как используются одновременно разные системы счисления."},
		{"1", "Вывод ошибки, так как строка не является математической операцией."},
		{"1 + 2 + 3", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."},
	}

	for _, tc := range testCase {

		result := calculator(tc.input)

		if result != tc.expected {
			t.Errorf("Для %s ожидается %s, получено %s", tc.input, tc.expected, result)
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
