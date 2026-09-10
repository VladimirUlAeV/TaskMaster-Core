package main

import (
	datastructur "TaskMaster/DataStructur"
	"bufio"	
	"fmt"
	"os"
	"strconv"
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
	mananger := datastructur.Mananger()
	var Counter int 
	fmt.Println("Приветствую вас в Task Master!")
	InputUser := bufio.NewScanner(os.Stdin)
	for{
		menu()
		fmt.Println("Выберите действие:")
		InputUser.Scan()
		InputText := InputUser.Text()
		switch InputText{
		case "1":
			mananger.AllTask()
		case "2":
			var err error
			fmt.Println("Введите тэг задачи(обязательно):")
			InputUser.Scan()
 			Tag := InputUser.Text()
			fmt.Println("Введите описание задачи:")
			InputUser.Scan()
 			Title := InputUser.Text()
			Counter, err = mananger.NewTask(Counter,Tag, Title)
			if err != nil {
			fmt.Println("Задача", "не создана по причине", err)
			}else{
				fmt.Println("Задача успешно добавлена")
			}
		case "3":
			fmt.Print("Введите айди задачи: ")
			InputUser.Scan()
			IdStr := InputUser.Text()
			Id, errA := strconv.Atoi(IdStr)
			if errA != nil{
				fmt.Println("Ошибка: пустое или не правильное значение")
			}
			err := mananger.TaskDone(Id)
			if err != nil{
				fmt.Println(err)
			}
		case "4":
			fmt.Print("Введите айди задачи: ")
			InputUser.Scan()
			IdStr := InputUser.Text()
			Id, errA := strconv.Atoi(IdStr)
			if errA != nil{
				fmt.Println("Ошибка: пустое или не правильное значение")
			}
			err := mananger.DeleteTask(Id)
			if err != nil{
				fmt.Println(err)
			}
		case "5":
		fmt.Println("Введите тэг для поиска: ")
		InputUser.Scan()
		Tag := InputUser.Text()
		err := mananger.SearchTask(Tag)
		if err != nil{
			fmt.Println(err)
		}
		case "0":
			fmt.Println("Спасибо за использование моей программы!")
			return
		default:
			fmt.Println("Команда не распознана, выберите еще раз")
		}
	}

}