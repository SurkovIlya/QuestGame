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
	fmt.Fscan(os.Stdin, &GoGame)
	if GoGame == "start" {
		gamePart.СreatePerson(GoGame)

	} else {
		fmt.Println(assistant.Start_no)
		fmt.Fscan(os.Stdin, &GoGame)

	}

	// var name string
	// var age int
	// fmt.Print("Введите имя: ")
	// fmt.Fscan(os.Stdin, &name)

	// fmt.Print("Введите возраст: ")
	// fmt.Fscan(os.Stdin, &age)

	// fmt.Println(name, age)
}
