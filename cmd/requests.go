package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getBaseUrl() string {
	return "https://api.github.com"
}

func buildRoute(uri string) string {
	return fmt.Sprintf("%s/%s", getBaseUrl(), uri)
}

type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

func createNewRequest(resource string, requestBody []byte) *http.Request {
	var request *http.Request
	var err error
	if nil == requestBody {
		fmt.Println(buildRoute(resource))
		request, err = http.NewRequest("GET", buildRoute(resource), nil)
	} else {
		request, err = http.NewRequest("POST", buildRoute(resource), bytes.NewBuffer(requestBody))
	}
	whereErrorsGoToDie(err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/vnd.github.v3+json")
	request.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv(EnvVariableThatHoldsMyPat)))
	request.Header.Set("User-Agent", os.Getenv(EnvVariableThatHoldsMyGhUser))
	request.Header.Set("Time-Zone", "America/Chicago")
	return request
}

func parseResponseBody(body io.Reader) []byte {
	contents, err := ioutil.ReadAll(body)
	whereErrorsGoToDie(err)
	return contents
}

func executeRequest(client HttpClient, request *http.Request) (*http.Response, []byte) {
	response, err := client.Do(request)
	whereErrorsGoToDie(err)
	body := parseResponseBody(response.Body)
	err = response.Body.Close()
	return response, body
}

func getResource(resource string, requestBody []byte) []byte {
	client := &http.Client{}
	request := createNewRequest(resource, requestBody)
	_, body := executeRequest(client, request)
	return body
}
