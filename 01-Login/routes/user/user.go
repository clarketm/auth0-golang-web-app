package user

import (
        "github.com/clarketm/auth0-golang-web-app/01-Login/app"
        templates "github.com/clarketm/auth0-golang-web-app/01-Login/routes"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
