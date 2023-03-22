package school

import (
	"errors"
	"fmt"
)

type Subject struct {
	name  string
	score float64
}

// funzione pubblica che restituisce un intero (fornito dall'utente, che indica il numero di materie) e un errore
func GetNSubjects() (int, error) {
	// creazione della variabile di appoggio, gli metto lo stesso nome che ha nel main
	nSubject := 0
	fmt.Println("Number of subjects: ")
	_, err := fmt.Scanln(&nSubject)

	if err != nil {
		return 0, err
	}

	if nSubject <= 0 {
		return 0, errors.New("number <= 0")
	}

	return nSubject, nil
}

// funzione pubblica che dato un numero n restituisce una slice di n subjects (fornite dall'utente)
func GetSubjects(n int) ([]Subject, error) {

	subjects := []Subject{}

	//inserimento delle materie e dei voti
	for i := 0; i < n; i++ {
		s := Subject{}

		fmt.Printf("Enter name subject %d \n", i+1)
		fmt.Scanln(&s.name)

		fmt.Printf("Enter score subject %d \n", i+1)
		_, err := fmt.Scanln(&s.score)

		// se il valore fornito dall'utente è un intero < 0 allora fesco dalla funzione
		if err != nil || s.score < 0 {
			return nil, errors.New("wrong format")
		}

		// alla lista subjects di tipo Subject viene aggiunta s (anch'essa di tipo Subject) che ho precedentemente riempito con i valori presi in input
		subjects = append(subjects, s)
	}

	return subjects, nil
}
