// declaration on the top to state that this is a utility package
// a utility package is supposed to provide some exportmembers to a package who imports it
package school

import (
	"errors"
	"fmt"
)

// export members can be of any type like constant, map, function, struct, array, slice etc.
// Go exports an exportmember if its name starts with capital letter
// all other exportmembers not starting with an uppercase letter is private to the package (can only be used within the package)

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

// funzione che dato un numero n restituisce una slice di n subjects (fornite dall'utente)
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

// funzione che data una slice di Subject restituisce la subject con lo score più alto
func GetMaxSubject(s []Subject) string {

	// creazione della variabile che conterrà il nome e il voto della materia massima
	max := Subject{
		name:  "",
		score: 0,
	}

	// ricerca del massimo score
	for _, subject := range s {
		if subject.score > max.score {
			max = subject
		}
	}

	return max.name
}

// funzione che data una slice di Subject restituisce la media
func GetMean(s []Subject) float64 {

	// somma dei voti
	sumScore := 0.0
	for _, subject := range s {
		sumScore += subject.score
	}

	// conversion of the return value of the len function from int to float64, using the float64 function
	len := float64(len(s))

	// calcolo della media
	avg := 0.0
	avg = sumScore / len
	return avg
}
