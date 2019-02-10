package models

import (
	"encoding/json"
	"testing"
)

var testJSON = `{"id":"507f191e810c19729de860ea", "userId":"Me", "name":"Fluffy", "type":"Snake", "breed":"Python", "dateOfBirth":"September", "image":"base64"}`
var testJSONBadID = `{"id":"bad", "userId":"Me", "name":"Fluffy", "type":"Snake", "breed":"Python", "dateOfBirth":"September", "image":"base64"}`

func TestValidPet(t *testing.T) {
	var pet Pet
	err := json.Unmarshal([]byte(testJSON), &pet)

	if err != nil {
		t.Errorf("Unexecpted error: %s", err)
	}

	if pet.UserID != "Me" {
		t.Errorf("Incorrect userId, expected: %s, actual: %s", pet.UserID, "Me")
	}

	if pet.Name != "Fluffy" {
		t.Errorf("Incorrect name, expected: %s, actual: %s", pet.Name, "Fluffy")
	}

	if pet.Type != "Snake" {
		t.Errorf("Incorrect type, expected: %s, actual: %s", pet.Type, "Snake")
	}

	if pet.Breed != "Python" {
		t.Errorf("Incorrect breed, expected: %s, actual: %s", pet.Breed, "Python")
	}

	if pet.DateOfBirth != "September" {
		t.Errorf("Incorrect dateOfBirth, expected: %s, actual: %s", pet.DateOfBirth, "September")
	}

	if pet.Image != "base64" {
		t.Errorf("Incorrect image, expected: %s, actual: %s", pet.Image, "base64")
	}
}
func TestInvalidId(t *testing.T) {
	var pet Pet
	err := json.Unmarshal([]byte(testJSONBadID), &pet)

	if err.Error() != `invalid ObjectId in JSON: "bad"` {
		t.Errorf("Expected error %s, but was %s", `invalid ObjectId in JSON: "bad"`, err)
	}
}
