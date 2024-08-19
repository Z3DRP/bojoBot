package applier

import (
	"context"
	"sync"
)

type Applier interface {
	Apply(ctx context.Context, job Job, submissionCountMutex *sync.Mutex, subCount *int, subLimit int)
}

type 

type PageCommands struct {
	ApplyButton  string
	FinishButton string
	// other elements
}
