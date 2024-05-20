package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	v := reflect.ValueOf(plant)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		key := field.Name
		if unitTag := tag.Get("unit"); unitTag != "" {
			key += "(unit=" + unitTag + ")"
		}
		if colorTag := tag.Get("color_scheme"); colorTag != "" {
			key = "Color(color_scheme=" + colorTag + ")"
		}
		fmt.Println(key + ":" + fmt.Sprintf("%v", v.Field(i).Interface()))
	}
}

func main() {
	plant := UnknownPlant{
		FlowerType: "rosa",
		LeafType:   "lanceolate",
		Color:      255,
	}

	anotherPlant := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}

	describePlant(plant)
	fmt.Println(strings.Repeat("-", 20))
	describePlant(anotherPlant)
}
