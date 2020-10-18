package dao

import (
	"log"

	. "meeting/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MeetingsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "meetings"
)

// Establish a connection to database
func (m *MeetingsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list meetings
func (m *MeetingsDAO) FindAll() ([]Meeting, error) {
	var meetings []Meeting
	err := db.C(COLLECTION).Find(bson.M{}).All(&meetings)
	return meetings, err
}

// Find meeting by its ids
func (m *MeetingsDAO) FindById(id string) (Meeting, error) {
	var meeting Meeting
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&meeting)
	return meeting, err
}

// Find a meeting by its start and end time
func (m *MeetingsDAO) FindByTime(startimee string,endtimee string) (Meeting, error) {
	var meeting Meeting
	var st=startimee
	var et=endtimee
	err := db.C(COLLECTION).Find(bson.M{"starttime":bson.M{"$gte": st } , "endtime":bson.M{"$lte":et}}).All(&meeting)
	return meeting, err
}
// Find  meeting by its email
func (m *MeetingsDAO) FindByPatimeeting(emailid string) (Meeting, error) {
	var meeting Meeting
	err := db.C(COLLECTION).Find(bson.M{"Participant ":bson.M{"email": emailid }}).All(&meeting)
	return meeting, err
}

// Insert a meeting into database
func (m *MeetingsDAO) Insert(meeting Meeting) error {
	err := db.C(COLLECTION).Insert(&meeting)
	return err
}

