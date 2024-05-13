package globalF

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal("Cannot load savedPerons.json:", err)
	}

	var ArrPerson BPersons

	err = json.Unmarshal(personData, &ArrPerson)
	if err != nil {
		log.Fatal("Cannot load savedPerons.json:", err)
	}

	ArrPerson.Persons = append(ArrPerson.Persons, NewPerson)

	SavePersons, err := json.Marshal(ArrPerson)
	if err != nil {
		fmt.Println(err)

	}

	err = os.WriteFile("./savedPerons.json", SavePersons, 0)
	if err != nil {
		fmt.Println(err)
	}

}
