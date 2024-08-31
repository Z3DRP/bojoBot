package boerr

import "fmt"

type GeneralError struct {
	Err string
}

func (ge *GeneralError) Error() string {
	return fmt.Sprintf("general error: %s", ge.Err)
}

type NoResults struct {
	S string
}

func (nr NoResults) Error() string {
	return nr.S
}

type SubmissionError struct {
	S     string
	JobId string
}

func (se SubmissionError) Error() string {
	return se.S
}

type NoListingFoundError struct {
	JobId       string
	ElementType string
}

func (nlfe *NoListingFoundError) Error() string {
	return fmt.Sprintf("No listings found jobId: %s on element %s", nlfe.JobId, nlfe.ElementType)
}

type BrowserError struct {
	S string
}

func (be BrowserError) Error() string {
	return fmt.Sprintf("A browser error ocurred: %s", be.S)
}

type PageActionError struct {
	S string
}

func (p PageActionError) Error() string {
	return fmt.Sprintf("An expected error occurred while performing: %s", p.S)
}
