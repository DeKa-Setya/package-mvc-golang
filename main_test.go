package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	BadRequestCode     = 400
	SuccessRequestCode = 200
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

func TestUserSignup(t *testing.T) {

	url := "http://localhost:8081/api/v1/account/add"

	tests := []TestStruct{
		{`{}`, BadRequestCode, "", 0},
		{`{"name":""}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456","saldo" : 1000000}`, SuccessRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"12456","saldo" : 1000000}`, BadRequestCode, "", 0}, // request should fail because of
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("UserSignup", tests, t)

}

func TestUserSignin(t *testing.T) {

	url := "http://localhost:8081/api/v1/login"

	tests := []TestStruct{
		{`{}`, BadRequestCode, "", 0},
		{`{"name":""}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, SuccessRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"12456"}`, BadRequestCode, "", 0}, // request should fail because of
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("UserSignin", tests, t)
}

type List struct{}

func (l List) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-Id", r.Header.Get("X-Request-Id"))

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte("inanzzz"))
}
func TestUserGetAccountTrans(t *testing.T) {
	url := "http://localhost:8081/api/v1/account/"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-Request-Id", "Test-Header")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X251bWJlciI6ODMyNzEyLCJuYW1lIjoiSXJ3YW4gU3lhaHB1dHJhIn0.0-Ivg0qYyl3-8iWxhOjQ55bAb9jcrEt25bJBuZCP408")
	res := httptest.NewRecorder()

	list := List{}
	list.Handle(res, req)

	result := res.Result()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error(err)
	}
	result.Body.Close()

	if http.StatusOK != result.StatusCode {
		t.Error("wanted", http.StatusOK, "got", result.StatusCode)
	}
	if "Test-Header" != result.Header.Get("X-Request-Id") {
		t.Error("wanted Test-Header got", result.Header.Get("X-Request-Id"))
	}
	if "inanzzz" != string(body) {
		t.Error("wanted inanzzz got", string(body))
	}
}

func TestUserGetAccountMutasi(t *testing.T) {
	url := "http://localhost:8081/api/v1/mutasi/"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-Request-Id", "Test-Header")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X251bWJlciI6ODMyNzEyLCJuYW1lIjoiSXJ3YW4gU3lhaHB1dHJhIn0.0-Ivg0qYyl3-8iWxhOjQ55bAb9jcrEt25bJBuZCP408")

	res := httptest.NewRecorder()

	list := List{}
	list.Handle(res, req)

	result := res.Result()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error(err)
	}
	result.Body.Close()

	if http.StatusOK != result.StatusCode {
		t.Error("wanted", http.StatusOK, "got", result.StatusCode)
	} else {
		t.Logf("responseBody : %s", result.Body)
	}
	if "Test-Header" != result.Header.Get("X-Request-Id") {
		t.Error("wanted Test-Header got", result.Header.Get("X-Request-Id"))
	}
	if "inanzzz" != string(body) {
		t.Error("wanted inanzzz got", string(body))
	}
}

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {

	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}
