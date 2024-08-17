package applications

import (
	"time"
)

type Application struct {
	id                    int
	company               string
	jobTitleId            int
	jobBoardId            int
	location              string
	pay                   float32
	applyDate             time.Time
	submittedSuccessfully bool
	runId                 int
}
