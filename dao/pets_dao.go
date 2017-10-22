package dao

import (
	"log"

	models "github.com/abert-on/pettrack-go-api/models"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// PetsDAO represents the database server
type PetsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION represents the database collection to use
const (
	COLLECTION = "pets"
)

// Connect establishes a connection to DB
func (p *PetsDAO) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

// FindAllByUserID gets a list of all pets for a user
func (p *PetsDAO) FindAllByUserID(userID string) ([]models.Pet, error) {
	var pets []models.Pet
	err := db.C(COLLECTION).Find(bson.M{"$eq": bson.M{"userId": userID}}).All(&pets)
	return pets, err
}

// FindByID finds a pet by ID
func (p *PetsDAO) FindByID(id string) (models.Pet, error) {
	var pet models.Pet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&pet)
	return pet, err
}

// Insert inserts a pet into the DB
func (p *PetsDAO) Insert(pet models.Pet) error {
	err := db.C(COLLECTION).Insert(&pet)
	return err
}

// Delete deletes an existing pet
func (p *PetsDAO) Delete(pet models.Pet) error {
	err := db.C(COLLECTION).Remove(&pet)
	return err
}

// Update updates an existing pet
func (p *PetsDAO) Update(pet models.Pet) error {
	err := db.C(COLLECTION).UpdateId(pet.ID, &pet)
	return err
}
