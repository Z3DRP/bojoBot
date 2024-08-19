package applier

import (
	"context"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/job"
)

type Applier interface {
	Apply(ctx context.Context, ops RunOptions, submissionCountMutex *sync.Mutex, subCount *int, subLimit int)
}

type RunOptions struct {
	JobInfo       job.JobTitle
	Url           string
	ApplySelector string
	FinishButton  string
	// other elements
}

func NewRunOptions(job job.JobTitle, url string, selctr string, finSelectr string) *RunOptions {
	return &RunOptions{
		JobInfo:       job,
		Url:           url,
		ApplySelector: selctr,
		FinishButton:  finSelectr,
	}
}
