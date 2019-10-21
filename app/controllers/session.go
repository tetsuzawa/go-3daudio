package controllers

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/tetsuzawa/go-3daudio/app/models"
)

func getUser(w http.ResponseWriter, r *http.Request) *models.User {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u *models.User

	//if un, ok := dbSessions[c.Value]; ok {
	//	u = dbUsers[un]
	//}

	s, err := models.GetSession(c.Value)
	if err != nil {
		u, err = models.GetUserByUserName(s.UserName)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	//un := dbSessions[c.Value]
	//_, ok := dbUsers[un]
	//return ok
	s, err := models.GetSession(c.Value)
	if err != nil {
		log.Println(err)
	}

	_, err = models.GetUserByUserName(s.UserName)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
