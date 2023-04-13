package app

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/just-purple/brownies/11-snake-init/terminal"
)

const APPLE = "🍎"
const SNAKE_BODY = "█"

type Pair struct {
	X, Y int
}

type App struct {
	PlayGround Pair
	Apples     []Pair
	Snake      Pair
	Direction  Pair
	Delay      time.Duration
}

func (a *App) handleInput(input keyboard.Key) {
	// TODO:
	//	- right now this function reverts the direction
	//	- edit the Direction based on the user input
	// 	- ex: if input == keyboard.KeyArrowDown { ... }

	a.Direction = Pair{X: a.Direction.X * -1, Y: a.Direction.Y * -1}
}

func (a *App) draw() {
	fmt.Println("Press esc to exit")

	// top border
	fmt.Print("┎")
	for x := 0; x < a.PlayGround.X; x++ {
		fmt.Print("-")
	}
	fmt.Println("┐")

	for y := 0; y < a.PlayGround.Y; y++ {
		// left border
		fmt.Print("|")

		// inside the frame
		for x := 0; x < a.PlayGround.X; x++ {
			if x == a.Snake.X && y == a.Snake.Y {
				fmt.Print(SNAKE_BODY)
				continue
			}

			fmt.Print(" ")
		}

		// right border
		fmt.Println("|")
	}

	// bottom border

	fmt.Print("┖")
	for x := 0; x < a.PlayGround.X; x++ {
		fmt.Print("-")
	}
	fmt.Println("┘")
}

func (a *App) evaluate() {
	// TODO:
	// - move the snake
	// - here there is a problem :)
	a.Snake.X = (a.Snake.X + a.Direction.X) % a.PlayGround.X
	a.Snake.Y = (a.Snake.Y + a.Direction.Y) % a.PlayGround.Y

	// TODO-THEN:
	// - this is the place where you check things
	// - is the snake dead?
	// - is the snake eating?
	// - where do I place the apple

}

func (a *App) Run() {
	terminal.HideCursor()
	inputChannel := make(chan keyboard.Key)

	// run this function over the input channel
	go func(ch chan keyboard.Key) {
		for {
			_, key, err := keyboard.GetSingleKey()
			if err != nil {
				close(ch)
				return
			}
			if key == keyboard.KeyEsc {
				close(ch)
				return
			}
			ch <- key
		}
	}(inputChannel)

stdinloop:
	for {
		// select the first available channel
		select {
		case input, ok := <-inputChannel:
			if !ok {
				break stdinloop
			}
			a.handleInput(input)

		case <-time.After(a.Delay):
			terminal.Clean()
			a.evaluate()
			a.draw()
		}
	}
	terminal.ShowCursor()
}
