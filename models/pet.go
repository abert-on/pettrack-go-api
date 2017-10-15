package models

import "gopkg.in/mgo.v2/bson"

// Represents a Pet. Uses bson keyword to tell mgo
// how to name properties in mongodb document
type Pet struct {
	ID bson.ObjectId 'bson:"_id" json:"id"'
	UserId string 'bson:"userId" json:"userId"'
	Name string 'bson:"name" json:"name"'
	Type string 'bson:"type" json:"type"'
	Breed string 'bson:"breed" json:"type"'
	DateOfBirth string 'bson:"dateOfBirth json:"dateOfBirth"'
	Image string 'bson:"image" json:"image"'
}