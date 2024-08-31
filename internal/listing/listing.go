package listing

type Listing struct {
	Id       string
	Company  string
	Location string
	Position string
	Pay      float64
	Url      string
}

func NewListing(id string, c string, l string, pos string, p float64, url string) *Listing {
	return &Listing{
		Id:       id,
		Company:  c,
		Location: l,
		Position: pos,
		Pay:      p,
		Url:      url,
	}
}
