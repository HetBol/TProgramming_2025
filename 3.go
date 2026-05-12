// Лабораторная работа 3
package main

import (
	"fmt"
	"math"
)

// Функция для вычисления y по заданным формулам
func calculateY(x, a, b float64) float64 {
	if x < 5 {
		// Первая формула: y = lg^2(a^2 + x) / (a + x)^2
		// math.Log10 — десятичный логарифм
		// math.Pow — возведение в степень
		numerator := math.Pow(math.Log10(math.Pow(a, 2)+x), 2)
		denominator := math.Pow(a+x, 2)
		return numerator / denominator
	} else {
		// Вторая формула: y = (a + bx)^3.5 / (1.8 + cos^3(ax))
		// math.Cos — косинус (аргумент в радианах)
		numerator := math.Pow(a+(b*x), 3.5)
		denominator := 1.8 + math.Pow(math.Cos(a*x), 3)
		return numerator / denominator
	}
}

func main() {
	// Данные из таблицы
	const a = -2.5
	const b = 3.4

	// --- Задача А: Табулирование функции (цикл с шагом) ---
	xStart := 3.5
	xEnd := 6.5
	dx := 0.6

	fmt.Println("=== Задача А: Цикл от Xн до Хк ===")
	fmt.Printf("%-10s | %-15s\n", "x", "y")
	fmt.Println("-------------------------------")

	for x := xStart; x <= xEnd+0.0001; x += dx {
		resY := calculateY(x, a, b)
		fmt.Printf("%-10.1f | %-15.6f\n", x, resY)
	}

	// --- Задача B: Вычисление для конкретных значений X ---
	// Используем слайс (массив) для хранения X1, X2, X3, X4, X5
	xValues := []float64{2.89, 3.54, 5.21, 6.28, 3.48}

	fmt.Println("\n=== Задача B: Набор значений X1-X5 ===")
	fmt.Printf("%-10s | %-15s\n", "Аргумент", "Результат y")
	fmt.Println("-------------------------------")

	for i, x := range xValues {
		resY := calculateY(x, a, b)
		fmt.Printf("X%d (%-5.2f) | %-15.6f\n", i+1, x, resY)
	}
}