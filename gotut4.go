package main

import "fmt"
//STRUCTURE
type person struct{
	name string
	age uint16
	profession string
}
func main(){

	person1 := person{name:"abhinav",
					  age:21,
					  profession:"intern"}
	fmt.Println(person1.name)				  

}

