package template

import (
	"text/template"
	"strings"
	"io"
)


func DayTemplate() *template.Template {
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
{{end}}

projects
{{with .ProjectSummaries -}}
{{- range . -}}
- [{{.Duration}}] {{.Ratio}} {{.Name}}
{{end -}}
{{end}}
@@@

・疑問点や気にかかっていること
{{with .Comments -}}
{{- range . -}}
{{.Description}}
{{end -}}
{{end}}

・明日の作業予定
{{with .Plans -}}
{{- range . -}}
{{.Description}}
{{end -}}
{{end -}}
`

	return template.Must(template.New("letter").Parse(strings.Replace(letter, "@", "`", -1)))
}

func WeekTemplate() *template.Template {
	const letter = `
## Report[{{.From}} 〜 {{.Till}}]

- total {{.DurationTotal}}

{{with .ProjectSummaries -}}
{{- range . -}}
### [{{.Duration}}] {{.Name}}
{{with .Items -}}
{{- range . -}}
- [{{.Duration}}] {{.Title}}
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

func WeekTemplateSupressDetail() *template.Template {
	const letter = `
@@@
## Report[{{.From}} 〜 {{.Till}}]

- total {{.DurationTotal}}

{{with .ProjectSummaries -}}
{{- range . -}}
### [{{.Duration}}] {{.Name}}
{{end -}}
{{end}}

### tags
{{with .TagSummaries -}}
{{range . -}}
- [{{.Duration}}] {{.Ratio}} {{.Name}}
{{end -}}
{{end -}}
@@@
`

	return template.Must(template.New("letter").Parse(strings.Replace(letter, "@", "`", -1)))
}

func TemplateExecute(template *template.Template, w io.Writer, content *OutputContent) (err error) {
	return template.Execute(w, content)
}
