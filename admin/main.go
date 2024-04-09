package main

import "fmt"

func main() {

	var name *string = new(string)

	*name = "Hello"

	fmt.Println(name, *name)
}
