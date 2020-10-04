package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, app.authenticate)
	mux := pat.New()
	mux.Post("/tree/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createTree))
	mux.Get("/tree/:id", dynamicMiddleware.ThenFunc(app.getTree))
	mux.Get("/list", dynamicMiddleware.ThenFunc(app.getTrees))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))
	mux.Get("/ping", dynamicMiddleware.ThenFunc(ping))
	return standardMiddleware.Then(mux)
}
