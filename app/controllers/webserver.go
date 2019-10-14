package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"github.com/tetsuzawa/go-3daudio/config"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

//var tpls = template.Must(template.New("").Funcs(fm).ParseFiles("app/views/hrtf.html", "app/views/analysis.html"))
var tpls = template.Must(template.New("").Funcs(fm).ParseGlob("app/views/*.html"))

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func viewHRTFHandler(w http.ResponseWriter, r *http.Request) {
	var id string

	if r.Method == http.MethodPost {
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

		f, h, err := r.FormFile("q")
		if err != nil {
			log.Println(err)
		} else {
			//TODO complex code
			func() {
				defer f.Close()
				bs, err := ioutil.ReadAll(f)
				if err != nil {
					log.Println(err)
				}
				dst, err := os.Create(filepath.Join("./resources/", h.Filename))
				if err != nil {
					log.Println(err)
				}
				defer dst.Close()
				_, err = dst.Write(bs)
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}

	if r.Method == http.MethodGet {
		id = r.URL.Query().Get("id")
		if id == "" {
			//APIError(w, "No id param", http.StatusBadRequest)
			//return
			id = "01DQ44KFF4D44TFZA9963GD1VS"
			//TODO id hard code
		}
	}

	hrtf, err := models.GetHRTF(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	//hrtf := models.NewHRTF(1, "tetsu", 20, 20, 0, 0.35555)
	err = tpls.ExecuteTemplate(w, "hrtf.html", hrtf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	err := tpls.ExecuteTemplate(w, "analysis.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{
		Error: errMessage,
		Code:  code,
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/sofa/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiSOFAHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		APIError(w, "No id param", http.StatusBadRequest)
		return
	}
	df, err := models.GetHRTF(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/api/sofa/", apiMakeHandler(apiSOFAHandler))
	http.HandleFunc("/hrtf/", viewHRTFHandler)
	http.HandleFunc("/analysis/", viewAnalysisHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
