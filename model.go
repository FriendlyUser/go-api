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
// pass in date as string, postgres processes it as date
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

// doc struct
type docs struct {
	ID           int     `json:"id"`
	PublicId     int     `json:"publicid"`
	DocName      string  `json:"docname"`
	DocTag       string  `json:"doctag"`
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

		uvicitems = append(uvicitems, j)
	}

	return uvicitems, nil
}


func (j *uvicjob) getUvic(db *sql.DB) error {
	return db.QueryRow("SELECT jobid, jobtitle,org,pos,loc,numapps,deadline,coop FROM uvic WHERE id=$1", 
		j.ID).Scan(&j.JobId, &j.Title,
			&j.Organization,&j.Position,&j.Location,&j.NumApps,&j.Deadline,&j.Coop)
}

func (j *uvicjob) updateUvic(db *sql.DB) error {
	_, err := db.Exec("UPDATE uvic SET jobid=$1,jobtitle=$2,org=$3,pos=$4,loc=$5,numapps=$6,deadline=$7,coop=$8	 WHERE id=$9", 
		j.JobId, j.Title, j.Organization, j.Position, j.Location, j.NumApps, j.Deadline,j.Coop)

	return err
}

func (j *uvicjob) deleteUvic(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM uvic WHERE id=$1", j.ID)

	return err
}

func (j *uvicjob) createUvic(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO uvic(jobid,jobtitle,org,pos,loc,numapps,deadline,coop) VALUES($1, $2, $3, $4, $5, $6,$7,$8) RETURNING id", j.JobId, j.Title, j.Organization, j.Position, j.Location,j.NumApps,j.Deadline,j.Coop).Scan(&j.ID)
	return err
}

func getUvicItems(db *sql.DB, start, count int) ([]uvicjob, error) {
	rows, err := db.Query("SELECT jobid,jobtitle,org,pos,loc,numapps,deadline,coop FROM uvic LIMIT $1 offset $2", count, start)

	if err != nil {
		return nil, err
	}

	// Will execute at the end of the scope
	defer rows.Close()

	uvicitems := []uvicjob{}

	for rows.Next() {
		var j uvicjob
		if err := rows.Scan(&j.ID, &j.JobId, &j.Title,
			&j.Organization,&j.Position,&j.Location,&j.NumApps,&j.Deadline,&j.Coop); err != nil {
			return nil, err
		}

		uvicitems = append(uvicitems, j)
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

func getAllDocsDB(db *sql.DB) ([]docsItems, error) {
	rows, err := db.Query("SELECT * FROM docs")
    
    if err != nil {
		return nil, err
	}

	// Will execute at the end of the scope
	defer rows.Close()

	docsItems := []docs{}
	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	// var jobsearchitems []jobsearchitem

	for rows.Next() {
		var j docs

		if err := rows.Scan(&j.ID, &j.PublicId, &j.DocName,
			&j.DocTag); err != nil {
			return nil, err
		}

		docsItems = append(docsItems, j)
	}

	return docsItems, nil
}

func (j *docs) getDocItem(db *sql.DB) error {
	return db.QueryRow("SELECT id,public_id,doc_name,doc_tag FROM docs WHERE id=$1", 
		j.ID).Scan(&j.PublicId, &j.DocName, &j.DocTag)
}

func (j *docs) updateDoc(db *sql.DB) error {
	_, err := db.Exec("UPDATE docs SET public_id=$1, doc_name=$2, doc_tag=$3 WHERE id=$4", 
		j.publicId, j.docName, j.docTag, j.ID)

	return err
}

func (j *docs) deleteDoc(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM docs WHERE id=$1", j.ID)

	return err
}

func (j *docs) createDoc(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO docs(public_id,doc_name,doc_tag) VALUES($1, $2, $3) RETURNING id", j.PublicId, j.DocName, j.DocTag).Scan(&j.ID)

	return err
}

// implement later
// func getJobSearchItems(db *sql.DB, start, count int) ([]jobsearchitem, error) {
// 	rows, err := db.Query("SELECT numjobs, avgkeywords, avgskills,city, searchterm, searchtime FROM jobinfo LIMIT $1 offset $2", count, start)

// 	if err != nil {
// 		return nil, err
// 	}

// 	// Will execute at the end of the scope
// 	defer rows.Close()

// 	jobsearchitems := []jobsearchitem{}
// 	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
// 	// var jobsearchitems []jobsearchitem

// 	for rows.Next() {
// 		var j jobsearchitem

// 		if err := rows.Scan(&j.ID, &j.NumJobs, &j.AvgKeywords,
// 			&j.AvgSkills,&j.City,&j.SearchTerm,&j.SearchTime); err != nil {
// 			return nil, err
// 		}

// 		jobsearchitems = append(jobsearchitems, j)
// 	}

// 	return jobsearchitems, nil
// }