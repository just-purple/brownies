package main

// importazione di pacchetti
import (
	"fmt"
	"strings"
)

func main() {
	
	// dichiarazione della variabile di tipo string con nome parola
	var parola string

	// presa in input della variabile parola
	fmt.Println("Dammi una parola")
	fmt.Scanln(&parola)

	// alla variabile parola contenente la stringa presa in input viene applicata la funzione ToLower del pacchetto strings
	parola = strings.ToLower(parola)

	// stampa della variabile parola con le modifiche apportate
	fmt.Print("In minuscolo: ", parola, "\n")

}
