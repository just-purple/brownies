package main

import (
	"fmt"

	"github.com/just-purple/brownies/6-pointing-out/storage"
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
	ints.Print()

	ints.Dump(FILE_PATH)

	taken := 0

	err = ints.Take(25, &taken)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("taken: ", taken)

}
