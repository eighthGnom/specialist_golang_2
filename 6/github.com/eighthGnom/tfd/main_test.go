package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	Input    int
	Expected int
	Answer   int
}

var Cases = []TestCase{
	{
		Input:    0,
		Expected: 1,
	},
	{
		Input:    1,
		Expected: 1,
	},
	{
		Input:    2,
		Expected: 2,
	},
	{
		Input:    10,
		Expected: 3628800,
	},
}

func TestMainFact(t *testing.T) {
	for id, testCase := range Cases {
		if testCase.Answer = factorial(testCase.Input); testCase.Answer != testCase.Expected {
			t.Errorf("Error while runing CASE#%d, input: %d, answer: %d, expected: %d",
				id, testCase.Input, testCase.Answer, testCase.Expected)
		}
	}
}

type HTTpTestCase struct {
	Name     string
	Numeric  int
	Expected []byte
}

var HTTpCases = []HTTpTestCase{
	{
		Name:     "first test",
		Numeric:  0,
		Expected: []byte("1"),
	},
	{
		Name:     "second test",
		Numeric:  1,
		Expected: []byte("1"),
	},
	{
		Name:     "third test",
		Numeric:  10,
		Expected: []byte("3628800"),
	},
}

func TestHTTPFactorial(t *testing.T) {
	handler := http.HandlerFunc(HTTPFactorial)
	for _, testCase := range HTTpCases {
		t.Run(testCase.Name, func(t *testing.T) {
			uri := fmt.Sprintf("/fatorial?num=%d", testCase.Numeric)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, uri, nil)
			handler.ServeHTTP(recorder, request)
			if string(recorder.Body.Bytes()) != string(testCase.Expected) {
				t.Errorf("Faling while runing test# %s: for: %d, expected: %s, got: %s",
					testCase.Name, testCase.Numeric, testCase.Expected, recorder.Body.Bytes())
			}
		})
	}
}
