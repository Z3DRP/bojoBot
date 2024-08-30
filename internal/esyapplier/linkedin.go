package esyapplier

import (
	"context"
	"time"

	"github.com/Z3DRP/bojoBot/internal/boerr"
	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/element"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type LinkdInEasyApply struct {
	Search *bojo.BojoSearch
}

func handleErr(err error, errors *[]error, createCstmErr func() error) bool {
	if err != nil {
		*errors = append(*errors, createCstmErr())
		return true
	}
	return false
}

// TODO ctx cancel will be handled by the service because it will have esy applys and cmpx applys or both
func (l *LinkdInEasyApply) HandleApply() <-chan bojo.SubmissionResults {
	// this func will take the context and the bo search
	// it will loop all jobs and call apply on each job.. and keep track of errors for single jobs
	// and it will keep track of the context timeout and job submisn count
	outChan := make(chan bojo.SubmissionResults)
	page := l.Search.Browser.MustPage(l.Search.URL)
	jobs, err := l.ParseJobs(page, l.Search.Criteria)
	if err != nil {
		outChan <- *bojo.NewSubmissionResult(nil, boerr.BrowserError{})
	}

	//TODO write jobs to channel then pass chanel to apply
	for _, job := range jobs {
		l.Search.SubmissionCountMtx.Lock()
		if l.Search.UseSubmissionCount && l.Search.SubmissionCount > l.Search.SubmissionLimit {
			l.Search.SubmissionCountMtx.Unlock()
			break
		}
		l.Search.SubmissionCountMtx.Unlock()

	}

	return outChan
}

func (l *LinkdInEasyApply) apply(ctx context.Context, bo *bojo.BojoSearch) (bool, []error) {
	errors := make([]error, 0)
	page := l.Search.Browser.MustPage(l.Search.URL)
	// page, err := lnkdin.Search.Browser.Page(lnkdin.Search.URL)
	easyApplyJobs := bo.Jobs["easyApply"]
	selector := element.NewLinkedInSelector()
	for _, job := range easyApplyJobs {
		selector.SetListingId(job.Id)
		// listingSelector := fmt.Sprintf(`[data-job-id="%s"]`, job.Id)
		// page.MustElement(element.LinkedInSelector).MustClick()
		elm, err := page.Element(selector.Listing)
		// if err != nil {
		// 	errors = append(errors, &boerr.GeneralError{Err: err.Error()})
		// 	continue
		// }
		if handleErr(err, &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}
		if elm == nil {
			errors = append(errors, &boerr.NoListingFoundError{
				JobId:       job.Id,
				ElementType: selector.ListingElement,
			})
			continue
		}

		if handleErr(elm.Click("left", 1), &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}

		if handleErr(page.WaitStable(time.Millisecond*200), &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}

		applyBtn, err := page.Element("jobs-apply-button--top-card")
		if handleErr(err, &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}
		// applyBtn.MustWaitStable().MustClick().MustWaitInvisible()
		// applyBtn.WaitStable(time.Millisecond*200).Click("left", 1)
		err = applyBtn.WaitStable(time.Millisecond * 200)
		if handleErr(applyBtn.WaitStable(time.Millisecond*200), &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}

		err = applyBtn.Click("left", 1)
		if handleErr(applyBtn.Click("left", 1), &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			continue
		}

		// NOTE logic above gets to the point where job is clicked, then job details rendered on right side of screen and modal should be
		// opened in the last applyBtn.Click() call
		// TODO handle modal input, next btn clicks, and submt btn clicks, modal exit click

		bo.SubmissionCountMtx.Lock()
		defer bo.SubmissionCountMtx.Lock()

		bo.SubmissionCount++
		if bo.UseSubmissionCount && bo.SubmissionCount >= bo.SubmissionLimit {
			// job limit reached
			return true, errors
		}
	}

	// // move locking unlocking into loop
	// bo.SubmissionCountMtx.Lock()
	// defer bo.SubmissionCountMtx.Unlock()

	// bo.SubmissionCount++
	// if bo.UseSubmissionCount && bo.SubmissionCount >= bo.SubmissionLimit {
	// 	return fmt.Errorf("job limit reached")
	// }

	return nil
}

func (lnkdin *LinkdInEasyApply) ParseJobs(pg *rod.Page) <-chan listing.Listing {
	// parse out easy jobs
	// then loop and write each to chan
	jobListings := make(map[string][]listing.LinkedinListing)
	return jobListings
}
