// in the main, remember to test with each path
package main

import (
	"fmt"

	"github.com/just-purple/brownies/0-cooked/5-storing-out/storage"
)

const FILE_PATH = "./storage.txt"
const WRONG_FILE_PATH = "./wrong_storage.txt"

// questo non c'è, il file txt non esiste
const NOT_EXISTING_PATH = "./wrong_path.txt"

func main() {
	ints, err := storage.Load("viola.txt")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	ints.Print()

	ints.Dump("viola.txt")

}
