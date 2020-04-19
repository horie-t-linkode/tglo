package tglo_core

import (
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"io"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/template"
	"github.com/masaki-linkode/tglo/pkg/tglo_core/time_util"
)

type TogglClient struct {
	ApiToken string
	WorkSpaceId int
	VerboseOut io.Writer
}

func (me *TogglClient) Process(from time.Time, till time.Time, showDetail bool) (content *template.OutputContent, err error) {

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
	
	projectSummaries := []*template.ProjectSummary{}
	for _, data := range summaryReport.Data {

		projectSummaryItems := []*template.ProjectSummaryItem{}
		for _, item := range data.Items {
			//verboseOut.Write([]byte(fmt.Sprintf("item %s %d\n", item.Title["time_entry"], item.Time)))
			projectSummaryItems = append(projectSummaryItems, template.NewProjectSummaryItem(item.Title["time_entry"], int64(item.Time)))
		}

		//verboseOut.Write([]byte(fmt.Sprintf("project %s %d\n", data.Title.Project, data.Time)))
		projectSummaries = append(projectSummaries, template.NewProjectSummary(data.Title.Project, int64(data.Time), showDetail, projectSummaryItems))
	}

	timeEntries, err := session.GetTimeEntries(from, till)
	if err != nil { return nil, err }

	timeEntryDetails := []*template.TimeEntryDetail{}
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		if te.Stop != nil { // タイマー実行中のエントリは対象から外す。
			start := te.Start.In(time_util.Jst())
			stop := te.Stop.In(time_util.Jst())
			From(te.Tags).ForEachT(func(tagname string) {
				tagSumMap[tagname] = tagSumMap[tagname] + te.Duration
			})
			timeEntryDetails = append(timeEntryDetails, template.NewTimeEntryDetail(int64(te.Duration), start, stop, projectMap[te.Pid], te.Description))
		}
	})

	tagSummaries := []*template.TagSummary{}
	From(tags).
		WhereT(func(tag toggl.Tag) bool {
			return tagSumMap[tag.Name] > 0
		}).
		SelectT(func(tag toggl.Tag) *template.TagSummary {
			return template.NewTagSummary(tag.Name, tagSumMap[tag.Name], int64(summaryReport.TotalGrand))
		}).
		ToSlice(&tagSummaries)

	content = template.NewOutputContent(from, till, int64(summaryReport.TotalGrand), timeEntryDetails, projectSummaries, tagSummaries)

	return content, nil
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