package esyapplier

import (
	"fmt"
	"time"

	"github.com/Z3DRP/bojoBot/internal/boerr"
	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/element"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// returns all three inputs on linkedIN search component
// document.querySelectorAll(".jobs-search-box__keyboard-text-input")

// search from linkedIn/jobs is executed by enter press

// attribute for input that holds value typed into search and enter pressed and page reloads
// "data-job-search-box-keywords-input-trigger"

// cls of submit button if you go to actual search page linkedIn/jobs/search instead of linkedIn/jobs
// .jobs-search-box__submit-button
type LinkdInEasyApply struct {
	Search *bojo.BojoSearch
}

func (l LinkdInEasyApply) loadListingChannel(lst []listing.Listing) <-chan listing.Listing {
	out := make(chan listing.Listing)
	go func() {
		for _, l := range lst {
			out <- l
		}
		close(out)
	}()
	return out
}

// TODO ctx cancel will be handled by the service because it will have esy applys and cmpx applys or both
func (l *LinkdInEasyApply) HandleApply() <-chan bojo.SubmissionResults {
	l.Search.Browser = rod.New().MustConnect()
	outChan := make(chan bojo.SubmissionResults)
	page := l.Search.Browser.MustPage(l.Search.URL)
	// TODO add search and click for job title
	// searchInputs, err := page.Elements(".jobs-search-box__keyboard-text-input")
	// searchInput, err := page.Element(".jobs-search-global-typeahead__input")
	page.MustElement(".jobs-search-global-typeahead__input").MustInput(l.Search.Criteria.JobName).MustType(input.Enter)
	page.MustWaitLoad()

	//TODO make parse jobs
	jobs, err := l.ParseJobs(page)
	if err != nil {
		outChan <- *bojo.NewSubmissionResult(nil, boerr.BrowserError{})
		close(outChan)
	}

	listingChan := l.loadListingChannel(jobs)
	resultChan := l.apply(listingChan)
	processedListings := make(map[string]bojo.BoResult)
	for rslt := range resultChan {
		if _, exists := processedListings[rslt.JobListing.Id]; !exists {
			processedListings[rslt.JobListing.Id] = rslt
		}
		// OR
		// outChan <- createNewSubmissionResult, but would need to hold a single listing and then an err value
		// then would have to remove the outChan <- *bojo below
		// i think this way I could merge the results in the service cls for complex and easy applies into one
	}

	l.Search.Browser.Close()
	outChan <- *bojo.NewSubmissionResult(processedListings, nil)
	return outChan
}

func (l *LinkdInEasyApply) apply(listings <-chan listing.Listing) <-chan bojo.BoResult {
	// page := l.Search.Browser.MustPage(l.Search.URL)
	// page, err := lnkdin.Search.Browser.Page(lnkdin.Search.URL)
	// easyApplyJobs := bo.Jobs["easyApply"]

	subChan := make(chan bojo.BoResult)

	go func() {
		selector := element.NewLinkedInSelector()
		createPage := func() *rod.Page {
			return l.Search.Browser.MustPage()
		}

		for jobListing := range listings {
			page := l.Search.EasyPool.MustGet(createPage)
			selector.SetListingId(jobListing.Id)
			// listingSelector := fmt.Sprintf(`[data-job-id="%s"]`, jobListing.Id)
			// page.MustElement(element.LinkedInSelector).MustClick()

			elm, err := page.Element(selector.Listing)

			if err != nil {
				subChan <- bojo.NewBoResult(jobListing, boerr.NoResults{S: fmt.Sprintf("listing id: %s not found on page", jobListing.Id)})
				continue
			}

			if err := elm.Click("left", 1); err != nil {
				subChan <- bojo.NewBoResult(jobListing, boerr.PageActionError{S: "click"})
				continue
			}
			page.MustWaitStable()
			// if handleErr(page.WaitStable(time.Millisecond*200), &errors, func() error { return &boerr.GeneralError{Err: err.Error()} }) {
			// 	continue
			// }

			applyBtn, err := page.Element("jobs-apply-button--top-card")
			if err != nil {
				subChan <- bojo.NewBoResult(jobListing, boerr.NoResults{S: "Could not find apply button"})
			}

			err = applyBtn.WaitStable(time.Microsecond * 200)
			if err != nil {
				subChan <- bojo.NewBoResult(jobListing, &boerr.GeneralError{Err: "An error occurred while waiting for page to stablize"})
			}
			// applyBtn.MustWaitStable().MustClick().MustWaitInvisible()
			// applyBtn.WaitStable(time.Millisecond*200).Click("left", 1)

			if err := applyBtn.Click("left", 1); err != nil {
				subChan <- bojo.NewBoResult(jobListing, boerr.PageActionError{S: "click"})
			}

			// NOTE logic above gets to the point where job is clicked, then job details rendered on right side of screen and modal should be
			// opened in the last applyBtn.Click() call
			// TODO handle modal input, next btn clicks, and submt btn clicks, modal exit click

			l.Search.SubmissionCountMtx.Lock()
			l.Search.SubmissionCount++
			l.Search.SubmissionCountMtx.Unlock()
			l.Search.EasyPool.Put(page)

			if l.Search.UseSubmissionCount && l.Search.SubmissionCount >= l.Search.SubmissionLimit {
				// job limit reached
				break
			}
		}

		close(subChan)
	}()

	// // move locking unlocking into loop
	// bo.SubmissionCountMtx.Lock()
	// defer bo.SubmissionCountMtx.Unlock()

	// bo.SubmissionCount++
	// if bo.UseSubmissionCount && bo.SubmissionCount >= bo.SubmissionLimit {
	// 	return fmt.Errorf("job limit reached")
	// }

	return subChan
}

func (lnkdin *LinkdInEasyApply) ParseJobs(pg *rod.Page) ([]listing.Listing, error) {
	// parse out easy jobs
	// then loop and write each to chan
	jobListings := make([]listing.Listing, 0)
	return jobListings, nil
}
