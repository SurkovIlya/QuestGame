package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
)

var varAnswe int

func TalkwhithMam(ownerUser assistant.Person) {
	fmt.Printf("%v, Ты уже встал? \n", ownerUser.Name)
	fmt.Printf("Выбирите варианты ответа: 1) Да; 2) (Проверить свою ловкость)Не твое собачье дело!; 3) Промолчать \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			fmt.Println("Вы: Да, я встал мам!")
			break
		} else if varAnswe == 2 {
			fmt.Println("Вы: Не твое собачье дело!")
			//добавить проврке на пиздюлины
			break
		} else if varAnswe == 3 {
			fmt.Println("Вы промолчали....")
			break
		} else {
			fmt.Println("Укажите корректный номер варианта ответа")
		}
	}

}
