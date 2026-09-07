package datastructur

import (
	"errors"
)

func Mananger()*taskManenger{
	return &taskManenger{
		Task: make([]taskBar, 0),
		ID: 0,
	}
}

func (t *taskManenger)NewTask(name string, tag string, title string)(*taskManenger, error){
	if name == "" || tag == ""{
		return &taskManenger{}, errors.New("не ввели обязательные данные")
	}
	t.ID += 1
	t.Task = append(t.Task,taskBar{name, tag, false, title})

	return &taskManenger{}, nil
}