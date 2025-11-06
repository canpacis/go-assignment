package greet_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/canpacis/go-assignment/internal/errs"
	"github.com/canpacis/go-assignment/services/greet"
)

type TestCase struct {
	Input           any
	ExpectedMessage string
	ExpectedStatus  int
}

func TestGreetService(t *testing.T) {
	cases := []TestCase{
		{
			Input:           "Alice",
			ExpectedMessage: "Hello Alice",
			ExpectedStatus:  http.StatusOK,
		},
		{
			Input:           "alice",
			ExpectedMessage: "Hello alice",
			ExpectedStatus:  http.StatusOK,
		},
		{
			Input:           "Mike",
			ExpectedMessage: "Hello Mike",
			ExpectedStatus:  http.StatusOK,
		},
		{
			Input:           "Nate",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           "Zane",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           "",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           " ",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           0,
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           "😀",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           "Âlice",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           " Alice",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
		{
			Input:           "'Alice'",
			ExpectedMessage: "Invalid Input",
			ExpectedStatus:  http.StatusBadRequest,
		},
	}

	ctx := context.Background()
	service := &greet.Service{}

	for _, testCase := range cases {
		resp, err := service.Greet(ctx, &greet.GreetPayload{Name: fmt.Sprintf("%s", testCase.Input)})

		if testCase.ExpectedStatus == http.StatusOK {
			if err != nil {
				t.Errorf("Expected no error but got %s with input %s", err, testCase.Input)
			}
			if resp.Message != testCase.ExpectedMessage {
				t.Errorf("Expected message to be %s but got %s with input %s", testCase.ExpectedMessage, resp.Message, testCase.Input)
			}
		} else {
			var apiErr *errs.Error
			if errors.As(err, &apiErr) {
				if apiErr.Status != testCase.ExpectedStatus {
					t.Errorf("Expected error status to be %d but got %d with input %s", testCase.ExpectedStatus, apiErr.Status, testCase.Input)
				}

				if apiErr.Message != testCase.ExpectedMessage {
					t.Errorf("Expected error message to be %s but got %s with input %s", testCase.ExpectedMessage, apiErr.Message, testCase.Input)
				}
			} else {
				t.Errorf("Expected error to be *errs.Error but got %T with input %s", err, testCase.Input)
			}
		}
	}
}
