package applier

import (
	"context"

	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/listing"
	"github.com/go-rod/rod"
)

type Applier interface {
	Apply(ctx context.Context, jp *bojo.BojoSearch)
	ParseJobs(pg *rod.Page, jobTitle string, expYrs int, expLvl string) (map[string][]listing.Listing, error)
}
