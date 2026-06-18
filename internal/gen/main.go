package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"slices"
	"strings"
	"text/template"
)

var (
	sumTypes = []string{
		"InputMedia",
		"InputPaidMedia",
		"InlineQueryResult",
		"MaybeInaccessibleMessage",
		"InputMessageContent",
	}

	builtinTypes = []string{
		"InputMedia",
		"InputFile",
		"InaccessibleMessage",
		"MaybeInaccessibleMessage",
	}
)

const genFilePerm = 0o660

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

//go:embed templates/*.tmpl
var templateFS embed.FS

var tmpls *template.Template

func main() {
	funcMap := template.FuncMap{
		"pascal": func(s string) string { return snakeToCamel(s, true) },
		"camel":  func(s string) string { return snakeToCamel(s, false) },
		"trimOptional": func(s string) string {
			return strings.TrimPrefix(s, "Optional. ")
		},
	}

	baseTmpl := template.New("").Funcs(funcMap)

	var err error

	tmpls, err = baseTmpl.ParseFS(templateFS, "templates/*.tmpl")

	if err != nil {
		panic(err)
	}

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

	// TODO: generate enums for sum types
	updateType := spec.Types[0]
	generateEnums(updateType, spec.Enums)
	generateHandlers(updateType)

	generateTypes(parser, spec.Types)
	generateRequests(parser, spec.Methods)
	generateMethods(parser, spec.Methods)
}

func generateHandlers(updateType Type) {
	var templateData struct {
		Fields []TypeField
	}

	if len(updateType.Fields) > 1 {
		templateData.Fields = updateType.Fields[1:]
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "handlers.tmpl", templateData); err != nil {
		panic(err)
	}

	formattedCode, err := format.Source(buf.Bytes())

	if err != nil {
		panic("gofmt error: " + err.Error())
	}

	err = os.WriteFile("./handlers/handlers.go", formattedCode, genFilePerm)

	if err != nil {
		panic(err)
	}
}

func generateEnums(updateType Type, enums []Enum) {
	type singleEnum struct {
		Name   string
		Values []string
	}

	preparedEnums := []singleEnum{}

	for _, e := range enums {
		preparedEnums = append(preparedEnums, singleEnum{
			Name:   e.Name,
			Values: e.Values,
		})
	}

	updateFields := []TypeField{}

	if len(updateType.Fields) > 1 {
		updateFields = updateType.Fields[1:]
	}

	data := struct {
		Enums        []singleEnum
		UpdateFields []TypeField
	}{
		Enums:        preparedEnums,
		UpdateFields: updateFields,
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "enums.tmpl", data); err != nil {
		panic(err)
	}

	formattedCode, err := format.Source(buf.Bytes())

	if err != nil {
		panic("gofmt error: " + err.Error())
	}

	if err := os.WriteFile("./enums.go", formattedCode, genFilePerm); err != nil {
		panic(err)
	}
}

func generateMethods(parser *Parser, methods []Method) {
	type methodTemplateData struct {
		Name        string
		PascalName  string
		Href        string
		Description []string
		Args        string
		ReturnType  string
		TypeString  string
		Data        string
		GenVoid     bool
	}

	preparedMethods := []methodTemplateData{}

	for _, m := range methods {
		pascalName := snakeToCamel(m.Name, true)
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

		genVoid := len(m.Fields) > 0 && !strings.HasPrefix(pascalName, "Get")

		preparedMethods = append(preparedMethods, methodTemplateData{
			Name:        m.Name,
			PascalName:  pascalName,
			Href:        m.Href,
			Description: m.Description,
			Args:        args,
			ReturnType:  returnType,
			TypeString:  typeString,
			Data:        data,
			GenVoid:     genVoid,
		})
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "methods.tmpl", preparedMethods); err != nil {
		panic(err)
	}

	formattedCode, err := format.Source(buf.Bytes())

	if err != nil {
		panic("gofmt error: " + err.Error())
	}

	if err := os.WriteFile("./methods.go", formattedCode, genFilePerm); err != nil {
		panic(err)
	}
}

type (
	fieldData struct {
		Name        string
		GoName      string
		Case        string
		Required    bool
		CheckForNil bool
	}

	structMultipartData struct {
		StructName string
		Fields     []fieldData
	}

	typeStructData struct {
		Doc         bool
		Description []string
		Href        string
		CamelName   string
		IsInterface bool
		Fields      []string
	}
)

