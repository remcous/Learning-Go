package main

import "fmt"

// Without a select statement will deadlock
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go func() {
// 		v := 1
// 		ch1 <- v
// 		v2 := <-ch2
// 		fmt.Println(v, v2)
// 	}()
// 	v2 := <-ch1
// 	v := 2
// 	ch2 <- v
// 	fmt.Println(v, v2)
// }

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}
