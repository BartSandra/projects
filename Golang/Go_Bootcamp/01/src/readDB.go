package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type XMLDatabase struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type DBReader interface {
	Read(file string) (XMLDatabase, error)
}

type JSONReader struct{}

func (jr JSONReader) Read(file string) (XMLDatabase, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return XMLDatabase{}, err
	}

	var database XMLDatabase
	err = json.Unmarshal(data, &database)
	if err != nil {
		return XMLDatabase{}, err
	}

	return database, nil
}

type XMLReader struct{}

func (jr XMLReader) Read(file string) (XMLDatabase, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return XMLDatabase{}, err
	}

	var database XMLDatabase
	err = xml.Unmarshal(data, &database)
	if err != nil {
		return XMLDatabase{}, err
	}

	return database, nil
}

func main() {
	jsonReader := JSONReader{}
	xmlReader := XMLReader{}

	file := os.Args[2]
	fileExtension := strings.ToLower(filepath.Ext(file))

	var database XMLDatabase
	var err error

	switch fileExtension {
	case ".json":
		database, err = jsonReader.Read(file)
	case ".xml":
		database, err = xmlReader.Read(file)
	default:
		fmt.Println("Invalid file extension")
		return
	}

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if fileExtension == ".json" {
		xmlData, err := xml.MarshalIndent(database, "", "    ")
		if err != nil {
			fmt.Println("Error converting to XML:", err)
			return
		}
		fmt.Println("XML:", string(xmlData))
	} else if fileExtension == ".xml" {
		jsonData, err := json.MarshalIndent(database, "", "    ")
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}
		fmt.Println("JSON:", string(jsonData))
	}
}
