package cmplxapplier

import (
	"context"
	"fmt"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/go-rod/rod"
)

type ComplexApply struct {
	Browser *rod.Browser
}

func (ca *ComplexApply) Apply(ctx context.Context, ops applier.RunOptions, subCountMtx *sync.Mutex, subCount *int, subLimit int) error {
	page := ca.Browser.MustPage(ops.Url)

	subCountMtx.Lock()
	defer subCountMtx.Unlock()

	*subCount++
	if ops.UseSubmissionCount && *subCount >= subLimit {
		return fmt.Errorf("submission limit")
	}
	return nil
}
