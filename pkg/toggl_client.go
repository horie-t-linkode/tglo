
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

func Process(apiToken string, workspaceId int, from time.Time, till time.Time, w io.Writer) (err error) {

	w.Write([]byte(fmt.Sprintln(from.Format(time.ANSIC))))
	w.Write([]byte(fmt.Sprintln(from.Format(time.RFC3339))))
	w.Write([]byte(fmt.Sprintln(till.Format(time.ANSIC))))
	w.Write([]byte(fmt.Sprintln(till.Format(time.RFC3339))))

	type Content struct {
		Date string
		DurationSum string
		DurationTagSum []string
		TimeEntries []string
	}
	content := Content{Date: from.Format("2006-01-02")}

	session := toggl.OpenSession(apiToken)

	account, err := session.GetAccount()
	if err != nil { return err }

	tags := account.Data.Tags
	tagSumMap := makeTagDurationSumMap(tags)
	From(tags).ForEachT(func(tag toggl.Tag) {
		w.Write([]byte(fmt.Sprintln(tag.Name)))
	})


	projects, err := session.GetProjects(workspaceId)
	if err != nil { return err }
	projectMap := makeProjectMap(projects)

	timeEntries, err := session.GetTimeEntries(from, till)
	if err != nil { return err }

	durationSum := int64(0)
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst())
		stop := te.Stop.In(jst())
		duration := time.Duration(te.Duration) * time.Second
		durationSum = durationSum + te.Duration
		From(te.Tags).ForEachT(func(tagname string) {
			tagSumMap[tagname] = tagSumMap[tagname] + te.Duration
		})
		s := fmt.Sprintf("- [%s] %02d:%02d - %02d:%02d %v %v", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
		content.TimeEntries = append(content.TimeEntries, s)
	})
	content.DurationSum = fmtDurationHHMM(time.Duration(durationSum) * time.Second)
	From(tags).
		WhereT(func(tag toggl.Tag) bool {
			return tagSumMap[tag.Name] > 0
		}).
		SelectT(func(tag toggl.Tag) string {
			return fmt.Sprintf("- [%s] %6.2f%% %s", fmtDurationHHMM(time.Duration(tagSumMap[tag.Name]) * time.Second), float64(tagSumMap[tag.Name]) * float64(100) / float64(durationSum), tag.Name)	
		}).
		ToSlice(&content.DurationTagSum)

	err = dayTemplate().Execute(w, content)
	if err != nil { return err }

	return nil
}

func Date(dateS string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02 MST", fmt.Sprintf("%s JST", dateS))
	if err != nil { return date, err }

	return date, nil
}

func Today() (date time.Time) {
	date = time.Now()
	date = date.Truncate( time.Hour ).Add( - time.Duration(date.Hour()) * time.Hour )
	return date
}

func NextDay(date time.Time) (time.Time) {
	nextDay := date.AddDate(0, 0, 1)
	return nextDay.Truncate( time.Hour ).Add( - time.Duration(nextDay.Hour()) * time.Hour )
}

func jst() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

func makeProjectMap(projects []toggl.Project) (map[int]string) {
	r := map[int]string{}
	From(projects).ToMapByT(&r,
        func(p toggl.Project) int { return p.ID },
        func(p toggl.Project) string { return p.Name },
	)
	return r
}

func makeTagDurationSumMap(tags []toggl.Tag) (map[string]int64) {
	r := map[string]int64{}
	From(tags).ToMapByT(&r,
		func(p toggl.Tag) string { return p.Name },
		func(p toggl.Tag) int64 { return int64(0)},
	)
	return r
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

tag
{{range $var := .DurationTagSum -}}
  {{$var}}
{{end -}}
@@@
・疑問点や気にかかっていること

・明日の作業予定
`

	return template.Must(template.New("letter").Parse(strings.Replace(letter, "@", "`", -1)))
}