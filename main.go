/*

The toggl command will display a user's Toggl account information.

Usage:
    toggl API_TOKEN

The API token can be retrieved from a user's account information page at toggl.com.

*/
package main

import (
	//"encoding/json"
	"os"
	"github.com/jason0x43/go-toggl"
	. "github.com/ahmetb/go-linq"
	"fmt"
	"time"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		println("usage:", os.Args[0], "API_TOKEN WORKSPACE_ID")
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	t := time.Now()
    t = t.Truncate( time.Hour ).Add( - time.Duration(t.Hour()) * time.Hour )
	println(t.Format(time.ANSIC))
	
	y := time.Now().AddDate(0, 0, -1)
	y = y.Truncate( time.Hour ).Add( - time.Duration(y.Hour()) * time.Hour )
	println(y.Format(time.ANSIC))

	apiToken := os.Args[1]
	workspaceId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println("error:", err)
		return
	}

	session := toggl.OpenSession(apiToken)

	projects, err := session.GetProjects(workspaceId)
	if err != nil {
		println("error:", err)
		return
	}

	timeEntries, err := session.GetTimeEntries(y, t)
	if err != nil {
		println("error:", err)
		return
	}
/*
	account, err := session.GetAccount()
	if err != nil {
		println("error:", err)
		return
	}
*/

	fmt.Printf("# %sの実績\n", y.Format("2006-01-02"))

	projectMap := makeProjectMap(projects)

//	From(account.Data.TimeEntries).ForEachT(func(te toggl.TimeEntry) {
	durationSum := int64(0)
	From(timeEntries).ForEachT(func(te toggl.TimeEntry) {
		start := te.Start.In(jst)
		stop := te.Stop.In(jst)
		duration := time.Duration(te.Duration) * time.Second
		durationSum = durationSum + te.Duration
		fmt.Printf("- [%s] %02d:%02d-%02d:%02d %v %v\n", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
	})
	fmt.Printf("total %s\n", fmtDurationHHMM(time.Duration(durationSum) * time.Second))
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