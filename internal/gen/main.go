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

	updateType := spec.Types[0]

	generateEnums(updateType, spec.Enums)
	generateTypes(parser, spec.Types)
	generateRequests(parser, spec.Methods)
	generateMethods(parser, spec.Methods)
	generateHandlers(updateType)

	exec.Command("gofmt", "-s", "-w", ".").Run()
}

const genFileMode = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
const genFilePerm = 0o660

func generateHandlers(updateType Type) {
	f, err := os.OpenFile("./handlers/handlers.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package handlers\n\n")
	f.WriteString(`
		import "github.com/TrixiS/goram"
		import "context"
	`)

	f.WriteString(`
		type routerHandlers[T any] struct {
			filters []Filter[T] // router-level filters for this update
			handlers []handler[T]
		}`,
	)

	f.WriteString("\n\ntype handlers struct {\n")

	for _, update := range updateType.Fields[1:] {
		fieldName := toPascalCase(update.Name, false)
		f.WriteString(fmt.Sprintf("%s routerHandlers[*goram.%s]\n", fieldName, update.Types[0]))
	}

	f.WriteString("}\n")

	for _, update := range updateType.Fields[1:] {
		updatePascalName := toPascalCase(update.Name, true)
		structFieldName := toPascalCase(update.Name, false)
		typeVar := "goram." + update.Types[0]

		f.WriteString(fmt.Sprintf(`
			// Add %s handler with provided filters
			func (r *Router) %s(handlerFunc Func[*%s], filters... Filter[*%s]) *Router {
				h := handler[*%s]{
				cb: handlerFunc,
				filters: filters,
				}

				r.handlers.%s.handlers = append(r.handlers.%s.handlers, h)
				return r
			}
		`,
			updatePascalName,
			updatePascalName,
			typeVar,
			typeVar,
			typeVar,
			structFieldName,
			structFieldName,
		))
	}

	for _, u := range updateType.Fields[1:] {
		handlersFieldName := toPascalCase(u.Name, false)
		updatePascalName := toPascalCase(u.Name, true)

		f.WriteString(fmt.Sprintf(`
			// Add router-level filter(s) to %s update
			func (r *Router) Filter%s(filters... Filter[*goram.%s]) *Router {
				r.handlers.%s.filters = append(r.handlers.%s.filters, filters...)
				return r
			}
		`,
			updatePascalName,
			updatePascalName,
			u.Types[0],
			handlersFieldName,
			handlersFieldName,
		))
	}

	f.WriteString("\n")

	for _, u := range updateType.Fields[1:] {
		pascalName := toPascalCase(u.Name, true)
		def := fmt.Sprintf(
			"func (r *Router) call%sHandlers(ctx context.Context, bot *goram.Bot, update *goram.%s, data Data) (bool, error) {",
			pascalName,
			u.Types[0],
		)

		f.WriteString(def)

		handlersFieldName := toPascalCase(u.Name, false)

		body := fmt.Sprintf(`
			for _, filter := range r.handlers.%s.filters {
				if !filter(ctx, bot, update, data) {
					return false, nil
				}
			}

			found, err := callHandlers(ctx, bot, r.handlers.%s.handlers, update, data)

			if found {
				return found, err
			}

			for _, child := range r.children {
				found, err := child.call%sHandlers(ctx, bot, update, data)

				if found {
					return found, err
				}
			}

			return false, nil
		`,
			handlersFieldName,
			handlersFieldName,
			pascalName,
		)

		f.WriteString(body)
		f.WriteString("}\n\n")
	}

	f.WriteString(
		"func (r *Router) feedUpdate(ctx context.Context, bot *goram.Bot, update *goram.Update, data Data) (bool, error) {\n",
	)

	for _, u := range updateType.Fields[1:] {
		fieldName := toPascalCase(u.Name, true)

		f.WriteString(fmt.Sprintf("if update.%s != nil {\n", fieldName))
		f.WriteString(fmt.Sprintf(
			"return r.call%sHandlers(ctx, bot, update.%s, data)\n",
			fieldName,
			fieldName,
		))

		f.WriteString("}\n")
	}

	f.WriteString("return false, nil\n")
	f.WriteString("}\n")

	f.Close()
}

