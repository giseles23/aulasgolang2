package main

import ( 
	"fmt"
	"sort"
)

func main(){
	ages := []int {14, 33, 16}
	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 1)
	fmt.Println(index)
	names := []string{"coraline", "terror", "fabiano"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "fabiano"))
x:=3
for x < 5 {
fmt.Println(x)
x++
}
for i:=1; i<7;  i++ {
fmt.Println("segundos ", i)
}
for i:=0; i <len(names); i++{
fmt.Println(names [i])
}

for index, value := range names {
	fmt.Println("o indice é ", index,"e o nome é ", value)
}

}
	

