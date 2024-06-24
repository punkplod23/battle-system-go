package character

import (
	"strings"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	player := CreatePlayer()

	if player.Name != "Hero" {
		t.Errorf("Expected player name to be 'Hero', got '%s'", player.Name)
	}

	if player.Charactertype != "HERO" {
		t.Errorf("Expected player type to be 'HERO', got '%s'", player.Charactertype)
	}

	if player.Attributes.HP != 100 || player.Attributes.MP != 100 {
		t.Errorf("Expected player attributes to be 100 hp and 100 mp, got %d hp and %d mp", player.Attributes.HP, player.Attributes.MP)
	}
}

func TestCreateEnemy(t *testing.T) {
	enemy := CreateEnemy()

	if !strings.HasPrefix(enemy.Name, "Enemy") {
		t.Errorf("Expected enemy name to start with 'Enemy', got '%s'", enemy.Name)
	}

	if enemy.Charactertype != "ENEMY" {
		t.Errorf("Expected enemy type to be 'ENEMY', got '%s'", enemy.Charactertype)
	}

	if enemy.Attributes.HP < 0 || enemy.Attributes.HP > 100 {
		t.Errorf("Expected enemy hp to be between 0 and 100, got %d", enemy.Attributes.HP)
	}

	if enemy.Attributes.MP < 0 || enemy.Attributes.MP > 100 {
		t.Errorf("Expected enemy mp to be between 0 and 100, got %d", enemy.Attributes.MP)
	}
}
