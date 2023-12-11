package models

import (
	"bytes"
	"text/template"
)

type Source struct {
	Uri          string
	Private_key  string
	Repository   string
	Username     string
	Password     string
	Access_token string
	Branch       string
	States       []string
}

var sourceTmpl = `
			{{ if .Uri }}\tURI: [[{{.Uri}}]]{{ end }}
			{{ if .Private_key }}\tprivate_key: {{.Private_key}}{{ end }}
			{{ if .Repository }}\trepository: [[{{.Repository}}]]{{ end }}
			{{ if .Access_token }}\taccess_token: {{.Access_token}}{{ end }}
			{{ if .Username }}\tusername: {{.Username}}{{ end }}
			{{ if .Password }}\tpassword: {{.Password}}{{ end }}
			{{ if .Branch }}\tbranch: {{.Branch}}{{ end }}
			{{ if .States -}}
			\tstates:	
			{{ range .States -}} 
			\t\t{{- .}} 
			{{ end }}
			{{- end -}}`

func (source Source) String() string {
	var buf bytes.Buffer
	parsedTemplate, _ := template.New("run").Parse(sourceTmpl)
	parsedTemplate.Execute(&buf, source)
	return buf.String()
}
