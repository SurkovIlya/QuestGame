package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
)

type Person struct {
	Name string
	Sex  string
	Prof assistant.Prof
}

var NewPerson Person

var ChoosePerson string

func СreatePerson(goGame string) {
	fmt.Println(assistant.Line)
	fmt.Println(assistant.Start_yes)
	for {
		if getcomand() {
			break
		} else {
			fmt.Println("Введите корректную команду")
		}
	}
	NP(ChoosePerson)
	fmt.Println(NewPerson)
}

func getcomand() bool {
	fmt.Fscan(os.Stdin, &ChoosePerson)
	if ChoosePerson == "create" {
		return true
	} else if ChoosePerson == "my peron" {
		return true
	} else {
		return false
	}
}

var chooseProf int
var chooseSex int

func NP(comand string) Person {
	if comand == "create" {
		fmt.Print("Введите имя: ")
		fmt.Fscan(os.Stdin, &NewPerson.Name)
		fmt.Print("Введите пол 1 - Женский, 2 - Мужской: ")
		for {
			fmt.Fscan(os.Stdin, &chooseSex)
			if chooseSex == 1 {
				NewPerson.Sex = "женский"
				break
			} else if chooseSex == 2 {
				NewPerson.Sex = "мужской"
				break
			} else {
				fmt.Println("Введите корректную команду")
			}
		}

		fmt.Printf("Введите число в зависимости от желаемой профессии: %v %v %v ", assistant.Profs[0], assistant.Profs[1], assistant.Profs[2])
		for {
			fmt.Fscan(os.Stdin, &chooseProf)
			if chooseProf == 1 {
				NewPerson.Prof = assistant.Strongman
				break
			} else if chooseProf == 2 {
				NewPerson.Prof = assistant.Agilman
				break
			} else if chooseProf == 3 {
				NewPerson.Prof = assistant.Wizard
				break
			} else {
				fmt.Println("Введите корректную команду")
			}
		}

	} else if comand == "my person" {
		fmt.Println("ПОКА ТАК")
	}
	return NewPerson
}
