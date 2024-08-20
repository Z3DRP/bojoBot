package bojo

import (
	"context"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/job"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type RunOptions struct {
	JobName       string
	ExperienceLvl string
	ExperienceYrs int
	// other elements
}

type BojoSearch struct {
	Browser            *rod.Browser
	Pool               rod.Pool[rod.Page]
	SubmissionLimit    int
	SubmissionCount    int
	SubmissionCountMtx *sync.Mutex
	UseSubmissionCount bool
	Jobs               []string
	Options            *RunOptions
}

func NewBojoSearch(bwsr *rod.Browser, job *job.JobTitle, subLimit int, useSub bool) *BojoSearch {
	opts := NewRunOptions(job.Name, job.ExperienceLevel, job.ExperienceYears)
	return &BojoSearch{
		Browser:            bwsr,
		Pool:               rod.NewPagePool(20),
		SubmissionLimit:    subLimit,
		SubmissionCountMtx: &sync.Mutex{},
		SubmissionCount:    0,
		UseSubmissionCount: useSub,
		Options:            opts,
	}
}

func NewRunOptions(name string, exLvl string, exYr int) *RunOptions {
	return &RunOptions{
		JobName:       name,
		ExperienceLvl: exLvl,
		ExperienceYrs: exYr,
	}
}
