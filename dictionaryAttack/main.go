package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	// ask for the sha
	var sha string
	fmt.Println("Entrer le sha:")
	fmt.Scanln(&sha)

	// choose the dictionary
	dict := 1
	fmt.Println("Dans quel dictionnaire chercher ? ([1]: dictionary.txt, 2: verbs.txt): ")
	fmt.Scanln(&dict)

	str := "sources/dictionary.txt"

	// open the file
	if dict == 2 {
		str = "sources/verbs.txt"
	}
	file, err := os.Open(str)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	// create a scanner
	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		sum := sha256.Sum256([]byte(fileScanner.Text()))
		str := fmt.Sprintf("%x", sum)

		if sha == str {
			fmt.Println("Le mot de passe is: ")
			fmt.Println(fileScanner.Text())
		}
	}

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}
