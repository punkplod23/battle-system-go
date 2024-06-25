package gameengine

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/punkplod23/battle-system-go/internal/app/character"
)

// GameEngine struct
type GameEngine struct {
	entities    []character.Character
	enemiecount int
}

func NewGameEngine() *GameEngine {
	// Setup Game
	g := new(GameEngine)
	min := 2
	max := 6
	// We minus the player from the total
	var maxRows int = (rand.Intn(max-min) + min)
	entities := make([]character.Character, maxRows)
	player := character.CreatePlayer()
	go player.Timer.StartCountdown()
	entities[0] = player
	g.enemiecount = maxRows - 1
	for pos, _ := range entities {
		if pos != 0 {
			entities[pos] = character.CreateEnemy()
			go entities[pos].Timer.StartCountdown()
		}

	}
	g.entities = entities
	return g
}

func (g *GameEngine) checkGameEndCondition() bool {

	if g.enemiecount <= 0 {
		fmt.Println("You Vanuqished the enemies")
		return false
	}

	if g.entities[0].Attributes.HP <= 0 {
		fmt.Println("You is Dead get over it")
		return false
	}

	return true
}

func (g *GameEngine) initEvent() {
	for pos, entity := range g.entities {
		if entity.Timer.CheckComplete() && entity.Attributes.HP > 0 {
			if pos == 0 {
				g.action(pos)
			} else {
				go g.action(pos)
				g.entities[pos].Timer.ResetTimer()
				go g.entities[pos].Timer.StartCountdown()
			}

		}
	}
}

func (g *GameEngine) printPlayer(position int) {

	message := g.entities[position].Name + " ready to attack"

	if position != 0 {
		message += " " + g.entities[0].Name
		fmt.Println(message)
		g.enemyAttack(position)
	}

	if position == 0 {
		fmt.Println(message)
		fmt.Printf("Please Select Enemy to Attack")
		g.heroSelectAttack()
	}

}

func (g *GameEngine) printCharacter(position int, show_position bool) {
	if show_position {
		fmt.Println("\r", "Posiiton:", position)
	}
	if show_position {
		fmt.Println("\r", "Name:", g.entities[position].Name)
	}

}

func (g *GameEngine) printEnemies() {
	for pos, enemy := range g.entities {
		if pos > 0 {
			if enemy.Attributes.HP > 0 {
				g.printCharacter(pos, true)
			}
		}
	}
}

func (g *GameEngine) enemyAttack(position int) {
	g.printCharacter(0, false)
	min := 2
	max := 30

	var damage int = (rand.Intn(max-min) + min)
	fmt.Println("\n", g.entities[position].Name, " attacks ", " inflicts ", damage, " points")
	g.entities[0].Attributes.HP = g.entities[0].Attributes.HP - damage

}

func (g *GameEngine) heroSelectAttack() {

	fmt.Println("Please Select Enemy position to attack", "\n")
	g.printEnemies()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", "Select Enemy by position")

		response, err := reader.ReadString('\n')
		if err != nil {
			g.heroSelectAttack()
		}

		responseint, err := strconv.Atoi(strings.ToLower(strings.TrimSpace(response)))
		if err != nil {
			g.heroSelectAttack()
		}

		g.attackEnemy(responseint)
		break
	}

}

func (g *GameEngine) attackEnemy(position int) {

	min := 2
	max := 50
	// We minus the player from the total
	var damage int = (rand.Intn(max-min) + min)
	fmt.Println("\n", "Hero Attacks", g.entities[position].Name, " inflicts ", damage, " points")
	g.entities[position].Attributes.HP = g.entities[position].Attributes.HP - damage
	if g.entities[position].Attributes.HP <= 0 {
		fmt.Println("\n"+g.entities[position].Name, " is dead")
		g.enemiecount = g.enemiecount - 1
	}
	g.entities[0].Timer.ResetTimer()
	go g.entities[0].Timer.StartCountdown()

}

func (g *GameEngine) action(position int) {
	g.printPlayer(position)
}

func (g *GameEngine) StartGame() {

	fmt.Println("Prepare to Fight", "\n")
	g.printCharacter(0, false)
	g.printEnemies()

	for g.checkGameEndCondition() {
		g.initEvent()

	}
}
