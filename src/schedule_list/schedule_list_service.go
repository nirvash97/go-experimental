package main

import (
	"fmt"
	"time"
)

type Schedule struct {
	date    time.Time
	display bool
	meeting []ScheduleMeeting
}

type ScheduleMeeting struct {
	meeting_time string
	status       string
}

func main() {
	fmt.Println("Hi Schedule")
	var scheduleList []Schedule

	scheduleDetail := Schedule{date: time.Now().UTC(), display: false, meeting: []ScheduleMeeting{
		{meeting_time: "10.00", status: "Available"},
		{meeting_time: "11.00", status: "Available"},
		{meeting_time: "12.00", status: "Blocked"}}}
	scheduleList = append(scheduleList, scheduleDetail)
	fmt.Println(scheduleList)

	fmt.Println("Date : " + scheduleList[0].date.Format(time.RFC3339))

}
