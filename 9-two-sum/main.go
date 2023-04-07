package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target

func main() {
	nums := []int{24, 2, -7, 6, 10}
	target := 8
	fmt.Println("array: ", nums, "target: ", target)
	fmt.Println("result: ", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	/*
		// Soluzione con complessina O(n^2)

			for i := 0; i < len(nums); i++ {
				for j := len(nums) - 1; j > i; j-- {
					if nums[i]+nums[j] == target {
						return []int{i, j}
					}
				}
			}
	*/

	// Soluzione con complessità O(n)

	// il ragionamento:
	// partendo dall'equazione  target = x + y , allora  target - x = y  e  target - y = x

	// l'algoritmo:
	// all'interno del ciclo che scorre gli elementi del vettore nums, ad ogni iterazione
	// controllo se il risultato della differenza ( target - num[index] = numeroRichiesto ) compare come chiave nella mappa
	// se è presente allora
	// la variabile indiceRichiesto memorizzerà il valore contenuto nella mappa alla chiave [target - num]
	// e i due indici che compongono la somma target sono l'indiceRichiesto e l'indice corrente (index)
	// altrimenti se non è presente inserisco il valore sottratto corrente ( num[index] ) come una nuova chiave nella mappa che avrà come valore associato l'indice di quel num nell'array, cioè l'indice corrente nel ciclo (index)

	// creazione della mappa che memorizzerà gli elementi dell'array nums già visitati
	// ciascuna chiave della mappa corrisponde ad un elemento num dell'array,
	// e il relativo valore della mappa contiene l'indice dell'elemento num nell'array
	visited := map[int]int{}
	// ciclo che scorre ogni elemento num dell'array nums
	for i, num := range nums {
		// mentre visiti ogni elemento dell'array lo aggiungi alla mappa
		// (complemento richiesto = somma target - numero corrente)
		// ad ogni iterazione controlla se il complemento è già presente nella mappa visited
		complement := target - num
		_, isPresent := visited[complement]
		if isPresent {
			// se presente, allora restituisce come risultato i due indici: il valore della mappa (indice dell'array nums) associato alla chiave complemento, e l'indice corrente del ciclo cioè l'indice dell'elemento num esaminato
			return []int{visited[complement], i}
		}
		// altrimenti aggiunge alla mappa l'elemento num corrente dell'array (come chiave) e il suo indice i (come valore)
		visited[num] = i
	}
	return nil
}
