package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/tetsuzawa/go-3daudio/config"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("app/views/hrtf.html"))

func viewHRTFHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "hrtf.html", nil)
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

	productCode := r.URL.Query().Get("id")
	if productCode == "" {
		APIError(w, "No id param", http.StatusBadRequest)
		return
	}
}

func StartWebServer() error {
	//http.HandleFunc("/api/sofa/", apiMakeHandler(apiSOFAHandler))
	http.HandleFunc("/hrtf/", viewHRTFHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
