package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	chooseMethod()
}

func chooseMethod() {
	reader := bufio.NewReader(os.Stdin)
	for k, m := range options {

		// m is a map[string]interface.
		// loop over keys and values in the map.

		fmt.Println(k, ":", m)
	}
	fmt.Println("Choose the operation you want to perform")

	input, _ := reader.ReadString('\n')
	// converting CRLF to LF
	input = strings.Replace(input, "'\n", "", -1)
	if input == "Todo" {
		Choosetodo()
	}

}

var options = map[string]string{
	"Todo":   "in memory todo app",
	"Chat":   "A socket based chat app",
	"Broker": "simple in memory message broker",
}
