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

func TestFirstThree(t *testing.T) {
	patterns := []struct {
		want  string
		given string
	}{
		{"abc", "abcde"},
		{"Abc", "Abcde"},
		{"012", "01234"},
	}
	for idx, p := range patterns {
		got := firstThree(p.given)

		if p.want != got {
			t.Errorf("Case(%d) want %v, got %v", idx, p.want, got)
		}
	}
}

func TestViewHRTFHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(viewHRTFHandler))
	defer ts.Close()
	t.Logf("test ViewHRTFHandler")

	req := httptest.NewRequest("GET", ts.URL+"/hrtf/?id=1", nil)
	res := httptest.NewRecorder()
	viewHRTFHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("invalid code. want %v, got %v", http.StatusOK, res.Code)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	t.Log(data)

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
