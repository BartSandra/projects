package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

func main() {
	linesFlag := flag.Bool("l", false, "Count lines")
	charsFlag := flag.Bool("m", false, "Count characters")
	wordsFlag := flag.Bool("w", false, "Count words")

	flag.Parse()

	if !(*linesFlag || *charsFlag || *wordsFlag) {
		*wordsFlag = true
	}

	fileNames := flag.Args()

	if len(fileNames) == 0 {
		fmt.Println("No files specified.")
		return
	}

	var wg sync.WaitGroup

	for _, fileName := range fileNames {
		wg.Add(1)

		go func(fileName string) {
			defer wg.Done()

			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Printf("Error reading file %s: %s\n", fileName, err)
				return
			}

			result := ""

			if *linesFlag {
				lineCount := strings.Count(string(content), "\n")
				result += fmt.Sprintf("%d\t", lineCount+1)
			}

			if *charsFlag {
				charCount := len(content)
				result += fmt.Sprintf("%d\t", charCount)
			}

			if *wordsFlag {
				wordCount := len(strings.Fields(string(content)))
				result += fmt.Sprintf("%d\t", wordCount)
			}

			fmt.Printf("%s%s\n", result, fileName)
		}(fileName)
	}

	wg.Wait()
}
