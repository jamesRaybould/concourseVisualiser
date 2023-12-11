package models

import (
	"bytes"
	"text/template"
)

type Run struct {
	Args []string
	Dir  string
	Path string
}

var runTmpl = `
		\t\tpath: {{ .Path }}
		{{ if .Dir }}\t\tdir: {{.Dir}}{{end}}
		{{ if .Args -}}
		\t\targs:	
		{{ range .Args -}} 
		\t\t\t{{- .}} 
		{{ end }}
		{{- end -}}`

func (run Run) String() string {
	var buf bytes.Buffer
	parsedTemplate, _ := template.New("run").Parse(runTmpl)
	parsedTemplate.Execute(&buf, run)
	return buf.String()
}
