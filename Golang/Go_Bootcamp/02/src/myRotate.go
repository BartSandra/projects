package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	archiveDir := flag.String("a", "", "Archive directory")
	flag.Parse()

	files := flag.Args()

	if *archiveDir == "" {
		fmt.Println("Archive directory must be specified")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go rotateFile(file, *archiveDir, &wg)
	}

	wg.Wait()
}

func rotateFile(file, archiveDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("Failed to get file info:", err)
		return
	}
	unixTimestamp := fileInfo.ModTime().Unix()

	srcFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Failed to open source file:", err)
		return
	}
	defer srcFile.Close()

	archiveFileName := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(file), unixTimestamp)

	archivePath := filepath.Join(archiveDir, archiveFileName)

	archiveFile, err := os.Create(archivePath)
	if err != nil {
		fmt.Println("Failed to create archive file:", err)
		return
	}
	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, srcFile)
	if err != nil {
		fmt.Println("Failed to copy file content:", err)
		return
	}

	err = os.Remove(file)
	/*if err != nil {
		fmt.Println("Failed to remove source file:", err)
		return
	}*/

	fmt.Println("Successfully rotated file", file, "to", archivePath)
}
