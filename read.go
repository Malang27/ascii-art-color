package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read est une fonction pour lire la police depuis des fichiers
func Read(num int, filename string) string {

	f, e := os.Open(filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line
}
