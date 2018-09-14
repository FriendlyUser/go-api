package main

import (
	"database/sql"
)

type jobsearchitem struct {
	ID          int     `json:"id"`
	NumJobs     int     `json:"numjobs"`
	AvgKeywords float64 `json:"avgkeywords"`
	AvgSkills   float64 `json:"avgskills"`
	City        string  `json:"city"`
	SearchTerm  string  `json:"searchterm"`
	SearchTime  string  `json:"searchtime"`
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
	err := db.Exec("INSERT INTO jobinfo(numjobs, avgkeywords,avgskills,city,searchterm,searchtime) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", 
		j.NumJobs, j.AvgKeywords, j.AvgSkills, j.City, j.SearchTerm,
		j.SearchTime).Scan(&j.ID)

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