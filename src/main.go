package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	showIntro()

	for {

		showCommands()

		command := getCommand()

		fmt.Println("you have choose:", command)

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("showing logs...")
		case 0:
			fmt.Println("exiting...")
			os.Exit(0)
		default:
			fmt.Println("invalid command, try again")
		}

	}

}

func clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func showIntro() {
	clear()

	hostname, _ := os.Hostname()

	fmt.Println("Welcome,", hostname)
}

func showCommands() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func getCommand() int {
	var command int

	fmt.Print("type here: ")
	fmt.Scan(&command)

	return command
}

func initMonitoring() {
	fmt.Println("monitoring...")

	var qty int
	fmt.Print("how many origins do you want to monitor? ")
	fmt.Scan(&qty)

	if qty <= 0 || reflect.TypeOf(qty).String() != "int" {
		fmt.Println("invalid quantity, try again")
		return
	}

	var domains []string

	for i := 0; i < qty; i++ {
		var origin string

		fmt.Print(strconv.Itoa(i+1) + "° origin name: ")

		fmt.Scan(&origin)

		domains = append(domains, origin)

	}

	majorLen := getMajorLen(domains)

	for i, domain := range domains {

		if i == 0 {
			fmt.Println(strings.Repeat("*", majorLen+25))
		}

		makeRequest(domain)

		fmt.Println(strings.Repeat("*", majorLen+25))
	}

}

func makeRequest(domain string) bool {
	res, err := http.Get(domain)

	if err != nil {
		fmt.Println("an error occurred:", err)
		return false
	}

	requestURL := res.Request.URL
	statusCode := res.StatusCode

	fmt.Println("requested URL:", requestURL)
	fmt.Println("status code:", statusCode)

	return true
}

func getMajorLen(strings []string) int {
	output := len(strings[0])

	for i := 1; i < len(strings); i++ {
		if len(strings[i]) > output {
			output = len(strings[i])
		}
	}

	return output
}
