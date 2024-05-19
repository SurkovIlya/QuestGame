package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/globalF"
)

var varAnswe int

func TalkwhithMam(ownerUser assistant.Person) assistant.Person {

	fmt.Printf("МАМА: %v, Ты уже встал? \n", ownerUser.Name)

	fmt.Printf("Выберите варианты ответа:\n 1) Да;\n 2) Не твое собачье дело! (Проверить свою ловкость. Нужное значение: ---> 21 <---) ;\n 3) Промолчать \n")
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
	return ownerUser
}

func WayToColledge(ownerUser assistant.Person) assistant.Person {

	fmt.Printf("%v \n %v \n", assistant.TxrProlog, assistant.TxtStreet)
	fmt.Printf(assistant.Btalk1)
	fmt.Printf("Попытаться прервать бездумный поток мыслей Бабки. (Требуемый уровень Интеллекта ---> 15 <---) \n")
	fmt.Printf("1) Да \n2) Нет \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			checkI := globalF.CheckInt(ownerUser.Prof.Specifications.Inelect)
			fmt.Printf("Уровень вашего интеллекта: %v \n", checkI)
			if checkI > 15 {
				fmt.Println("Вы: Прошу прощение... Я очень тороплюсь. \n ....Вы успешно распрощались с бабкой и направились дальше.")
				fmt.Println("... Вы успешно добираетесь до учебы вовремя!")
				ownerUser.Lvlgame.Lvl = 1
				ownerUser.Lvlgame.Way = 11
				globalF.SavePerson(ownerUser)
				fmt.Printf("Прогресс успешно сохранен! \n \n")
				fmt.Printf("\n \n КОНЕЦ ПЕРВОЙ ГЛАВЫ \n \n")
				fmt.Printf("Введие любую цифру для продолжения\n")
				fmt.Fscan(os.Stdin, &varAnswe)

			} else {
				fmt.Println(assistant.ErrorCheck)
				fmt.Printf("Вы: Я быть торопливый человек... \n")
				fmt.Printf(assistant.Btalk2)
				fmt.Printf("...Вы берете пакеты и идете с бабушкой до ее дома...\n %v ... Неожиданно вкрикнуа бабушка", assistant.Btalk3)
				fmt.Println("... Вы видите место назначение, но вдруг из-за угла выскакивают 2 человека...")
				assistant.RespG()
				fmt.Println("... Ва провалили проверку на ИНТЕЛЛЕКТ. И решаете сразу драться ...")
				fmt.Println(ownerUser.Atak)
				fmt.Println(assistant.Gopnik.Atak)
				if ownerUser.Atak > assistant.Gopnik.Atak {
					fmt.Printf("... Вы нанесли критический  урон!... Гопники убегают \n")
					fmt.Printf("... ВЫ УВЕЛИЧИЛИ СВОИ ХАРАКТЕРИСТИКИ \n Сила: %v + %v \n Ловкость: %v + %v \n Здоровье: %v +%v \n", ownerUser.Prof.Specifications.Str, 3, ownerUser.Prof.Specifications.Agil, 1, ownerUser.Prof.Specifications.Hp, 50)
					ownerUser.Prof.Specifications.Str = ownerUser.Prof.Specifications.Str + 3
					ownerUser.Prof.Specifications.Agil = ownerUser.Prof.Specifications.Agil + 1
					ownerUser.Prof.Specifications.Hp = ownerUser.Prof.Specifications.Hp + 50
					fmt.Printf("... Вы помогли бабушке и защитили от гопников. Вы получаете игровой предмет СТАРИННЫЙ КЛЮЧ \n \n")
					ownerUser.Inventar = struct {
						Name    string "json:\"name\""
						UpdHp   int    "json:\"updhp\""
						UpdMana int    "json:\"updmana\""
						UpdAtak int    "json:\"updatakf\""
					}{"oldKey", 0, 0, 0}
					ownerUser.Lvlgame.Lvl = 1
					ownerUser.Lvlgame.Way = 10
					globalF.SavePerson(ownerUser)
					fmt.Printf("Прогресс успешно сохранен! \n \n")
					fmt.Printf("... Вы взгленули на ключ и сразу поняли что нужно делать ...\n \n КОНЕЦ ПЕРВОЙ ГЛАВЫ \n \n")
					fmt.Printf("Введие любую цифру для продолжения\n")
					fmt.Fscan(os.Stdin, &varAnswe)

				}

			}
			break
		} else if varAnswe == 2 {
			fmt.Printf(assistant.Btalk2)
			fmt.Printf("...Вы берете пакеты и идете с бабушкой до ее дома...\n %v ... Неожиданно вкрикнуа бабушка \n", assistant.Btalk3)
			fmt.Printf("... Вы видите место назначение, но вдруг из-за угла выскакивают 2 человека...Увидев вас почувствовали, что лучше пройти мимо...\n")
			fmt.Printf(assistant.Btalk4)
			fmt.Printf("... Вы помогли бабушке донести пакеты... Вы получили предмет ПИРОЖОК С КАРТОШОЙ \n \n")
			ownerUser.Inventar = struct {
				Name    string "json:\"name\""
				UpdHp   int    "json:\"updhp\""
				UpdMana int    "json:\"updmana\""
				UpdAtak int    "json:\"updatakf\""
			}{"пирожок с картошкой", 150, 50, 0}
			fmt.Println("... Вы успешно добираетесь до учебы вовремя!")
			ownerUser.Lvlgame.Lvl = 1
			ownerUser.Lvlgame.Way = 12
			globalF.SavePerson(ownerUser)
			fmt.Printf("Прогресс успешно сохранен! \n \n")
			fmt.Printf("\n \n КОНЕЦ ПЕРВОЙ ГЛАВЫ \n \n")
			fmt.Printf("Введие любую цифру для продолжения\n")
			fmt.Fscan(os.Stdin, &varAnswe)
			break
		} else {
			fmt.Println("Укажите корректный номер варианта ответа")
		}
	}
	return ownerUser
}
