package gamePart

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/globalF"
)

var NewPerson assistant.Person

var ChoosePerson string

func Persons(goGame string) assistant.Person {
	var User assistant.Person
	fmt.Println(assistant.Line)
	fmt.Println(assistant.Start_yes)
	for {
		if getcomand() {
			break
		} else {
			fmt.Println("Введите корректную команду")
		}
	}
	if ChoosePerson == "create" {
		NP(ChoosePerson)
		User = assistant.ProgrssUser(NewPerson)
		globalF.SavePerson(User)
	} else if ChoosePerson == "person" {
		listPerson := SelectPerson()
		fmt.Println("Выбирите персонажа:")
		var indexPerson int
		for i, per := range listPerson.Persons {
			fmt.Printf("%v) %v \n", i, per)
		}

		for {
			fmt.Fscan(os.Stdin, &indexPerson)
			if indexPerson < len(listPerson.Persons) {
				User = assistant.ProgrssUser(listPerson.Persons[indexPerson])
				break
			} else {
				fmt.Println("Введите корректный номер персонажа")
			}
		}
	} else {
		fmt.Println("Ошибка обработки данны")
	}
	return User
}

func getcomand() bool {
	fmt.Fscan(os.Stdin, &ChoosePerson)
	if ChoosePerson == "create" {
		return true
	} else if ChoosePerson == "person" {
		return true
	} else {
		return false
	}
}

var chooseProf int
var chooseSex int

func NP(comand string) assistant.Person {

	fmt.Println("Введите имя: ")
	fmt.Fscan(os.Stdin, &NewPerson.Name)
	fmt.Println("Введите пол 1 - Женский, 2 - Мужской: ")
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

	fmt.Printf("Выберите желаемую проффесию: %v %v %v ", assistant.Profs[0], assistant.Profs[1], assistant.Profs[2])
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
	return NewPerson
}

func SelectPerson() globalF.BPersons {
	personData, err := os.ReadFile("savedPerons.json")
	if err != nil {
		log.Fatal("У вас нет сохраненных персанажей")

	}

	var ArrSelectPerson globalF.BPersons

	err = json.Unmarshal(personData, &ArrSelectPerson)
	if err != nil {
		log.Fatal("У вас нет сохраненных персанажей")
	}
	return ArrSelectPerson
}