func generateEnums(updateType Type, enums []Enum) {
	f, err := os.OpenFile("./enums.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package goram\n\n")

	for _, e := range enums {
		decl := fmt.Sprintf("type %s string\n", e.Name)
		f.WriteString(decl)

		f.WriteString("const (\n")

		for _, v := range e.Values {
			name := e.Name + toPascalCase(v, true)
			assig := fmt.Sprintf("%s %s = \"%s\"\n", name, e.Name, v)
			f.WriteString(assig)
		}

		f.WriteString("\n)\n")
	}

	f.WriteString("type UpdateType string\n\n")
	f.WriteString("const (\n")

	for _, field := range updateType.Fields[1:] { // skip update_id
		name := "Update" + toPascalCase(field.Name, true)
		f.WriteString(fmt.Sprintf("// %s\n", field.Description[len("Optional. "):]))
		f.WriteString(fmt.Sprintf(`%s UpdateType = "%s"`, name, field.Name))
		f.WriteString("\n\n")
	}

	f.WriteString(")\n")

	f.Close()
}

func generateMethods(parser *Parser, methods []Method) {
	f, err := os.OpenFile("./methods.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package goram")

	f.WriteString(`
		import (
			"context"
		)
	`)

	for _, m := range methods {
		pascalName := toPascalCase(m.Name, true)
		structName := pascalName

		if !strings.HasSuffix(structName, "Request") {
			structName += "Request"
		}

		parsedSpecType := parser.ParseSpecTypes(m.Returns)
		typeString := parsedSpecType.TypeString()
		returnType := fmt.Sprintf("(r %s, err error)", typeString)
		args := ""
		data := "nil"

		if len(m.Fields) == 0 {
			args = "(ctx context.Context)"
		} else {
			args = fmt.Sprintf("(ctx context.Context, request *%s)", structName)
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
				`res, err := makeRequest[%s](ctx, b.options.Client, b.baseUrl, "%s", b.options.FloodHandler, %s)

				if err != nil {
					return r, err
				}

				return res.Result, nil
			}
				`,
				typeString,
				m.Name,
				data,
			),
		)
	}

	f.Close()
}

