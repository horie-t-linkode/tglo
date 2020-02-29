/*

The toggl command will display a user's Toggl account information.

Usage:
    toggl API_TOKEN

The API token can be retrieved from a user's account information page at toggl.com.

*/
package main

import (
	"os"
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"strconv"
	"text/template"
)

func main() {
	argc := len(os.Args)
	if ! (argc == 3 || argc == 4) {
		println("usage:", os.Args[0], "API_TOKEN WORKSPACE_ID")
		println("usage:", os.Args[0], "API_TOKEN WORKSPACE_ID yyyy-mm-dd")
		return
	}

	apiToken := os.Args[1]
	workspaceId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println("error:", err)
		return
	}

	var date time.Time
	if argc == 4 {
		date, err = time.Parse("2006-01-02 MST", fmt.Sprintf("%s JST", os.Args[3]))
		if err != nil {
			println("error:", err)
			return
		}
	} else {
		date = time.Now()
    	date = date.Truncate( time.Hour ).Add( - time.Duration(date.Hour()) * time.Hour )
	}

	println(date.Format(time.ANSIC))
	println(date.Format(time.RFC3339))

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	
	nextDate := date.AddDate(0, 0, 1)
	nextDate = nextDate.Truncate( time.Hour ).Add( - time.Duration(nextDate.Hour()) * time.Hour )
	println(nextDate.Format(time.ANSIC))

	const letter = `
# {{.Date}}の実績
total {{.DurationSum}}
{{range $var := .Descriptions -}}
  {{$var}}
{{end -}}
`
	tmpl := template.Must(template.New("letter").Parse(letter))

	type Content struct {
		Date string
		DurationSum string
		Descriptions []string
	}
	content := Content{Date: date.Format("2006-01-02")}

	session := toggl.OpenSession(apiToken)

	projects, err := session.GetProjects(workspaceId)
	if err != nil {
		println("error:", err)
		return
	}
	projectMap := makeProjectMap(projects)

	timeEntries, err := session.GetTimeEntries(date, nextDate)
	if err != nil {
		println("error:", err)
		return
	}

	durationSum := int64(0)
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst)
		stop := te.Stop.In(jst)
		duration := time.Duration(te.Duration) * time.Second
		durationSum = durationSum + te.Duration
		s := fmt.Sprintf("- [%s] %02d:%02d - %02d:%02d %v %v", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
		content.Descriptions = append(content.Descriptions, s)
	})
	content.DurationSum = fmtDurationHHMM(time.Duration(durationSum) * time.Second)

	err = tmpl.Execute(os.Stdout, content)
	if err != nil {
		println("executing template:", err)
		return
	}
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