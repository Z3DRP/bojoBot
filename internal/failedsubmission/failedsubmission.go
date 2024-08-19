package failedsubmission

type FailedSubmission struct {
	Id             int
	CompletedRunId int
	ApplicationUrl string
	MissingFields  string
}