func generateRequests(parser *Parser, methods []Method) {
	f, err := os.OpenFile("./requests.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package goram\n\n")
	f.WriteString(`
		import (
			"mime/multipart"
			"encoding/json"
			"io"
			"strconv"
		)
	`)

	for _, m := range methods {
		if len(m.Fields) == 0 {
			continue
		}

		pascalName := toPascalCase(m.Type.Name, true)
		structName := pascalName
		suffix := ""

		if !strings.HasSuffix(structName, "Request") {
			structName += "Request"
			suffix = "Request"
		}

		f.WriteString(fmt.Sprintf("// see Bot.%s(ctx, &%s{})\n", pascalName, structName))
		typeString := generateTypeString(parser, &m.Type, suffix, false, false)
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
		fmt.Sprintf("func (r *%s) writeMultipart(w *multipart.Writer) {", structName),
	)

	for _, field := range m.Fields {
		parsedTypeField := parser.ParseTypeField(&field)

		if parsedTypeField.ParsedSpecType.GoType == "InputFile" {
			w.WriteString(fmt.Sprintf(`
				if r.%s.FileId != "" {
					w.WriteField("%s", r.%s.FileId)
				} else if r.%s.Reader != nil {
					fw, _ := w.CreateFormFile("%s", r.%s.Reader.Name()) 
					io.Copy(fw, r.%s.Reader)
				}
			`,
				parsedTypeField.GoName,
				field.Name,
				parsedTypeField.GoName,
				parsedTypeField.GoName,
				field.Name,
				parsedTypeField.GoName,
				parsedTypeField.GoName,
			))
		} else if parsedTypeField.ParsedSpecType.GoType == "InputMedia" &&
			parsedTypeField.ParsedSpecType.ParsedType != ParsedTypeArray {

			w.WriteString(fmt.Sprintf(`
				{
					inputFile := r.%s.getMedia()

					if inputFile.Reader != nil {
						fw, _ := w.CreateFormFile("%s", inputFile.Reader.Name())
						io.Copy(fw, inputFile.Reader)
						r.%s.setMedia("attach://%s")	
					}

					b, _ := json.Marshal(r.%s)
					fw, _ := w.CreateFormField("%s")
					fw.Write(b)
				}
			`,
				parsedTypeField.GoName,
				field.Name,
				parsedTypeField.GoName,
				field.Name,
				parsedTypeField.GoName,
				field.Name,
			))
		} else if parsedTypeField.ParsedSpecType.GoType == "InputMedia" &&
			parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeArray &&
			parsedTypeField.ParsedSpecType.Levels == 1 {

			w.WriteString(fmt.Sprintf(`
				for i := 0; i < len(r.%s); i++ {
					inputMedia := r.%s[i]
					fieldName := "%s" + strconv.Itoa(i)
					inputFile := inputMedia.getMedia()
					if inputFile.Reader != nil {
						fw, _ := w.CreateFormFile(fieldName, inputFile.Reader.Name())
						io.Copy(fw, inputFile.Reader)
						inputMedia.setMedia("attach://" + fieldName)
					}					
				}
				{
					b, _ := json.Marshal(r.%s)
					fw, _ := w.CreateFormField("%s")
					fw.Write(b)
				}
			`,
				parsedTypeField.GoName,
				parsedTypeField.GoName,
				field.Name,
				parsedTypeField.GoName,
				field.Name,
			))
		} else if parsedTypeField.ParsedSpecType.GoType == "ChatId" {
			w.WriteString(fmt.Sprintf(
				`w.WriteField("%s", r.%s.String())
				`,
				field.Name,
				parsedTypeField.GoName,
			))
		} else if parsedTypeField.ParsedSpecType.GoType == "string" &&
			parsedTypeField.ParsedSpecType.ParsedType == ParsedTypePrimitive {

			if !parsedTypeField.Field.Required {
				w.WriteString(fmt.Sprintf(`if r.%s != "" {`, parsedTypeField.GoName))
			}

			w.WriteString(fmt.Sprintf(`
				w.WriteField("%s", r.%s)
				`,
				field.Name,
				parsedTypeField.GoName,
			))

			if !parsedTypeField.Field.Required {
				w.WriteString("}\n")
			}
		} else if parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeEnum {
			w.WriteString(fmt.Sprintf(
				`w.WriteField("%s", string(r.%s))
				`,
				field.Name,
				parsedTypeField.GoName,
			))
		} else {
			checkForNil := !parsedTypeField.Field.Required && (parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeStruct ||
				parsedTypeField.ParsedSpecType.Levels > 0 ||
				parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeInterface)

			if checkForNil {
				w.WriteString(fmt.Sprintf(`if r.%s != nil {`, parsedTypeField.GoName))
			}

			w.WriteString(fmt.Sprintf(`
				{
					b, _ := json.Marshal(r.%s)
					fw, _ := w.CreateFormField("%s")
					fw.Write(b)
				}
			`, parsedTypeField.GoName, field.Name))

			if checkForNil {
				w.WriteString("}\n")
			}
		}
	}

	w.WriteString("}\n")
}

var builtinTypes = []string{"InputMedia", "InputFile"}

func generateTypes(parser *Parser, types []Type) {
	f, err := os.OpenFile("./types.go", genFileMode, genFilePerm)

	if err != nil {
		panic(err)
	}

	f.WriteString("package goram\n\n")

	for _, t := range types {
		if slices.Contains(builtinTypes, t.Name) {
			continue
		}

		f.WriteString(generateTypeString(parser, &t, "", true, true))

		const inputMediaPrefix = "InputMedia"
		const inputPaidMediaPrefix = "InputPaidMedia"

		if strings.HasPrefix(t.Name, inputMediaPrefix) ||
			(strings.HasPrefix(t.Name, inputPaidMediaPrefix) && len(t.Fields) > 0) {
			generateInputMediaMethods(f, &t)
		}
	}

	f.Close()
}

