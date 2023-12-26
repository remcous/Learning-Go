package main

// import "fmt"

type MyFuncOpts struct{
	FirstName string
	LastName string
	Age int
}

func MyFunc(opts MyFuncOpts) error{
	// do something here
	return nil
}

func main(){
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age: 50,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName: "Smith",
	})
}