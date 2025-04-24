package main
 
 import (
    "fmt")

func analisarNotas()(float64, string){
	var nota1, nota2 float64
	fmt.Println("nota 1: ")
	fmt.Scan(&nota1)
	fmt.Println("nota 2: ")
	fmt.Scan(&nota2)
	var media float64 = (nota1 + nota2)/2
	if (media>=6){
		return media, "aprovado"
	} else if (media<6){
		return media, "reprovado"
	} else {
	return media, "nota invalida"
	}
}
func main(){
media, resultado := analisarNotas()
fmt.Printf("média: %.1f \n ", media)
fmt.Println("resultado:", resultado)
}
