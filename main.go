package main

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/gamePart"
)

func main() {

	var GoGame string
	var ownerUser assistant.Person
	fmt.Println(assistant.Hi)
	fmt.Println(assistant.Start)
	for {
		fmt.Fscan(os.Stdin, &GoGame)
		if GoGame == "start" {
			ownerUser = gamePart.Persons(GoGame)
			break
		} else {
			fmt.Println("Введите корректную команду")
		}
	}

	if ownerUser.Lvlgame.Lvl == 0 {
		fmt.Println(assistant.TxtPart1)
		ownerUser = gamePart.TalkwhithMam(ownerUser)
		ownerUser = gamePart.WayToColledge(ownerUser)
		if ownerUser.Lvlgame.Way == 10 {
			gamePart.SecretWay(ownerUser)
		} else if ownerUser.Lvlgame.Way == 11 || ownerUser.Lvlgame.Way == 12 {
			gamePart.ColledgeDay(ownerUser)
		}
	} else if ownerUser.Lvlgame.Lvl == 1 {
		if ownerUser.Lvlgame.Way == 10 {
			gamePart.SecretWay(ownerUser)
		} else if ownerUser.Lvlgame.Way == 11 || ownerUser.Lvlgame.Way == 12 {
			gamePart.ColledgeDay(ownerUser)
		}
	}

}
