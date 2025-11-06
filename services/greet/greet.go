package greet

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/canpacis/go-assignment/internal/errs"
)

// Represents the Greeter Service.
type Service struct {
	// Service may include service specific
	// fields and injected dependencies
}

// Represents the payload passed to the Greet method.
type GreetPayload struct {
	Name string
}

func (gp *GreetPayload) Validate() error {
	if len(strings.TrimSpace(gp.Name)) == 0 {
		return &errs.Error{Message: "Invalid Input", Status: http.StatusBadRequest}
	}
	return nil
}

// Represents the response returned by the Greet method.
type GreetResponse struct {
	Message string `json:"message"`
}

// Greet method handles the business logic of the Greeter service.
func (*Service) Greet(ctx context.Context, payload *GreetPayload) (*GreetResponse, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	firstLetter := strings.Split(payload.Name, "")[0]
	regexpr := regexp.MustCompile("[a-mA-M]")

	if !regexpr.MatchString(firstLetter) {
		return nil, &errs.Error{Message: "Invalid Input", Status: http.StatusBadRequest}
	}

	message := fmt.Sprintf("Hello %s", payload.Name)

	return &GreetResponse{Message: message}, nil
}
