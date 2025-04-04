package main

import ( 
	"fmt"
)

func main(){  // Vetores fornecidos
    ages := []int{16, 14, 43}
    nomes := []string{"fabiano", "coraline", "terror"}

    // Loop para associar os índices dos vetores
    for i := 0; i < len(ages); i++ {
     // Associa o índice do vetor "ages" com o índice do vetor "nomes"
        fmt.Printf("Nome: %s, Idade: %d\n", nomes[i], ages[i])
    }
}
	

