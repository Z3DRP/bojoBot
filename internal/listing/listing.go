package listing

type ListingT struct {
	JobName           string
	Url               string
	PreviouslyApplied bool
	Comapny           string
	Discription       string
}

type Listing interface {
	// need to put something
}
