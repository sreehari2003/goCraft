package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Todo struct {
	id    string
	time  time.Time
	title string
	desc  string
}

// contains all the todo data
var data = []Todo{}

func addTodo(d *[]Todo) {
	uid, err := exec.Command("uuidgen").Output()
	if err != nil {
		fmt.Printf("error creating todo")
	}
	fmt.Println("Enter the title of your todo")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	var i = Todo{}
	i.title = input

	fmt.Println("Enter the description of your todo")
	input, _ = reader.ReadString('\n')
	i.desc = input
	i.id = string(uid)

	*d = append(*d, i)
}
func printTodo(data []Todo) {
	if len(data) == 0 {
		fmt.Printf("Your todo list is empty, please add some todo\n")
	} else {
		for i := 0; i < len(data); i++ {
			fmt.Println(data[i].id)
			fmt.Println(data[i].title)
			fmt.Println(data[i].desc)
			fmt.Println(data[i].time)

		}
	}
}

func Choosetodo() {
	reader := bufio.NewScanner(os.Stdin)
	for k, m := range todoOption {
		fmt.Println(k, ":", m)
	}
	fmt.Println("Choose the operation you want to perform on todo")
	reader.Scan()
	input := reader.Text()

	if input == "PRINT" {
		printTodo(data)
	} else if input == "INSERT" {
		addTodo(&data)
	} else {
		_, ok := todoOption[input]
		if !ok {
			fmt.Println("Inavlid input")
		}
	}
}

var todoOption = map[string]string{
	"PRINT":  "Print all todo",
	"INSERT": "Add a todo",
	"QUIT":   "To Quit the app",
}
