package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/Sileadim/skilltree_backend/pkg/forms" // New import
	"github.com/Sileadim/skilltree_backend/pkg/models"
)

func (app *application) getTrees(w http.ResponseWriter, r *http.Request) {

	trees, err := app.trees.List()
	fmt.Println(trees)
	if err != nil {
		app.notFound(w)
		return
	}
	var m []map[string]interface{}
	for _, t := range trees {
		mapRepr, err := t.ToMap()
		if err != nil {
			app.serverError(w, err)
		}
		m = append(m, mapRepr)
	}
	byteRepresentation, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}
	print(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteRepresentation)

}

func (app *application) getTree(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	t, err := app.trees.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	byteRepresentation, err := t.ToJSON()
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteRepresentation)

}

func (app *application) createTree(w http.ResponseWriter, r *http.Request) {

	t := map[string]interface{}{}
	fmt.Println(r)
	err := json.NewDecoder(r.Body).Decode(&t)
	fmt.Println(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	content, err := json.Marshal(t["content"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(t)
	id, err := app.trees.Insert(t["name"].(string), t["uuid"].(string), string(content))
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(id)
	fmt.Fprintf(w, "Created tree with id: %v", id)

}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {

	t := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Decoding error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.users.Insert(t["name"].(string), t["email"].(string), t["password"].(string))
	if err != nil {
		fmt.Println(err.Error())
		if errors.Is(err, models.ErrDuplicateEmail) {
			app.duplicateEmailError(w, err)

		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	fmt.Fprintf(w, "Signed up user %v", t["name"].(string))
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Decoding error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(t)
	id, name, err := app.users.Authenticate(t["email"].(string), t["password"].(string))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.incorrectCrendentialsError(w, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := map[string]interface{}{"name": name, "id": id}

	byteRepresentation, err := json.Marshal(data)
	if err != nil {
		app.serverError(w, err)
	}

	fmt.Println(id)
	app.session.Put(r, "authenticatedUserID", id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteRepresentation)

	//fmt.Fprintf(w, "Logged in user %v", t["name"].(string))
	//w.Write([]byte("OK"))

}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	fmt.Fprintf(w, "Logged out")
}
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping!")
	w.Write([]byte("OK"))
}
