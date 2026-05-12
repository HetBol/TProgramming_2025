package main

import "fmt"

// описание структуры
type Computer struct {
	Brand    string  // марка
	Model    string  // модель
	HDDSpace int     // объем жесткого диска в ГБ
}

// метод-конструктор
// возвращает указатель на новую структуру Computer
func NewComputer(brand, model string, hdd int) *Computer {
	return &Computer{
		Brand:    brand,
		Model:    model,
		HDDSpace: hdd,
	}
}

// метод (геттер) для получения объема диска
func (c *Computer) GetHDD() int {
	return c.HDDSpace
}

//  метод (сеттер) для установки нового объема диска
func (c *Computer) SetHDD(newSpace int) {
	if newSpace > 0 {
		c.HDDSpace = newSpace
	} else {
		fmt.Println("Ошибка: объем диска должен быть положительным")
	}
}

// полезный метод для вывода информации
func (c *Computer) PrintInfo() {
	fmt.Printf("Компьютер: %s %s, Диск: %d ГБ\n", c.Brand, c.Model, c.HDDSpace)
}

func main() {
	// используем конструктор для создания объекта
	myComp := NewComputer("Asus", "ROG", 512)

	// выводим начальную информацию
	fmt.Println("Начальные данные:")
	myComp.PrintInfo()

	// используем сеттер, чтобы изменить объем диска
	fmt.Println("\nОбновляем объем диска...")
	myComp.SetHDD(1024)

	// используем геттер, чтобы получить значение и вывести его
	currentHDD := myComp.GetHDD()
	fmt.Printf("Новый объем диска через геттер: %d ГБ\n", currentHDD)

	// финальный вывод
	myComp.PrintInfo()
}