package scheduledruns

import "time"

type ScheduledRuns struct {
	id                  int
	name                string
	creationDate        time.Time
	jobTitleId          int
	jobBoardId          int
	runDay              int
	runDayOfWeek        string
	runMonth            string
	runTime             string
	runType             string
	recurring           bool
	easyApplyOnly       bool
	durrationMinutes    float32
	numberOfSubmissions int
	everyHour           int
	everyMin            int
}
