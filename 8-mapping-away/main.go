//  inside the main.go create a map contacts := map[string]string which is a map of name→phone-number pairs

package main

import (
	"errors"
	"fmt"
)

func validFlag(flag int) error {
	if flag < -1 || flag > 3 {
		return errors.New("invalid input")
	}
	return nil
}
func validPhone(phone string) error {
	if len(phone) != 10 {
		return errors.New("invalid phone-number")
	}
	return nil
}
func presentName(name string, contacts map[string]string) error {
	_, isPresent := contacts[name]
	if !isPresent {
		return errors.New("name not present")
	}
	return nil
}

var prompt = `Type
	0 to exit
	1 to add a name and a phone-number to the contacts
	2 to remove a name from the contacts
	3 to retrieve a phone-umber by specifying a name
`

func main() {
	contacts := map[string]string{}
	contacts["ale"] = "3493932811"
	contacts["leonora"] = "3401409831"
	contacts["viola"] = "3486368251"

	var flag int
	fmt.Println(prompt)
	// endless for in which the user is given 4 options
	for {
		// inserimento della scelta dell'utente
		fmt.Println("\n Option: ")
		fmt.Scanln(&flag)
		// controllo dell'input
		err := validFlag(flag)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		// type 0 to exit
		if flag == 0 {
			return
		}
		// type 1 to add a name and a phone-number to the contacts
		if flag == 1 {
			name := ""
			phone := ""
			fmt.Println("Name: ")
			fmt.Scanln(&name)
			// nel caso di inserimento di un nuovo elemento nella mappa
			// se provo a inserire una chiave gia esistente, il valore associato a quella chiave viene aggiornato all'ultimo inserimento
			_, isPresent := contacts[name]
			if isPresent {
				fmt.Println("Update ", contacts[name], "phone-number: ")
			} else {
				fmt.Println("Phone-number: ")
			}
			fmt.Scanln(&phone)
			// controllo che il numero inserito rispetti la lunghezza di un numero di telefono
			err := validPhone(phone)
			if err != nil {
				fmt.Println("Error:", err.Error())
				return
			}
			contacts[name] = phone
		}
		// type 2 to remove a name from the contacts
		if flag == 2 {
			name := ""
			fmt.Println("Name: ")
			fmt.Scanln(&name)
			// controllo che il nome inserito sia nella mappa
			err := presentName(name, contacts)
			if err != nil {
				fmt.Println("Error:", err.Error())
				return
			}
			delete(contacts, name)
		}
		// type 3 to retrieve a phone-number by specifying a name
		if flag == 3 {
			name := ""
			fmt.Println("Name: ")
			fmt.Scanln(&name)
			// controllo che il nome sia presente
			err := presentName(name, contacts)
			if err != nil {
				fmt.Println("Error:", err.Error())
				return
			}
			fmt.Println("Phone-number:\n", contacts[name])
		}
		// stampa della mappa
		if flag == -1 {
			fmt.Println(contacts)
		}
	}
}
