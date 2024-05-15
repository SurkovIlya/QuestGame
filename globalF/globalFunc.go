package globalF

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
)

type BPersons struct {
	Persons []assistant.Person
}

// var SavePerson []Person

func SavePerson(NewPerson assistant.Person) {

	personData, err := os.ReadFile("savedPerons.json")
	if err != nil {

		os.Create("savedPerons.json")
	}

	var ArrPerson BPersons
	if NewPerson.Id == 0 {
		NewPerson.Id = time.Now().Unix()
	}

	err = json.Unmarshal(personData, &ArrPerson)
	if err != nil {

		ArrPerson.Persons = append(ArrPerson.Persons, NewPerson)
		SavePersons, err := json.Marshal(ArrPerson)
		if err != nil {
			fmt.Println(err)

		}

		err = os.WriteFile("./savedPerons.json", SavePersons, 0666)
		if err != nil {
			fmt.Println(err)
		}
	} else {

		for i, value := range ArrPerson.Persons {
			if value.Id == NewPerson.Id {
				ArrPerson.Persons = append(ArrPerson.Persons[:i], ArrPerson.Persons[i+1:]...)
				ArrPerson.Persons = append(ArrPerson.Persons, NewPerson)
				SavePersons, err := json.Marshal(ArrPerson)
				if err != nil {
					fmt.Println(err)

				}

				err = os.WriteFile("./savedPerons.json", SavePersons, 0666)
				if err != nil {
					fmt.Println(err)
				} else {
					ArrPerson.Persons = append(ArrPerson.Persons, NewPerson)

					SavePersons, err := json.Marshal(ArrPerson)
					if err != nil {
						fmt.Println(err)

					}

					err = os.WriteFile("./savedPerons.json", SavePersons, 0666)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}

	}
}
