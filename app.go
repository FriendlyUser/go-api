package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//"html/template"
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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

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

func (app *App) getJobSearchItems(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count < 1 || count > 10 {
		count = 10
	}

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
	/**
	entry := "client/dist/index.html"

	// open and parse a template text file
	// TODO: Refactor and clean
	if tpl, err := template.New("index").ParseFiles(entry); err != nil {
		log.Fatal(err)
	} else {
		tpl.Lookup("index").ExecuteTemplate(w, "index.html", nil)
	}
	*/
}

func (app *App) initializeRoutes() {
	static := "./client/dist/static/"

	api := app.Router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/jobs", app.getJobSearchItems).Methods("GET")
	api.HandleFunc("/jobs", app.createJobSearchItem).Methods("POST")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.getJobSearchItem).Methods("GET")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.updateJobSearchItem).Methods("PUT")
	api.HandleFunc("/jobs/{id:[0-9]+}", app.deleteJobSearchItem).Methods("DELETE")

	//app.Router.PathPrefix("/dist/").Handler(http.FileServer(http.Dir(static)))
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))

	app.Router.HandleFunc("/", app.indexHandler).Methods("GET")
}

// Initialize : connects to postgresql
func (app *App) Initialize(connectionString string) {

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run : runs the app
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
