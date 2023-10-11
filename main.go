package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Vérifier si des arguments ont été fournis en ligne de commande
	if len(os.Args) == 1 {
		return
	}

	// Paramètres par défaut
	font := "standard.txt"
	color := "white"
	str := ""
	sliceData := ""

	// 1. Rassembler les informations (entrée et drapeaux)
	for i, v := range os.Args[1:] {

		// Gérer le drapeau de couleur
		if len(v) >= 8 && v[0:8] == "--color=" {
			color = v[8:]
		} else if len(v) > 8 && v[0:8] == "--slice=" {
			sliceData = v[8:]
		} else if len(v) > 1 && v[0:2] == "--" {
			continue
		} else if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			str += v
			if i != len(os.Args[1:])-1 {
				str += " "
			}
		}
	}

	// 2. Définir la plage de texte à colorer (s'il y en a une)
	num1, num2, slice := 0, 0, false
	if sliceData != "" {
		num1, num2, slice = IndexSlice(sliceData)
	}

	// 3. Écrire le texte ligne par ligne avec la couleur nécessaire
	args := strings.Split(str, "\\r\\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for j, letter := range word {
				if slice {
					if j >= num1 && j <= num2-1 {
						// Imprimer avec la couleur spécifiée
						fmt.Print(Color(color), Read(1+int(letter-' ')*9+i, font))
					} else {
						// Imprimer en blanc
						fmt.Print(Color("white"), Read(1+int(letter-' ')*9+i, font))
					}
				} else {
					// Imprimer avec la couleur spécifiée
					fmt.Print(Color(color), Read(1+int(letter-' ')*9+i, font))
				}
			}
			fmt.Printf("\n")
		}
	}
}
