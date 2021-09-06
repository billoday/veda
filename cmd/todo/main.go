/*
Copyright © 2021 Bill O'Day <billoday@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package todo

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var (
	category string

	rootCmd = &cobra.Command{
	Use:   "",
	Short: "get",
	Long: `Display all ToDos stored in VEDA system`,
	Run: func(cmd *cobra.Command, args []string) {
		getTodos()
	},
}
	newCmd = &cobra.Command{
		Use:   "new [title of ToDo]",
		Short: "new",
		Long: `Create a new ToDo in VEDA system`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newTodo(strings.Join(args, " "), category)
		},
	}
	doneCmd = &cobra.Command{
		Use:   "done [todo_id]",
		Short: "done",
		Long: `Set ToDo to done in VEDA system`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			todoId, err := strconv.Atoi(args[0])
			check(err)
			finishTodo(todoId)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		check(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	newCmd.Flags().StringVarP(&category, "category", "c", "DEFAULT", "Category to Use")
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(doneCmd)
}


func displayToDos(allTodos categorizedTodo, includeDone bool) {
	fd := int(os.Stdout.Fd())
	viewW, _, _ := GetSize(fd)
	halfWidth := viewW / 2
	fmt.Printf("\n")
	for key, value := range allTodos.idByCategory {
		if len(value) > 0 {
			printCategory(key, halfWidth)
		}
		for _, tempTodoId := range value {
			tempTodo := allTodos.byId[tempTodoId]
			if !includeDone {
				if !tempTodo.isDone {
					printTodo(tempTodo, halfWidth)
				}
			} else {
				printTodo(tempTodo, halfWidth)
			}
		}
	}
}

func getTodos() {
	allTodos := loadFile()
	displayToDos(allTodos, false)
}

func newTodo(todoText string, category string) {
	allTodos := loadFile()
	todoId := max(allTodos.byId) + 1
	created := time.Now()
	completed := NullTime{Valid: false}
	tags := make([]string, 0)
	tempTodo := ToDo{
		todoId:    todoId,
		created:   created,
		isDone:    false,
		completed: completed,
		todoTitle: todoText,
		todoTags:  tags,
	}
	var tempList []int
	tempList, catExists := allTodos.idByCategory[category]
	if !catExists {
		tempList = make([]int, 0)
	}
	allTodos.idByCategory[category] = append(tempList, todoId)
	allTodos.byId[todoId] = tempTodo
	// fmt.Printf("New id %d - self reported %d \n", todoId, tempTodo.todoId)
	saveFile(allTodos)
	displayToDos(allTodos, false)
}

func finishTodo(todoId int) {
	allTodos := loadFile()
	todo, idExists := allTodos.byId[todoId]
	if !idExists {
		log.Panicln("Invalid Id")
	} else {
		completed := time.Now()
		todo.isDone = true
		todo.completed = NullTime{completed, true}
		allTodos.byId[todoId] = todo
		saveFile(allTodos)
		displayToDos(allTodos, true)
	}
}

func printCategory(category string, halfWidth int) {
	fmt.Println("")
	fmt.Println(color.Ize(color.Blue, wordWrap(category, halfWidth)))
}

func printTodo(todo ToDo, halfWidth int) {
	var completedText string
	tagsText := ""
	if todo.isDone {
		completedText = " [X] " + todo.completed.Time.Format(time.RFC3339)
	} else {
		completedText = " [ ] "
	}
	if len(todo.todoTags) > 0 {
		tagsText = strings.Join(todo.todoTags, ", ")
	}
	fullTextTodo := "" +
		color.Ize(color.Red, strconv.Itoa(todo.todoId)) +
		color.Ize(color.Green, completedText) +
		color.Ize(color.Gray, tagsText) +
		color.Ize(color.White, todo.todoTitle)
	fmt.Println(wordWrap(fullTextTodo, halfWidth))
}

func GetSize(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}

func wordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}

func max(numbers map[int]ToDo) int {
	var maxNumber int
	for maxNumber = range numbers {
		break
	}
	for n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}
