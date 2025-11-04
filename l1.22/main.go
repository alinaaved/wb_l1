package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) (*big.Int, error) {
	if b.Sign() == 0 {
		return nil, fmt.Errorf("деление на ноль")
	}
	return new(big.Int).Div(a, b), nil
}

func Parse(s string) *big.Int {
	z, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("не число: " + s)
	}
	return z
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var sa, sb string
	fmt.Print("Введите числа a и b через пробел: ")
	fmt.Fscan(in, &sa, &sb)

	a := Parse(sa)
	b := Parse(sb)

	fmt.Printf("Сложение:  %s + %s = %s\n", a, b, add(a, b))
	fmt.Printf("Вычитание: %s - %s = %s\n", a, b, sub(a, b))
	fmt.Printf("Умножение: %s * %s = %s\n", a, b, mul(a, b))

	if q, err := div(a, b); err == nil {
		fmt.Printf("Деление (целое): %s : %s = %s\n", a, b, q)
	} else {
		fmt.Println("Деление (целое): ошибка:", err)
	}
}
