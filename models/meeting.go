package models

import "gopkg.in/mgo.v2/bson"
import "time"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Participant struct {
	Name         string          `bson:"name" json:"name"`
	Email        string          `bson:"email" json:"email"`
	RSVP         string          `bson:"rsvp" json:"rsvp"`
}
type Meeting struct {
	ID            bson.ObjectId   `bson:"_id" json:"id"`
	Title         string          `bson:"title" json:"title"`
	Participants  Participant         `bson:"participants{}" json:"participants{}"`
	StartTime     time.Time          `bson:"starttime" json:"starttime"`
	EndTime       time.Time           `bson:"endtime" json:"endtime"`
	Timestamp     time.Time       `bson:"timestamp" json:"timestamp"`
}


