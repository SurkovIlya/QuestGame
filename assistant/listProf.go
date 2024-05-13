package assistant

type Person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Prof Prof   `json:"prof"`
}

type Prof struct {
	Name           string `json:"nameProf"`
	Specifications struct {
		Hp      int `json:"hp"`
		Mana    int `json:"mana"`
		Str     int `json:"str"`
		Agil    int `json:"agil"`
		Inelect int `json:"inelect"`
	} `json:"specifications"`
}

var Strongman = Prof{Name: "Силач", Specifications: struct {
	Hp      int "json:\"hp\""
	Mana    int "json:\"mana\""
	Str     int "json:\"str\""
	Agil    int "json:\"agil\""
	Inelect int "json:\"inelect\""
}{Hp: 980, Mana: 350, Str: 18, Agil: 7, Inelect: 7}}

var Wizard = Prof{Name: "Волшебник", Specifications: struct {
	Hp      int "json:\"hp\""
	Mana    int "json:\"mana\""
	Str     int "json:\"str\""
	Agil    int "json:\"agil\""
	Inelect int "json:\"inelect\""
}{Hp: 800, Mana: 700, Str: 10, Agil: 9, Inelect: 20}}

var Agilman = Prof{Name: "Ловкач", Specifications: struct {
	Hp      int "json:\"hp\""
	Mana    int "json:\"mana\""
	Str     int "json:\"str\""
	Agil    int "json:\"agil\""
	Inelect int "json:\"inelect\""
}{
	Hp: 600, Mana: 500, Str: 8, Agil: 22, Inelect: 12}}
