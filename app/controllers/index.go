package controllers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

/*
type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var dbUsers = make(map[string]user)      // user ID, user
var dbSessions = make(map[string]string) // session ID, user ID
*/

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

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

	var u models.User

	// process form submission
	if r.Method == http.MethodPost {

		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// ########## username taken? ##########
		//if _, ok := dbUsers[un]; ok {
		//	http.Error(w, "Username already taken", http.StatusForbidden)
		//	return
		//}

		user, err := models.GetUserByUserName(un)
		if err == nil {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// ########## username taken? ##########

		// ########## create session ##########
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// ########## create session ##########

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// store user in dbUsers
		u := user{un, bs, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := tpls.ExecuteTemplate(w, "signup.html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
