package main

import ("fmt"
	    "math";"math/rand")
func abhi () {
    fmt.Println("a random no. bw 1-100",rand.Intn(100))

}
func main() {
	fmt.Println("square root of 4 is",math.Sqrt(4))
	abhi()
}