package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	// Add a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Tree struct {
	ID int `json:"id"`
	//Owner   int
	Title   string    `json:"title"`
	Uuid    string    `json:"uuid"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

func (tree *Tree) ToJSON() ([]byte, error) {

	m, err := tree.ToMap()
	if err != nil {
		return nil, err
	}
	fmt.Println(m)
	byteRepresentation, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return byteRepresentation, nil

}

func (tree *Tree) ToMap() (map[string]interface{}, error) {

	m := map[string]interface{}{}
	m["id"] = tree.ID
	m["title"] = tree.Title
	m["uuid"] = tree.Uuid
	fmt.Println(m)
	m["content"] = &map[string]interface{}{}
	err := json.Unmarshal([]byte(tree.Content), m["content"])
	if err != nil {
		return nil, err
	}
	return m, nil
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
