package main

import (
	"fmt"
)

func validar(str string) bool {
	var pila string
	for i := range str {
		if str[i:i+1] == "(" {
			pila += "("
		}
		if str[i:i+1] == ")" {
			if len(pila) > 0 {
				pila = pila[:len(pila)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(pila)>0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(validar("()"))
	fmt.Println(validar(")(()))"))
	fmt.Println(validar("("))
	fmt.Println(validar("(())((()())())"))
}
