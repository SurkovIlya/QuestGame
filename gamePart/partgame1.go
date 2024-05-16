package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/globalF"
)

var varAnswe int

func TalkwhithMam(ownerUser assistant.Person) {
	fmt.Printf("МАМА: %v, Ты уже встал? \n", ownerUser.Name)

	fmt.Printf("Выбирите варианты ответа:\n 1) Да;\n 2) Не твое собачье дело! (Проверить свою ловкость. Нужное значение: ---> 21 <---) ;\n 3) Промолчать \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			fmt.Println("Вы: Да, я встал мам!")
			break
		} else if varAnswe == 2 {
			checkA := globalF.CheckAgil(ownerUser.Prof.Specifications.Agil)
			fmt.Printf("Уровень вашей ловкости: %v \n", checkA)
			if checkA > 21 {
				fmt.Printf("Вы: Не твое собачье дело!\n")
			} else {
				fmt.Println(assistant.ErrorCheck)
				fmt.Printf("Вы: Не твое собачье дело!\n")
				fmt.Println("... Резко открывается дверь и в комнату влетает ваш отец и дает тебе ПОДЗАТЫЛЬНИК")
				ownerUser.Prof.Specifications.Hp = ownerUser.Prof.Specifications.Hp - 15
				fmt.Println("!!!!Вы получили 15 урона!!!!")
			}
			break
		} else if varAnswe == 3 {
			fmt.Println("Вы промолчали....")
			break
		} else {
			fmt.Println("Укажите корректный номер варианта ответа")
		}
	}

}

func WayToColledge(ownerUser assistant.Person) {
	fmt.Printf("%v \n %v \n", assistant.TxrProlog, assistant.TxtStreet)
	fmt.Printf(assistant.Btalk1)
	fmt.Printf("Попытаться прервать бездумный поток мыслей Бабки. (Требуемый уровень Интелекта ---> 15 <---)")
	fmt.Printf("1) Да \n 2) Нет \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			checkI := globalF.CheckInt(ownerUser.Prof.Specifications.Inelect)
			fmt.Printf("Уровень вашего интелекта: %v \n", checkI)
			if checkI > 15 {
				fmt.Println("Вы: Прошу прощение... Я очень тороплюсь. \n ....Вы успешно распрощались с бабкой и направились дальше.")
				// успешно ... можно вывести текст интерфейса и начать новый уровень(сохранить персонажа и прогрес игры)
			} else {
				fmt.Println(assistant.ErrorCheck)
				fmt.Printf("Вы: Я быть торопливый человек... \n")
				fmt.Printf(assistant.Btalk2)
				fmt.Printf("...Вы берете пакеты и идете с бабушкой до ее дома...\n %v ... Неожиданно вкрикнуа бабушка", assistant.Btalk3)
				fmt.Println("... Вы видите место назначение, но вдруг из-за угла выскакивают 2 человека...")
				assistant.RespG()
				fmt.Println("... Ва провалили проверку на ИНТЕЛЕКТ. И решаете сразу драться ...")
				if ownerUser.Atak > assistant.Gopnik.Atak {
					fmt.Printf("... Вы нонесли кретический урон!... Гопники убегают \n")
					fmt.Printf("... ВЫ УВЕЛИЧИЛИ СВОИ СТАТЫ \n Сила: %v + %v \n Ловкость: %v + %v \n Здоровье: %v +%v \n", ownerUser.Prof.Specifications.Str, 3, ownerUser.Prof.Specifications.Agil, 1, ownerUser.Prof.Specifications.Hp, 50)
					ownerUser.Prof.Specifications.Str = ownerUser.Prof.Specifications.Str + 3
					ownerUser.Prof.Specifications.Agil = ownerUser.Prof.Specifications.Agil + 1
					ownerUser.Prof.Specifications.Hp = ownerUser.Prof.Specifications.Hp + 50
					fmt.Println("... Вы помогли бабушке и защитили от гопников. Вы получаете игровой предмет СТАРИННЫЙ КЛЮЧ")
					ownerUser.Inventar = struct {
						Name    string "json:\"name\""
						UpdHp   int    "json:\"updhp\""
						UpdMana int    "json:\"updmana\""
						UpdAtak int    "json:\"updatakf\""
					}{"oldKey", 0, 0, 0}
				}
				// персонаж получил ключ. далее идет короткая версия сюжета и финал
			}
			break
		} else if varAnswe == 2 {
			fmt.Printf(assistant.Btalk2)
			// тут будет разбор с бабкой несу сумки и т.д. За помощь должен качнуться доп сюжет и ответвление
			break
		} else {
			fmt.Println("Укажите корректный номер варианта ответа")
		}
	}

}
