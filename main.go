package main

import ( 
	"fmt"
"strings"
 "sort")

func main(){

greeting := "hello meu amigão!"
fmt.Println(strings.ToUpper(greeting))
fmt.Println(strings.Contains(greeting, "amigão"))
fmt.Println(strings.ReplaceAll(greeting, "hello", "olá"))
fmt.Println(strings.Index(greeting, "amigão"))
ages := []int {3, 5, 9, 1, 45}
fmt.Println(ages)
sort.Ints(ages)
fmt.Println(ages)
index := sort.SearchInts(ages, 1)
fmt.Println(index)
names := []string{"coraline", "terror", "fabiano"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names, "terror"))
}

