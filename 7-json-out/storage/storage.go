// re-implement Dump function using the json library

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// Storage struct contains a slice of int
type Storage struct {
	// store, stock
	store []int
}

// Load function returns a storage given a path to a file
// la funzione Load prende il nome del file da leggere come parametro e restituisce la Storage contenente la slice di interi letti dal file e un eventuale errore
func Load(path_jason string) (*Storage, error) {
	// lettura del contenuto del file
	file, err := ioutil.ReadFile(path_jason)
	// in caso di file rotto torna una Storage vuota e un errore
	if err != nil {
		return nil, err
	}
	// utilizzo del pacchetto encoding/json per parsare (analizzare un flusso continuo di dati in ingresso in modo da determinare la correttezza della sua struttura grazie ad una data grammatica formale)
	// il contenuto del file json come un array di interi.
	// In caso di errore durante la lettura del file o durante la decodifica del contenuto, la funzione restituisce un errore
	result := Storage{}
	// la funzione Unmarshal prende (sempre?) per riferimento la struttura dati (slice) su cui scrivere il contenuto del file json
	if err := json.Unmarshal(file, &result.store); err != nil {
		return &Storage{}, err
	}
	// la funzione Load restituisce per riferimento la struct Storage (quindi la sua locazione in memoria)
	return &result, nil
}

// Print method prints the storage
func (s *Storage) Print() {
	// loop for-each per iterare attraverso gli elementi della slice s e utilizza fmt.Printf per stampare ciascun elemento, separato da spazio
	for _, v := range s.store {
		fmt.Printf("%d ", v)
	}
	// stampa una nuova riga vuota alla fine della storage
	fmt.Println()
}

// Dump method, for the Storage struct, accepts a path to a file and stores it into it (if the file is not empty overwrite it)
func (s *Storage) Dump(path string) error {
	// conversione della slice di interi (s.store) in JSON
	file, err := json.Marshal(s.store)
	if err != nil {
		return err
	}
	// scrittura del JSON nel file, sovrascrivendo l'eventuale contenuto precedente se il file non è vuoto
	err = ioutil.WriteFile(path, file, 0644)
	// in caso di errore durante la conversione o la scrittura, la funzione restituisce un errore
	if err != nil {
		return err
	}

	return nil
}

// method Inc accepts no parameter and increments all the value inside of the storage by 1
func (s *Storage) Inc() {
	for i := 0; i < len(s.store); i++ {
		s.store[i]++
	}
}

// method Take takes the element at the specified index:
// it takes an integer named position and an integer named taken by reference
func (s *Storage) Take(position int, taken *int) error {
	if position < 0 || position > len(s.store) {
		return errors.New("invalid position")
	}
	*taken = s.store[position]
	// it returns only an error which is not nil if the specified position does not exists
	return nil
}

// method Add accepts an integers and adds it to the storage
// il tipo Storage viene preso per riferimento
// il numero che prende in ingresso viene passato per valore perchè tanto si tratta di un numero
func (s *Storage) Add(number int) {
	// il numero da aggiungere alla fine della storage viene passato per riferimento (locazione in memoria)
	s.store = append(s.store, number)
}

// method Remove remove the element at the specified index:
// it takes an integer named position
func (s *Storage) Remove(position int) error {
	if position < 0 || position > len(s.store) {
		return errors.New("invalid position")
	}
	// If you want to keep your array ordered:
	// all of the elements at the right of the deleting index you have to shift by one to the left
	// dentro la funzione non posso mettere s.store poichè è un puntatore, quindi creo una variabile di appoggio slice
	// dopo la rimozione metto la slice aggiornata dentro s.store
	slice := s.store
	slice = append(slice[:position], slice[position+1:]...)
	s.store = slice
	// it returns only an error which is not nil if the specified position does not exists
	return nil
}

// method FunkyReduce
func (s *Storage) FunkyReduce() {
	i := 0
	// loop with entry condition so that if the last number is even leave it like it is
	for i+1 < len(s.store) {
		// add to each number in even position the following number (in odd position)
		s.store[i] += s.store[i+1]
		// then remove each number in odd position
		s.Remove(i + 1)
		i++
	}
}
