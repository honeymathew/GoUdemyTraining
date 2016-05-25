package main

import "fmt"

func main() {

	a := 20

	fmt.Println(a)
	fmt.Println(&a)//memmory address of a

	var b *int = &a//b points to a memory location where an integer value is stored
	fmt.Println(b)//will print the address of a
	fmt.Println(*b)//value


	//b is a pointer to an int ie pointing to a memory address that stores an integer
}
