package main

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"fmt"
	"os"
)

func main() {
	// struct1 := Struct1{
	// 	Astr: "xxxxxxx",
	// 	AStruct: []Struct2{
	// 		{Something: "1"},
	// 		{Something: "2"},
	// 	},
	// 	Amap: map[string]Struct2{
	// 		"A": {},
	// 		"B": {},
	// 	},
	// }

	// // if err := xxx.Validate(); err != nil {
	// // 	fmt.Printf("%v\n", err)
	// // 	ebs, _ := json.Marshal(err)
	// // 	fmt.Printf("%v\n", string(ebs))
	// // }

	// validator := playground.New()
	// if err := validator.Struct(struct1); err != nil {
	// 	ebs, _ := json.Marshal(err)
	// 	fmt.Printf("%v\n", string(ebs))

	// 	for _, err := range err.(playground.ValidationErrors) {
	// 		fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
	// 		fmt.Println(err.Field())     // by passing alt name to ReportError like below
	// 		fmt.Println(err.StructNamespace())
	// 		fmt.Println(err.StructField())
	// 		fmt.Println(err.Tag())
	// 		fmt.Println(err.ActualTag())
	// 		fmt.Println(err.Kind())
	// 		fmt.Println(err.Type())
	// 		fmt.Println(err.Value())
	// 		fmt.Println(err.Param())
	// 		fmt.Println()
	// 	}
	// }

	// config := "../openapi-parser-templates/go/server/echo/handlebars/project.yml"
	// config := "../openapi-parser-templates/go/server/echo/go/project.yml"
	schemaPath := "../codegen-templates/openapi.yaml"

	// fmt.Println("=================================================================")
	// fmt.Println("=================================================================")

	ldr := loader.New()
	rc, err := ldr.Load(schemaPath)

	if err != nil {
		os.Exit(handleError(err))
	}

	g := generator.New()

	if err := g.Generate(rc); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v", err.Error())
	return 1
}
