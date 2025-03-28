package main

import (
	"fmt"
)

func main(){
	var usuario string
	var senha int
cont := 0
for {
	fmt.Println("usuario: ")
	fmt.Scan(&usuario)
	fmt.Println("senha: ")
	fmt.Scan(&senha)
	if usuario != "admin" || senha != 1234 {
		fmt.Println("senha ou usuário incorretos tente denovo")
	} else {cont ++
		fmt.Println("pode entrar")
	 } 
	 if cont == 1 {
		break
	 }
	}
}

