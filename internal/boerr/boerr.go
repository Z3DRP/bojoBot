package boerr

import "fmt"

type GeneralError struct {
	Err string
}

func (ge *GeneralError) Error() string {
	return fmt.Sprintf("general error: %s", ge.Err)
}

type NoResults struct {
	s string
}

func (nr NoResults) Error() string {
	return nr.s
}

type SubmissionError struct {
	s     string
	JobId string
}

func (se SubmissionError) Error() string {
	return se.s
}

type NoListingFoundError struct {
	JobId       string
	ElementType string
}

func (nlfe *NoListingFoundError) Error() string {
	return fmt.Sprintf("No listings found jobId: %s on element %s", nlfe.JobId, nlfe.ElementType)
}
