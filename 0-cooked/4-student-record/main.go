package main

import (
	"fmt"

	"github.com/just-purple/brownies/0-cooked/4-student-record/school"
)

func main() {

	nSubject, err := school.GetNSubjects()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// creazione di una slice delle materie di tipo Subject, che verrà riempita grazie alla funzione getSubjects
	subjects, err := school.GetSubjects(nSubject)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// passo alla funzione NewStudentRecord una slice di nome subjects
	// s sarà di tipo StudentRecord
	s := school.NewStudentRecord(subjects...)

	max := s.MaxScoreSubject()
	// per poter leggere nel main il campo name della struttura Subject, va messo con la lettera maiuscola, in questo modo può essere esportato
	fmt.Println("Highest score subject: ", max.Name)

	// metodo ovvero funzione che prende come riferimento la struct StudentRecord
	// la caratteristica del metodo è che gli passi il riferimento con: nomeriferimento (s) . nomemetodo (Mean)
	mean := s.Mean()
	fmt.Println("Mean of the student: ", mean)
}
