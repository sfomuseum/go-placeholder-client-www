package algnhsa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleEvent(ctx context.Context, event events.APIGatewayProxyRequest, handler http.Handler, opts *Options) (events.APIGatewayProxyResponse, error) {
	r, err := newHTTPRequest(ctx, event, opts.UseProxyPath)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	log.Println("SERVED")
	
	for k, v := range w.Header() {
		log.Println("W HEADER", k, v)
	}

	for k, v := range w.HeaderMap {
		log.Println("W HEADER MAP", k, v)
	}

	log.Println("RESPONSE")
	
	return newAPIGatewayResponse(w, opts.binaryContentTypeMap)
}

var defaultOptions = &Options{}

// ListenAndServe starts the AWS Lambda runtime (aws-lambda-go lambda.Start) with a given handler.
func ListenAndServe(handler http.Handler, opts *Options) {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	if opts == nil {
		opts = defaultOptions
	}
	opts.setBinaryContentTypeMap()
	lambda.Start(func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return handleEvent(ctx, event, handler, opts)
	})
}
