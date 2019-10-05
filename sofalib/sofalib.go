package sofalib

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "http://hoge.fuga.com/fuge/"

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	APIClient := &APIClient{key, secret, &http.Client{}}
	return APIClient
}

func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), (10))
	log.Println(timestamp)
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	apiURL, err := url.Parse(urlPath)
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type HRTF struct {
	Name string
	Age  string
	Data map[float64]float64
}

func (api *APIClient) GetHRTF() ([]HRTF, error) {
	url := "hoge/gethrtf"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)

	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetHRTF err=%s", err.Error())
		return nil, err
	}
	var hrtf []HRTF
	err = json.Unmarshal(resp, &hrtf)
	if err != nil {
		log.Printf("action=GetHRTF err=%s", err.Error())
		return nil, err
	}
	return hrtf, nil
}
