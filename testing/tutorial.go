package main

import "fmt"

type fullname struct {
	first  string
	second string
}

var name fullname = fullname{"Aidan", "Strong"}
var name2 fullname = fullname{"Jenny", "Chandler"}

var nameArray = []fullname{name, name2}

func main() {

	fmt.Println(len(nameArray))
	fmt.Println(cap(nameArray))
}
