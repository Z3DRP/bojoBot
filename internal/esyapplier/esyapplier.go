package esyapplier

import (
	"context"
	"fmt"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/go-rod/rod"
)

type EasyApply struct {
	Browser *rod.Browser
}

func (ea *EasyApply) Apply(ctx context.Context, ops applier.RunOptions, subCountMtx *sync.Mutex, subCount *int, subLimit int) error {
	page := ea.Browser.MustPage(ops.Url)
	page.MustElement(ops.ApplyButton).MustClick()

	subCountMtx.Lock()
	defer subCountMtx.Unlock()

	*subCount++
	if ops.UseSubmissionCount && *subCount >= subLimit {
		return fmt.Errorf("job limit reached")
	}

	return nil
}