func generateRequests(parser *Parser, methods []Method) {
	type requestTemplateData struct {
		PascalName    string
		StructName    string
		StructData    typeStructData
		MultipartData structMultipartData
	}

	preparedRequests := []requestTemplateData{}

	for _, m := range methods {
		if len(m.Fields) == 0 {
			continue
		}

		pascalName := snakeToCamel(m.Type.Name, true)
		structName := pascalName
		suffix := ""

		if !strings.HasSuffix(structName, "Request") {
			structName += "Request"
			suffix = "Request"
		}

		t := parser.ParseType(&m.Type)
		camelName := snakeToCamel(t.Type.Name, true)

		if suffix != "" {
			camelName += suffix
		}

		structFields := make([]string, len(t.Fields))

		for _, field := range t.Fields {
			structFields = append(structFields, field.StructField(false, true))
		}

		sData := typeStructData{
			Doc:         false,
			Description: t.Type.Description,
			Href:        t.Type.Href,
			CamelName:   camelName,
			IsInterface: len(t.Fields) == 0,
			Fields:      structFields,
		}

		multipartFields := make([]fieldData, 0, len(m.Fields))

		for _, field := range m.Fields {
			parsedTypeField := parser.ParseTypeField(&field)
			spec := parsedTypeField.ParsedSpecType

			currentCase := ""
			checkForNil := false

			if spec.GoType == "InputFile" {
				currentCase = "InputFile"
			} else if spec.GoType == "InputSticker" && spec.ParsedType == ParsedTypeArray {
				currentCase = "InputStickerArray"
			} else if spec.GoType == "InputSticker" {
				currentCase = "InputSticker"
			} else if spec.GoType == "InputMedia" && spec.ParsedType != ParsedTypeArray {
				currentCase = "InputMedia"
			} else if spec.GoType == "InputMedia" && spec.ParsedType == ParsedTypeArray && spec.Levels == 1 {
				currentCase = "InputMediaArray"
			} else if spec.GoType == "ChatID" {
				currentCase = "ChatID"
			} else if spec.GoType == "string" && spec.ParsedType == ParsedTypePrimitive {
				currentCase = "StringPrimitive"
			} else if spec.ParsedType == ParsedTypeEnum {
				currentCase = "Enum"
			} else {
				currentCase = "Default"
				checkForNil = !parsedTypeField.Field.Required &&
					(spec.ParsedType == ParsedTypeStruct ||
						spec.Levels > 0 ||
						spec.ParsedType == ParsedTypeInterface)
			}

			multipartFields = append(multipartFields, fieldData{
				Name:        field.Name,
				GoName:      parsedTypeField.GoName,
				Case:        currentCase,
				Required:    parsedTypeField.Field.Required,
				CheckForNil: checkForNil,
			})
		}

		mData := structMultipartData{
			StructName: structName,
			Fields:     multipartFields,
		}

		preparedRequests = append(preparedRequests, requestTemplateData{
			PascalName:    pascalName,
			StructName:    structName,
			StructData:    sData,
			MultipartData: mData,
		})
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "requests.tmpl", preparedRequests); err != nil {
		panic(err)
	}

	formattedCode, err := format.Source(buf.Bytes())

	if err != nil {
		panic("gofmt error: " + err.Error())
	}

	if err := os.WriteFile("./requests.go", formattedCode, genFilePerm); err != nil {
		panic(err)
	}
}

