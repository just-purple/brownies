package main

import (
	"fmt"
)

// creazione del tipo subject con due tipi all'interno
type Subject struct {
	name  string
	score int
}

func main() {
	// prende in input il numero di materie
	nSubject := 0
	fmt.Println("Number of subjects")
	//_,err:= modo per controllare l'input. se err!=nill l'input non è corretto cioè non è del tipo desiderato
	_, err := fmt.Scanln(&nSubject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// controllo dell'input sui numeri negativi
	if nSubject <= 0 {
		fmt.Println("Error: number <= 0")
		return
	}
	// creazione di una slice delle materie, la parentesi quadra indica l'unidimensionalità, di tipo Subject
	subjects := []Subject{
		/*
			Subject{
				name: "bo",
			},
			{
				name: "mate",
			},
		*/
	}
	//riempimento delle materie e dei voti
	for i := 0; i < nSubject; i++ {
		s := Subject{}

		fmt.Printf("Enter name subject %d \n", i+1)
		fmt.Scanln(&s.name)

		fmt.Printf("Enter score subject %d \n", i+1)
		// _ sta al posto di una variabile, che va necessariamente messa ma che poi non mi serve e non la uso
		_, err := fmt.Scanln(&s.score)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		subjects = append(subjects, s)
	}
	//calcolo della somma dei voti
	sumScore := 0
	max := Subject{
		name:  "",
		score: 0,
	}
	for _, subject := range subjects {
		sumScore += subject.score
		if subject.score > max.score {
			// viene assegnato a max.name e max.score i rispettivi di subject (sono lo stesso tipo)
			max = subject
		}
	}
	avg := sumScore / nSubject
	fmt.Println("Average: ", avg)
	fmt.Println("Max subject: ", max.name)
}
