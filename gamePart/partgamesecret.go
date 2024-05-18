package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/globalF"
)

func SecretWay(ownerUser assistant.Person) {
	ownerUser.Lvlgame.Lvl = 99

	fmt.Printf(" ... Оказавшись наедине с ключом по среди пустынной улицы вы начали видеть обрывки будущего ... \n ничто больше не может быть скрыто от вас ... \n Вы становитесь сверх существом и перемещаетесь на необитаемый остров размышлять о вечном ... \n \n")
	fmt.Printf("... КОНЕЦ ИГРЫ ... \n \n 1) Завершить \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			globalF.DelPerson(ownerUser)
			break
		} else {
			fmt.Println("Так нужно!")
		}

	}

}
