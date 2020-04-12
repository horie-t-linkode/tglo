
package pkg

import (
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"github.com/snabb/isoweek"
	"text/template"
	"io"
)

type TogglClient struct {
	ApiToken string
	WorkSpaceId int
	Verbose bool
}

func (me *TogglClient) ProcessDay(from time.Time, till time.Time, w io.Writer) (err error) {
	return me.process(from, till, w, dayTemplate(), true)
}

func (me *TogglClient) ProcessWeek(from time.Time, till time.Time, w io.Writer, showDetail bool) (err error) {
	return me.process(from, till, w, weekTemplate(), showDetail)
}

func (me *TogglClient) process(from time.Time, till time.Time, w io.Writer, template *template.Template, showDetail bool) (err error) {

	if me.Verbose {
		toggl.EnableLog()
		w.Write([]byte(fmt.Sprintln(from.Format(time.ANSIC))))
		w.Write([]byte(fmt.Sprintln(from.Format(time.RFC3339))))
		w.Write([]byte(fmt.Sprintln(till.Format(time.ANSIC))))
		w.Write([]byte(fmt.Sprintln(till.Format(time.RFC3339))))	
	} else {
		toggl.DisableLog()
	}

	session := toggl.OpenSession(me.ApiToken)

	account, err := session.GetAccount()
	if err != nil { return err }

	tags := account.Data.Tags
	tagSumMap := makeTagDurationSumMap(tags)
	//From(tags).ForEachT(func(tag toggl.Tag) {
	//	w.Write([]byte(fmt.Sprintln(tag.Name)))
	//})

	projects, err := session.GetProjects(me.WorkSpaceId)
	if err != nil { return err }
	projectMap := makeProjectMap(projects)

	summaryReport, err := session.GetSummaryReport(me.WorkSpaceId, from.Format(time.RFC3339), till.Format(time.RFC3339))
	if err != nil { return err }

	//w.Write([]byte(fmt.Sprintf("total grand %d\n", summaryReport.TotalGrand)))
	
	projectSummaries := []*ProjectSummary{}
	for _, data := range summaryReport.Data {

		projectSummaryItems := []*ProjectSummaryItem{}
		for _, item := range data.Items {
			//w.Write([]byte(fmt.Sprintf("item %s %d\n", item.Title["time_entry"], item.Time)))
			projectSummaryItems = append(projectSummaryItems, newProjectSummaryItem(item.Title["time_entry"], int64(item.Time)))
		}

		//w.Write([]byte(fmt.Sprintf("project %s %d\n", data.Title.Project, data.Time)))
		projectSummaries = append(projectSummaries, newProjectSummary(data.Title.Project, int64(data.Time), showDetail, projectSummaryItems))
	}

	timeEntries, err := session.GetTimeEntries(from, till)
	if err != nil { return err }

	//durationSum := int64(0)
	timeEntryDetails := []*TimeEntryDetail{}
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst())
		stop := te.Stop.In(jst())
		//duration := time.Duration(te.Duration) * time.Second
		//durationSum = durationSum + te.Duration
		From(te.Tags).ForEachT(func(tagname string) {
			tagSumMap[tagname] = tagSumMap[tagname] + te.Duration
		})
		//s := fmt.Sprintf("- [%s] %02d:%02d - %02d:%02d %v %v", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
		timeEntryDetails = append(timeEntryDetails, newTimeEntryDetail(int64(te.Duration), start, stop, projectMap[te.Pid], te.Description))
	})

	//content.DurationSum = fmtDurationHHMM(time.Duration(durationSum) * time.Second)
	tagSummaries := []*TagSummary{}
	From(tags).
		WhereT(func(tag toggl.Tag) bool {
			return tagSumMap[tag.Name] > 0
		}).
		SelectT(func(tag toggl.Tag) *TagSummary {
			return newTagSummary(tag.Name, tagSumMap[tag.Name], int64(summaryReport.TotalGrand))
			//return fmt.Sprintf("- [%s] %6.2f%% %s", fmtDurationHHMM(time.Duration(tagSumMap[tag.Name]) * time.Second), float64(tagSumMap[tag.Name]) * float64(100) / float64(durationSum), tag.Name)	
		}).
		ToSlice(&tagSummaries)

	content := newOutputContent(from, till, int64(summaryReport.TotalGrand), timeEntryDetails, projectSummaries, tagSummaries)

	err = template.Execute(w, content)
	if err != nil { return err }

	return nil
}

func (me *TogglClient) Date(dateS string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02 MST", fmt.Sprintf("%s JST", dateS))
	if err != nil { return date, err }

	return date, nil
}

func (me *TogglClient) Today() (date time.Time) {
	date = time.Now()
	date = date.Truncate( time.Hour ).Add( - time.Duration(date.Hour()) * time.Hour )
	return date
}

func (me *TogglClient) Yesterday() (time.Time) {
	r := time.Now().AddDate(0, 0, -1)
	return r.Truncate( time.Hour ).Add( - time.Duration(r.Hour()) * time.Hour )
}

func daysAgo(date time.Time, days int) (time.Time) {
	r := date.AddDate(0, 0, -1 * days)
	return r.Truncate( time.Hour ).Add( - time.Duration(r.Hour()) * time.Hour )
}

func (me *TogglClient) After24Hours(date time.Time, days time.Duration) (time.Time) {
	return date.Add((days * 24 * 60 * 60 - 1) * time.Second)
}

func startDayOfWeek(date time.Time) (time.Time) {
	isoYear, isoWeek := date.ISOWeek()

	year, month, day := isoweek.StartDate(isoYear, isoWeek)
	r := time.Date(year, month, day, 0, 0, 0, 0, jst())

	return r
}

func (me *TogglClient) StartDayOfThisWeek() (time.Time) {
	return startDayOfWeek(me.Today())
}

func (me *TogglClient) StartDayOfLastWeek() (time.Time) {
	return startDayOfWeek(daysAgo(me.Today(), 7))
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