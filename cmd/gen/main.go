package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	Required    bool     `json:"required"`
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

	parser := NewParser(&spec)
	parser.IgnoredTypeNames = []string{"InputFile"}

	generateEnums(spec.Enums)
	generateTypes(parser, spec.Types)
	generateRequests(parser, spec.Methods)
	generateMethods(parser, spec.Methods)

	exec.Command("gofmt", "-s", "-w", "./pkg").Run()
}

const genFileMode = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
const genFilePerm = 0o660

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

func generateMethods(parser *Parser, methods []Method) {
	f, err := os.OpenFile("./pkg/bot/methods.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package bot")

	f.WriteString(`
		import (
			"context"
			"github.com/TrixiS/goram/pkg/types"
		)
	`)

	for _, m := range methods {
		pascalName := toPascalCase(m.Name)
		parsedSpecType := parser.ParseSpecTypes(m.Returns)
		importedTypeString := parsedSpecType.ImportedTypeString("types")
		returnType := fmt.Sprintf("(r %s, err error)", importedTypeString)
		args := ""
		data := "nil"

		if len(m.Fields) == 0 {
			args = "(ctx context.Context)"
		} else {
			args = fmt.Sprintf("(ctx context.Context, request *types.%sRequest)", pascalName)
			data = "request"
		}

		for _, d := range m.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			f.WriteString(comment)
		}

		f.WriteString(fmt.Sprintf("// %s\n", m.Href))

		f.WriteString(
			fmt.Sprintf("func (b *Bot) %s%s %s {\n", pascalName, args, returnType),
		)

		f.WriteString(
			fmt.Sprintf(
				`res, err := makeRequest[%s](ctx, b.options.Client, b.baseURL, "%s", %s)

				if err != nil {
					return r, err
				}

				return res.Result, nil
			}
				`,
				importedTypeString,
				m.Name,
				data,
			),
		)
	}

	f.Close()
}

func generateRequests(parser *Parser, methods []Method) {
	f, err := os.OpenFile("./pkg/types/requests.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package types\n\n")
	f.WriteString("import \"mime/multipart\"\n")
	f.WriteString("import \"encoding/json\"\n\n")

	for _, m := range methods {
		if len(m.Fields) == 0 {
			continue
		}

		pascalName := toPascalCase(m.Type.Name)
		structName := pascalName + "Request"
		f.WriteString(fmt.Sprintf("// use Bot.%s(ctx, &%s{})\n", pascalName, structName))
		typeString := GenerateTypeString(parser, &m.Type, "Request", false, false)
		f.WriteString(typeString)
		f.WriteString("\n")
		generateRequestWriteMultipart(f, parser, &m, structName)
	}

	f.Close()
}

func generateRequestWriteMultipart(
	w io.StringWriter,
	parser *Parser,
	m *Method,
	structName string,
) {
	w.WriteString(
		fmt.Sprintf("func (r *%s) WriteMultipart(w *multipart.Writer) {", structName),
	)

	for _, field := range m.Fields {
		parsedTypeField := parser.ParseTypeField(&field)

		if parsedTypeField.ParsedSpecType.GoType == "InputFile" {
			// TODO: switch on type of input file here
			// TODO: implement uploading
			w.WriteString(fmt.Sprintf(`
				if s, ok := r.%s.(string); ok {
					w.WriteField("%s", s)
				} else {
					// fw, _ := w.CreateFormFile("%s", "%s") 
				}
			`, parsedTypeField.GoName, field.Name, field.Name, "todo.jpeg"))
		} else if parsedTypeField.ParsedSpecType.GoType == "ChatID" {
			w.WriteString(fmt.Sprintf(`
				w.WriteField("%s", r.%s.String())
			`, field.Name, parsedTypeField.GoName))
		} else if parsedTypeField.ParsedSpecType.GoType == "string" &&
			parsedTypeField.ParsedSpecType.ParsedType == ParsedTypePrimitive {

			w.WriteString(fmt.Sprintf(`
					w.WriteField("%s", r.%s)
				`, field.Name, parsedTypeField.GoName))
		} else if parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeEnum {
			w.WriteString(fmt.Sprintf(`
					w.WriteField("%s", string(r.%s))
				`, field.Name, parsedTypeField.GoName))
		} else {
			checkForNil := parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeStruct ||
				parsedTypeField.ParsedSpecType.Levels > 0 ||
				parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeInterface

			if checkForNil {
				w.WriteString(fmt.Sprintf(`if r.%s != nil {`, parsedTypeField.GoName))
			}

			// TODO: write to a field via encoder, do not cast to string
			w.WriteString(fmt.Sprintf(`%s, _ := json.Marshal(r.%s)
					w.WriteField("%s", string(%s))
				`, field.Name, parsedTypeField.GoName, field.Name, field.Name))

			if checkForNil {
				w.WriteString("}\n")
			}
		}
	}

	w.WriteString("}\n")
}

