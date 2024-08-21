package element

type LinkedInSelector struct {
	Listing       string // data-job-id on a div
	Job           string //title need find
	HasApplied    string // li in bottom of div
	EasyApplyFlag string // li in bottom of div
	EasyApplyBtn  string // btn with txt content of 'Easy Apply'
	JobType       string // remote, onsite, hybrid need to find
	ModalNext     string // btn on modal to move to next page
	ModalSubmit   string // btn on modal to submit
	ResumeSelect  string // btn to select resume
	LoginInpt     string
	PwdInpt       string
	LoginBtn      string
}
