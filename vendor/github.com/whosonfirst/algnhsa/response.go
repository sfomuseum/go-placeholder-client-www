package algnhsa

import (
	"encoding/base64"
	"net/http/httptest"
	"log"
	"github.com/aws/aws-lambda-go/events"
)

const acceptAllContentType = "*/*"

func newAPIGatewayResponse(w *httptest.ResponseRecorder, binaryContentTypes map[string]bool) (events.APIGatewayProxyResponse, error) {
	event := events.APIGatewayProxyResponse{}

	rsp := w.Result()
	
	log.Println("HI")
	// Set status code.
	event.StatusCode = rsp.StatusCode

	// w.Result().Header.Set("Content-Type", "text/html")
	
	// Set headers.
	event.MultiValueHeaders = rsp.Header

	// Set body.
	contentType := rsp.Header.Get("Content-Type")
	log.Println("CONTENT TYPE", contentType)

	log.Println("EVENT HEADERS", event.Headers)
	log.Println("EVENT MULTI HEADERS", event.MultiValueHeaders)
	
	if binaryContentTypes[acceptAllContentType] || binaryContentTypes[contentType] {
		event.Body = base64.StdEncoding.EncodeToString(w.Body.Bytes())
		event.IsBase64Encoded = true
	} else {
		event.Body = w.Body.String()
	}

	return event, nil
}
