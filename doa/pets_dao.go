package dao

import (
	"log"

	. "github.com/abert-on/pets-restapi/models"

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

func (p *PetsDAO) connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

func (p *PetsDAO) FindAll() ([]Pet, error) {
	var pets []Pet
	err := db.C(COLLECTION).Find(bson.M{}).All(&pets)
	return pets, err
}

func (p *PetsDAO) FindById(id string) (Pet, error) {
	var pet Pet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&pet)
	return pet, err
}

func (p *PetsDAO) Insert(pet Pet) error {
	err := db.C(COLLECTION).Insert(&pet)
	return err
}

func (p *PetsDAO) Delete(pet Pet) error {
	err := db.C(COLLECTION).Remove(&pet)
	return err
}

func (p *PetsDAO) Update(pet Pet) error {
	err := db.C(COLLECTION).UpdateId(pet.ID, &pet)
	return err
}
