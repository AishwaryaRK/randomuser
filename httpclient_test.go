package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type MockRoundTripper struct {
	mockResponse *http.Response
	mockError    error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.mockResponse, m.mockError
}

func TestGetHttpJsonResponse(t *testing.T) {
	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"results":[{"name":{"title":"mr", "first":"hiroki"},"location":{"city":"tokyo"}}]}`)),
	}
	mockedHttpClient := &http.Client{Transport: &MockRoundTripper{mockResponse: mockResp, mockError: nil}}

	var actual Response
	err := GetHttpJsonResponse(&HttpClient{client: mockedHttpClient}, "https://randomuser.me/api/", &actual)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	user := User{Name: Name{Title: "mr", First: "hiroki"}, Location: Location{City: "tokyo"}}
	expected := Response{Results: []User{user}}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestGetHttpJsonResponseInvalidUrl(t *testing.T) {
	var actual Response
	err := GetHttpJsonResponse(NewHttpClient(5), "invalid url", &actual)
	if err == nil {
		t.Fatalf("expected error, invalid url")
	}
	fmt.Println(err)
}
