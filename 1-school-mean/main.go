package main

import (
	"fmt"
)

// creazione del tipo Subject strutturato da due tipi predefiniti al suo interno
type Subject struct {
	name  string
	score int
}

func main() {

	// presa in input del numero di materie
	nSubject := 0
	fmt.Println("Number of subjects")
	//_,err:= è il modo per controllare che l'input sia del tipo desiderato, se err!=nill allora l'input non è corretto
	// la funzione Scanln ritorna due cose: un intero e un errore, per questo prima del := ci sono due elementi: _ , err
	// dato che l'intero non ci serve, ma va necessariamente specificato, al suo posto mettiamo l'underscore
	_, err := fmt.Scanln(&nSubject)
	if err != nil {
		fmt.Println(err.Error())
		// dopo ogni controllo con uscita è buona norma mettere la return
		return
	}
	// controllo dell'input in caso di inserimento di numeri negativi o dello 0
	if nSubject <= 0 {
		fmt.Println("Error: number <= 0")
		return
	}

	// creazione di una slice/array/lista delle materie di tipo Subject, la parentesi quadra indica l'unidimensionalità
	subjects := []Subject{}

	//riempimento delle materie e dei voti
	for i := 0; i < nSubject; i++ {
		s := Subject{}

		// la funzione Printf si formatta in base a un indicatore di formato, e a differenza del Println non viene aggiunta la riga a capo in automatico (per questo viene specificato \n)
		fmt.Printf("Enter name subject %d \n", i+1)
		fmt.Scanln(&s.name)

		fmt.Printf("Enter score subject %d \n", i+1)
		// _ sta al posto di una variabile che va necessariamente messa (poichè la funzione Scanln ritorna un intero e un error),
		// ma che poi non serve e non viene utilizzata, quindi al suo posto viene messo l'underscore
		_, err := fmt.Scanln(&s.score)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// alla lista subjects di tipo Subject viene accodata la variabile s (anch'essa di tipo Subject) che ho precedentemente riempito con i valori presi in input
		subjects = append(subjects, s)
	}

	// calcolo della somma dei voti
	sumScore := 0
	// dato che devo memorizzare sia il nome che il voto della materia massimo, la variabile max la posso creare di tipo struttura Subject
	max := Subject{
		name:  "",
		score: 0,
	}
	// in questo for non viene specificato l'indice di partenza (infatti c'è _), viene presa in considerazione la slice subjects
	for _, subject := range subjects {
		sumScore += subject.score
		if subject.score > max.score {
			// a max.name e max.score vengono assegnati i rispettivi subject.name e subject.score,
			// è possibile fare questa operazione più sintetica perchè max e subject sono dello stesso tipo
			max = subject
		}
	}

	avg := sumScore / nSubject
	fmt.Println("Average: ", avg)
	fmt.Println("Max subject: ", max.name)

}
