package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"title": strings.Title,
	"first": func(s string) string {
		return strings.ToLower(string(s[0]))
	},
	"unexport": func(s string) string {
		return strings.ToLower(string(s[0])) + string(s[1:])
	},
}

const tmpl = `// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}

package optional

// {{ .Name }} is an optional {{ .Type }}
type {{ .Name }} struct {
	{{ .Type | unexport }} {{ .Type }}
	present bool
}

// Empty{{ .Name | title }} returns an empty optional.{{ .Name }}
func Empty{{ .Name | title }}() {{ .Name }} {
	return {{ .Name }}{}
}

// Of{{ .Type | title }} creates an optional.{{ .Name }} from a {{ .Type }}
func Of{{ .Type | title }}({{ .Type | first }} {{ .Type }}) {{ .Name }} {
	return {{ .Name }}{ {{ .Type | unexport }}: {{ .Type | first }}, present: true}
}

// Set sets the {{ .Type }} value
func (o *{{ .Name }}) Set({{ .Type | first }} {{ .Type }}) {
	o.{{ .Type | unexport }} = {{ .Type | first }}
	o.present = true
}

// {{ .Type | title }} returns the {{ .Type }} value
func (o *{{ .Name }}) {{ .Type | title }}() {{ .Type }} {
	return o.{{ .Type | unexport }}
}

// Present returns whether or not the value is present
func (o *{{ .Name }}) Present() bool {
	return o.present
}

// OrElse returns the {{ .Type }} value or a default value if the value is not present
func (o *{{ .Name }}) OrElse({{ .Type | first }} {{ .Type }}) {{ .Type }} {
	if o.present {
		return o.{{ .Type | unexport}}
	}
	return {{ .Type | first }}
}`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("type is required")
	}

	typ := args[0]
	var name string
	var fileName string

	exported := strings.Title(typ) == typ

	if len(args) == 2 {
		name = args[1]
		fileName = strings.ToLower(name + ".go")
	} else if len(args) > 2 {
		name = args[1]
		fileName = args[2]
	} else {
		if exported {
			name = "Optional" + strings.Title(typ)
		} else {
			name = "optional" + strings.Title(typ)
		}
		fileName = strings.ToLower(fmt.Sprintf("optional_%s.go", typ))
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(tmpl))

	data := struct {
		Timestamp time.Time
		Type      string
		Name      string
	}{
		time.Now().UTC(),
		typ,
		name,
	}

	var buf bytes.Buffer

	err := t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(p)
	if err != nil {
		log.Fatal(err)
	}
}
