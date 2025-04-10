package main

import ( 
	"fmt"
)

func main(){ 
    numeros := [5]int {}
    fmt.Println("digite 5 números")
    fmt.Scan(&numeros[0])
        fmt.Scan(&numeros[1])
            fmt.Scan(&numeros[2])
                fmt.Scan(&numeros[3])
                    fmt.Scan(&numeros[4])
    fmt.Println("a soma é ", numeros[0]+ numeros[1] + numeros[2] + numeros[3] + numeros[4])
}
	

