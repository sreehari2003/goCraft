package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	chooseMethod()
}

func chooseMethod() {
	reader := bufio.NewScanner(os.Stdin)
	for k, m := range options {

		// m is a map[string]interface.
		// loop over keys and values in the map.

		fmt.Println(k, ":", m)
	}
	fmt.Println("Choose the operation you want to perform")

	reader.Scan()
	input := reader.Text()
	// converting CRLF to LF
	if input == "Todo" {
		Choosetodo()
	}

}

var options = map[string]string{
	"Todo":   "in memory todo app",
	"Chat":   "A socket based chat app",
	"Broker": "simple in memory message broker",
}
