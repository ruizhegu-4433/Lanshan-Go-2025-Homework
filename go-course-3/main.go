package main

import (
	"fmt"
	"go-course-3/character"
	"go-course-3/commend"
	"go-course-3/monster"
	"math/rand"
)

func main() {
	var monsters []commend.A_B
	characters := &character.Attribute{
		Hp:         100,
		Attraction: 10,
		Health:     5,
		Skill:      15,
	}
	var choose int
	fmt.Println("choose monster(1 is slime,2 is dagon)")
	fmt.Scanf("%d", &choose)
	fmt.Scanln()
	switch choose {
	case 1:
		slime := &monster.Slime{
			Att: monster.Attributes{
				Hp:         30,
				Attraction: 5,
			},
			Skill: 10,
		}
		monsters = append(monsters, slime)
	case 2:
		dagon := &monster.Dagon{
			Att: monster.Attributes{
				Hp:         200,
				Attraction: 15,
			},
			Skill: 25,
		}
		monsters = append(monsters, dagon)
	}
	count := 2
	for {
		var player commend.Attracter
		var fighter commend.Beattracter
		var Restore character.C_actioner
		var chooses int

		player = characters
		fighter = monsters[0]
		fmt.Println("")
		fmt.Println("========= yurr turn ========")
		fmt.Println("")
		fmt.Println("1.restore Hp,2.Attract,3.skill Attract(only use twice) ")
		fmt.Scanf("%d", &chooses)
		fmt.Scanln()
		switch chooses {
		case 1:
			Restore = characters
			Restore.Restore(Restore.Gethealth())
		case 2:
			player.Attract(fighter)
		case 3:
			if count > 0 {
				player.Skillattract(fighter)
			} else {
				fmt.Println("no count")
				continue
			}
			count--

		}
		if fighter.GetHp() <= 0 {
			fmt.Println("you win!")
			break
		}
		fmt.Println("")
		fmt.Println("======== monster's turn =============")
		fmt.Println("")
		player = monsters[0]
		fighter = characters

		num := rand.Intn(100)
		if num >= 1 && num <= 50 {
			player.Skillattract(fighter)
		} else {
			player.Attract(fighter)
		}

		if fighter.GetHp() <= 0 {
			fmt.Println("you not win!")
			break
		}
	}
	fmt.Println("press any key to Esc")
	var a int
	fmt.Scanf("%d", a)
}
