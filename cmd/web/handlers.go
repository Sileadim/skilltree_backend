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

	//trees, err := app.trees.List()

	m := map[string]interface{}{}
	fmt.Println(m)
	//for t, i := range(trees){
	//
	//
	//
	//}

}

func (app *application) getTree(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	t, err := app.trees.Get(id)
	fmt.Println(t)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	byteRepresentation, err := t.ToJSON()
	print("byteRepresentation")
	fmt.Println(byteRepresentation)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteRepresentation)

}

func (app *application) createTree(w http.ResponseWriter, r *http.Request) {

	t := map[string]interface{}{}
	fmt.Println(r)
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
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

	id, err := app.trees.Insert(t["title"].(string), t["uuid"].(string), string(content))
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(id)
	fmt.Fprintf(w, "Created tree with id: %v", id)

}

//func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
//    err := r.ParseForm()
//    if err != nil {
//        app.clientError(w, http.StatusBadRequest)
//        return
//    }
//
//    form := forms.New(r.PostForm)
//    form.Required("name", "email", "password")
//    form.MaxLength("name", 255)
//	form.MaxLength("email", 255)
//    form.MatchesPattern("email", forms.EmailRX)
//    form.MinLength("password", 10)
//
//    //if !form.Valid() {
//    //    app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
//    //    return
//    //}
//
//    // Try to create a new user record in the database. If the email already exists
//    // add an error message to the form and re-display it.
//    err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
//    if err != nil {
//        if errors.Is(err, models.ErrDuplicateEmail) {
//            form.Errors.Add("email", "Address is already in use")
//            app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
//        } else {
//            app.serverError(w, err)
//        }
//        return
//    }
//
//    // Otherwise add a confirmation flash message to the session confirming that
//    // their signup worked and asking them to log in.
//    app.session.Put(r, "flash", "Your signup was successful. Please log in.")
//
//    // And redirect the user to the login page.
//    http.Redirect(w, r, "/user/login", http.StatusSeeOther)
//}
//
//
//func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
//    err := r.ParseForm()
//    if err != nil {
//        app.clientError(w, http.StatusBadRequest)
//        return
//    }
//
//    // Check whether the credentials are valid. If they're not, add a generic error
//    // message to the form failures map and re-display the login page.
//    form := forms.New(r.PostForm)
//    id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
//    if err != nil {
//        if errors.Is(err, models.ErrInvalidCredentials) {
//            form.Errors.Add("generic", "Email or Password is incorrect")
//            app.render(w, r, "login.page.tmpl", &templateData{Form: form})
//        } else {
//            app.serverError(w, err)
//        }
//        return
//    }
//
//    // Add the ID of the current user to the session, so that they are now 'logged
//    // in'.
//    app.session.Put(r, "authenticatedUserID", id)
//
//    // Redirect the user to the create snippet page.
//    http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
//}
//
//func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
//    // Remove the authenticatedUserID from the session data so that the user is
//    // 'logged out'.
//    app.session.Remove(r, "authenticatedUserID")
//    // Add a flash message to the session to confirm to the user that they've been
//    // logged out.
//    app.session.Put(r, "flash", "You've been logged out successfully!")
//    http.Redirect(w, r, "/", http.StatusSeeOther)
//}
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
