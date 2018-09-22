package main

import (
	"database/sql"
)

// indeed job posting struct
type jobsearchitem struct {
	ID          int     `json:"id"`
	NumJobs     int     `json:"numjobs"`
	AvgKeywords float64 `json:"avgkeywords"`
	AvgSkills   float64 `json:"avgskills"`
	City        string  `json:"city"`
	SearchTerm  string  `json:"searchterm"`
	SearchTime  string  `json:"searchtime"`
}

// uvic job posting format
// id is the classic serial id
// jobid is used in the uvic system
//pass in date as string, postgres processes it as date
type uvicjob struct {
	ID           int     `json:"id"`
	JobId        int     `json:"jobid"`
	Title        string  `json:"jobtitle"`
	Organization string  `json:"org"`
	Position     string  `json:"pos"`
	Location     string  `json:"loc"`
	NumApps      string  `json:"numapps"`
	Deadline     string  `json:"deadline"`
	Coop         bool    `json:"coop"`
}

func getAllUvic(db *sql.DB) ([]uvicjob, error) {
	rows, err := db.Query("SELECT * FROM uvic")
    
    if err != nil {
		return nil, err
	}

	// Will execute at the end of the scope
	defer rows.Close()

	uvicitems := []uvicjob{}
	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	// var jobsearchitems []jobsearchitem

	for rows.Next() {
		var j uvicjob

		if err := rows.Scan(&j.ID, &j.JobId, &j.Title,
			&j.Organization,&j.Position,&j.Location,&j.NumApps,&j.Deadline,&j.Coop); err != nil {
			return nil, err
		}

		uvicitems = append(uvicjob, j)
	}

	return uvicitems, nil
}


func getAllJobsSearch(db *sql.DB) ([]jobsearchitem, error) {
	rows, err := db.Query("SELECT * FROM jobinfo")
    
    if err != nil {
		return nil, err
	}

	// Will execute at the end of the scope
	defer rows.Close()

	jobsearchitems := []jobsearchitem{}
	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	// var jobsearchitems []jobsearchitem

	for rows.Next() {
		var j jobsearchitem

		if err := rows.Scan(&j.ID, &j.NumJobs, &j.AvgKeywords,
			&j.AvgSkills,&j.City,&j.SearchTerm,&j.SearchTime); err != nil {
			return nil, err
		}

		jobsearchitems = append(jobsearchitems, j)
	}

	return jobsearchitems, nil
}

func (j *jobsearchitem) getJobSearchItem(db *sql.DB) error {
	return db.QueryRow("SELECT numjobs, avgkeywords, avgskills,city, searchterm, searchtime FROM jobinfo WHERE id=$1", 
		j.ID).Scan(&j.NumJobs, &j.AvgKeywords, &j.AvgSkills, &j.City, &j.SearchTerm, &j.SearchTime)
}

func (j *jobsearchitem) updateJobSearchItem(db *sql.DB) error {
	_, err := db.Exec("UPDATE jobinfo SET numjobs=$1, avgkeywords=$2, avgskills=$3, city=$4, searchterm=$5, searchtime=$6,	 WHERE id=$7", 
		j.NumJobs, j.AvgKeywords, j.AvgSkills, j.City, j.SearchTerm, j.SearchTime, j.ID)

	return err
}

func (j *jobsearchitem) deleteJobSearchItem(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM jobinfo WHERE id=$1", j.ID)

	return err
}

func (j *jobsearchitem) createJobSearchItem(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO jobinfo(numjobs, avgkeywords,avgskills,city,searchterm,searchtime) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", j.NumJobs, j.AvgKeywords, j.AvgSkills, j.City, j.SearchTerm,j.SearchTime).Scan(&j.ID)

	return err
}

func getJobSearchItems(db *sql.DB, start, count int) ([]jobsearchitem, error) {
	rows, err := db.Query("SELECT numjobs, avgkeywords, avgskills,city, searchterm, searchtime FROM jobinfo LIMIT $1 offset $2", count, start)

	if err != nil {
		return nil, err
	}

	// Will execute at the end of the scope
	defer rows.Close()

	jobsearchitems := []jobsearchitem{}
	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	// var jobsearchitems []jobsearchitem

	for rows.Next() {
		var j jobsearchitem

		if err := rows.Scan(&j.ID, &j.NumJobs, &j.AvgKeywords,
			&j.AvgSkills,&j.City,&j.SearchTerm,&j.SearchTime); err != nil {
			return nil, err
		}

		jobsearchitems = append(jobsearchitems, j)
	}

	return jobsearchitems, nil
}