// this file is an entry point of the function because it has main function

package main

import (
	"fmt"
	// in go.mod file you can find the module
	// this path appeared after command+s e moccoli
	"github.com/just-purple/brownies/0-cooked/3-school-module/school"
)

func main() {

	// package declaration: package name (name of the directory - school) can be different than package declaration
	// package declaration is used to create package reference variable
	// fmt.Println("github.com/just-purple/brownies/3-school-module/main.go ==> main()")

	nSubject, err := school.GetNSubjects()
	// controlli sull'input riferiti ai risultati della funzione getNSubjects
	// vengono effettuati nel main così con la return il programma può finire
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

	// stampa della slice inserita con la funzione getSubjects
	fmt.Println("Slice provided by the user:", subjects)

	fmt.Println("Highest score subject: ", school.GetMaxSubject(subjects))

	fmt.Println("Average: ", school.GetMean(subjects))
}
