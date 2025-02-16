package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type Spec struct {
	Enums []Enum `json:"enums"`
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

	generateEnums(spec.Enums)
	exec.Command("gofmt", "-s", "-w", "./pkg").Run()
}

func generateEnums(enums []Enum) {
	f, err := os.OpenFile("./pkg/types/enums.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o660)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("package types\n\n")

	for _, e := range enums {
		decl := fmt.Sprintf("type %s string\n", e.Name)
		f.WriteString(decl)

		f.WriteString("const (\n")

		for _, v := range e.Values {
			name := makeEnumValueName(e.Name, v)
			assig := fmt.Sprintf("%s %s = \"%s\"\n", name, e.Name, v)
			f.WriteString(assig)
		}

		f.WriteString("\n)\n")
	}
}

func makeEnumValueName(enumName string, v string) string {
	runes := []rune(v)

	builder := &strings.Builder{}
	builder.Grow(len(enumName) + len(runes))

	builder.WriteString(enumName)
	builder.WriteString(strings.ToUpper(string(runes[0])))

	upper := false

	for _, r := range runes[1:] {
		if r == '_' {
			upper = true
			continue
		}

		if upper {
			upper = false
			builder.WriteString(strings.ToUpper(string(r)))
		} else {
			builder.WriteString(string(r))
		}
	}

	return builder.String()
}
