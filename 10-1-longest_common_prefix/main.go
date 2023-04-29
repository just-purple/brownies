package main

import (
	"fmt"
)

func main() {
	strs := []string{"cane", "casa", "camion"}
	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {

	prefix := ""
	var app byte
	str := ""
	for j := 0; j < 6; j++ {
		for i := 0; i < len(strs); i++ {
			str = strs[i]
			if i == 0 {
				app = str[j]
			} else {
				if str[j] != app {
					return prefix
				}
				prefix += string(app)
			}
		}
	}
	return prefix
}
