package main

import (
	"fmt"

	"github.com/just-purple/brownies/0-cooked/6-pointing-out/storage"
)

const FILE_PATH = "./storage.txt"
const WRONG_FILE_PATH = "./wrong_storage.txt"

// questo non c'è, il file txt non esiste
const NOT_EXISTING_PATH = "./wrong_path.txt"

func main() {

	ints, err := storage.Load(FILE_PATH)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	ints.Print()

	ints.Inc()
	fmt.Println("After increments by 1: ")
	ints.Print()

	ints.Dump(FILE_PATH)
	fmt.Println("Check storage.txt")

	// sintassi dei puntatori:
	// simbolo * permette di accedere al valore contenuto in una locazione di memoria cui il puntatore fa riferimento;
	// simbolo & restituisce il riferimento alla locazione di memoria di una variabile.
	// taken viene passato alla funzione Take per riferimento, in questo modo
	// non viene gestito il valore di taken bensì l'indirizzo della locazione di memoria in cui si trova
	taken := 0
	err = ints.Take(0, &taken)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("Taken index 0: ", taken)

	ints.Add(100)
	fmt.Println("After add 100 at the end: ")
	ints.Print()

	err = ints.Remove(1)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("After remove value at index 1: ")
	ints.Print()

	ints.FunkyReduce()
	fmt.Println("After FunkyReduce: ")
	ints.Print()

	ints.Dump("sumup.txt")
	fmt.Println("Now check sumup.txt")

}
