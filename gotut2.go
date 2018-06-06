package main

import ("fmt")

func main(){
//pointers
x :=5
a :=&x
fmt.Println(x)//5
fmt.Println(a)//memory addresas
fmt.Println(*a)//value at that memory address
*a = 15
fmt.Println(x)
*a = *a**a
fmt.Println(x)
}