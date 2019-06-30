package httpclient

import (
	"encoding/json"
	"github.com/sinhashubham95/shareO-server/cache"
	"io"
	"io/ioutil"
	"net/http"
	netURL "net/url"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 5,
	}
}

// GET request
func GET(url string, headerParams, queryParams map[string]string, cacheEnabled bool,
	data interface{}) error {
	totalURL := addQueryParameters(url, queryParams)
	// check in the cache if cache is enabled
	if cacheEnabled {
		err := cache.GetInterface(totalURL, data)
		if err == nil {
			return nil
		}
	}
	request, err := http.NewRequest("GET", totalURL, nil)
	if err != nil {
		return err
	}
	addHeaders(request, headerParams)
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	err = getIODataAsInterface(response.Body, data)
	if err != nil {
		return err
	}
	// save in cache if cache is enabled
	if cacheEnabled {
		cache.SetInterface(totalURL, data)
	}
	return nil
}

func addQueryParameters(url string, queryParameters map[string]string) string {
	if queryParameters == nil {
		return url
	}
	var sb strings.Builder
	sb.WriteString(url)
	if len(queryParameters) > 0 {
		sb.WriteString("?")
	}
	for key, value := range queryParameters {
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(netURL.QueryEscape(value))
		sb.WriteString("&")
	}
	return sb.String()
}

func addHeaders(request *http.Request, headers map[string]string) {
	if headers == nil {
		return
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}

func getIODataAsInterface(reqBody io.ReadCloser, dto interface{}) error {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return err
	}
	jsonErr := json.Unmarshal(body, dto)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
