package character

import (
	"math/rand"

	"github.com/punkplod23/battle-system-go/internal/app/timer"
)

// Character struct
type Character struct {
	Name          string
	Charactertype string
	Attributes    Attributes
	Timer         *timer.Timer
}

// Attributes for Character
type Attributes struct {
	HP int
	MP int
}

// letterRunes List for creating random names for Enemies
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes returns a random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CreatePlayer returns Character of charactertype HERO
func CreatePlayer() Character {
	timer := timer.NewTimer(7)
	attributes := Attributes{HP: 100, MP: 100}
	character := Character{Name: "Hero", Charactertype: "HERO", Attributes: attributes, Timer: timer}
	return character
}

// CreateEnemy returns Character of charactertype ENEMY
func CreateEnemy() Character {
	min := 7
	max := 10
	timer := timer.NewTimer(rand.Intn(max-min) + min)
	attributes := Attributes{HP: rand.Intn(100), MP: rand.Intn(100)}
	character := Character{Name: "Enemy" + RandStringRunes(10), Charactertype: "ENEMY", Attributes: attributes, Timer: timer}
	return character
}
