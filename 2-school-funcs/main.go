package main

import (
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
	return nSubject, err
}

// funzione che dato un numero n restituisce una slice di n subjects (fornite dall'utente)
func getSubjects(n int) []Subject {

	subjects := []Subject{}

	//riempimento delle materie e dei voti
	for i := 0; i < n; i++ {
		s := Subject{}

		fmt.Printf("Enter name subject %d \n", i+1)
		fmt.Scanln(&s.name)

		// per l'inserimento dei voti vorrei eseguire un ciclo con condizione di uscita
		// se il valore fornito dall'utente è un intero >= 0 allora finisce il ciclo e si va avanti con la vita
		for {
			fmt.Printf("Enter score subject %d \n", i+1)
			_, err := fmt.Scanln(&s.score)
			// controllo dell'input ricorsivo fino a che l'utente non inserisce un valore adeguato
			if err != nil || s.score < 0 {
				fmt.Println("Error, try again")
			}
			else {
				break
			}
		}

		// alla lista subjects di tipo Subject viene aggiunta s (anch'essa di tipo Subject) che ho precedentemente riempito con i valori presi in input
		subjects = append(subjects, s)
	}

	return subjects
}

// funzione che data una slice di Subject restituisce la subject con lo score più alto
func getMaxSubject(s []Subject) string {

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
		fmt.Println("Error:", err)
		return
	}
	if nSubject <= 0 {
		fmt.Println("Error: number <= 0")
		return
	}

	// creazione di una slice delle materie di tipo Subject, che verrà riempita grazie alla funzione getSubjects
	subjects := getSubjects(nSubject)

	// stampa della slice inserita con la funzione getSubjects
	fmt.Println("Slice provided by the user:", subjects)

	fmt.Println("Max subject: ", getMaxSubject(subjects))

	fmt.Println("Average: ", getMean(subjects))

}
