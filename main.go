package main

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/gamePart"
)

func main() {
	var GoGame string
	fmt.Println(assistant.Hi)
	fmt.Println(assistant.Start)
	for {
		fmt.Fscan(os.Stdin, &GoGame)
		if GoGame == "start" {
			gamePart.Persons(GoGame)
			break
		} else {
			fmt.Println("Введите корректную команду")
		}
	}
}
