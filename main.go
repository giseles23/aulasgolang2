package main

import (
	"fmt"
)

func main(){
	var usuario string
	var senha int
fmt.Println("digite seu nome de usuário")
fmt.Scan(&usuario)
fmt.Println("digite sua senha")
fmt.Scan(&senha)
if usuario == "admin" && senha == 1234 {
	fmt.Println("pode entrar")
} else {
	fmt.Println("acesso negado")
}
}