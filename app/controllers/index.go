package controllers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c1 = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c1)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c1.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c1.Value] = un
		dbUsers[un] = u
	}

	// visit count
	/*
		c2, err := r.Cookie("visit-count")
		if err == http.ErrNoCookie {
			c2 = &http.Cookie{
				Name:  "visit-count",
				Value: "0",
			}
		}
		cnt, err := strconv.Atoi(c2.Value)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		cnt++
		c2.Value = strconv.Itoa(cnt)
		http.SetCookie(w, c2)
	*/

	err = tpls.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
