package main

import (
	"fmt"
	"strings"
)

func main() {

	var parola string

	fmt.Println("Dammi una parola")
	fmt.Scanln(&parola)

	parola = strings.ToLower(parola)

	fmt.Print("In minuscolo: ", parola, "\n")

}
