package bojo

import (
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

type BojoSearch struct {
	Browser            *rod.Browser
	Pool               rod.Pool[rod.Page]
	SubmissionLimit    int
	SubmissionCount    int
	SubmissionCountMtx *sync.Mutex
	UseSubmissionCount bool
	Jobs               map[string][]listing.LinkedinListing
	URL                string
	Criteria           *SearchCriteria
}

func NewBojoSearch(bwsr *rod.Browser, job *job.JobTitle, subLimit int, useSub bool) *BojoSearch {
	opts := NewSearchCriteria(job.Name, job.ExperienceLevel, job.ExperienceYears)
	return &BojoSearch{
		Browser:            bwsr,
		Pool:               rod.NewPagePool(20),
		SubmissionLimit:    subLimit,
		SubmissionCountMtx: &sync.Mutex{},
		SubmissionCount:    0,
		UseSubmissionCount: useSub,
		Criteria:           opts,
	}
}

func NewSearchCriteria(name string, exLvl string, exYr int) *SearchCriteria {
	return &SearchCriteria{
		JobName:       name,
		ExperienceLvl: exLvl,
		ExperienceYrs: exYr,
	}
}
