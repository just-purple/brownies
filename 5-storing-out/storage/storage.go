// storage package handle the storage of integers numbers
package storage

import (
	// "bufio"
	"os"
	// "strconv"
)

// Storage struct contains a slice of int
type Storage struct {
	storage []int
}

// Load function returns a storage given a path to a file
func Load(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// creazione della slice di interi da restituire
	var numbers []int

	// se tutto è andato a buon fine la funzione Load restituisce una slice di interi e nessun errore
	return numbers, nil
}

// Dump method, for the Storage struct, accepts a path to a file and stores it into it (if the file is not empty overwrite it)
