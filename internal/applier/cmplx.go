package applier

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-rod/rod"
)

type ComplexApply struct {
}

func (ca *ComplexApply) Apply(ctx context.Context, ops RunOptions, subCountMtx *sync.Mutex, subCount *int, subLimit int) error {
	page := ca.Browser.MustPage(ops.Url)

	subCountMtx.Lock()
	defer subCountMtx.Unlock()

	*subCount++
	if ops.UseSubmissionCount && *subCount >= subLimit {
		return fmt.Errorf("submission limit")
	}
	return nil
}
