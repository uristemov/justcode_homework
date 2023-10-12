package main

import (
	"encoding/json"
	"fmt"
	"github.com/dave/jennifer/jen"
	"log"
	"strings"
)

func generateStructFromJSON(jsonData string, structName string) (*jen.File, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(jsonData)).Decode(&data); err != nil {
		return nil, err
	}

	file := jen.NewFile("main")

	generateStructFields(file, structName, data)

	return file, nil
}

func generateStructFields(f *jen.File, structName string, data map[string]interface{}) {
	f.Type().Id(structName).StructFunc(func(g *jen.Group) {
		for key, value := range data {
			fieldName := jen.Id(key)

			switch val := value.(type) {
			case map[string]interface{}:
				g.Add(fieldName).Add(generateNestedStruct(val))
			case []interface{}:
				if len(val) == 0 {
					g.Add(fieldName).Index().Interface()
				} else {
					g.Add(fieldName).Index().Add(reflectType(val[0]))
				}
			default:
				g.Add(fieldName).Add(reflectType(value))
			}
		}
	})
}

func generateNestedStruct(data map[string]interface{}) *jen.Statement {
	return jen.StructFunc(func(g *jen.Group) {
		for key, value := range data {
			fieldName := jen.Id(key)

			switch val := value.(type) {
			case map[string]interface{}:
				g.Add(fieldName).Add(generateNestedStruct(val))
			case []interface{}:
				if len(val) == 0 {
					g.Add(fieldName).Index().Interface()
				} else {
					g.Add(fieldName).Index().Add(reflectType(val[0]))
				}
			default:
				g.Add(fieldName).Add(reflectType(value))
			}
		}
	})
}

func reflectType(value interface{}) *jen.Statement {
	switch value.(type) {
	case string:
		return jen.String()
	case float64:
		return jen.Float64()
	default:
		return jen.Interface()
	}
}

func main() {
	jsonData := `{
		"field1": 1,
		"nested": {
			"field2": "value2",
			"field3": {
				"field4": "value4"
			}
		},
		"arrayField": ["item1", "item2"]
	}`

	structName := "GeneratedStruct"

	file, err := generateStructFromJSON(jsonData, structName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", file)
}
