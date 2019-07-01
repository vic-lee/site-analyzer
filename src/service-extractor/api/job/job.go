package job

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TVSeriesExtractionJob describes a TVSeries data extraction job.
// There are 3 possible statuses: Ready | Processing | Completed
type TVSeriesExtractionJob struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Jobs container
var jobs []TVSeriesExtractionJob

// Routes return a router with routes associated with TVJobs
func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/jobs", getJobs).Methods("GET")
	r.HandleFunc("/jobs/{id}", getJob).Methods("GET")
	r.HandleFunc("/jobs", postJob).Methods("POST")

	var PORT = 4000
	fmt.Printf("Router running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
	return r
}

// homeHandler responds a helper message that directs client to use
// the actual API at `/jobs`.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&map[string]interface{}{
		"Message": "This is the home route of the extractor service API. " +
			"To make RESTful calls visit `/jobs.`",
	})
}

// getJob handles querying job with a specific ID.
func getJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range jobs {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			break
		}
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&map[string]interface{}{
		"Message": "No job exists yet.",
	})
}

// getJobs reponds all jobs in the system.
func getJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(jobs) == 0 {
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"Message": "No job exists yet.",
		})
	} else {
		json.NewEncoder(w).Encode(jobs)
	}
}

// postJob creates a new job, if a job with the same Name does not already exist.
func postJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if name, ok := r.URL.Query()["name"]; !ok {

	} else {
		for _, item := range jobs {
			if item.Name == name[0] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		j := TVSeriesExtractionJob{
			ID:     rand.Intn(1000000000),
			Name:   name[0],
			Status: "Ready",
		}
		jobs = append(jobs, j)
		json.NewEncoder(w).Encode(&j)
	}
}
