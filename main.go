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
}