package globalF

import (
	"encoding/json"
	"fmt"
	"os"

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

	err = json.Unmarshal(personData, &ArrPerson)
	if err != nil {
		ArrPerson.Persons = append(ArrPerson.Persons, NewPerson)
	}

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
