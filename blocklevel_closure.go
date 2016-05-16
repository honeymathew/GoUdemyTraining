package main

import "fmt"

func main() {
	x :=40
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "nice"
		fmt.Println(y)
	}
	//fmt.Println(y) //y not acssesible here

}
