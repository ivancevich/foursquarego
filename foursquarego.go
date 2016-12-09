package foursquarego

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FoursquareApi struct {
	clientID     string
	clientSecret string
	queryQueue   chan query
	HttpClient   *http.Client
}

type query struct {
	url         string
	form        url.Values
	data        *foursquareResponse
	method      int
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

type Omit interface{}

const (
	API_URL = "https://api.foursquare.com/v2/"
	VERSION = "20150813"
	MODE    = "foursquare"
	_GET    = iota
	_POST   = iota
)

func NewFoursquareApi(clientID string, clientSecret string) *FoursquareApi {
	queue := make(chan query)
	a := &FoursquareApi{
		clientID:     clientID,
		clientSecret: clientSecret,
		queryQueue:   queue,
		HttpClient:   http.DefaultClient,
	}
	go a.throttledQuery()
	return a
}

func (a *FoursquareApi) apiGet(urlStr string, form url.Values, data *foursquareResponse) error {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = form.Encode()
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResp apiResponse
	if err := json.Unmarshal(contents, &apiResp); err != nil {
		return err
	}

	if resp.StatusCode != 200 || apiResp.Meta.Code != 200 {
		return newApiError(resp, apiResp.Meta)
	}

	*data = apiResp.Response
	return nil
}

func cleanValues(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}

func (a *FoursquareApi) execQuery(urlStr string, form url.Values, data *foursquareResponse, method int) error {
	form.Set("v", VERSION)
	form.Set("m", MODE)
	if form.Get("oauth_token") == "" {
		form.Set("client_id", a.clientID)
		form.Set("client_secret", a.clientSecret)
	}
	switch method {
	case _GET:
		return a.apiGet(urlStr, form, data)
	default:
		return errors.New("HTTP method not supported")
	}
	return errors.New("ack")
}

func (a *FoursquareApi) throttledQuery() {
	for q := range a.queryQueue {
		err := a.execQuery(q.url, q.form, q.data, q.method)
		q.response_ch <- response{q.data, err}
	}
}

func (a *FoursquareApi) Close() {
	close(a.queryQueue)
}
