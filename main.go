package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	id    string
	time  time.Time
	title string
	desc  string
}

func main() {
	data := []Todo{}
	chooseMethod()
	printTodo(data)
	addTodo(&data)
}

func addTodo(d *[]Todo) {
	fmt.Println("Enter the title of your todo")
}
func printTodo(data []Todo) {
	if len(data) == 0 {
		fmt.Printf("Your todo list is empty, please try again later\n")
	} else {
		for i := 0; i < len(data); i++ {
			fmt.Println(data[i].id)
			fmt.Println(data[i].title)
			fmt.Println(data[i].desc)
			fmt.Println(data[i].time)

		}
	}
}
func chooseMethod() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose the operation you want to perform")
	input, _ := reader.ReadString('\n')
	// converting CRLF to LF
	input = strings.Replace(input, "'\n", "", -1)

	if input == "TODO" {
		// render all logic for todo
	}

}
