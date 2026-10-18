package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		num1, num2 float64
	)
	fmt.Println("《 Простой калькулятор на Go 》")
	num1 = number1()
	num2 = number2()
	sum(num1, num2)
	difference(num1, num2)
	product(num1, num2)
	quotient(num1, num2)
	power(num1, num2)
}

func number1() float64 {
	var num1 float64
	fmt.Print("✦› Введите первое число: ")
	fmt.Scan(&num1)
	if num1 == 0 {
		fmt.Println("✦› ❌ Ошибка: введено некорректное значение! Пожалуйста, введите число.")
		return number1()
	}
	return num1
}

func number2() float64 {
	var num2 float64
	fmt.Print("✦› Введите второе число: ")
	fmt.Scan(&num2)
	if num2 == 0 {
		fmt.Println("✦› ❌ Ошибка: введено некорректное значение! Пожалуйста, введите число.")
		return number2()
	}
	return num2
}

func sum(num1, num2 float64) float64 {
	fmt.Println("┌✦› Сумма чисел:", num1, "+", num2, "=", num1+num2)
	return num1 + num2
}

func difference(num1, num2 float64) float64 {
	fmt.Println("├✦› Разность чисел:", num1, "-", num2, "=", num1-num2)
	return num1 - num2
}

func product(num1, num2 float64) float64 {
	fmt.Println("├✦› Произведение чисел:", num1, "*", num2, "=", num1*num2)
	return num1 * num2
}

func quotient(num1, num2 float64) float64 {
	fmt.Println("├✦› Частное чисел:", num1, "/", num2, "=", num1/num2)
	return num1 / num2
}

func power(num1, num2 float64) float64 {
	fmt.Println("└✦› Возведение в степень:", num1, "^", num2, "=", math.Pow(num1, num2))
	return math.Pow(num1, num2)
}
