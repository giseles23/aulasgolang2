package main
 
 import (
    "fmt")

func main(){
capitais := map[string]string{
"sp" : "são paulo",
"rj" : "rio janeiro",
"es" : "espirito santo",
}

capitais["ba"] = "bahia"

for k, v := range capitais {
	fmt.Println("sigla, nome", k,v)
}
delete(capitais, "sp")
for k, v := range capitais {
	fmt.Println("sigla, nome", k,v)
}
}