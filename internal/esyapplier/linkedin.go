package esyapplier

import (
	"context"
	"fmt"
	"time"

	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type LinkdInEasyApply struct {
	Search *bojo.BojoSearch
}

func (lnkdin *LinkdInEasyApply) Apply(ctx context.Context, bo *bojo.BojoSearch) error {
	page := lnkdin.Search.Browser.MustPage(lnkdin.Search.URL)
	easyApplyJobs := bo.Jobs["easyApply"]
	for _, job := range easyApplyJobs {
		listingSelector := fmt.Sprintf(`[data-job-id="%s"]`, job.Id)
		// page.MustElement(element.LinkedInSelector).MustClick()
		elm, err := page.Element(listingSelector)
		if err != nil {
			// do something with err
			continue
		}
		elm.Click("left", 1)
		page.WaitStable(time.Millisecond * 200)
		applyBtn, err := page.Element("jobs-apply-button--top-card")
		if err != nil {
			continue
		}
		// applyBtn.MustWaitStable().MustClick().MustWaitInvisible()
		// applyBtn.WaitStable(time.Millisecond*200).Click("left", 1)
		err = applyBtn.WaitStable(time.Millisecond * 200)
		if err != nil {
			continue
		}
		err = applyBtn.Click("left", 1)
		if err != nil {
			continue
		}

	}

	// move locking unlocking into loop
	bo.SubmissionCountMtx.Lock()
	defer bo.SubmissionCountMtx.Unlock()

	*&bo.SubmissionCount++
	if bo.UseSubmissionCount && *&bo.SubmissionCount >= bo.SubmissionLimit {
		return fmt.Errorf("job limit reached")
	}

	return nil
}

func ParseJobs(pg *rod.Page, jobTitle string, jobExperinceYrs int, jobExperinceLvl string) (map[string][]listing.LinkedinListing, error) {
	// parses out complex and easy apply jobs
	jobListings := make(map[string][]listing.LinkedinListing)
	return jobListings, nil
}
