package dao

import (
	"log"

	. "github.com/abert-on/pettrack-go-api/models"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type PetsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "pets"
)

// Establish a connection to DB
func (p *PetsDAO) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

// Get a list of all pets
func (p *PetsDAO) FindAll() ([]Pet, error) {
	var pets []Pet
	err := db.C(COLLECTION).Find(bson.M{}).All(&pets)
	return pets, err
}

// find a pet by ID
func (p *PetsDAO) FindById(id string) (Pet, error) {
	var pet Pet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&pet)
	return pet, err
}

// insert a pet into the DB
func (p *PetsDAO) Insert(pet Pet) error {
	err := db.C(COLLECTION).Insert(&pet)
	return err
}

// delete an existing pet
func (p *PetsDAO) Delete(pet Pet) error {
	err := db.C(COLLECTION).Remove(&pet)
	return err
}

// update an existing pet
func (p *PetsDAO) Update(pet Pet) error {
	err := db.C(COLLECTION).UpdateId(pet.ID, &pet)
	return err
}