func generateTypes(parser *Parser, types []Type) {
	f, err := os.OpenFile("./pkg/types/types.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package types\n\n")

	for _, t := range types {
		f.WriteString(GenerateTypeString(parser, &t, "", true, true))
	}

	f.Close()
}

func GenerateTypeString(parser *Parser, t *Type, suffix string, doc bool, tagJSON bool) string {
	builder := strings.Builder{}

	// TODO: do not use fmt here
	if doc {
		for _, d := range t.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			builder.WriteString(comment)
		}

		builder.WriteString(fmt.Sprintf("// %s\n", t.Href))
	}

	pascalName := toPascalCase(t.Name)

	if suffix != "" {
		pascalName += suffix
	}

	if len(t.Fields) == 0 {
		decl := fmt.Sprintf("type %s interface{}\n", pascalName)
		builder.WriteString(decl)
		return builder.String()
	}

	decl := fmt.Sprintf("type %s struct {\n", pascalName)
	builder.WriteString(decl)

	for i, field := range t.Fields {
		parsedTypeField := parser.ParseTypeField(&field)
		builder.WriteString(parsedTypeField.StructField(tagJSON))

		if i < len(t.Fields)-1 {
			builder.WriteString("\n")
		}
	}

	builder.WriteString("\n}\n")
	return builder.String()
}

type ParsedType int

const (
	ParsedTypePrimitive ParsedType = iota
	ParsedTypeStruct
	ParsedTypeArray
	ParsedTypeInterface
	ParsedTypeEnum
)

type ParsedSpecType struct {
	ParsedType ParsedType
	GoType     string
	Levels     int // for arrays
}

func (p *ParsedSpecType) ImportedTypeString(pkg string) string {
	builder := strings.Builder{}

	if p.ParsedType == ParsedTypeStruct {
		builder.WriteRune('*')
	} else if p.ParsedType == ParsedTypeArray {
		for i := 0; i < p.Levels; i++ {
			builder.WriteString("[]")
		}
	}

	if p.ParsedType != ParsedTypePrimitive {
		builder.WriteString(pkg)
		builder.WriteRune('.')
	}

	builder.WriteString(p.GoType)
	return builder.String()
}

type ParsedTypeField struct {
	Field          *TypeField
	GoName         string
	ParsedSpecType ParsedSpecType
	Ignored        bool
}

func (p *ParsedTypeField) StructField(tagJSON bool) string {
	builder := strings.Builder{}

	builder.WriteString(p.GoName)
	builder.WriteRune(' ')

	if p.ParsedSpecType.ParsedType == ParsedTypeStruct {
		builder.WriteRune('*')
	} else if p.ParsedSpecType.ParsedType == ParsedTypeArray {
		for i := 0; i < p.ParsedSpecType.Levels; i++ {
			builder.WriteString("[]")
		}
	}

	builder.WriteString(p.ParsedSpecType.GoType)

	if tagJSON {
		builder.WriteRune(' ')
		builder.WriteString("`json:\"")

		if p.Ignored {
			builder.WriteRune('-')
		} else {
			builder.WriteString(p.Field.Name)

			if !p.Field.Required {
				builder.WriteRune(',')
				builder.WriteString("omitempty")
			}
		}

		builder.WriteString("\"`")
	}

	builder.WriteString(" // ")
	builder.WriteString(p.Field.Description)

	return builder.String()
}

