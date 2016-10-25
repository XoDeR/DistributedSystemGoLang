package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

/*
// Process launch code snippet
	// Either:

	// The process ends and its error (if any) is received through done
	// 60 seconds have passed and the program is killed.

	cmd := exec.Command("sleep", "5")
	// cmd := exec.Command("cat", "/dev/urandom")

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(60 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
		}
		log.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Printf("process done with error = %v", err)
		} else {
			log.Print("process done gracefully without error")
		}
	}

*/

func RunProcess(processName string, param1 string) {
	errch := make(chan error, 1)
	cmd := exec.Command(processName, param1)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", stdout)

	go func() {
		errch <- cmd.Wait()
	}()

	go func() {
		for _, char := range "|/-\\" {
			fmt.Printf("\r%s...%c", "Running traceroute", char)
			time.Sleep(100 * time.Millisecond)
		}
		scanner := bufio.NewScanner(stdout)
		fmt.Println("")
		for scanner.Scan() {
			line := scanner.Text()
			log.Println(line)
		}
	}()

	select {
	case <-time.After(time.Second * 1):
		log.Println("Timeout hit..")
		return
	case err := <-errch:
		if err != nil {
			log.Println("traceroute failed:", err)
		}
	}
}

func main() {

	// Windows console is not seen -- platform specific issue

	cmd := exec.Command("cmd", "/C", "f:\\GoLang\\src\\microservices.counter\\test0002\\test0002.exe")
	fmt.Printf("%s", cmd)
	//RunProcess("cmd", "/K") //cd ..") //\\GoLang\\src\\microservices.counter\\test0002")
	//RunProcess("cmd /K " + "test0002.exe")

	// pause
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(err)
	fmt.Println(line)
	fmt.Println(hasMoreInLine)
}
