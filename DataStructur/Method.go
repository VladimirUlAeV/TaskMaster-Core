package datastructur

import "errors"

func Mananger()*taskManenger{
	return &taskManenger{
		Task: make([]taskBar, 0),
		ID: 0,
	}
}

func (t *taskManenger)NewTask(name string, tag string, title string)(*taskManenger, error){
	if Name == "" && tag == ""{
		return &taskManenger{}, errors.New("не ввели обязательные данные")
	}
	t.ID++
	t.Task = append([]taskBar{
		Name: name,

	})
	return &taskManenger{},nil
}