type Parser struct {
	EnumNames        []string
	InterfaceNames   []string
	IgnoredTypeNames []string
}

func NewParser(spec *Spec) *Parser {
	p := Parser{}

	for _, e := range spec.Enums {
		p.EnumNames = append(p.EnumNames, e.Name)
	}

	for _, t := range spec.Types {
		if len(t.Fields) == 0 {
			p.InterfaceNames = append(p.InterfaceNames, t.Name)
		}
	}

	return &p
}

func (p *Parser) ParseSpecTypes(types []string) ParsedSpecType {
	psc := ParsedSpecType{}

	if len(types) == 1 && slices.Contains(p.EnumNames, types[0]) {
		psc.GoType = types[0]
		psc.ParsedType = ParsedTypeEnum
	} else if len(types) == 1 && slices.Contains(p.InterfaceNames, types[0]) {
		psc.GoType = types[0]
		psc.ParsedType = ParsedTypeInterface
	} else if len(types) == 2 && types[0] == "InputFile" {
		psc.GoType = "InputFile"
		psc.ParsedType = ParsedTypeInterface
	} else {
		p.parseSpecType(&psc, types[0])
	}

	return psc
}

func (p *Parser) ParseTypeField(t *TypeField) *ParsedTypeField {
	ptf := &ParsedTypeField{
		Field:  t,
		GoName: toPascalCase(t.Name),
	}

	if t.Name == "reply_markup" && len(t.Types) == 4 {
		ptf.ParsedSpecType.GoType = "Markup"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else if len(t.Types) == 4 && t.Name == "media" {
		ptf.ParsedSpecType.GoType = "[]MediaGroupInputMedia"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else if t.Name == "message_ids" {
		ptf.ParsedSpecType.GoType = "int"
		ptf.ParsedSpecType.ParsedType = ParsedTypeArray
		ptf.GoName = "MessageIDs"
	} else if strings.Contains(t.Name, "message") && strings.Contains(t.Name, "id") {
		ptf.ParsedSpecType.GoType = "int"
	} else if len(t.Types) == 2 && (t.Name == "id" || t.Name == "chat_id") {
		ptf.ParsedSpecType.GoType = "ChatID"
	} else if t.Name == "parse_mode" {
		ptf.ParsedSpecType.GoType = "ParseMode"
		ptf.ParsedSpecType.ParsedType = ParsedTypeEnum
	} else {
		ptf.ParsedSpecType = p.ParseSpecTypes(t.Types)
	}

	ptf.Ignored = ptf.ParsedSpecType.ParsedType != ParsedTypeArray &&
		slices.Contains(p.IgnoredTypeNames, ptf.ParsedSpecType.GoType)

	return ptf
}

func (g *Parser) parseSpecType(p *ParsedSpecType, fieldType string) {
	if fieldType == "Integer" {
		p.GoType = "int64"
		return
	} else if fieldType == "String" {
		p.GoType = "string"
		return
	} else if fieldType == "Boolean" {
		p.GoType = "bool"
		return
	} else if fieldType == "Float" {
		p.GoType = "float64"
		return
	}

	const arrayPrefix = "Array of "

	if strings.HasPrefix(fieldType, arrayPrefix) {
		p.Levels += 1
		p.ParsedType = ParsedTypeArray
		g.parseSpecType(p, fieldType[len(arrayPrefix):])
		return
	}

	if p.Levels == 0 {
		p.ParsedType = ParsedTypeStruct
	}

	p.GoType = toPascalCase(fieldType)
}

func toPascalCase(v string) string {
	runes := []rune(v)

	builder := &strings.Builder{}

	upper := false
	c := len(runes)

	for i, r := range runes {
		if r == '_' {
			upper = true
			continue
		}

		if upper || i == 0 || (i == c-2 && r == 'i' && runes[i+1] == 'd') || // detect tailing ID
			(i == c-1 && r == 'd' && (runes[i-1] == 'i' || runes[i-1] == 'I')) {

			upper = false
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
