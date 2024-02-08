package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit,omitempty"`
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

func readXMLDatabase(xmlFile string, reader XMLReader) XMLDatabase {
	database, err := reader.Read(xmlFile)
	if err != nil {
		log.Fatal("Failed to read the original XML database: ", err)
	}
	return database
}

func readJSONDatabase(jsonFile string, reader JSONReader) XMLDatabase {
	database, err := reader.Read(jsonFile)
	if err != nil {
		log.Fatal("Failed to read the new JSON database: ", err)
	}
	return database
}

func compareDatabases(originalDB, newDB XMLDatabase) {
	for _, recipe := range originalDB.Cakes {
		found := false
		for _, newRecipe := range newDB.Cakes {
			if recipe.Name == newRecipe.Name {
				found = true
				compareRecipes(recipe, newRecipe)
				break
			}
		}
		if !found {
			fmt.Printf("REMOVED cake \"%s\"\n", recipe.Name)
		}
	}

	for _, newRecipe := range newDB.Cakes {
		found := false
		for _, recipe := range originalDB.Cakes {
			if recipe.Name == newRecipe.Name {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("ADDED cake \"%s\"\n", newRecipe.Name)
		}
	}
}

func compareRecipes(originalRecipe, newRecipe Cake) {
	if originalRecipe.Time != newRecipe.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", originalRecipe.Name, newRecipe.Time, originalRecipe.Time)
	}

	originalIngredients := make(map[string]Ingredient)
	for _, ingredient := range originalRecipe.Ingredients {
		originalIngredients[ingredient.Name] = ingredient
	}

	for _, newIngredient := range newRecipe.Ingredients {
		originalIngredient, ok := originalIngredients[newIngredient.Name]
		if ok {
			delete(originalIngredients, newIngredient.Name)

			if originalIngredient.Count != newIngredient.Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Name, originalRecipe.Name, newIngredient.Count, originalIngredient.Count)
			}

			if originalIngredient.Unit != newIngredient.Unit {
				if originalIngredient.Unit != "" && newIngredient.Unit != "" {
					fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Name, originalRecipe.Name, newIngredient.Unit, originalIngredient.Unit)
				}
			}
		} else {
			if newIngredient.Unit == "" {
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, originalRecipe.Name)
			} else {
				fmt.Printf("ADDED ingredient \"%s\" with unit \"%s\" for cake \"%s\"\n", newIngredient.Name, newIngredient.Unit, originalRecipe.Name)
			}
		}
	}

	for _, deletedIngredient := range originalIngredients {
		if deletedIngredient.Unit == "" {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", deletedIngredient.Name, originalRecipe.Name)
		} else {
			fmt.Printf("REMOVED ingredient \"%s\" with unit \"%s\" for cake \"%s\"\n", deletedIngredient.Name, deletedIngredient.Unit, originalRecipe.Name)
		}
	}
}

func main() {
	xmlFile := flag.String("old", "", "Path to the original XML database")
	jsonFile := flag.String("new", "", "Path to the new JSON database")
	flag.Parse()

	if *xmlFile == "" {
		log.Fatal("Path to the original XML database is required")
	}
	if *jsonFile == "" {
		log.Fatal("Path to the new JSON database is required")
	}

	xmlReader := XMLReader{}
	jsonReader := JSONReader{}

	originalDB := readXMLDatabase(*xmlFile, xmlReader)
	newDB := readJSONDatabase(*jsonFile, jsonReader)

	compareDatabases(originalDB, newDB)
}
