package main

import ( 
	"fmt"
)

func main(){ 
age:= 45
fmt.Println(age <= 50)
fmt.Println(age >= 50)
fmt.Println(age == 50)
fmt.Println(age != 50)

if age < 30 {
    fmt.Println("menor que 30 anos")
} else if age <40 { 
    fmt.Println("menor que 40 anos")
} else { fmt.Println("não é menor que 40")
}
names := []string {"fabiano", "batman", "gisele", "cascão", "holmes", "bem 10"}
for index, valor:= range names {
    if index == 1 {
        fmt.Println("conntinua após a posição", index,  "e nome", valor)
    continue
    }
    if index > 2 {
        fmt.Println("sair antes de", valor)
        break
    }
fmt.Println("nome:", valor)
}
}



