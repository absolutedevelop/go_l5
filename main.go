package main 

import "fmt"
import "strings"

func main(){

	var name string = "Fortunate"

	fmt.Println("The name is",name)
	fmt.Println("Len of name is",len(name))

	names := []string{"John", "Ashly"}
	fmt.Println("Names are",strings.Join(names,", "))



}