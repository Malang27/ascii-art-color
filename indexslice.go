package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// IndexSlice est une fonction qui prend les données du drapeau --slice= et les convertit en entiers
func IndexSlice(a string) (int, int, bool) {
	index := -1
	for i, v := range a {
		if v == ':' {
			if i == 0 || i == len(a)-1 {
				fmt.Println("Veuillez sélectionner une plage dans un format correct : <int64:int64>")
				os.Exit(0)
			}
			index = i
			break
		}
	}
	num1, e1 := strconv.Atoi(a[:index])
	num2, e2 := strconv.Atoi(a[index+1:])
	if e1 != nil {
		log.Fatal(e1)
		os.Exit(1)
	} else if e2 != nil {
		log.Fatal(e2)
		os.Exit(1)
	}

	return num1, num2, true
}
