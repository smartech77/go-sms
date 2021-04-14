package siminn

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// SiminnSMS
// ..
type SiminnSMS struct {
	URL      string // The URL for the sms service
	Username string // A username you were given
	Password string // A password you were given
	SendFrom string // The name you want displayed as the sender
}

// SendSMS
// This function send a text message using Siminn SMS api ( see docs in README )
// At this moment it's a bit hacky because the API does not use HTTP Response
// Codes as suggested by the HTTP speficifaction, this is no ones fault
// sometimes things just turn out that way, we will accomidate for this.
func (s *SiminnSMS) SendSMS(ctx context.Context, content string, number string) (error, bool, int) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, s.URL+"push", nil)
	if err != nil {
		return err, false, 0
	}

	newQuery := r.URL.Query()
	newQuery.Set("l", s.Username)
	newQuery.Set("p", s.Password)
	newQuery.Set("A", s.SendFrom)
	newQuery.Set("T", content)
	newQuery.Set("msisdn", number)
	r.URL.RawQuery = newQuery.Encode()

	r = r.WithContext(ctx)

	response, sendError := client.Do(r)
	if sendError != nil {
		return sendError, false, 0
	}

	bodyBytes, responseReadErr := io.ReadAll(response.Body)
	if responseReadErr != nil {
		return responseReadErr, false, response.StatusCode
	}

	// super hack because it always returns 200
	if strings.Contains(string(bodyBytes), "ERROR") {
		responseError := errors.New(string(bodyBytes))
		return responseError, false, response.StatusCode
	}

	// Again, hacky because of the way response are being made
	if strings.Contains(string(bodyBytes), "SUCCESS") {
		return nil, true, response.StatusCode
	}

	// A status code that is not 200, is not an intended error all we
	// can do here is return the code and body for debugging
	if response.StatusCode != 200 {
		responseError := errors.New("Siminn returned a non 200 error:" + strconv.Itoa(response.StatusCode) + ": " + string(bodyBytes))
		return responseError, false, response.StatusCode
	}

	// always default to fail
	return nil, false, response.StatusCode
}
