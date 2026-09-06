package main

import (
	datastructur "TaskMaster/DataStructur"
	"bufio"	
	"fmt"
	"os"
)

func menu(){
	fmt.Println("1 — Показать все задачи")
	fmt.Println("2 — Добавить задачу")
	fmt.Println("3 — Отметить задачу выполненной")
	fmt.Println("4 — Удалить задачу")
	fmt.Println("5 — Поиск по тегу")
	fmt.Println("0 — Выход")
} 

func main(){
	Mananger := datastructur.Mananger()
	var err error
	fmt.Println("Приветствую вас в Task Master!")
	InputUser := bufio.NewScanner(os.Stdin)
	for{
		menu()
		fmt.Println("Выберите действие:")
		InputUser.Scan()
		InputText := InputUser.Text()
		switch InputText{
		case "1":

		case "2":
			fmt.Println("Введите название задачи(обязательно):")
			InputUser.Scan()
 			Name := InputUser.Text()
			fmt.Println("Введите тэг задачи(обязательно):")
			InputUser.Scan()
 			Tag := InputUser.Text()
			fmt.Println("Введите описание задачи:")
			InputUser.Scan()
 			Title := InputUser.Text()
			Mananger, err = datastructur.NewTask(Name, Tag, Title)
			if err != nil {
			fmt.Println("Задача не создана по причине ", err)
			} 

		case "3":

		case "4":

		case "5":

		case "0":
			fmt.Println("Спасибо за использование моей программы!")
			return
		default:
			fmt.Println("Команда не распознана, выберите еще раз")
		}
	}

}