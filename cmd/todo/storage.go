package todo

import (
	"encoding/json"
	"github.com/adrg/xdg"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const datafile = "veda/ToDo.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileName() string {
	var dataFilePath string
	dataFilePath, err := xdg.DataFile(datafile)
	check(err)
	// log.Printf("Using %s\n", dataFilePath)
	return dataFilePath
}

func loadFile() categorizedTodo {
	dataFilePath := getFileName()
	var rawData raw
	fileData, err := os.Open(dataFilePath)
	if err != nil {
		log.Printf(err.Error())
		rawData = raw{
			Todos: []todoRaw{},
		}
	} else {
		defer func(fileData *os.File) {
			err := fileData.Close()
			check(err)
		}(fileData)
		//log.Printf("File Loaded")
		byteValue, err := ioutil.ReadAll(fileData)
		err = json.Unmarshal(byteValue, &rawData)
		check(err)
	}
	blankCats := make(map[string][]int)
	blankIds := make(map[int]ToDo)
	allTodos := categorizedTodo{
		idByCategory: blankCats,
		byId:         blankIds,
	}
	for i := 0; i < len(rawData.Todos); i++ {
		category := rawData.Todos[i].Category
		// check if category exists in allTodos
		_, catExists := allTodos.idByCategory[category]
		tempTodo := rawTodoToTodo(rawData.Todos[i])
		if catExists {
			allTodos.idByCategory[category] = append(allTodos.idByCategory[category], tempTodo.todoId)
		} else {
			tempList := make([]int, 0)
			allTodos.idByCategory[category] = append(tempList, tempTodo.todoId)
		}
		allTodos.byId[tempTodo.todoId] = tempTodo
	}
	return allTodos
}

func saveFile(allTodos categorizedTodo) {
	rawTodoList := categorizedTodosToRaw(allTodos)
	dataFilePath := getFileName()
	file, _ := os.OpenFile(dataFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)
	jsonString, err := json.Marshal(rawTodoList)
	check(err)
	ioutil.WriteFile(dataFilePath, jsonString, os.ModePerm)
	//encoder := json.NewEncoder(file)
	//err := encoder.Encode(rawTodoList)
	check(err)
}

func rawTodoToTodo(rawTodo todoRaw) ToDo {
	createdTime, err := time.Parse(time.RFC3339, rawTodo.Created)
	check(err)
	var completedTime NullTime
	if rawTodo.IsDone {
		completedTime.Time, err = time.Parse(time.RFC3339, rawTodo.Completed)
		check(err)
		completedTime.Valid = true
	} else {
		completedTime.Valid = false

	}
	return ToDo{
		todoId: rawTodo.TodoId,
		created: createdTime,
		isDone: rawTodo.IsDone,
		completed: completedTime,
		todoTitle: rawTodo.TodoTitle,
		todoTags: rawTodo.TodoTags,
	}
}

func categorizedTodosToRaw(allTodos categorizedTodo) raw {
	rawTodoList := make([]todoRaw, 0)
	for key, value := range allTodos.idByCategory {
		for _, tempTodoId := range value {
			tempTodo := allTodos.byId[tempTodoId]
			var rawCompleted string
			if tempTodo.completed.Valid {
				rawCompleted = tempTodo.completed.Time.Format(time.RFC3339)
			} else {
				rawCompleted = ""
			}
			rawTodoList = append(
				rawTodoList,
				todoRaw{
					TodoId:    tempTodo.todoId,
					Created:   tempTodo.created.Format(time.RFC3339),
					IsDone:    tempTodo.isDone,
					Completed: rawCompleted,
					TodoTitle: tempTodo.todoTitle,
					TodoTags:  tempTodo.todoTags,
					Category:  key,
				},
			)
		}
	}
	return raw{
		Todos: rawTodoList,
	}
}
