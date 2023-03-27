package main

import (
	"fmt"

	"github.com/just-purple/brownies/0-cooked/7-json-out/storage"
)

const FILE_PATH = "./storage.txt"
const JASON_PATH = "./storage.json"

func main() {
	// Many programming languages have their own way of storing data internally, that other languages don’t understand.
	// To allow these languages to interact (communicate between one program and another), the data needs to be converted to a common format they can all understand.
	// One of data interchange formats is JSON (JavaScript Object Notation), a popular way to transmit data over the internet as well as between programs in the same system.
	ints, err := storage.Load(JASON_PATH)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("Storage after Load with file json: ")
	ints.Print()

	ints.Inc()
	fmt.Println("After increments by 1: ")
	ints.Print()

	ints.Dump(JASON_PATH)
	fmt.Println("Check storage.json")

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
