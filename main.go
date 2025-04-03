package main

import (
	"fmt"
)

func main(){
	var ages = [4]int{17, 16, 20, 40}
	nomes := [4]string{"luisa mel", "fabiano", "gisele", "rafael"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "garfield"
	fmt.Println(nomes)
	// oi fabiano
	var score = []int{100, 200, 300, 400}
	fmt.Println(score)
	score[0]= 101
	fmt.Println(score, len(score), cap(score))
rangeOne := score[1:4]
fmt.Println(rangeOne)
fmt.Println(rangeOne, len(rangeOne), cap(rangeOne))
rangeTwo := score[2:]
fmt.Println(rangeTwo, len(rangeTwo), cap(rangeTwo))
rangeThree := score[:3]
fmt.Println(rangeThree, len(rangeThree), cap(rangeThree))

var nomes2 = []string {"gisele2", "chico", "cachorro"}
fmt.Println(nomes2)
nomes2 = append(nomes2, "nome", "blabla", "klgjdsd")
fmt.Println(nomes2, len(nomes2), cap(nomes2))
nomes2[4] = "fascinado"
fmt.Println(nomes2, len(nomes2), cap(nomes2))
}

