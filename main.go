package main

import (
	"fmt"
)

func main(){
	var usuario string
	var senha string
cont := 0
for {
	fmt.Println("usuario: ")
	fmt.Scan(&usuario)
	fmt.Println("senha: ")
	fmt.Scan(&senha)
	if usuario != "admin" && senha == "1234" {
		fmt.Println(" usuário incorreto, tente novamente")
	}
	if usuario == "admin" && senha != "1234" {
		fmt.Println("senha incorreta, tente novamente")
	}
	if usuario != "admin" && senha != "1234" {
		fmt.Println("usuário e senha incorretos, tente novamente")
	}
	if usuario == "admin" && senha == "1234" {cont ++
		fmt.Println("pode entrar")
	 } 
	 if cont == 1 {
		break
	 }
	}
}

