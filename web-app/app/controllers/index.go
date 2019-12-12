package controllers

import (
	"github.com/tetsuzawa/go-3daudio/web-app/app/models"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	err := tpls.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewSignupHandler(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u models.UserServicer

	// process form submission
	if r.Method == http.MethodPost {

		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		ro := r.FormValue("role")

		// ########## username taken? ##########
		_, err := models.GetUserByUserName(un)
		if err == nil {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// ########## username taken? ##########

		// ########## create session ##########
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, c)

		s := models.NewSession(c.Value, un, time.Now())
		if err = s.Create(); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// ########## create session ##########

		// ########## create user id ##########
		uID := uuid.NewV4()
		// ########## create user id ##########

		// ########## store user in dbUsers ##########
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := models.NewUser(uID.String(), un, string(bs), f, l, ro)
		if err = u.Create(); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// ########## store user in dbUsers ##########

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := tpls.ExecuteTemplate(w, "signup.html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