func generateTypes(parser *Parser, types []Type) {
	type (
		typeStructData struct {
			Doc         bool
			Description []string
			Href        string
			CamelName   string
			IsInterface bool
			Fields      []string
		}

		sumTypeStructData struct {
			Description []string
			Href        string
			Name        string
			Fields      []string
		}

		typeTemplateItem struct {
			IsSumType       bool
			GenMediaMethods bool
			TypeRaw         *Type
			StructData      typeStructData
			SumTypeData     sumTypeStructData
		}
	)

	preparedTypes := []typeTemplateItem{}

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

		const (
			inputMediaPrefix     = "InputMedia"
			inputPaidMediaPrefix = "InputPaidMedia"
		)

		genMediaMethods := strings.HasPrefix(t.Name, inputMediaPrefix) ||
			(strings.HasPrefix(t.Name, inputPaidMediaPrefix) && len(t.Fields) > 0)

		if checkSumType(&t) {
			uniqueFields := []*ParsedTypeField{}

			for _, s := range t.SubTypes {
				parsedSub := parser.ParsedTypes[s]

			parsedFieldsLoop:
				for _, parsedField := range parsedSub.Fields {
					for _, field := range uniqueFields {
						if field.Field.Name == parsedField.Field.Name {
							continue parsedFieldsLoop
						}
					}

					uniqueFields = append(uniqueFields, parsedField)
				}
			}

			fieldLines := make([]string, 0, len(uniqueFields))

			for i, field := range uniqueFields {
				fieldLines = append(fieldLines, field.StructField(true, i > 0))
			}

			sumData := sumTypeStructData{
				Description: parseResult.Type.Description,
				Href:        parseResult.Type.Href,
				Name:        parseResult.Type.Name,
				Fields:      fieldLines,
			}

			preparedTypes = append(preparedTypes, typeTemplateItem{
				IsSumType:   true,
				SumTypeData: sumData,
			})

			continue
		}

		structFields := make([]string, 0, len(parseResult.Fields))

		for _, field := range parseResult.Fields {
			structFields = append(structFields, field.StructField(true, true))
		}

		sData := typeStructData{
			Doc:         true,
			Description: parseResult.Type.Description,
			Href:        parseResult.Type.Href,
			CamelName:   snakeToCamel(parseResult.Type.Name, true),
			IsInterface: len(parseResult.Fields) == 0,
			Fields:      structFields,
		}

		typeCopy := t

		preparedTypes = append(preparedTypes, typeTemplateItem{
			IsSumType:       false,
			GenMediaMethods: genMediaMethods,
			TypeRaw:         &typeCopy,
			StructData:      sData,
		})
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "types.tmpl", preparedTypes); err != nil {
		panic(err)
	}

	formattedCode, err := format.Source(buf.Bytes())

	if err != nil {
		panic("gofmt error: " + err.Error())
	}

	if err := os.WriteFile("./types.go", formattedCode, genFilePerm); err != nil {
		panic(err)
	}
}

func checkSumType(t *Type) bool {
	return len(t.SubTypes) > 0 && !slices.Contains(sumTypes, t.Name)
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
	data := struct {
		GoName      string
		TypeString  string
		TagJSON     bool
		FieldName   string
		Required    bool
		Doc         bool
		Description string
	}{
		GoName:      p.GoName,
		TypeString:  p.ParsedSpecType.TypeString(),
		TagJSON:     tagJSON,
		FieldName:   p.Field.Name,
		Required:    p.Field.Required,
		Doc:         doc,
		Description: p.Field.Description,
	}

	buf := bytes.Buffer{}

	if err := tmpls.ExecuteTemplate(&buf, "structField.tmpl", data); err != nil {
		panic(err)
	}

	return buf.String()
}

type Parser struct {
	EnumNames      []string
	InterfaceNames []string
	ParsedTypes    map[string]TypeParseResult
}

func (p *Parser) ParseTypeField(t *TypeField) *ParsedTypeField {
	ptf := &ParsedTypeField{
		Field:  t,
		GoName: snakeToCamel(t.Name, true),
	}

	if t.Name == "reply_markup" && len(t.Types) == 4 {
		ptf.ParsedSpecType.GoType = "Markup"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
	} else if len(t.Types) == 4 && t.Name == "media" {
		ptf.ParsedSpecType.GoType = "InputMedia"
		ptf.ParsedSpecType.ParsedType = ParsedTypeArray
		ptf.ParsedSpecType.Levels = 1
	} else if len(t.Types) == 1 && t.Types[0] == "String" && t.Name == "sticker" && !strings.HasPrefix(t.Description, "File identifier") {
		ptf.ParsedSpecType.GoType = "InputFile"
		ptf.ParsedSpecType.ParsedType = ParsedTypeInterface
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

	p.GoType = snakeToCamel(fieldType, true)
}

func snakeToCamel(s string, upper bool) string {
	if s == "id" {
		return "ID"
	}

	if s == "url" {
		return "URL"
	}

	builder := strings.Builder{}
	builder.Grow(len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '_' {
			const (
				idSuffix  = "_id"
				urlSuffix = "_url"
			)

			if checkSuffix(s, i, idSuffix) {
				builder.WriteString("ID")
				break
			}

			if checkSuffix(s, i, urlSuffix) {
				builder.WriteString("URL")
				break
			}

			upper = true
			continue
		}

		if upper {
			if char >= 'a' && char <= 'z' {
				char -= 32
			}

			upper = false
		}

		builder.WriteByte(char)
	}

	return builder.String()
}

func checkSuffix(s string, i int, suffix string) bool {
	return len(s)-i == len(suffix) && s[i:] == suffix
}
