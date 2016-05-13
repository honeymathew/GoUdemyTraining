package main

import "fmt"

var a int = 20 //package level scope ie,scope of variable 'a' is whole package
func main() {

	fmt.Println(a)//accessible because scope of 'a' is package level scope
	foo()
	y := 10 //block level scope. varible 'y' is accessible only with in this block
	fmt.Println(y)

	//fmt.Println(z)//not accesible...
	//z := 30 //...bcz scp of 'z' starts here
}

func foo() {

	fmt.Println(a)//accessible because scope of 'a' is package level scope
	//fmt.Print(y)
	//y is not accessible here because its scope is block level scope
}
