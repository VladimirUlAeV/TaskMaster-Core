package datastructur

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func Mananger()*taskManenger{
	return &taskManenger{
		Task: make([]taskBar, 0),
		HowManyTask: 0,
	}
}

func (t *taskManenger)NewTask(Counter int,tag string, title string) (int, error){
	if tag == ""{
		return 0 ,errors.New("не ввели обязательные данные")
	}
	Counter += 1
	t.HowManyTask += 1
	t.Task = append(t.Task,taskBar{Counter, tag, false, title})
	return  Counter ,nil
}

func (t *taskManenger)AllTask(){
	if t.HowManyTask == 0{
		fmt.Println("На данный момент нет задач")
	}else{
		spew.Dump(t.Task)
		fmt.Println("Всего задач:", t.HowManyTask)
	}
}

func (t *taskManenger)TaskDone(ID int)error{
	for index, v := range t.Task {
    	if v.Id == ID{
			t.Task[index].Status = true
			fmt.Println("Статус задачи",ID,"успешно обновлен")
			return nil
		}else{
			return errors.New("Нету такой задачи")
		}
	}
	return errors.New("Вы ещё не создали задачу")
}