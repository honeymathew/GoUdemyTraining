package main

import "fmt"

func main() {
	//fmt.Println(x) //cant access bcz in block level scope order is imprtnt.
	fmt.Println(y)
	//x := 10
}
 var y =20    //in package level scope, order is not matter.'x' is declared here but accessed
             // in main.bcz scope of variable 'x' is package level