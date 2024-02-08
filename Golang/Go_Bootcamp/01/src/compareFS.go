package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	oldFile := flag.String("old", "", "Path to the old filesystem snapshot")
	newFile := flag.String("new", "", "Path to the new filesystem snapshot")
	flag.Parse()

	if *oldFile == "" {
		log.Fatal("Path to the old filesystem snapshot is required")
	}
	if *newFile == "" {
		log.Fatal("Path to the new filesystem snapshot is required")
	}

	addedFiles, removedFiles := compareFilesystems(*oldFile, *newFile)

	for _, file := range addedFiles {
		fmt.Printf("ADDED %s\n", file)
	}
	for _, file := range removedFiles {
		fmt.Printf("REMOVED %s\n", file)
	}
}

// compareFilesystems compares the old and new filesystem snapshots
func compareFilesystems(oldFile, newFile string) ([]string, []string) {
	var addedFiles, removedFiles []string

	oldFiles := make(map[string]bool)
	readSnapshotFile(oldFile, func(line string) {
		oldFiles[line] = true
	})

	readSnapshotFile(newFile, func(line string) {
		if _, ok := oldFiles[line]; !ok {
			addedFiles = append(addedFiles, line)
		} else {
			delete(oldFiles, line)
		}
	})

	for file := range oldFiles {
		removedFiles = append(removedFiles, file)
	}

	sort.Strings(addedFiles)
	sort.Strings(removedFiles)

	return addedFiles, removedFiles
}

// readSnapshotFile reads a snapshot file line by line
func readSnapshotFile(filename string, callback func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read file:", err)
	}
}
