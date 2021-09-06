package todo

import (
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

type raw struct {
	Todos []todoRaw `json:"todos"`
}

type todoRaw struct {
	TodoId  int    `json:"todo_id"`
	Created   string `json:"created"`
	IsDone    bool   `json:"is_done"`
	Completed string   `json:"completed"`
	TodoTitle string   `json:"todo_title"`
	TodoTags []string `json:"todo_tags"`
	Category string   `json:"category"`
}

type ToDo struct {
	todoId int
	created time.Time
	isDone bool
	completed NullTime
	todoTitle string
	todoTags []string
}

type categorizedTodo struct {
	idByCategory map[string][]int
	byId map[int]ToDo
}

