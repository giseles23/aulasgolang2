package main
 
 import (
    "fmt")

func dadosPessoa(idade int) (int, string){
var condicao string
if idade >= 18 {
    condicao = "você é maior de idade"
} else {
    condicao = "você é menor de idade"
}
return idade, condicao
}

func main(){
    var idadeusuario int
fmt.Println("qual sua idade?")
fmt.Scan(&idadeusuario)
idade, condicao := dadosPessoa(idadeusuario)
fmt.Println("você tem", idade, "anos e", condicao)
}