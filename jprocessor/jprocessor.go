package jprocessor

import (
	"context"
	"sync"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/go-rod/rod"
)

type JobProcessor struct {
	Browser *rod.Browser
	SubLimit int
	SubCount int
	SubCountMtx *sync.Mutex
}

func NewJobProcessor(bwsr *rod.Browser, subLimit int) *JobProcessor {
	return &JobProcessor{
		Browser: bwsr,
		SubLimit: subLimit,
		SubCountMtx: &sync.Mutex{},
	}
}

func (jp *JobProcessor) ProcessJobs(ctx context.Context, listings []Listing, applier.Applier) error {
	
}