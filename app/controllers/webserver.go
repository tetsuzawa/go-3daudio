package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/tetsuzawa/go-3daudio/app/models"
	"github.com/tetsuzawa/go-3daudio/config"
)

//var tpls = template.Must(template.New("").Funcs(fm).ParseFiles("app/views/hrtf.html", "app/views/analysis.html"))
var tpls = template.Must(template.New("").Funcs(fm).ParseGlob("app/views/templates/*.html"))

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
	http.HandleFunc("/", viewIndexHandler)
	http.HandleFunc("/api/sofa/", apiMakeHandler(apiSOFAHandler))
	http.HandleFunc("/hrtf/", viewHRTFHandler)
	http.HandleFunc("/analysis/", viewAnalysisHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
