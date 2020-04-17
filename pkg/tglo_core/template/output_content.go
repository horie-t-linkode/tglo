package template

import (
	"fmt"
	"time"
)

func fmtDurationHHMM(d time.Duration) string {
    d = d.Round(time.Minute)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    return fmt.Sprintf("%02d:%02d", h, m)
}

type OutputContent struct {
	From string
	Till string
	DurationTotal string
	TimeEntryDetails []*TimeEntryDetail
	ProjectSummaries []*ProjectSummary
	TagSummaries []*TagSummary
}

type TimeEntryDetail struct {
	Duration string
	From string
	Till string
	ProjectName string
	Description string
}

type ProjectSummary struct {
	Name string
	Duration string
	ShowDetail bool
	Items []*ProjectSummaryItem
}

type ProjectSummaryItem struct {
	Title string
	Duration string
}

type TagSummary struct {
	Name string
	Duration string
	Ratio string
}

func NewOutputContent(from time.Time, till time.Time, durationTotal int64, timeEntryDetails []*TimeEntryDetail, projectSummaries []*ProjectSummary, tagSummaries []*TagSummary) (*OutputContent) {
	me := &OutputContent{
		From: from.Format("2006-01-02"),
		Till: till.Format("2006-01-02"),
		DurationTotal: fmtDurationHHMM(time.Duration(durationTotal/1000) * time.Second),
		TimeEntryDetails: timeEntryDetails,
		ProjectSummaries: projectSummaries,
		TagSummaries: tagSummaries,
	}

	return me
}

func NewTimeEntryDetail(duration int64, from time.Time, till time.Time, projectName string, description string) (*TimeEntryDetail) {

	me := &TimeEntryDetail {
		Duration: fmtDurationHHMM(time.Duration(duration) * time.Second),
		From: fmt.Sprintf("%02d:%02d", from.Hour(), from.Minute()),
		Till: fmt.Sprintf("%02d:%02d", till.Hour(), till.Minute()),
		ProjectName: projectName,
		Description: description,
	}

	return me
}

func NewProjectSummary(name string, duration int64, showDetail bool, items []*ProjectSummaryItem) (*ProjectSummary) {
	me := &ProjectSummary {
		Name: name,
		Duration: fmtDurationHHMM(time.Duration(duration/1000) * time.Second),
		ShowDetail: showDetail,
		Items: items,
	}

	return me
}

func NewProjectSummaryItem(title string, duration int64) (*ProjectSummaryItem) {
	me := &ProjectSummaryItem {
		Title: title,
		Duration: fmtDurationHHMM(time.Duration(duration/1000) * time.Second),
	}

	return me
}

func NewTagSummary(name string, duration int64, durationTotal int64) (*TagSummary) {
	me := &TagSummary {
		//fmt.Sprintf("- [%s] %6.2f%%", fmtDurationHHMM(time.Duration(name) * time.Second), float64(duration) * float64(100) / float64(durationTotal), name)	
		Name: name,
		Duration: fmtDurationHHMM(time.Duration(duration) * time.Second),
		Ratio: fmt.Sprintf("%6.2f%%", float64(duration) * float64(100) / float64(durationTotal/1000)),
	}

	return me
}