package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"
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
	SubTypes    []string    `json:"subtypes"`
	SubTypeOf   []string    `json:"subtype_of"`
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

	// TODO: also generate enums for sum types
	generateEnums(updateType, spec.Enums)
	generateHandlers(updateType)

	generateTypes(parser, spec.Types)
	generateRequests(parser, spec.Methods)
	generateMethods(parser, spec.Methods)

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
		fieldName := camelCase(update.Name, false)
		f.WriteString(fmt.Sprintf("%s routerHandlers[*goram.%s]\n", fieldName, update.Types[0]))
	}

	f.WriteString("}\n")

	for _, update := range updateType.Fields[1:] {
		updatePascalName := camelCase(update.Name, true)
		structFieldName := camelCase(update.Name, false)
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
		handlersFieldName := camelCase(u.Name, false)
		updatePascalName := camelCase(u.Name, true)

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
		pascalName := camelCase(u.Name, true)
		handlersFieldName := camelCase(u.Name, false)

		fmt.Fprintf(f, `
			func (r *Router) call%sHandlers(ctx context.Context, bot *goram.Bot, update *goram.%s, data Data) (bool, error) {
				queue := make([]*Router, 0, len(r.children) + 1)			
				queue = append(queue, r)

			queueLoop:
				for len(queue) > 0 {
					current := queue[0]
					queue = queue[1:]

					for _, filter := range current.handlers.%s.filters {
						ok, err := filter(ctx, bot, update, data)

						if err != nil {
							return ok, err
						}

						if !ok {
							continue queueLoop
						}
					}

					found, err := callHandlers(ctx, bot, current.handlers.%s.handlers, update, data)

					if err != nil || found {
						return found, err
					}

					if len(current.children) > 0 {
						queue = append(queue, current.children...)
					}
				}

				return false, nil
			}
		`,
			pascalName,
			u.Types[0],
			handlersFieldName,
			handlersFieldName,
		)
	}

	f.WriteString(
		"func (r *Router) feedUpdate(ctx context.Context, bot *goram.Bot, update *goram.Update, data Data) (bool, error) {\n",
	)

	for _, u := range updateType.Fields[1:] {
		fieldName := camelCase(u.Name, true)

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
			name := e.Name + camelCase(v, true)
			assig := fmt.Sprintf("%s %s = \"%s\"\n", name, e.Name, v)
			f.WriteString(assig)
		}

		f.WriteString("\n)\n")
	}

	f.WriteString("type UpdateType string\n\n")
	f.WriteString("const (\n")

	for _, field := range updateType.Fields[1:] { // skip update_id
		name := "Update" + camelCase(field.Name, true)
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
		pascalName := camelCase(m.Name, true)
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

		fmt.Fprintf(f, "// %s\n", m.Href)
		fmt.Fprintf(f, "func (b *Bot) %s%s %s {\n", pascalName, args, returnType)
		fmt.Fprintf(
			f,
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
		)

		if len(m.Fields) == 0 || strings.HasPrefix(pascalName, "Get") {
			continue
		}

		fmt.Fprintf(
			f,
			`
			// Does the same as Bot.%s, but parses response body only in case of an error. 
			// Therefore works faster if you dont need the response value.
			func (b *Bot) %sVoid%s error {
				return makeVoidRequest(ctx, b.options.Client, b.baseUrl, "%s", b.options.FloodHandler, %s)
			}
			`,
			pascalName,
			pascalName,
			args,
			m.Name,
			data,
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

		pascalName := camelCase(m.Type.Name, true)
		structName := pascalName
		suffix := ""

		if !strings.HasSuffix(structName, "Request") {
			structName += "Request"
			suffix = "Request"
		}

		fmt.Fprintf(f, "// see Bot.%s(ctx, &%s{})\n", pascalName, structName)
		generateTypeStruct(f, parser.ParseType(&m.Type), suffix, false, false)
		f.WriteString("\n")
		generateRequestWriteMultipart(f, parser, &m, structName)
	}

	f.Close()
}

