package controllers

import (
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid"
	"github.com/tetsuzawa/go-3daudio/app/models"
)

type hrtfData struct {
	User models.User
	HRTF *models.HRTF
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func viewHRTFHandler(w http.ResponseWriter, r *http.Request) {
	var id string

	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "academic" {
		http.Error(w, "Sorry. Hrtf database is for academic use only for now.", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		//if posted from form. add data to db
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		id = ulid.MustNew(ulid.Now(), entropy).String()

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		name := r.Form.Get("name")
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		azimuth, err := strconv.Atoi(r.Form.Get("azimuth"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		elevation, err := strconv.Atoi(r.Form.Get("elevation"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data, err := strconv.Atoi(r.Form.Get("data"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		hrtf := models.NewHRTF(id, name, uint(age), float64(azimuth), float64(elevation), float64(data))
		if err = hrtf.Create(); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//save file posted from form
		f, h, err := r.FormFile("q")
		if err != nil {
			log.Println(err)
		} else {
			//TODO complex code
			bs, err := ioutil.ReadAll(f)
			if err != nil {
				log.Println(err)
			}
			f.Close()
			dst, err := os.Create(filepath.Join("./resources/", h.Filename))
			if err != nil {
				log.Println(err)
			}
			_, err = dst.Write(bs)
			if err != nil {
				log.Println(err)
			}
			dst.Close()
		}
	}

	if r.Method == http.MethodGet {
		id = r.URL.Query().Get("id")
		if id == "" {
			//APIError(w, "No id param", http.StatusBadRequest)
			//return
			c, err := r.Cookie("id")
			if err != nil {
				//TODO error handling
				//TODO id hard code
				log.Println(err)
				id = "01DQ44KFF4D44TFZA9963GD1VS"
			} else {
				id = c.Value
			}
		}
	}

	hrtf, err := models.GetHRTF(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "id",
		Value:    id,
		Path:     "/",
		HttpOnly: true,
	})

	//hrtf := models.NewHRTF(1, "tetsu", 20, 20, 0, 0.35555)
	hd := hrtfData{
		User: u,
		HRTF: hrtf,
	}
	err = tpls.ExecuteTemplate(w, "hrtf.html", hd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
