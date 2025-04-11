package main

import ( 
	"fmt"
)

func sayGreeting(nome string) {
fmt.Println("olá!", nome)
}

func AddNumber(numero1 int, numero2 int) int {
return numero1 + numero2
}

func main(){ 
    sayGreeting("fabiano")
resultado := AddNumber(10, 23)
fmt.Println(resultado)
}
