package application

import (
	"time"
)

type Application struct {
	Id         int
	Company    string
	JobTitleId int
	JobBoardId int
	Location   string
	Pay        float32
	ApplyDate  time.Time
	RunId      int
	Url        string
}
