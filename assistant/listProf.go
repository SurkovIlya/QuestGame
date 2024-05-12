package assistant

type Prof struct {
	Name string
	Specifications
}
type Specifications struct {
	Hp      int
	Mana    int
	Str     int
	Agil    int
	Inelect int
}

var Strongman = Prof{Name: "Силач", Specifications: Specifications{Hp: 980, Mana: 350, Str: 18, Agil: 7, Inelect: 7}}
var Wizard = Prof{Name: "Волшебник", Specifications: Specifications{Hp: 800, Mana: 700, Str: 10, Agil: 9, Inelect: 20}}
var Agilman = Prof{Name: "Ловкач", Specifications: Specifications{Hp: 600, Mana: 500, Str: 8, Agil: 22, Inelect: 12}}
