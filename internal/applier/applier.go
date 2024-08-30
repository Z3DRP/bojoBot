package applier

import (
	"context"

	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type Applier interface {
	HandleJobSearch()
	apply(ctx context.Context, jp *bojo.BojoSearch)
	ParseJobs(pg *rod.Page, c bojo.SearchCriteria) (map[string][]listing.Listing, error)
}
