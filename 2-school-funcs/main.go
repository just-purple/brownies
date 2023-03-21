package main

import (
	"errors"
	"fmt"
)

type Subject struct {
	name  string
	score int
}

// funzione che restituisce un intero (fornito dall'utente, che indica il numero di materie) e un errore
func getNSubjects() (int, error) {
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

// funzione che dato un numero n restituisce una slice di n subjects (fornite dall'utente), e un errore
func getSubjects(n int) ([]Subject, error) {

	subjects := []Subject{} // equivale a var subjects []Subject ?

	//riempimento delle materie e dei voti
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
func getMaxSubject(s []Subject) string {

	// creazione della variabile che conterrà il nome e il voto della materia massima
	max := Subject{
		name:  "",
		score: 0,
	}

	// ricerca del massimo score

	// il ciclo for può essere utilizzato con un range di slice per iterare attraverso gli elementi di una slice.
	// il codice eseguito in questo blocco verrà ripetuto per ogni elemento nella slice s, ma ad esempio con l'operatore [:3] si restituiscono solo i primi tre valori della slice
	// se non si ha bisogno dell'indice è possibile ometterlo utilizzando _ come nome della variabile
	// se invece l'indice viene indicato, range s restituisce una coppia di valori per ogni elemento nella slice: il primo è l'indice dell'elemento e il secondo è il valore dell'elemento stesso
	// ad esempio, la prima iterazione del ciclo for avrà i valori index = 0 e value = math, e così via
	for _, subject := range s {
		if subject.score > max.score {
			max = subject
		}
	}

	return max.name
}

// funzione che data una slice di Subject restituisce la media
func getMean(s []Subject) int {

	// somma dei voti
	sumScore := 0
	for _, subject := range s {
		sumScore += subject.score
	}

	// calcolo della media
	avg := sumScore / len(s)
	return avg
}

func main() {
	nSubject, err := getNSubjects()
	// controlli sull'input riferiti ai risultati della funzione getNSubjects
	// vengono effettuati nel main così con la return il programma può finire
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// creazione di una slice delle materie di tipo Subject, che verrà riempita grazie alla funzione getSubjects
	subjects, err := getSubjects(nSubject)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// stampa della slice inserita con la funzione getSubjects
	fmt.Println("Slice provided by the user:", subjects)

	fmt.Println("Max subject: ", getMaxSubject(subjects))

	fmt.Println("Average: ", getMean(subjects))

}