func generateInputMediaMethods(w io.StringWriter, t *Type) {
	w.WriteString(fmt.Sprintf(`
		func (i *%s) setMedia(fileId string) {
			i.Media.FileId = fileId
		}
	`, t.Name))

	w.WriteString(fmt.Sprintf(`
		func (i *%s) getMedia() InputFile {
			return i.Media
		}
	`, t.Name))
}

func generateTypeString(parser *Parser, t *Type, suffix string, doc bool, tagJSON bool) string {
	builder := strings.Builder{}

	if doc {
		for _, d := range t.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			builder.WriteString(comment)
		}

		builder.WriteString(fmt.Sprintf("// %s\n", t.Href))
	}

	pascalName := toPascalCase(t.Name, true)

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

func (p *ParsedSpecType) TypeString() string {
	builder := strings.Builder{}

	if p.ParsedType == ParsedTypeStruct {
		builder.WriteRune('*')
	} else if p.ParsedType == ParsedTypeArray {
		for i := 0; i < p.Levels; i++ {
			builder.WriteString("[]")
		}
	}

	builder.WriteString(p.GoType)
	return builder.String()
}

type ParsedTypeField struct {
	Field          *TypeField
	GoName         string
	ParsedSpecType ParsedSpecType
}

func (p *ParsedTypeField) StructField(tagJSON bool) string {
	builder := strings.Builder{}

	builder.WriteString(p.GoName)
	builder.WriteRune(' ')
	builder.WriteString(p.ParsedSpecType.TypeString())

	if tagJSON {
		builder.WriteRune(' ')
		builder.WriteString("`json:\"")
		builder.WriteString(p.Field.Name)

		if !p.Field.Required {
			builder.WriteRune(',')
			builder.WriteString("omitempty")
		}

		builder.WriteString("\"`")
	}

	builder.WriteString(" // ")
	builder.WriteString(p.Field.Description)

	return builder.String()
}

type Parser struct {
	EnumNames      []string
	InterfaceNames []string
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
		GoName: toPascalCase(t.Name, true),
	}

	if t.Name == "reply_markup" && len(t.Types) == 4 {
		ptf.ParsedSpecType.GoType = "Markup"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else if len(t.Types) == 4 && t.Name == "media" {
		ptf.ParsedSpecType.GoType = "InputMedia"
		ptf.ParsedSpecType.ParsedType = ParsedTypeArray
		ptf.ParsedSpecType.Levels = 1
	} else if len(t.Types) == 2 && (t.Name == "id" || t.Name == "chat_id" || t.Name == "from_chat_id") {
		ptf.ParsedSpecType.GoType = "ChatId"
	} else if t.Types[0] == "Integer" &&
		(strings.HasSuffix(t.Name, "id") || t.Name == "offset") &&
		t.Name != "message_id" {

		ptf.ParsedSpecType.GoType = "int64"
	} else if t.Name == "allowed_updates" {
		ptf.ParsedSpecType.GoType = "UpdateType"
		ptf.ParsedSpecType.ParsedType = ParsedTypeArray
		ptf.ParsedSpecType.Levels = 1
	} else if t.Name == "parse_mode" {
		ptf.ParsedSpecType.GoType = "ParseMode"
		ptf.ParsedSpecType.ParsedType = ParsedTypeEnum
	} else if t.Name == "media" && t.Types[0] == "String" {
		ptf.ParsedSpecType.GoType = "InputFile"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else {
		ptf.ParsedSpecType = p.ParseSpecTypes(t.Types)
	}

	return ptf
}

func (g *Parser) parseSpecType(p *ParsedSpecType, fieldType string) {
	if fieldType == "Integer" {
		p.GoType = "int"
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

	p.GoType = toPascalCase(fieldType, true)
}

func toPascalCase(v string, title bool) string {
	runes := []rune(v)
	builder := strings.Builder{}
	upper := false

	for i, r := range runes {
		if r == '_' {
			upper = true
			continue
		}

		if upper || (title && i == 0) {
			upper = false
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
