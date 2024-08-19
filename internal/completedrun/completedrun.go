package completedrun

import "time"

type CompletedRun struct {
	Id                    int
	ExecutionDate         time.Time
	Start                 string
	Finish                string
	ApplicationsSubmitted int
	FailedSubmissions     bool
	RunId                 int
}
