package controllers

import (
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/tetsuzawa/go-3daudio/web-app/app/models"
	"golang.org/x/crypto/bcrypt"
)

func viewLoginHandler(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// is there a username?
		//u, ok := dbUsers[un]
		u, err := models.GetUserByUserName(un)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		s := models.NewSession(c.Value, un, time.Now())
		err = s.Create()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := tpls.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func viewLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
	}
	// delete the session
	//delete(dbSessions, c.Value)
	s, err := models.GetSession(c.Value)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := s.Delete(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// remove the cookie
	c = &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
