package datastructur

type taskBar struct {
	Id int
	Tag    string
	Status bool // False - не выполнина
	Title  string
}

type taskManenger struct {
	Task []taskBar
	HowManyTask   int
}
