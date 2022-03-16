package main

import (
	"fmt"

	// kin-openapi recommended here https://openapi.tools/#data-validators
	"github.com/getkin/kin-openapi/openapi3"
)

func init() {
	openapi3.SchemaFormatValidationDisabled = true
}

func main() {
	openapi3.DefineStringFormat("timestamp", `[1-9][0-9]*`)
	loader := openapi3.NewLoader()

	doc, err := loader.LoadFromFile("github_openapi.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println("Loading done Succesful")

	if err = doc.Validate(loader.Context); err != nil {
		panic(err)
	}
	fmt.Println("Validation Succesful")

	data := map[string]interface{}{
		"key": "foo",
		"id":  float64(5),
	}
	err = doc.Components.Schemas["key-simple"].Value.VisitJSON(data)
	if err != nil {
		fmt.Println(err)
	}

	// Check Enum validation
	tm := map[string]interface{}{
		"role":  "member",
		"state": "foo",
		"url":   "boo",
	}
	err = doc.Components.Schemas["team-membership"].Value.VisitJSON(tm)
	if err != nil {
		fmt.Println(err)
	}

	// Check Validation of nested structure
	view_traffic := map[string]interface{}{
		"count":   float64(1234),
		"uniques": float64(564),
		"views": []interface{}{
			map[string]interface{}{"count": float64(56), "uniques": float64(67), "timestamp": "2022-03-10T20:02:00Z"},
		},
	}
	err = doc.Components.Schemas["view-traffic"].Value.VisitJSON(view_traffic)
	if err != nil {
		fmt.Println(err)
	}
}
