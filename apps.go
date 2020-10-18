package main

import (	
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "meeting/config"
	. "meeting/dao"
	. "meeting/models"
)
import "time"
var config = Config{}
var dao = MeetingsDAO{}

// GET list all meetings
func AllMeetingsEndPoint(w http.ResponseWriter, r *http.Request) {
	meetings, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, meetings)
}

// GET a meeting by its ID
func FindMeetingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meeting, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
		return
	}
	respondWithJson(w, http.StatusOK, meeting)
}

// GET a meeting by  timestamp
func FindMeetingTimeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	
	meeting, err := dao.FindByTime(params["starttime"],params["endtime"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
		return
	}
	respondWithJson(w, http.StatusOK, meeting)
}

// GET a meeting by its email
func FindPatimeetingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meeting, err := dao.FindByPatimeeting(params["emailid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
		return
	}
	respondWithJson(w, http.StatusOK, meeting)
}

// Create new meetings
func CreateMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meeting Meeting
	if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	meeting.ID = bson.NewObjectId()
	meeting.Timestamp=time.Now()
	if err := dao.Insert(meeting); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, meeting)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Eestablish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/meetings", AllMeetingsEndPoint).Methods("GET")
	r.HandleFunc("/meetings", CreateMeetingEndPoint).Methods("POST")
	r.HandleFunc("/meetings/{id}", FindMeetingEndpoint).Methods("GET")
	r.HandleFunc("/meetings/start={starttime}&end={endtime}", FindMeetingTimeEndpoint).Methods("GET")
	r.HandleFunc("/meetings/?participant={emailid}", FindPatimeetingEndpoint).Methods("GET")
	
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
