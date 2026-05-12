package main

import (
	"fmt"
	"math"
)

// Функция для расчета значения Y
func calculateY(x, a, b float64) float64 {
	// Условие из задания: x < 5 и x >= 5
	if x < 5 {
		// Первая формула: y = lg^2(a^2 + x) / (a + x)^2
		numerator := math.Pow(math.Log10(math.Pow(a, 2)+x), 2)
		denominator := math.Pow(a+x, 2)
		return numerator / denominator
	} else {
		// Вторая формула: y = (a + bx)^3.5 / (1.8 + cos^3(ax))
		numerator := math.Pow(a+(b*x), 3.5)
		denominator := 1.8 + math.Pow(math.Cos(a*x), 3)
		return numerator / denominator
	}
}

func main() {
	// Входные данные из таблицы
	const a = -2.5
	const b = 3.4

	// --- Задача А: Использование цикла for для табулирования ---
	xStart := 3.5
	xEnd := 6.5
	dx := 0.6

	fmt.Println("=== Задача А: Табулирование функции ===")
	fmt.Printf("| %-10s | %-15s |\n", "x", "y")
	fmt.Println("|------------|-----------------|")

	// В Go цикл while заменяется на for
	for x := xStart; x <= xEnd+0.0001; x += dx {
		y := calculateY(x, a, b)
		fmt.Printf("| %-10.1f | %-15.6f |\n", x, y)
	}

	// --- Задача B: Использование массива (слайса) и цикла range ---
	xValues := []float64{2.89, 3.54, 5.21, 6.28, 3.48}

	fmt.Println("\n=== Задача B: Расчет по массиву значений ===")
	fmt.Printf("| %-10s | %-15s |\n", "Аргумент", "Результат")
	fmt.Println("|------------|-----------------|")

	// Используем range для перебора элементов массива
	for i, x := range xValues {
		y := calculateY(x, a, b)
		fmt.Printf("| X%-9d | %-15.6f |\n", i+1, y)
	}
}