// this file is an entry point of the function because it has main function

package main

import (
	"fmt"
	// this path appeared after command+s e moccoli
	"github.com/just-purple/brownies/3-school-module/function"
)

func main() {
	// package declaration: package name (name of the directory - function) can be different than package declaration
	// package declaration is used to create package reference variable
	fmt.Println("github.com/just-purple/brownies/3-school-module/main.go ==> main()")

	nSubject, err := function.GetNSubjects()
	// controlli sull'input riferiti ai risultati della funzione getNSubjects
	// vengono effettuati nel main così con la return il programma può finire
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// creazione di una slice delle materie di tipo Subject, che verrà riempita grazie alla funzione getSubjects
	subjects, err := function.GetSubjects(nSubject)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// stampa della slice inserita con la funzione getSubjects
	fmt.Println("Slice provided by the user:", subjects)

	fmt.Println("Subject with the highest score: ", function.GetMaxSubject(subjects))

	fmt.Println("Average: ", function.GetMean(subjects))
}
