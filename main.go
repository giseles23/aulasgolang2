package main

import (
	"fmt"
)

func main(){
	var num1 int
	var num2 int
fmt.Println("digite um número ")
fmt.Scan(&num1) 
fmt.Println("digite outro número ")
fmt.Scan(&num2)
fmt.Println("a soma do primeiro número com o segundo é", num1 + num2)
fmt.Println("a multiplicação deles dois é ", num1 * num2)
fmt.Println("a subtraçao do primeiro número com o segundo é", num1 - num2)
fmt.Println("a divisão do número 1 pelo número 2 é ", num1 / num2)
fmt.Println("o resto dessa divisão é ", num1 % num2)
num1++
fmt.Println("a incrementação no número 1 é ", num1)
if num1 > 0 && num2 > 0 {
	fmt.Println("números positivos")
}
}