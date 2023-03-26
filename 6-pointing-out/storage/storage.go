package storage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Storage struct contains a slice of int
type Storage struct {
	// store, stock
	store []*int
}

// Load function returns a storage given a path to a file
func Load(path_file string) (*Storage, error) {

	file, err := os.Open(path_file)
	// in caso di file rotto torna una Storage vuota e un errore
	if err != nil {
		return nil, err
	}
	// it's important to close the file after reading it
	defer file.Close()

	// viene creato uno scanner con la funzione bufio.NewScanner() per leggere il file riga per riga
	scanner := bufio.NewScanner(file)
	result := &Storage{}
	// lancio di scan alla prima riga
	scanner.Scan()

	// la funzione TrimSpace toglie gli eventiali spazi all'inizio e alla fine a una stringa (scanner.Text())
	trimmed := strings.TrimSpace(scanner.Text())
	line := strings.Split(trimmed, " ")

	// conversione da string a int, se fallisce restituisce errore
	for _, v := range line {
		number, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result.store = append(result.store, &number)
	}
	// se scanner.Scan() è true significa che c'è una seconda riga
	if scanner.Scan() {
		return nil, errors.New("storage file must be one line")
	}

	return result, nil
}

// Print method prints the storage
func (s *Storage) Print() {
	// loop for-each per iterare attraverso gli elementi della slice s e utilizza fmt.Printf per stampare ciascun elemento, separato da spazio
	for _, v := range s.store {
		fmt.Printf("%d ", *v)
	}
	// stampa una nuova riga vuota alla fine della storage
	fmt.Println()
}

// Dump method, for the Storage struct, accepts a path to a file and stores it into it (if the file is not empty overwrite it)
func (s *Storage) Dump(path string) error {
	// la funzione create se riceve un path esistente lo vuota, altrimenti crea un nuovo file a quel path
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// Always remember to close the open file descriptor when you finish working with the file
	// so that the system can reuse it
	defer file.Close()

	// immagazzinare la s nel file creato, e sovrascrivere la s in un file non vuoto
	dump := ""
	for _, v := range s.store {
		// conversione, la funzione Sprintf prende i e lo mette dentro %d quindi se vogliamo uno spazio fra un numero e l'altro bastera mettere uno spazio dopo %d
		toAppend := fmt.Sprintf("%d ", *v)
		// append sulla stringa dump
		dump += toAppend
	}
	// mette una riga a capo in fondo al file
	dump += "\n"

	_, err = file.WriteString(dump)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Inc() {

	for _, v := range s.store {
		*v++
	}
}

func (s *Storage) Take(position int, taken *int) error {
	if position < 0 || position > len(s.store) {
		return errors.New("invalid position")
	}
	*taken = *s.store[position]
	return nil
}
