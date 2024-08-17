package models

import "time"

type CompletedRuns struct {
	id                    int
	executionDate         time.Time
	start                 string
	finish                string
	applicationsSubmitted int
	failedSubmissions     bool
	runId                 int
}
