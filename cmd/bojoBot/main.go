package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Z3DRP/bojoBot/internal/applier"
	"github.com/Z3DRP/bojoBot/internal/bojo"
	"github.com/Z3DRP/bojoBot/internal/dac"
	"github.com/Z3DRP/bojoBot/internal/scheduledrun"
	"github.com/Z3DRP/bojoBot/jprocessor"
	"github.com/go-rod/rod"
)

func createContext(srun scheduledrun.ScheduledRun) (context.Context, context.CancelFunc) {
	if srun.DurationMinutes > 0 {
		timeout := time.Duration(srun.DurationMinutes) * time.Minute
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		return ctx, cancel
	}

	ctx, cancel := context.WithCancel(context.Background())
	return ctx, cancel
}

func main() {
	runArgs := make([]string, 0)
	for _, arg := range os.Args {
		runArgs = append(runArgs, arg)
	}

	dbPath := runArgs[0]
	srunId, err := strconv.ParseInt(runArgs[1], 10, 32)
	if err != nil {
		log.Fatal("Unable to parse schedule run id from caller")
	}

	db, err := sql.Open("sqlite3", dbPath)
	//TODO setup custom log in log package
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	srun, err := dac.GetScheduledRun(db, int(srunId))
	if err != nil {
		log.Fatal("Unable to convert scheduled run id into int from int64")
	}

	ctx, cancel := createContext(*srun)
	defer cancel()

	job, err := dac.GetJobTitle(db, int(srun.JobTitleId))
	if err != nil {
		log.Fatal("Unable to find job title for scheduled run")
	}

	jobBoard, err := dac.GetJobBoard(db, int(srun.JobBoardId))
	if err != nil {
		log.Fatal("Unable to find job board for scheduled run")
	}

	useSubmissionLimit := srun.NumberOfSubmissions > 0
	browser := rod.New().MustConnect()
	defer browser.MustClose()
	search := bojo.NewBojoSearch(ctx, cancel, browser, job, srun.NumberOfSubmissions, useSubmissionLimit)

}
