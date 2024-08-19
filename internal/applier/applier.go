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
	JobInfo            job.JobTitle
	Url                string
	ApplyButton        string
	FinishButton       string
	UseSubmissionCount bool
	// other elements
}
