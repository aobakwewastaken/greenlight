package main

import (
	"fmt"
	"net/http"
)

func (app *Application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *Application) errResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"

	app.errResponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request) {

	message := "the requested resource could not be found"

	app.errResponse(w, r, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {

	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)

	app.errResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *Application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {

	app.errResponse(w, r, http.StatusUnprocessableEntity, errors)
}
