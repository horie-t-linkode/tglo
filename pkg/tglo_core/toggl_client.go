
package tglo_core

import (
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"github.com/snabb/isoweek"
	"io"
)

type TogglClient struct {
	ApiToken string
	WorkSpaceId int
	VerboseOut io.Writer
}

func (me *TogglClient) Process(from time.Time, till time.Time, showDetail bool) (content *OutputContent, err error) {

	if me.VerboseOut != nil {
		toggl.EnableLog()
		me.VerboseOut.Write([]byte(fmt.Sprintln(from.Format(time.ANSIC))))
		me.VerboseOut.Write([]byte(fmt.Sprintln(from.Format(time.RFC3339))))
		me.VerboseOut.Write([]byte(fmt.Sprintln(till.Format(time.ANSIC))))
		me.VerboseOut.Write([]byte(fmt.Sprintln(till.Format(time.RFC3339))))	
	} else {
		toggl.DisableLog()
	}

	session := toggl.OpenSession(me.ApiToken)

	account, err := session.GetAccount()
	if err != nil { return nil, err }

	tags := account.Data.Tags
	tagSumMap := makeTagDurationSumMap(tags)
	//From(tags).ForEachT(func(tag toggl.Tag) {
	//	verboseOut.Write([]byte(fmt.Sprintln(tag.Name)))
	//})

	projects, err := session.GetProjects(me.WorkSpaceId)
	if err != nil { return nil, err }
	projectMap := makeProjectMap(projects)

	summaryReport, err := session.GetSummaryReport(me.WorkSpaceId, from.Format(time.RFC3339), till.Format(time.RFC3339))
	if err != nil { return nil, err }

	//verboseOut.Write([]byte(fmt.Sprintf("total grand %d\n", summaryReport.TotalGrand)))
	
	projectSummaries := []*ProjectSummary{}
	for _, data := range summaryReport.Data {

		projectSummaryItems := []*ProjectSummaryItem{}
		for _, item := range data.Items {
			//verboseOut.Write([]byte(fmt.Sprintf("item %s %d\n", item.Title["time_entry"], item.Time)))
			projectSummaryItems = append(projectSummaryItems, newProjectSummaryItem(item.Title["time_entry"], int64(item.Time)))
		}

		//verboseOut.Write([]byte(fmt.Sprintf("project %s %d\n", data.Title.Project, data.Time)))
		projectSummaries = append(projectSummaries, newProjectSummary(data.Title.Project, int64(data.Time), showDetail, projectSummaryItems))
	}

	timeEntries, err := session.GetTimeEntries(from, till)
	if err != nil { return nil, err }

	timeEntryDetails := []*TimeEntryDetail{}
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst())
		stop := te.Stop.In(jst())
		From(te.Tags).ForEachT(func(tagname string) {
			tagSumMap[tagname] = tagSumMap[tagname] + te.Duration
		})
		timeEntryDetails = append(timeEntryDetails, newTimeEntryDetail(int64(te.Duration), start, stop, projectMap[te.Pid], te.Description))
	})

	tagSummaries := []*TagSummary{}
	From(tags).
		WhereT(func(tag toggl.Tag) bool {
			return tagSumMap[tag.Name] > 0
		}).
		SelectT(func(tag toggl.Tag) *TagSummary {
			return newTagSummary(tag.Name, tagSumMap[tag.Name], int64(summaryReport.TotalGrand))
		}).
		ToSlice(&tagSummaries)

	content = newOutputContent(from, till, int64(summaryReport.TotalGrand), timeEntryDetails, projectSummaries, tagSummaries)

	return content, nil
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

func Yesterday() (time.Time) {
	r := time.Now().AddDate(0, 0, -1)
	return r.Truncate( time.Hour ).Add( - time.Duration(r.Hour()) * time.Hour )
}

func daysAgo(date time.Time, days int) (time.Time) {
	r := date.AddDate(0, 0, -1 * days)
	return r.Truncate( time.Hour ).Add( - time.Duration(r.Hour()) * time.Hour )
}

func After24Hours(date time.Time, days time.Duration) (time.Time) {
	return date.Add((days * 24 * 60 * 60 - 1) * time.Second)
}

func startDayOfWeek(date time.Time) (time.Time) {
	isoYear, isoWeek := date.ISOWeek()

	year, month, day := isoweek.StartDate(isoYear, isoWeek)
	r := time.Date(year, month, day, 0, 0, 0, 0, jst())

	return r
}

func StartDayOfThisWeek() (time.Time) {
	return startDayOfWeek(Today())
}

func StartDayOfLastWeek() (time.Time) {
	return startDayOfWeek(daysAgo(Today(), 7))
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