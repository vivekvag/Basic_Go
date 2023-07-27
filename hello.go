package main

import "fmt"

// Basics of Go
// func main() {
// 	fmt.Println("Hello World!")
// }

// Datatypes of GO
// func main() {

// 	var a string = "initial"
// 	fmt.Println(a)

// 	var b, c int = 1, 2
// 	fmt.Println(b, c)

// 	var d = true
// 	fmt.Println(d)

// 	var e int
// 	fmt.Printf("Datatype: %T & Value: %v \n", e, e) //%T for Datatype and %v for value

// 	f := "sky"
// 	fmt.Println(f)
// }

// Looping in Go
// func main() {
// 	sum := 0
// 	for i := 0; i <= 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// Range in GO
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// Map Iteration
	// m := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// Channel Iteration
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
