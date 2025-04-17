package main
 
 import (
    "fmt")


func main(){
    alunoidade := make(map[string]int)
	alunoidade ["fabiano"] = 16
	alunoidade ["gisele"] = 18
	alunoidade ["abobora"] = 3
	alunoidade ["xerequinha"] = 15
	fmt.Println("idade da xana", alunoidade["xerequinha"])
	notasalunos := map[string]float64{
		"fabiano" : 7.9,
		"abobora" : 10,
		"xerequinha" : 3.8,
		"gisele" : 10,
	} 
	for nome, nota := range notasalunos{
fmt.Printf("%s tirou a nota %.1f\n", nome, nota )
	}
}