package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var countBytes, countLines, countWords, countCharacters bool
	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count lines")
	flag.BoolVar(&countWords, "w", false, "Count words")
	flag.BoolVar(&countCharacters, "m", false, "Count Characters")
	flag.Parse()
	if !countBytes && !countLines && !countWords && !countCharacters {
		countBytes = true
		countLines = true
		countWords = true
		countCharacters = true
	}
	args := flag.Args()
	var reader io.Reader
	if len(args) == 1 {
		fileName := args[0]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Printf("Error in closing the file")
				return
			}
		}(file)
		reader = file
	} else {
		reader = os.Stdin
	}
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	fileContent := string(content)
	lines := strings.Count(fileContent, "\n")
	characters := strings.Count(fileContent, "")
	words := strings.Fields(fileContent)
	bytes := len(fileContent)
	if countBytes {
		fmt.Printf("File size: %d bytes\n", bytes)
	}
	if countLines {
		fmt.Printf("Number of lines: %d lines\n", lines)
	}
	if countWords {
		fmt.Printf("Number of words: %d words\n", len(words))
	}
	if countCharacters {
		fmt.Printf("Number of characters: %d characters\n", characters)
	}
}
