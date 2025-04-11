package main

import
 ( 
	"fmt"
)

var opcao int
var saldo float64
var opcao2 int
func Inicio(){
    fmt.Println("Digite seu saldo")
    fmt.Scan(&saldo)
}
func Opcoes(){

    fmt.Println("escolha sua opção: 'sacar'(1) ou 'depositar'(2) - 'ver saldo'(3)")
    fmt.Scan(&opcao)
}
func Sacar() {
    var valor float64
   if opcao == 1 {
fmt.Println("digite quanto você quer sacar")
fmt.Scan(&valor)
   
if valor <= saldo {
    saldo = saldo - valor
    fmt.Printf("saldo suficiente, saldo atual é %.2f\n", saldo)
} else { fmt.Println("saldo insuficiente")
}
}
}
func Depositar(){
    var valor float64
    if opcao == 2 {
 fmt.Println("digite quanto você quer depositar")
 fmt.Scan(&valor)
 saldo = valor + saldo
 fmt.Printf("seu saldo atual é %.2f\n", saldo)
 
    }
}
func VerSaldo(){
    if opcao == 3 {
        fmt.Printf("seu saldo atual é %.2f\n", saldo)

    }
}

func SairOUContinuar(){
for {
    Opcoes()
Sacar()
Depositar()
VerSaldo()
fmt.Println("'sair'(1) 'continuar'(2)")
fmt.Scan(&opcao2)
if opcao2 == 1 {
    break
} else if opcao2 == 2 {
    continue
}
}
}

func main(){ 
    Inicio()
    SairOUContinuar()
}
