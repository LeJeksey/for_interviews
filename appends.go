package main

import (
	"fmt"
	"strings"
)

// Если мы будем добавлять в слайс новые элементы не вызывая при этом расширения - append перепишут последние значения
func appends() {
	x := make([]int, 0, 4)

	x = append(x, 1, 2, 3)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Println(strings.Repeat("-", 20))

	y := append(x, 4)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
	fmt.Println(strings.Repeat("-", 20))

	z := append(x, 5)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
	fmt.Printf("%v len=%d cap=%d\n", z, len(z), cap(z))
}

// НО, если мы будем добавлять в слайс новые элементы вызывая при этом расширения -
// append создаст для новых слайсов новые массивы
func butAppends() {
	x := make([]int, 0, 3) // выделяем 3 элемента, чтобы произошло расширение и перенос массива в памяти

	x = append(x, 1, 2, 3)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Println(strings.Repeat("-", 20))

	y := append(x, 4)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
	fmt.Println(strings.Repeat("-", 20))

	z := append(x, 5)
	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
	fmt.Printf("%v len=%d cap=%d\n", z, len(z), cap(z))
}

// Но не нужно использовать append для создания нового слайса/массива,
// append принято использовать только для добавления элементов в существующий слайс
