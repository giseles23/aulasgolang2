package main
 
 import (
    "fmt")

func main(){
estoque := make(map[string]int)
estoque["coxinha"]=10
estoque["pão de queijo"]=15
estoque["refrigerante"]=20

for produtos, quantidade := range estoque {
	fmt.Printf("%s tem %d unidades sobrando \n", produtos, quantidade)
}
fmt.Println("------------------- vendas --------------------")
estoque["coxinha"]-=2
estoque["pão de queijo"]-=1
for produtos, quantidade := range estoque {
	fmt.Printf("%s agora tem %d unidades sobrando \n", produtos, quantidade)
}
}