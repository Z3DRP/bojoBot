package jprocessor

import (
	"context"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type JobProcessor struct {
	Browser            *rod.Browser
	SubLimit           int
	SubCount           int
	SubCountMtx        *sync.Mutex
	UseSubmissionCount bool
}

func NewJobProcessor(bwsr *rod.Browser, subLimit int, useSub bool) *JobProcessor {
	return &JobProcessor{
		Browser:            bwsr,
		SubLimit:           subLimit,
		SubCountMtx:        &sync.Mutex{},
		UseSubmissionCount: useSub,
	}
}

func (jp *JobProcessor) ProcessJobs(ctx context.Context, listings []listing.Listing, ap applier.Applier) error {

}

func (jp *JobProcessor) ParseJobs(pg *rod.Page) (map[string][]listing.Listing, error) {
	// parses out complex and easy apply jobs
	jobListings := make(map[string][]listing.Listing)
	return jobListings, nil
}
