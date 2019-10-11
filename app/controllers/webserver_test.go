package controllers

import (
	"bytes"
	"github.com/tetsuzawa/go-3daudio/app/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestViewHRTFHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(viewHRTFHandler))
	defer ts.Close()
	t.Logf("test ViewHRTFHandler")

	//req := httptest.NewRequest("GET", ts.URL+"/hrtf/?id=1", nil)
	req := httptest.NewRequest("GET", ts.URL+"/hrtf/", nil)
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	req.RequestURI = ""

	c := http.DefaultClient

	t.Logf("log. header: %v", req.Header)
	t.Logf("log. body: %v", req.Body)
	t.Logf("log. requestURI: %v", req.RequestURI)
	res, err := c.Do(req)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	t.Log(data)

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Error by res.StatusCode. want %v, got %v", http.StatusOK, status)
	}

	tpls := template.Must(template.New("").Funcs(fm).ParseFiles("app/views/hrtf.html", "app/views/analysis.html"))
	hrtf, err := models.GetHRTF("1")
	if err != nil {
		t.Fatalf("Error by models.GetHRTF(). %v", err)
	}

	var buff = new(bytes.Buffer)

	err = tpls.ExecuteTemplate(buff, "hrtf.html", hrtf)
	if err != nil {
		t.Fatalf("Error by tpls.ExecuteTemplate(). %v", err)
	}

	want := buff.String()

	if string(data) != want {
		t.Errorf("Error by rr.Body.String().\n \nwant \n%v\n,\n \ngot \n%v", want, string(data))
	}

}
