package datastructur

import ()

type taskBar struct{
	Name string
	Tag string
	Status bool  // False - не выполнина 
	title string
}

type taskManenger struct{
	Task []taskBar
	ID int
}