
package pkg

import (
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"text/template"
	"strings"
	"io"
)

func ProcessDay(apiToken string, workspaceId int, dateS string, w io.Writer) (err error) {

	date, err := time.Parse("2006-01-02 MST", fmt.Sprintf("%s JST", dateS))
	if err != nil { return err }

	w.Write([]byte(fmt.Sprintln(date.Format(time.ANSIC))))
	w.Write([]byte(fmt.Sprintln(date.Format(time.RFC3339))))
	
	nextDay := nextDay(date)
	w.Write([]byte(fmt.Sprintln(nextDay.Format(time.ANSIC))))

	type Content struct {
		Date string
		DurationSum string
		TimeEntries []string
	}
	content := Content{Date: date.Format("2006-01-02")}

	session := toggl.OpenSession(apiToken)

	projects, err := session.GetProjects(workspaceId)
	if err != nil { return err }
	projectMap := makeProjectMap(projects)

	timeEntries, err := session.GetTimeEntries(date, nextDay)
	if err != nil { return err }

	durationSum := int64(0)
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst())
		stop := te.Stop.In(jst())
		duration := time.Duration(te.Duration) * time.Second
		durationSum = durationSum + te.Duration
		s := fmt.Sprintf("- [%s] %02d:%02d - %02d:%02d %v %v", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
		content.TimeEntries = append(content.TimeEntries, s)
	})
	content.DurationSum = fmtDurationHHMM(time.Duration(durationSum) * time.Second)

	err = dayTemplate().Execute(w, content)
	if err != nil { return err }

	return nil
}

func nextDay(date time.Time) (time.Time) {
	nextDay := date.AddDate(0, 0, 1)
	return nextDay.Truncate( time.Hour ).Add( - time.Duration(nextDay.Hour()) * time.Hour )
}

func jst() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

func makeProjectMap(projects []toggl.Project) map[int]string {
	projectMap := map[int]string{}
	From(projects).ToMapByT(&projectMap,
        func(p toggl.Project) int { return p.ID },
        func(p toggl.Project) string { return p.Name },
	)
	return projectMap
}

func fmtDurationHHMM(d time.Duration) string {
    d = d.Round(time.Minute)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    return fmt.Sprintf("%02d:%02d", h, m)
}

func dayTemplate() *template.Template {
	const letter = `
@@@
# {{.Date}}の実績
total {{.DurationSum}}
{{range $var := .TimeEntries -}}
  {{$var}}
{{end -}}
@@@
・疑問点や気にかかっていること

・明日の作業予定
`

	return template.Must(template.New("letter").Parse(strings.Replace(letter, "@", "`", -1)))
}