package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"os"
	"log"
)

func convert(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	foundPorts := false
	for scanner.Scan() {
		var newLine string
		if strings.Contains(scanner.Text(), "ports:") {
			newLine = "  expose:"
			foundPorts = true
		} else if foundPorts && strings.Contains(scanner.Text(), "-") {
			newLine = strings.Split(scanner.Text(), ":")[0]+`"`
		} else {
			foundPorts = false
			newLine = scanner.Text()
		}
		fmt.Fprintln(out, newLine)
	}

}

func main(){
	config, err := os.Open(os.Args[1])


	if err != nil{
		log.Fatal("Please supply a file path i can open as the first argument")
	}

	convert(config, os.Stdout)
}