package applier

import (
	"context"

	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/job"
)

type Applier interface {
	Apply(ctx context.Context, jp *bojo.BojoSearch, ops RunOptions)
}

func ParseJobs(pg *rod.Page, jobTitle string, jobExperinceYrs int, jobExperinceLvl string) (map[string][]listing.Listing, error) {
	// parses out complex and easy apply jobs
	jp.SubmissionCount = 0
	jobListings := make(map[string][]listing.Listing)
	return jobListings, nil
}
