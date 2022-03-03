package main

import "fmt"

func dataFunction(name, pw string) bool {

	if name == "vvalen" && pw == "soto" {
		return true
	} else {
		return false
	}
}

func main() {
	if dataFunction("vvalen", "soto") == true {
		fmt.Println("joined!")
	} else {
		fmt.Println("error data")
	}
}
