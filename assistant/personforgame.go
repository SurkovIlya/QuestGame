package assistant

type Person struct {
	Id      int64 `json:"id"`
	Lvlgame struct {
		Lvl int `json:"lvl"`
		Way int `json:"way"`
	} `json:"lvlgame"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Atak     int    `json:"ataka"`
	Prof     Prof   `json:"prof"`
	Inventar struct {
		Name    string `json:"name"`
		UpdHp   int    `json:"updhp"`
		UpdMana int    `json:"updmana"`
		UpdAtak int    `json:"updatakf"`
	} `json:"inventar"`
}

var UserP Person

func ProgrssUser(char Person) Person {
	UserP = char

	return UserP
}
