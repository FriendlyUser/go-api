package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App : application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}


// Create Table Query
// search time should be date, but since I control the inputted data, text is fine.
const prodTable = `CREATE TABLE IF NOT EXISTS jobinfo
(
    id SERIAL,
    numjobs INT,
    avgkeywords NUMERIC(5,2) NOT NULL DEFAULT 0.00,
    avgskills NUMERIC(5,2) NOT NULL DEFAULT 0.00,
    city TEXT NOT NULL,
    searchterm TEXT NOT NULL,
    searchtime DATE NOT NULL,
    CONSTRAINT jobinfo_pkey PRIMARY KEY (id)
)`
// coop postings and career center postings
const uvicJobQuery = `CREATE TABLE IF NOT EXISTS uvic
(
	id SERIAL,
	jobid INT,
	jobtitle TEXT NOT NULL,
	org TEXT NOT NULL,
	pos TEXT NOT NULL ,
	loc TEXT NOT NULL,
	numapps INT,
	deadline DATE NOT NULL,
	coop BOOL DEFAULT TRUE,
	CONSTRAINT uvic_pkey PRIMARY KEY (id)
)`

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Indeed Job Postings 

// get a job search item
func (app *App) getJobSearchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	j := jobsearchitem{ID: id}

	if err := j.getJobSearchItem(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Job Search Item not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, j)
}

func (app *App) getAllJobs(w http.ResponseWriter, r *http.Request) {

	jobSearchItems, err := getAllJobsSearch(app.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, jobSearchItems)
}

func (app *App) getJobSearchItems(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	//if count < 1 || count > 10 {
	//	count = 10
	//}

	if start < 0 {
		start = 0
	}

	jobSearchItems, err := getJobSearchItems(app.DB, start, count)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, jobSearchItems)
}

func (app *App) createJobSearchItem(w http.ResponseWriter, r *http.Request) {
	var j jobsearchitem
	decoder := json.NewDecoder(r.Body)
    fmt.Printf("%v\n", r.Body)
	if err := decoder.Decode(&j); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// defer : will be executed when the scope ends
	defer r.Body.Close()

	if err := j.createJobSearchItem(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, j)
}

func (app *App) updateJobSearchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var j jobsearchitem
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&j); err != nil {
		respondWithError(w, http.StatusBadGateway, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	j.ID = id

	if err := j.updateJobSearchItem(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, j)
}

func (app *App) deleteJobSearchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	j := jobsearchitem{ID: id}
	if err := j.deleteJobSearchItem(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	
	entry := "client/dist/index.html"

	// open and parse a template text file
	// TODO: Refactor and clean
	if tpl, err := template.New("index").ParseFiles(entry); err != nil {
		log.Fatal(err)
	} else {
		tpl.Lookup("index").ExecuteTemplate(w, "index.html", nil)
	}
	
}

// UVIC Routes

func (app *App) getAllUvic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	//j := jobsearchitem{ID: id}
	//if err := j.deleteJobSearchItem(app.DB); err != nil {
	//	respondWithError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *App) getUvic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	//j := jobsearchitem{ID: id}
	//if err := j.deleteJobSearchItem(app.DB); err != nil {
	//	respondWithError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


func (app *App) createUvic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	//j := jobsearchitem{ID: id}
	//if err := j.deleteJobSearchItem(app.DB); err != nil {
	//	respondWithError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


func (app *App) getUvicItems(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	//if count < 1 || count > 10 {
	//	count = 10
	//}

	if start < 0 {
		start = 0
	}

	uvicItems, err := getUvicItems(app.DB, start, count)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, uvicItems)
}

func (app *App) updateUvic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid jobsearchitem ID")
		return
	}

	j := uvicjob{ID: id}
	if err := j.updateUvic(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *App) deleteUvic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid uvicjob ID")
		return
	}

	j := uvicjob{ID: id}
	if err := j.deleteUvic(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *App) initializeRoutes() {
	static := "./client/dist/static/"

	api := app.Router.PathPrefix("/api/").Subrouter()
	// indeed postings
    api.HandleFunc("/jobs", app.getAllJobs).Methods("GET")
	api.HandleFunc("/jobs/{start:[0-9]+}/{count:[0-9]+}", app.getJobSearchItems).Methods("GET")
	api.HandleFunc("/jobs", app.createJobSearchItem).Methods("POST")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.getJobSearchItem).Methods("GET")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.updateJobSearchItem).Methods("PUT")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.deleteJobSearchItem).Methods("DELETE")

	// uvic postings
	api.HandleFunc("/uvic", app.getAllUvic).Methods("GET")
	api.HandleFunc("/uvic/{start:[0-9]+}/{count:[0-9]+}", app.getUvicItems).Methods("GET")
	api.HandleFunc("/uvic", app.createUvic).Methods("POST")
	api.HandleFunc("/uvic/{id:[0-9]+}", app.getUvic).Methods("GET")
	api.HandleFunc("/uvic/{id:[0-9]+}", app.updateUvic).Methods("PUT")
	api.HandleFunc("/uvic/{id:[0-9]+}", app.deleteUvic).Methods("DELETE")

	// serving front-end 
	app.Router.PathPrefix("/dist/").Handler(http.FileServer(http.Dir(static)))
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))

	app.Router.HandleFunc("/", app.indexHandler).Methods("GET")
}

// Initialize : connects to postgresql
func (app *App) Initialize(DATABASE_URL string) {

	var err error
	app.DB, err = sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
    // same as tableCreationQuery, could just make the app reload all the time
    if _, err := app.DB.Exec(prodTable); err != nil {
		log.Fatal(err)
	}
	 
	if _, err := app.DB.Exec(uvicJobQuery); err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run : runs the app
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
