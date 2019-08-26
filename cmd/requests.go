package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Returns the API's base URL
func getBaseUrl() string {
	return "https://api.github.com"
}

// Given a specific URI, returns a full route to the resource
func buildRoute(uri string) string {
	return fmt.Sprintf("%s/%s", getBaseUrl(), uri)
}

// This is a simple interface so the client can be easily mocked for testing
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

// Given a URI and a payload, this constructs a full request with all the
// necessary headers
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

// Converts the provided io.Reader (theoretically from a response.Body) into a
// byte slice by slurping its contents
func parseResponseBody(body io.Reader) []byte {
	contents, err := ioutil.ReadAll(body)
	whereErrorsGoToDie(err)
	return contents
}

// Given a client and request, runs the request and returns a tuple of the
// response and its body as a []byte
func executeRequest(client HttpClient, request *http.Request) (*http.Response, []byte) {
	response, err := client.Do(request)
	whereErrorsGoToDie(err)
	body := parseResponseBody(response.Body)
	err = response.Body.Close()
	return response, body
}

// Given a URI (not a URL) and a payload (which could be empty), constructs
// a client, executes the request, and returns the resource
func getResource(resource string, requestBody []byte) []byte {
	client := &http.Client{}
	request := createNewRequest(resource, requestBody)
	_, body := executeRequest(client, request)
	return body
}
