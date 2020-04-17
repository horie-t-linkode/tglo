package time_util

import (
	"fmt"
	"time"
	"github.com/snabb/isoweek"
)

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
	r := time.Date(year, month, day, 0, 0, 0, 0, Jst())

	return r
}

func StartDayOfThisWeek() (time.Time) {
	return startDayOfWeek(Today())
}

func StartDayOfLastWeek() (time.Time) {
	return startDayOfWeek(daysAgo(Today(), 7))
}

func Jst() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}