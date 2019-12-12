package controllers

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/tetsuzawa/go-3daudio/web-app/app/models"
)

var dbSessionsCleaned time.Time

func getUser(w http.ResponseWriter, r *http.Request) models.UserServicer {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			Path:     "/",
			HttpOnly: true,
		}
	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	//var u models.UserServicer

	s, err := models.GetSession(c.Value)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get session"))
		return models.UserServicer{}
	}

	u, err := models.GetUserByUserName(s.UserName)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get user"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("got user in controller:", u)
	return *u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		return false
	}
	s, err := models.GetSession(c.Value)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = models.GetUserByUserName(s.UserName)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func cleanSessions() {
	t := time.Now()
	t.Add(-30 * time.Second)
	ss, err := models.GetOldSessions(t)
	if err != nil {
		log.Println(err)
	}
	dbSessionsCleaned = time.Now()

	for _, s := range ss {
		if err := s.Delete(); err != nil {
			log.Println(err)
		}
	}
}
