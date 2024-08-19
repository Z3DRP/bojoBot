package scheduledrun

import "time"

type ScheduledRun struct {
	Id                  int
	Name                string
	CreationDate        time.Time
	JobTitleId          int
	JobBoardId          int
	RunDay              int
	RunDayOfWeek        string
	RunMonth            string
	RunType             string
	RunTime             string
	Recurring           bool
	EasyApplyOnly       bool
	DurationMinutes     float32
	NumberOfSubmissions int
	EveryHour           int
	EveryMin            int
}
