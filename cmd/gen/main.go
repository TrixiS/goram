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

type Method struct {
	Type
	Returns []string `json:"returns"`
}

type Spec struct {
	Enums   []Enum   `json:"enums"`
	Types   []Type   `json:"types"`
	Methods []Method `json:"methods"`
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

	nonPtrTypes := make([]string, len(spec.Enums))

	for i, e := range spec.Enums {
		nonPtrTypes[i] = e.Name
	}

	for _, t := range spec.Types {
		if len(t.Fields) == 0 {
			nonPtrTypes = append(nonPtrTypes, t.Name)
		}
	}

	generateEnums(spec.Enums)
	generateTypes(spec.Types, nonPtrTypes)
	generateRequests(spec.Methods, nonPtrTypes)
	generateMethods(spec.Methods)
	exec.Command("gofmt", "-s", "-w", "./pkg").Run()
}

const genFileMode = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
const genFilePerm = 0o660

func generateMethods(methods []Method) {
	f, err := os.OpenFile("./pkg/bot/methods.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package bot\n\n")

	f.WriteString(`import (
		"context"
		"github.com/TrixiS/goram/pkg/types"
	)
		`)

	for _, m := range methods {
		name := toPascalCase(m.Name)
		t := getFieldTypeString("", m.Returns, true, "types.")
		returnType := fmt.Sprintf("(r %s, err error)", t)
		args := ""
		data := "nil"

		if len(m.Fields) == 0 {
			args = "(ctx context.Context)"
		} else {
			args = fmt.Sprintf("(ctx context.Context, request *types.%sRequest)", name)
			data = "request"
		}

		for _, d := range m.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			f.WriteString(comment)
		}

		f.WriteString(fmt.Sprintf("// %s\n", m.Href))

		f.WriteString(
			fmt.Sprintf("func (b *Bot) %s%s %s {\n", name, args, returnType),
		)

		f.WriteString(
			fmt.Sprintf(
				`res, err := makeRequest[%s](ctx, b.options.Client, b.baseURL, "%s", %s)

					if err != nil {
						return r, err
					}

					if !res.OK {
						return r, res.error("%s")
					}

					return res.Result, nil
				}

				`,
				t,
				m.Name,
				data,
				m.Name,
			),
		)
	}

	f.Close()
}

func generateRequests(methods []Method, nonPtrTypes []string) {
	f, err := os.OpenFile("./pkg/types/requests.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package types\n\n")

	for _, m := range methods {
		if len(m.Fields) == 0 {
			continue
		}

		methodPascalName := toPascalCase(m.Type.Name)
		structName := methodPascalName + "Request"
		f.WriteString(fmt.Sprintf("// use Bot.%s(ctx, &%s{})\n", methodPascalName, structName))
		typeString := generateTypeString(m.Type, structName, nonPtrTypes, false)
		f.WriteString(typeString)
		f.WriteString("\n")
	}

	f.Close()
}

func generateTypes(types []Type, nonPtrTypes []string) {
	f, err := os.OpenFile("./pkg/types/types.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package types\n\n")

	for _, t := range types {
		f.WriteString(generateTypeString(t, t.Name, nonPtrTypes, true))
	}

	f.Close()
}

func generateTypeString(t Type, typeName string, nonPtrTypes []string, doc bool) string {
	builder := strings.Builder{}

	if doc {
		for _, d := range t.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			builder.WriteString(comment)
		}

		builder.WriteString(fmt.Sprintf("// %s\n", t.Href))
	}

	if len(t.Fields) == 0 {
		decl := fmt.Sprintf("type %s interface{}\n", typeName)
		builder.WriteString(decl)
		return builder.String()
	}

	decl := fmt.Sprintf("type %s struct {\n", typeName)
	builder.WriteString(decl)

	for i, field := range t.Fields {
		fieldName := toPascalCase(field.Name)
		fieldType := getFieldTypeString(
			field.Name,
			field.Types,
			!slices.Contains(nonPtrTypes, field.Types[0]),
			"",
		)

		structField := fmt.Sprintf(
			"%s %s `json:\"%s,omitempty\"` // %s",
			fieldName,
			fieldType,
			field.Name,
			field.Description,
		)

		builder.WriteString(structField)

		if i < len(t.Fields)-1 {
			builder.WriteString("\n")
		}
	}

	builder.WriteString("\n}\n")

	return builder.String()
}

func getFieldTypeString(fieldName string, fieldTypes []string, isPtr bool, prefix string) string {
	if fieldName == "reply_markup" && len(fieldTypes) == 4 {
		return prefix + "Markup"
	}

	if fieldName == "media" && len(fieldTypes) == 4 {
		return "[]" + prefix + "MediaGroupInputMedia"
	}

	if fieldName == "message_id" {
		return "int"
	}

	t := fieldTypes[0]

	if t == "Integer" {
		if (fieldName == "id" || fieldName == "chat_id") && len(fieldTypes) == 2 {
			return prefix + "ChatID"
		}

		return "int64"
	}

	return convertFieldTypeString(t, isPtr, prefix)
}

func convertFieldTypeString(fieldType string, isPtr bool, prefix string) string {
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
		return "[]" + convertFieldTypeString(fieldType[len(arrayPrefix):], false, prefix) // no []*T
	}

	if isPtr {
		return "*" + prefix + fieldType
	}

	return prefix + fieldType
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
		} else if (i < len(runes)-2 && r == 'i' && runes[i+1] == 'd' && i+1 == len(runes)-1) ||
			(i > 0 && i == len(runes)-1 && r == 'd' && runes[i-1] == 'i') { // detect id

			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
