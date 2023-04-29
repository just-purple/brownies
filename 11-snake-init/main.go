package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/just-purple/brownies/11-snake-init/app"
)

const JSON_PATH = "./score.json"

type UserScore struct {
	Name  string
	Score int
}

func main() {

	// TODO: press any key to start

	// MAYBE: this could return a score :)
	app := &app.App{
		PlayGround:    app.Pair{X: 20, Y: 6},
		StartingSnake: app.Pair{X: 5, Y: 5},
		Direction:     app.Pair{X: 1, Y: 0},
		Delay:         150 * time.Millisecond,
		StartingLen:   5,
	}

	score, err := app.Run()
	if err != nil {
		return
	}
	topScores, err := getTop(JSON_PATH)
	if err != nil {
		fmt.Println("errore:", err)
	}

	if len(topScores) < 10 || score > topScores[9].Score {
		var name string
		fmt.Println("what's your name?")
		fmt.Scanln(&name)

		new := UserScore{
			Name:  name,
			Score: score - app.StartingLen,
		}

		topScores = orderedAppend(topScores, new)
		if len(topScores) > 10 {
			topScores = topScores[:10]
		}
	}
	fmt.Println(topScores)
	err = setTop(topScores, JSON_PATH)
	if err != nil {
		fmt.Println("errore:", err)
	}

	// MAYBE: make the user add a name and store the top 10 scores somewhere :)

	// TODO: press any key to restart
}

// funzione che restituisce la lista dei top scores contenuti nel file json
func getTop(path string) ([]UserScore, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []UserScore{}, nil
	}

	result := []UserScore{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// funzione che inserisce fra i top scores un nuovo score,
// dopodichè restituisce la lista con i top scores aggiornata
func orderedAppend(topScores []UserScore, new UserScore) []UserScore {
	if len(topScores) == 0 {
		return []UserScore{new}
	}
	// aggiungi nel posto giusto e se la lunghezza supera 10 ritorna solo i primi 10
	i := 0
	for i = 0; i < len(topScores); i++ {
		if new.Score > topScores[i].Score {
			break
		}
	}
	app := append(topScores[:i], new)
	if i < len(topScores) {
		app = append(app, topScores[i:]...)
	}
	return app
}

// funzione che data la lista con i top scores la scrive nel file json
func setTop(topScores []UserScore, path string) error {
	content, err := json.Marshal(topScores)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, content, 0644)
	return err
}
