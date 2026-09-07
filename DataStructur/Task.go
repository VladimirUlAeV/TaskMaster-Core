package datastructur

type taskBar struct {
	Name   string
	Tag    string
	Status bool // False - не выполнина
	Title  string
}

type taskManenger struct {
	Task []taskBar
	ID   int
}
