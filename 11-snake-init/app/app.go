package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/just-purple/brownies/11-snake-init/terminal"
)

const APPLE = "@"
const SNAKE_BODY = "█"

type Pair struct {
	X, Y int
}

type App struct {
	PlayGround    Pair
	Apple         *Pair
	StartingSnake Pair
	Direction     Pair
	Delay         time.Duration
	StartingLen   int

	snake []*Pair
}

func (a *App) handleInput(input keyboard.Key) {
	if input == keyboard.KeyArrowDown {
		a.Direction = Pair{X: 0, Y: 1}
		return
	}
	if input == keyboard.KeyArrowUp {
		a.Direction = Pair{X: 0, Y: -1}
		return
	}
	if input == keyboard.KeyArrowLeft {
		a.Direction = Pair{X: -1, Y: 0}
		return
	}
	if input == keyboard.KeyArrowRight {
		a.Direction = Pair{X: 1, Y: 0}
		return
	}
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
			found := false
			for _, v := range a.snake {
				if v.X == x && v.Y == y {
					fmt.Print(SNAKE_BODY)
					found = true
					break
				}
			}
			if a.Apple.X == x && a.Apple.Y == y {
				fmt.Print(APPLE)
				found = true
			}

			if !found {
				fmt.Print(" ")
			}
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

func (a *App) isDead() bool {
	for i := 1; i < len(a.snake); i++ {
		if a.snake[0].X == a.snake[i].X && a.snake[0].Y == a.snake[i].Y {
			return true
		}
	}
	return false
}

func (a *App) isEating() bool {
	if a.Apple == nil {
		return false
	}
	if a.snake[0].X == a.Apple.X && a.snake[0].Y == a.Apple.Y {
		return true
	}
	return false
}

func (a *App) placeApple() {
	// array (slice di pair) con tutte le posizioni vuote (dove non c'è il serpente)
	// E POI RANdom selezioni un indice dell'array e poi piazzi la mela lì
	// in questo modo generi un numero random fra quelli nell'array e verso la fine quando il serpente è grande

	var tmp []Pair
	for y := 0; y < a.PlayGround.Y; y++ {
		for x := 0; x < a.PlayGround.X; x++ {
			for _, v := range a.snake {
				if v.X != x && v.Y != y {
					tmp = append(tmp, Pair{
						X: x,
						Y: y,
					})
				}
			}
		}
	}
	iRand := rand.Intn(len(tmp))
	a.Apple = &tmp[iRand]
}

func (a *App) evaluate() {
	// TODO:
	// - move the snake
	// - here there is a problem :)
	// alla posizione viene sommata la direzione, per questo nella funzione handleinput per gestire la direzione data in input le x e le y sono 0, 1, -1 con riferimento all'asse cartesiano con origine nel vertice in alto a sinistra del campo

	// tmp := make([]*Pair, len(a.snake))
	// tmp[0] = &Pair{
	// 	X: (a.snake[0].X + a.Direction.X + a.PlayGround.X) % a.PlayGround.X,
	// 	Y: (a.snake[0].Y + a.Direction.Y + a.PlayGround.Y) % a.PlayGround.Y,
	// }

	// for i := 0; i < len(tmp)-1; i++ {
	// 	tmp[i+1] = a.snake[i]
	// }
	// a.snake = tmp

	for i := len(a.snake) - 1; i > 0; i-- {
		a.snake[i] = &Pair{
			X: a.snake[i-1].X,
			Y: a.snake[i-1].Y,
		}
	}

	v := a.snake[0]
	v.X = (v.X + a.Direction.X + a.PlayGround.X) % a.PlayGround.X
	v.Y = (v.Y + a.Direction.Y + a.PlayGround.Y) % a.PlayGround.Y
}

func (a *App) Run() (int, error) {
	a.snake = make([]*Pair, a.StartingLen)
	for i := 0; i < a.StartingLen; i++ {
		a.snake[i] = &Pair{
			X: a.StartingSnake.X - a.Direction.X*i,
			Y: a.StartingSnake.Y - a.Direction.Y*i,
		}

	}
	a.placeApple()
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

	for {
		// select the first available channel
		select {
		case input, ok := <-inputChannel:
			if !ok {
				terminal.ShowCursor()
				return len(a.snake), nil //TODO errors.New("game stopped")
			}
			a.handleInput(input)

		case <-time.After(a.Delay):
			terminal.Clean()
			a.evaluate()
			if a.isDead() {
				terminal.ShowCursor()
				return len(a.snake), nil
			}
			if a.isEating() {
				// il serpente cresce di un pezzettino
				a.snake = append(a.snake, &Pair{})
				// piazzo la mela
				a.placeApple()
			}
			a.draw()
		}
	}
}
