package main

import "fmt"

// update fails because it changes address pointed to by px which is passed by
// value to function. After function ends the original is unchanged
func failedUpdate(px *int) {
	x := 20
	px = &x
}

// Uses dereference to change value pointed to by px, so this change stays
func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // Prints 10, update doesn't change outer value of x

	update(&x)
	fmt.Println(x) // Prints 20, update changes value of x by dereferencing pointer
}
