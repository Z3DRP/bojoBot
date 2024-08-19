package dac

import (
	"database/sql"
	"fmt"

	"github.com/Z3DRP/bojoBot/internal/application"
	"github.com/Z3DRP/bojoBot/internal/completedrun"
	"github.com/Z3DRP/bojoBot/internal/failedsubmission"
	"github.com/Z3DRP/bojoBot/internal/job"
	"github.com/Z3DRP/bojoBot/internal/jobboard"
	"github.com/Z3DRP/bojoBot/internal/resume"
	"github.com/Z3DRP/bojoBot/internal/scheduledrun"
)

func GetScheduledRun(db *sql.DB, sid int) (*scheduledrun.ScheduledRun, error) {
	row := db.QueryRow("SELECT id, name, creationDate, jobTitleId, jobBoardId, runDay, runDayOfWeek, runMonth, runTime, runType, recurring, easyApplyOnly, durrationMinutes, numberOfSubmissions, everyHour, everyMin WHEREE Id=?", sid)

	var run scheduledrun.ScheduledRun
	err := row.Scan(
		&run.Id,
		&run.Name,
		&run.CreationDate,
		&run.JobTitleId,
		&run.JobBoardId,
		&run.RunDay,
		&run.RunDayOfWeek,
		&run.RunMonth,
		&run.RunTime,
		&run.RunType,
		&run.Recurring,
		&run.EasyApplyOnly,
		&run.DurationMinutes,
		&run.NumberOfSubmissions,
		&run.EveryHour,
		&run.EveryMin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no results")
		}
		return nil, err
	}

	return &run, nil
}

func GetResume(db *sql.DB, rid int) (*resume.Resume, error) {
	row := db.QueryRow("SELECT id, name, jobTitleId, filePath FROM Resumes WHERE Id=?", rid)

	var resume resume.Resume
	switch err := row.Scan(
		&resume.Id,
		&resume.Name,
		&resume.JobTitleId,
		&resume.FilePath); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("no results")
	case nil:
		return &resume, nil
	default:
		return nil, err
	}
}

func GetJobBoard(db *sql.DB, jid int) (*jobboard.JobBoard, error) {
	row := db.QueryRow("SELECT id, name, url, hasEasyApply FROM JobBoards WHERE id=?", jid)

	var board jobboard.JobBoard
	err := row.Scan(
		&board.Id,
		&board.Name,
		&board.Url,
		&board.HasEasyApply)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no results")
		}
		return nil, err
	}
	return &board, nil
}

func GetJobTitle(db *sql.DB, jid int) (*job.JobTitle, error) {
	row := db.QueryRow("SELECT id, name, experienceLevel, experienceYears FROM JobTitles WHERE id=?", jid)

	var title job.JobTitle
	err := row.Scan(&title.Id, &title.Name, &title.ExperienceLevel, &title.ExperienceYears)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no results")
		}
		return nil, err
	}
	return &title, nil
}

func InsertCompletedRun(db *sql.DB, crun completedrun.CompletedRun) (int, error) {
	insertStmt := `INSERT INTO CompletedRuns (executionDate, start, finish, applicationsSubmitted, failedSubmissions, runId) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rslt, err := stmt.Exec(crun.ExecutionDate, crun.Start, crun.Finish, crun.ApplicationsSubmitted, crun.FailedSubmissions, crun.RunId)
	if err != nil {
		return 0, err
	}

	insertedId, err := rslt.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(insertedId), nil
}

func InsertApplication(db *sql.DB, app application.Application) (int, error) {
	insertStmt := `INSERT INTO Applications (company, jobTitleId, jobBoardId, location, pay, applyDate, url, runId) values(?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rslt, err := stmt.Exec(app.Company, app.JobTitleId, app.JobBoardId, app.Location, app.Pay, app.ApplyDate, app.Url, app.RunId)
	if err != nil {
		return 0, err
	}

	insertId, err := rslt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertId), nil
}

func InsertFailedSubmission(db *sql.DB, fsub failedsubmission.FailedSubmission) (int, error) {
	insertStmt := `INSERT INTO FailedSubmissions (completedRunId, applicationUrl, missingFields) VALUES(?, ?, ?)`
	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return 0, nil
	}
	defer stmt.Close()

	rslt, err := stmt.Exec(fsub.CompletedRunId, fsub.ApplicationUrl, fsub.MissingFields)
	if err != nil {
		return 0, err
	}

	insertId, err := rslt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertId), nil
}
