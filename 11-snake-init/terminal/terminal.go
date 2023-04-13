package terminal

import "fmt"

func HideCursor() {
	fmt.Print("\033[?25l") // Hide the cursor
}

func ShowCursor() {
	fmt.Print("\033[?25h") // show the cursor
}

func Clean() {
	fmt.Print("\033c")
}
