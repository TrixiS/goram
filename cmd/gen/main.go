package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"unicode"
)

type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TypeField struct {
	Name        string   `json:"name"`
	Types       []string `json:"types"`
	Description string   `json:"description"`
}

type Type struct {
	Name        string      `json:"name"`
	Href        string      `json:"href"`
	Description []string    `json:"description"`
	Fields      []TypeField `json:"fields"`
}

type Spec struct {
	Enums []Enum `json:"enums"`
	Types []Type `json:"types"`
}

func main() {
	f, err := os.Open("./spec.json")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	spec := Spec{}

	if err := json.NewDecoder(f).Decode(&spec); err != nil {
		panic(err)
	}

	noPtrTypes := make([]string, len(spec.Enums))

	for i, e := range spec.Enums {
		noPtrTypes[i] = e.Name
	}

	for _, t := range spec.Types {
		if len(t.Fields) == 0 {
			noPtrTypes = append(noPtrTypes, t.Name)
		}
	}

	generateEnums(spec.Enums)
	generateTypes(noPtrTypes, spec.Types)
	exec.Command("gofmt", "-s", "-w", "./pkg").Run()
}

const genFileMode = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
const genFilePerm = 0o660

func generateTypes(nonPtrTypes []string, types []Type) {
	b, err := os.ReadFile("./cmd/gen/types/types.go")

	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./pkg/types/types.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	_, err = f.Write(b)

	if err != nil {
		panic(err)
	}

	for _, t := range types {
		for _, d := range t.Description {
			comment := fmt.Sprintf("// %s\n", d)
			f.WriteString(comment)
		}

		if len(t.Fields) == 0 {
			decl := fmt.Sprintf("type %s interface{}\n", t.Name)
			f.WriteString(decl)
			continue
		}

		decl := fmt.Sprintf("type %s struct {\n", t.Name)
		f.WriteString(decl)

		for i, field := range t.Fields {
			if len(field.Types) > 1 {
				fmt.Println(t.Name, field.Types)
			}

			fieldName := toPascalCase(field.Name)
			fieldType := getFieldTypeString(
				field.Name,
				field.Types,
				!slices.Contains(nonPtrTypes, fieldName),
			)

			structField := fmt.Sprintf(
				"%s %s `json:\"%s\"` // %s",
				fieldName,
				fieldType,
				field.Name,
				field.Description,
			)

			f.WriteString(structField)

			if i < len(t.Fields)-1 {
				f.WriteString("\n")
			}
		}

		f.WriteString("\n}\n")
	}

	f.Close()
}

func getFieldTypeString(fieldName string, fieldTypes []string, isPtr bool) string {
	t := fieldTypes[0]

	if t == "Integer" {
		if (fieldName == "id" || fieldName == "chat_id") && len(fieldTypes) == 2 {
			return "ChatID"
		}

		return "int64"
	}

	return convertFieldTypeString(t, isPtr)
}

func convertFieldTypeString(fieldType string, isPtr bool) string {
	if fieldType == "Integer" {
		return "int64"
	}

	if fieldType == "String" {
		return "string"
	}

	if fieldType == "Boolean" {
		return "bool"
	}

	if fieldType == "Float" {
		return "float64"
	}

	const arrayPrefix = "Array of "

	if strings.HasPrefix(fieldType, arrayPrefix) {
		return "[]" + convertFieldTypeString(fieldType[len(arrayPrefix):], false) // no []*T
	}

	if isPtr {
		return "*" + fieldType
	}

	return fieldType
}

func generateEnums(enums []Enum) {
	f, err := os.OpenFile("./pkg/types/enums.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package types\n\n")

	for _, e := range enums {
		decl := fmt.Sprintf("type %s string\n", e.Name)
		f.WriteString(decl)

		f.WriteString("const (\n")

		for _, v := range e.Values {
			name := e.Name + toPascalCase(v)
			assig := fmt.Sprintf("%s %s = \"%s\"\n", name, e.Name, v)
			f.WriteString(assig)
		}

		f.WriteString("\n)\n")
	}

	f.Close()
}

func toPascalCase(v string) string {
	runes := []rune(v)

	builder := &strings.Builder{}
	builder.WriteRune(unicode.ToUpper(runes[0]))

	upper := false

	for i, r := range runes[1:] {
		if r == '_' {
			upper = true
			continue
		}

		i++

		if upper {
			upper = false
			builder.WriteRune(unicode.ToUpper(r))
		} else if (i < len(runes)-2 && r == 'i' && runes[i+1] == 'd') || (i > 0 && r == 'd' && runes[i-1] == 'i') { // detect id
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
