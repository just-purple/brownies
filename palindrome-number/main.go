package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x = ", x)
	fmt.Println(isPalindrome(x))

}

func isPalindrome(x int) bool {
	/*
		input_num := x
		var remainder int
		res := 0
		for x > 0 {
			remainder = x % 10
			res = (res * 10) + remainder
			x = x / 10
		}
		// se il numero fornito è zero salta il for e cade dentro l'if, mentre se il numero è minore di 0 finisce nell'else
		if input_num == res {
			return true
		} else {
			return false
		}
	*/
	if x < 0 {
		return false
	}
	text := fmt.Sprintf("%d", x)
	midlen := len(text) / 2
	j := len(text) - 1
	for i := 0; i < midlen; i++ {
		if text[i] != text[j] {
			return false
		}
		j--
	}
	return true
}