func generateRequestWriteMultipart(
	w io.Writer,
	parser *Parser,
	m *Method,
	structName string,
) {
	fmt.Fprintf(w, "func (r *%s) writeMultipart(w *multipart.Writer) {", structName)

	for _, field := range m.Fields {
		parsedTypeField := parser.ParseTypeField(&field)

		if parsedTypeField.ParsedSpecType.GoType == "InputFile" {
			fmt.Fprintf(w, `if r.%s.FileID != "" {
					w.WriteField("%s", r.%s.FileID)
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
			)
		} else if parsedTypeField.ParsedSpecType.GoType == "InputMedia" &&
			parsedTypeField.ParsedSpecType.ParsedType != ParsedTypeArray {

			fmt.Fprintf(w, `{
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
			)
		} else if parsedTypeField.ParsedSpecType.GoType == "InputMedia" &&
			parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeArray &&
			parsedTypeField.ParsedSpecType.Levels == 1 {

			fmt.Fprintf(w, `for i := 0; i < len(r.%s); i++ {
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
			)
		} else if parsedTypeField.ParsedSpecType.GoType == "ChatID" {
			fmt.Fprintf(w,
				"w.WriteField(\"%s\", r.%s.String())\n",
				field.Name,
				parsedTypeField.GoName,
			)
		} else if parsedTypeField.ParsedSpecType.GoType == "string" &&
			parsedTypeField.ParsedSpecType.ParsedType == ParsedTypePrimitive {

			if !parsedTypeField.Field.Required {
				fmt.Fprintf(w, "if r.%s != \"\" {", parsedTypeField.GoName)
			}

			fmt.Fprintf(w, "w.WriteField(\"%s\", r.%s)\n", field.Name, parsedTypeField.GoName)

			if !parsedTypeField.Field.Required {
				w.Write([]byte("}\n"))
			}
		} else if parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeEnum {
			fmt.Fprintf(w, "w.WriteField(\"%s\", string(r.%s))\n", field.Name, parsedTypeField.GoName)
		} else {
			checkForNil := !parsedTypeField.Field.Required &&
				(parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeStruct ||
					parsedTypeField.ParsedSpecType.Levels > 0 ||
					parsedTypeField.ParsedSpecType.ParsedType == ParsedTypeInterface)

			if checkForNil {
				fmt.Fprintf(w, "if r.%s != nil ", parsedTypeField.GoName)
			}

			fmt.Fprintf(w, `{
					b, _ := json.Marshal(r.%s)
					fw, _ := w.CreateFormField("%s")
					fw.Write(b)
				}
				`,
				parsedTypeField.GoName,
				field.Name,
			)
		}
	}

	w.Write([]byte("}\n"))
}

var builtinTypes = []string{
	"InputMedia",
	"InputFile",
	"InaccessibleMessage",
	"MaybeInaccessibleMessage",
}

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

		parseResult := parser.ParsedTypes[t.Name]

		if len(t.SubTypeOf) > 0 {
			if len(t.SubTypeOf) > 1 {
				panic(t)
			}

			parentParseResult := parser.ParsedTypes[t.SubTypeOf[0]]

			if checkSumType(parentParseResult.Type) {
				continue
			}
		}

		if checkSumType(&t) {
			generateSumTypeStruct(f, parser.ParsedTypes, parseResult)
			continue
		}

		generateTypeStruct(f, parseResult, "", true, true)

		const inputMediaPrefix = "InputMedia"
		const inputPaidMediaPrefix = "InputPaidMedia"

		if strings.HasPrefix(t.Name, inputMediaPrefix) ||
			(strings.HasPrefix(t.Name, inputPaidMediaPrefix) && len(t.Fields) > 0) {
			generateInputMediaMethods(f, &t)
		}
	}

	f.Close()
}

var ignoredSumTypes = []string{
	"InputMedia",
	"InputPaidMedia",
	"InlineQueryResult",
	"MaybeInaccessibleMessage",
	"InputMessageContent",
}

func checkSumType(t *Type) bool {
	return len(t.SubTypes) > 0 && !slices.Contains(ignoredSumTypes, t.Name)
}

func generateSumTypeStruct(
	w io.StringWriter,
	parsedTypes map[string]TypeParseResult,
	t TypeParseResult,
) {
	fields := []*ParsedTypeField{}

	for _, s := range t.Type.SubTypes {
		parsedType := parsedTypes[s]

	parsedFieldsLoop:
		for _, parsedField := range parsedType.Fields {
			for _, field := range fields {
				if field.Field.Name == parsedField.Field.Name {
					continue parsedFieldsLoop
				}
			}

			fields = append(fields, parsedField)
		}
	}

	w.WriteString("\n")

	for _, d := range t.Type.Description {
		w.WriteString("//\n//")
		w.WriteString(d)
		w.WriteString("\n")
	}

	w.WriteString("//\n// ")
	w.WriteString(t.Type.Href)
	w.WriteString("\n")

	w.WriteString("type ")
	w.WriteString(t.Type.Name)
	w.WriteString(" struct {")

	for i, field := range fields {
		w.WriteString("\n")
		w.WriteString(field.StructField(true, i > 0))
	}

	w.WriteString("\n}\n\n")
}

func generateInputMediaMethods(w io.StringWriter, t *Type) {
	w.WriteString(fmt.Sprintf(`
		func (i *%s) setMedia(fileID string) {
			i.Media.FileID = fileID
		}
	`, t.Name))

	w.WriteString(fmt.Sprintf(`
		func (i *%s) getMedia() InputFile {
			return i.Media
		}
	`, t.Name))
}

func generateTypeStruct(
	w io.StringWriter,
	t TypeParseResult,
	suffix string,
	doc bool,
	tagJSON bool,
) {
	if doc {
		for _, d := range t.Type.Description {
			comment := fmt.Sprintf("// %s\n//\n", d)
			w.WriteString(comment)
		}

		w.WriteString(fmt.Sprintf("// %s\n", t.Type.Href))
	}

	camelName := camelCase(t.Type.Name, true)

	if suffix != "" {
		camelName += suffix
	}

	if len(t.Fields) == 0 {
		decl := fmt.Sprintf("type %s interface{}\n", camelName)
		w.WriteString(decl)
		return
	}

	decl := fmt.Sprintf("type %s struct {\n", camelName)
	w.WriteString(decl)

	for i, field := range t.Fields {
		w.WriteString(field.StructField(tagJSON, true))

		if i < len(t.Fields)-1 {
			w.WriteString("\n")
		}
	}

	w.WriteString("\n}\n")
}

type TypeParseResult struct {
	Type   *Type
	Fields []*ParsedTypeField
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

	switch p.ParsedType {
	case ParsedTypeStruct:
		builder.WriteRune('*')
	case ParsedTypeArray:
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

func (p *ParsedTypeField) StructField(tagJSON bool, doc bool) string {
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

	if doc {
		builder.WriteString(" // ")
		builder.WriteString(p.Field.Description)
	}

	return builder.String()
}

type Parser struct {
	EnumNames      []string
	InterfaceNames []string
	ParsedTypes    map[string]TypeParseResult
}

func (p *Parser) ParseTypeField(t *TypeField) *ParsedTypeField {
	ptf := &ParsedTypeField{
		Field:  t,
		GoName: camelCase(t.Name, true),
	}

	if t.Name == "reply_markup" && len(t.Types) == 4 {
		ptf.ParsedSpecType.GoType = "Markup"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else if len(t.Types) == 4 && t.Name == "media" {
		ptf.ParsedSpecType.GoType = "InputMedia"
		ptf.ParsedSpecType.ParsedType = ParsedTypeArray
		ptf.ParsedSpecType.Levels = 1
	} else if len(t.Types) == 2 && (t.Name == "id" || t.Name == "chat_id" || t.Name == "from_chat_id") {
		ptf.ParsedSpecType.GoType = "ChatID"
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
	} else if len(t.Types) > 0 && t.Types[0] == "MaybeInaccessibleMessage" {
		ptf.ParsedSpecType.GoType = "Message"
		ptf.ParsedSpecType.ParsedType = ParsedTypeStruct
	} else {
		ptf.ParsedSpecType = p.ParseSpecTypes(t.Types)
	}

	return ptf
}

func NewParser(spec *Spec) *Parser {
	p := &Parser{
		ParsedTypes: make(map[string]TypeParseResult, len(spec.Types)),
	}

	for _, e := range spec.Enums {
		p.EnumNames = append(p.EnumNames, e.Name)
	}

	for _, t := range spec.Types {
		if len(t.Fields) == 0 && !checkSumType(&t) {
			p.InterfaceNames = append(p.InterfaceNames, t.Name)
		}
	}

	for _, t := range spec.Types {
		p.ParsedTypes[t.Name] = p.ParseType(&t)
	}

	return p
}

func (p *Parser) ParseType(t *Type) TypeParseResult {
	fields := make([]*ParsedTypeField, len(t.Fields))

	for i, field := range t.Fields {
		fields[i] = p.ParseTypeField(&field)
	}

	return TypeParseResult{
		Type:   t,
		Fields: fields,
	}
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

func (g *Parser) parseSpecType(p *ParsedSpecType, fieldType string) {
	switch fieldType {
	case "Integer":
		p.GoType = "int"
		return
	case "String":
		p.GoType = "string"
		return
	case "Boolean":
		p.GoType = "bool"
		return
	case "Float":
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

	p.GoType = camelCase(fieldType, true)
}

func camelCase(v string, title bool) string {
	if v == "id" {
		return "ID"
	}

	runes := []rune(v)
	builder := strings.Builder{}
	upper := false

	for i, r := range runes {
		if i == len(runes)-3 && r == '_' && runes[i+1] == 'i' && runes[i+2] == 'd' {
			builder.WriteString("ID")
			break
		}

		if r == '_' {
			upper = true
			continue
		}

		if upper || (title && i == 0 && r > 'Z') {
			upper = false
			builder.WriteRune(r - 32)
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
