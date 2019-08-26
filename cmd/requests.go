package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const EnvVariableThatHoldsMyPat = "GH_DEV_PAT"
const EnvVariableThatHoldsMyGhUser = "GH_USERNAME"
const EnvVariableThatHoldsMyRepoOwner = "GH_USERNAME"
const EnvVariableThatHoldsMyRepoName = "GH_REPO"

// GET  /
// GET  /repos/:owner/:repo/collaborators
// POST /repos/:owner/:repo/pulls
// POST /repos/:owner/:repo/pulls/:pull_number/reviews

func getBaseUrl() string {
	return "https://api.github.com"
}

func buildRoute(uri string) string {
	return fmt.Sprintf("%s/%s", getBaseUrl(), uri)
}

type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

func createNewRequest(resource string, buf *bytes.Buffer) *http.Request {
	var request *http.Request
	var err error
	if nil == buf {
		request, err = http.NewRequest("GET", buildRoute(resource), nil)
	} else {
		request, err = http.NewRequest("POST", buildRoute(resource), buf)
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

func getResource(client HttpClient, resource string) []byte {
	var buf *bytes.Buffer
	request := createNewRequest(resource, buf)
	_, body := executeRequest(client, request)
	return body
}
