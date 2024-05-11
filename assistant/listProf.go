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
var Wizard = Prof{Name: "Волшебник", Specifications: Specifications{Hp: 980, Mana: 350, Str: 18, Agil: 7, Inelect: 7}}
var Agilman = Prof{Name: "Ловкач", Specifications: Specifications{Hp: 980, Mana: 350, Str: 18, Agil: 7, Inelect: 7}}
