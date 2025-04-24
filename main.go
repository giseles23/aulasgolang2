package main
 
 import (
    "fmt")

func main(){
votos := make(map[string]int)
votos["ana"]+=1
votos["marcelo"]+=1
votos["ana"]+=1
votos["gisele"]+=1
votos["ana"]+=1
votos["marcelo"]+=1

for candidato, total := range votos{
	fmt.Printf("%s recebeu %d votos \n", candidato, total)
}
}