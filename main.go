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
)

func main() {
	if len(os.Args) != 2 {
		println("usage:", os.Args[0], "API_TOKEN")
		return
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	session := toggl.OpenSession(os.Args[1])

	account, err := session.GetAccount()
	if err != nil {
		println("error:", err)
		return
	}

	//projectMap := map[int]string{}
	projectMap := makeProjectMap(account)
/*
	From(account.Data.Projects).ToMapByT(&projectMap,
        func(p toggl.Project) int { return p.ID },
        func(p toggl.Project) string { return p.Name },
    )
*/

	From(account.Data.TimeEntries).ForEachT(func(te toggl.TimeEntry) {	
		//data, _ := json.MarshalIndent(te.(toggl.TimeEntry).Description, "", "    ")
		//println("account:", string(data))

		start := te.Start.In(jst)
		stop := te.Stop.In(jst)
		duration := time.Duration(te.Duration) * time.Second
		fmt.Printf("- [%s] %02d:%02d-%02d:%02d %v %v\n", fmtDurationHHMM(duration), start.Hour(), start.Minute(), stop.Hour(), stop.Minute(), projectMap[te.Pid], te.Description)
	})
}

func makeProjectMap(account toggl.Account) map[int]string {
	projectMap := map[int]string{}
	From(account.Data.Projects).ToMapByT(&projectMap,
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