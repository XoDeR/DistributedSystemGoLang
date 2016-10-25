package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("Test Launcher for MicroServices Counter")

	// create server
	exec.Command("f:\\GoLang\\src\\microservices.counter\\aServer\\aServer.exe")

	// create client 1
	exec.Command("f:\\GoLang\\src\\microservices.counter\\aClient\\aClient.exe", "-id 1")

	// create client 2
	exec.Command("f:\\GoLang\\src\\microservices.counter\\aClient\\aClient.exe", "-id 2")

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err1 := bio.ReadLine()
	fmt.Println(err1)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
