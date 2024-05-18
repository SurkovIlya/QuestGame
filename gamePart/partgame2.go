package gamePart

import (
	"fmt"
	"os"

	"github.com/SurkovIlya/MiniQuest-on-GO/assistant"
	"github.com/SurkovIlya/MiniQuest-on-GO/globalF"
)

func ColledgeDay(ownerUser assistant.Person) {
	var varAnswe int
	ownerUser.Lvlgame.Lvl = 2
	fmt.Printf("... Вы находитесь в главном корпусе Коледжа ... \n ... Сегодня достаточно людно, подумали вы. Взглянув на расписание пошли в аудиторию впитывать знания. \n")
	fmt.Printf("... На пути в аудиторию 301 вы встречаете своего друга ... \n")
	fmt.Printf(assistant.Ftalk1, ownerUser.Name)
	fmt.Printf("Выберите вариант ответа:\n 1) да, погнали в столовку\n 2) нет, спасибо я пойду на пару\n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			fmt.Printf("... Вы уходите с другом в столовую... Во время приема пищи завязывается диалог ...\n \n")
			TalkInCafe(ownerUser)
			break
		} else if varAnswe == 2 {
			fmt.Printf("... Закончились пары и впереди Тест вы ощущаете сильный голод и начинаете жалеть, что не пошли на обед с другом ... \n Поискать перекусить? \n 1) да \n 2) нет \n")
			for {
				fmt.Fscan(os.Stdin, &varAnswe)
				if varAnswe == 1 {
					fmt.Printf("... Вы ищете что покушать ...\n")
					if ownerUser.Inventar.Name == "пирожок с картошкой" {
						fmt.Printf("... Вы находите пирожок. Перекусив вы успешно подготавливаетесь в тесту ...\n ... Вы успешно сдали тест, плюс преподователь вспоминает что утром вы ей помогли и дает вам право не ходить больше на ее занятия... \n ... Вы давольный возвращаетесь домой... \n \n")
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
					} else {
						fmt.Printf("... ЕДЫ НЕ ОБНАРУЖЕНО. Вы плохо подготовились к тесту, но заметив что преподаватель который его принимает это бабушка которую вы встретили утром решили претвориться заболевшим и уходите домой. \n ... Родители не узнают о тесте и вы спокойно его сдаете на следующий день ...")
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
					break
				} else if varAnswe == 2 {
					fmt.Printf("... Вы решили не искать еду и начать готовиться к тесту... \n ... В процессе подготовки вы стали настолько голодны что вподаете в кому ...\n")
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
					break
				} else {
					fmt.Printf("Укажите корректный номер варианта ответа")
				}

			}
			break
		} else {
			fmt.Println("Укажите корректный номер варианта ответа")
		}
	}

}

func TalkInCafe(ownerUser assistant.Person) {
	fmt.Printf(assistant.Ftalk2)
	fmt.Printf("... Отобедав вы направились на занятия... И вот наступает Тест...\n")
	fmt.Printf("Зайдя в аудиторию, вы понимаете, что преподаватель, о котором говорил ваш друг — это та бабушка утром ...\n... Вы пересекитесь с преподавателем взглядом...\n \n")
	fmt.Printf("... Сюжетный поворот... Выполнить проверку? \n 1) Да \n 2) Нет \n")
	for {
		fmt.Fscan(os.Stdin, &varAnswe)
		if varAnswe == 1 {
			break
		} else {
			fmt.Printf("... Тут у вас нет выбора!!!! ....")
			break
		}

	}

	if ownerUser.Lvlgame.Way == 12 {
		fmt.Printf(" ... Утром вы ей помогли и она это запомнила ...\n ... Вы получаете зачет автоматом и отправляетесь домой ... \n \n")
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

	} else if ownerUser.Lvlgame.Way == 11 {

		fmt.Printf("... Выс жестко валят на тесте и отправляют на пересдачу... \n")
		fmt.Printf("... Придя домой родители узнат о том что вы заволили тест ... \n ... Вы получаете подзатыльник от отца ... \n КРИТИЧЕСКИЙ УРОН \n ... Вы отключаетесь ... \n \n ... КОНЕЦ ИГРЫ ... \n \n 1) Завершить \n")
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
}
