package models

import (
	"bytes"
	"text/template"
)

type Config struct {
	Platform string
	Caches   []Caches
	Inputs   []Inputs
	Outputs  []Outputs
	Params   map[string]string
	Run      Run
}

var configTmpl = `
		\tplatform: {{ .Platform }}
		{{ if .Caches -}}
		\tcaches:
		{{ range .Caches }} 
			{{- .}}
		{{ end }}
		{{- end -}}
		{{ if .Inputs -}}
		\tinputs:
		{{ range .Inputs }} 
			{{- .}}
		{{ end }}
		{{- end -}}
		{{ if .Outputs -}}
		\toutputs:	
		{{ range .Outputs }} 
			{{- .}} 
		{{ end }}
		{{- end -}}
		\trun:
		{{- .Run }}
		{{ if .Params -}}
		\tparams:
		{{ range $index, $element := .Params }} 
		\t\t{{- $index }}: {{.}}
		{{ end }}
		{{- end -}}`

func (config Config) String() string {
	var buf bytes.Buffer

	parsedTemplate, _ := template.New("task").Parse(configTmpl)
	parsedTemplate.Execute(&buf, config)
	return buf.String()
}
