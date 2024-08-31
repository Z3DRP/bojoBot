package bojo

import (
	"context"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/job"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type SearchCriteria struct {
	JobName       string
	ExperienceLvl string
	ExperienceYrs int
}

type BoResult struct {
	JobListing listing.Listing
	Err        error
}

type SubmissionResults struct {
	ProcessedJobs map[string]BoResult
	Err           error
}

type BojoSearch struct {
	Browser            *rod.Browser
	EasyPool           rod.Pool[rod.Page]
	ComplxPool         rod.Pool[rod.Page]
	Ctx                context.Context
	CtxCancel          context.CancelFunc
	EasyOnly           bool
	SubmissionLimit    int
	SubmissionCount    int
	SubmissionCountMtx *sync.Mutex
	UseSubmissionCount bool
	Jobs               map[string][]listing.Listing
	URL                string
	Criteria           SearchCriteria
}

func NewBojoSearch(ctx context.Context, cncl context.CancelFunc, bwsr *rod.Browser, job *job.Job, subLimit int, useSub bool) *BojoSearch {
	return &BojoSearch{
		Browser:            bwsr,
		EasyPool:           rod.NewPagePool(10),
		ComplxPool:         rod.NewPagePool(15),
		Ctx:                ctx,
		CtxCancel:          cncl,
		SubmissionLimit:    subLimit,
		SubmissionCountMtx: &sync.Mutex{},
		SubmissionCount:    0,
		UseSubmissionCount: useSub,
		Criteria:           NewSearchCriteria(job.Name, job.ExperienceLevel, job.ExperienceYears),
	}
}

func NewSearchCriteria(name string, exLvl string, exYr int) SearchCriteria {
	return SearchCriteria{
		JobName:       name,
		ExperienceLvl: exLvl,
		ExperienceYrs: exYr,
	}
}

func NewSubmissionResult(jobs map[string]BoResult, err error) *SubmissionResults {
	return &SubmissionResults{
		ProcessedJobs: jobs,
		Err:           err,
	}
}

func NewBoResult(l listing.Listing, e error) BoResult {
	return BoResult{
		JobListing: l,
		Err:        e,
	}
}
