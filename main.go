package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"bufio"
)

// Run a terminal command
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil{
		log.Fatalf("Command Failed: %v", err)
	}
}

// Get current working directory
func getCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get working directory: %v", err)
	}
	return dir
}


func main() {

	runCommand("docker", "rm", "-f", "apache-server")

	runCommand("docker", "run", "-d", "--name", "apache-server", "-p", "8080:80","-v", 
	fmt.Sprintf("%s/static:/usr/local/apache2/htdocs:ro", getCWD()), "httpd:2.4")

	log.Println("Apache server started at http://localhost:8080		IP Address:")

	
b
	fmt.Println("\nPress 'Enter' to close")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	runCommand("docker", "stop", "apache-server")
}
