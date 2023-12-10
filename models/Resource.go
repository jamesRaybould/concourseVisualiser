package models

import (
	"bytes"
	"strings"
	"text/template"
)

type Resource struct {
	Name        string
	Type        string
	Icon        string
	Check_Every string
	Source      Source
}

func (resource Resource) EscapedName() string {
	return strings.ReplaceAll(resource.Name, "-", "_")
}

var resourceTmpl = `
	object "{{.Name}}" as {{.EscapedName}} #94CEDB{
		name: {{.Name}}
		type: {{.Type}}
		{{ if .Check_Every -}}check_every: {{.Check_Every}}{{- end}}
		source:
			{{- .Source }}
	}
`

func (resource Resource) String() (string){
	var buf bytes.Buffer
	parsedTemplate, _ := template.New("run").Parse(resourceTmpl)
	parsedTemplate.Execute(&buf, resource)
	return buf.String()
}
