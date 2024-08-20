package esyapplier

import (
	"context"
	"fmt"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/Z3DRP/bojoBot/internal/bojo"
)

type LinkdInEasyApply struct {
	Search *bojo.BojoSearch
}

func (ea *EasyApply) Apply(ctx context.Context, bo bojo.BojoSearch) error {
	page := ea.Browser.MustPage(ops.Url)
	// get object that holds the selectors and html elements for linked easy apply
	page.MustElement(ops.ApplyButton).MustClick()

	bo.SubmissionCountMtx.Lock()
	defer bo.SubmissionCountMtx.Unlock()

	*&bo.SubmissionCount++
	if bo.UseSubmissionCount && *&bo.SubmissionCount >= bo.SubmissionLimit {
		return fmt.Errorf("job limit reached")
	}

	return nil
}
