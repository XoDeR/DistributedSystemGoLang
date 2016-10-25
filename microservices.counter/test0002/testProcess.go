package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Test Launcher for MicroServices Counter")

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
