package tglo_core

import (
	"text/template"
	"strings"
)


func dayTemplate() *template.Template {
	const letter = `
@@@
# {{.From}}の実績
total {{.DurationTotal}}

{{with .TimeEntryDetails -}}
{{- range . -}}
- [{{.Duration}}] {{.From}} - {{.Till}} {{.ProjectName}} {{.Description}}
{{end -}}
{{end}}

tags
{{with .TagSummaries -}}
{{- range . -}}
- [{{.Duration}}] {{.Ratio}} {{.Name}}
{{end -}}
{{end -}}

@@@
・疑問点や気にかかっていること

・明日の作業予定
`

	return template.Must(template.New("letter").Parse(strings.Replace(letter, "@", "`", -1)))
}

func weekTemplate() *template.Template {
	const letter = `
## Report[{{.From}} 〜 {{.Till}}]

- total {{.DurationTotal}}

{{with .ProjectSummaries -}}
{{- range . -}}
### [{{.Duration}}] {{.Name}}
{{if eq .ShowDetail true}}
{{with .Items -}}
{{- range . -}}
- [{{.Duration}}] {{.Title}}
{{end -}}
{{end -}}
{{end -}}
{{end -}}
{{end}}

### tags
{{with .TagSummaries -}}
{{range . -}}
- [{{.Duration}}] {{.Ratio}} {{.Name}}
{{end -}}
{{end -}}
`

	return template.Must(template.New("letter").Parse(letter))
}