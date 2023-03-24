// storage package handle the storage of integers numbers
package storage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Storage struct contains a slice of int
type Storage struct {
	// store, stock
	Store []int
}

// Load function returns a storage given a path to a file
func Load(path_file string) (Storage, error) {

	file, err := os.Open(path_file)

	// create a Storage to hold the file contents
	result := Storage{}

	// in caso di file rotto torna una Storage vuota e un errore
	if err != nil {
		return result, err
	}
	// it's important to close the file after reading it
	defer file.Close()

	// viene creato uno scanner con la funzione bufio.NewScanner() per leggere il file riga per riga
	scanner := bufio.NewScanner(file)
	// per ogni riga, viene utilizzata la funzione strconv.Atoi() per convertire la stringa in un intero
	// Se la conversione fallisce, la funzione restituisce un errore
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result.Store = append(result.Store, number)
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	// se tutto è andato a buon fine la funzione Load restituisce una Storage e nessun errore
	return result, nil
}

// Print method prints the storage
func (s Storage) Print() {

	// loop for-each per iterare attraverso gli elementi della slice s e utilizza fmt.Printf per stampare ciascun elemento, separato da spazio
	for _, i := range s.Store {
		fmt.Printf("%d ", i)
	}
	// stampa una nuova riga vuota alla fine della storage
	fmt.Println()
}

// Dump method, for the Storage struct, accepts a path to a file and stores it into it (if the file is not empty overwrite it)
