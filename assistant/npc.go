package assistant

import "fmt"

type NpcStat struct {
	Name string `json:"name"`
	Atak int    `json:"ataka"`
	Hp   int    `json:"hp"`
}

var Gopnik NpcStat

func RespG() {
	Gopnik = NpcStat{Name: "Гопник", Atak: 1, Hp: 150}
	fmt.Printf(Gtalk)
}

// Babka := NpcStat{Name: Бабушка, Atak: 1, Hp: 150}

// gamePart.WayToColledge(Gopnik)
