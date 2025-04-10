package main

import ( 
	"fmt"
)

func main(){ 
    var idade int
fmt.Println("qual sua idade?")
fmt.Scan(&idade)
if idade <18 {
    fmt.Println("voce é menor de idade")
} else if idade >=18 && idade <= 60 { 
    fmt.Println("você é adulto")
} else {
fmt.Println("voce é idoso")
}
}



