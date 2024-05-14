package assistant

type Person struct {
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
	UserP.Atak = UserP.Prof.Specifications.Str + UserP.Inventar.UpdAtak + 15
	return UserP
}
