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
	JobId    string
	Company  string
	Location string
	URL      string
	Position string
	Pay      float64
	Err      error
}

type SubmissionResults struct {
	Err           error
	ProcessedJobs map[string]BoResult
}

type BojoSearch struct {
	Browser            *rod.Browser
	Pool               rod.Pool[rod.Page]
	Ctx                context.Context
	CtxCancel          context.CancelFunc
	EasyOnly           bool
	SubmissionLimit    int
	SubmissionCount    int
	SubmissionCountMtx *sync.Mutex
	UseSubmissionCount bool
	Jobs               map[string][]listing.LinkedinListing
	URL                string
	Criteria           SearchCriteria
}

func NewBojoSearch(ctx context.Context, cncl context.CancelFunc, bwsr *rod.Browser, job *job.Job, subLimit int, useSub bool) *BojoSearch {
	return &BojoSearch{
		Browser:            bwsr,
		Pool:               rod.NewPagePool(20),
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
