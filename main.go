package main 

import "fmt"

func main(){
	a, b := 10, 3 
fmt.Println("a soma é ", a + b)
fmt.Println("a multiplicação é ", a * b)
fmt.Println("subtrair é ", a - b)
fmt.Println("divisão é ", a / b)
fmt.Println("o resto da divisão é", a % b)

a++
fmt.Println("incrementar 1 ao valor de a o faz valer ", a)
if a > 0 && b > 0 {
	fmt.Println("números positivos")
}
}