package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	isFile      bool
	isDir       bool
	isSymLink   bool
	fileExt     string
	rootDirPath string
)

func main() {
	flag.BoolVar(&isFile, "f", false, "Print only files")
	flag.BoolVar(&isDir, "d", false, "Print only directories")
	flag.BoolVar(&isSymLink, "sl", false, "Print only symbolic links")
	flag.StringVar(&fileExt, "ext", "", "Print only files with a certain extension")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		rootDirPath = args[0]
	} else {
		fmt.Println("Please provide a directory path")
		return
	}

	err := filepath.WalkDir(rootDirPath, visitFile)
	if err != nil {
		fmt.Println(err)
	}
}

func visitFile(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	fileMode := d.Type()
	if !isFile && !isDir && !isSymLink {
		printEntry(path, fileMode)
		return nil
	}

	if isFile && fileMode&fs.ModeType == 0 && (fileExt == "" || filepath.Ext(path) == "."+fileExt) {
		printEntry(path, fileMode)
		return nil
	}

	if isDir && fileMode.IsDir() {
		printEntry(path, fileMode)
		return nil
	}

	if isSymLink && fileMode&fs.ModeType == fs.ModeSymlink {
		linkPath, err := os.Readlink(path)
		if err != nil {
			printEntry(fmt.Sprintf("%s -> [broken]", path), fileMode)
			return nil
		}

		linkAbsPath := filepath.Join(filepath.Dir(path), linkPath)
		linkFileInfo, err := os.Lstat(linkAbsPath)
		if err != nil {
			printEntry(fmt.Sprintf("%s -> [broken]", path), fileMode)
			return nil
		}

		linkFileMode := linkFileInfo.Mode()
		if linkFileMode.IsDir() {
			printEntry(fmt.Sprintf("%s -> %s/", path, linkPath), fileMode)
			return nil
		}

		printEntry(fmt.Sprintf("%s -> %s", path, linkPath), fileMode)
		return nil
	}

	return nil
}

func printEntry(path string, fileMode fs.FileMode) {
	if fileMode&fs.ModeType == fs.ModeSymlink {
		linkPath, err := os.Readlink(path)
		if err != nil {
			fmt.Println(path)
		} else {
			fmt.Printf("%s -> %s\n", path, linkPath)
		}
	} else if fileMode.IsDir() {
		fmt.Printf("%s/\n", path)
	} else {
		fmt.Println(path)
	}
}
