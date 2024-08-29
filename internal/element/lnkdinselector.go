package element

import "fmt"

type linkedInSelector struct {
	Listing        string // data-job-id on a div
	Job            string //title need find
	HasApplied     string // li in bottom of div
	EasyApplyFlag  string // li in bottom of div
	EasyApplyBtn   string // btn with txt content of 'Easy Apply'
	JobType        string // remote, onsite, hybrid need to find
	ModalNext      string // btn on modal to move to next page
	ModalSubmit    string // btn on modal to submit
	ResumeSelect   string // btn to select resume
	LoginInpt      string
	PwdInpt        string
	LoginBtn       string
	ListingElement string
	JobElement     string
}

func NewLinkedInSelector() *linkedInSelector {
	return &linkedInSelector{
		Listing:        "",
		Job:            "title need find",
		HasApplied:     "li cls",
		EasyApplyFlag:  "li cls",
		JobType:        "li cls",
		ModalNext:      "nxt btn",
		ModalSubmit:    "smt btn",
		ResumeSelect:   "file btn",
		LoginInpt:      "usr inpt",
		PwdInpt:        "pwd inpt",
		LoginBtn:       "login",
		ListingElement: "div",
		JobElement:     "div",
	}
}

func (ls *linkedInSelector) SetListingId(jobID string) {
	ls.Listing = fmt.Sprintf(`[data-job-id="%s"]`, jobID)
}
