package http

import (
	"context"
	"errors"
	"log"
	gohttp "net/http"
	"time"
)

func PlaceholderReadyHandler(placeholder_url string) (gohttp.Handler, error) {

	placeholder_ready := false
	var placeholder_error error

	go func() {

		ctx := context.Background()

		d := time.Now().Add(30 * time.Second)

		ctx, cancel := context.WithDeadline(ctx, d)
		defer cancel()

		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				placeholder_error = errors.New("Placeholder ready check timed out")
				return
			case <-ticker.C:

				log.Println("Check", placeholder_url)

				rsp, err := gohttp.Get(placeholder_url)

				if err == nil {
					rsp.Body.Close()
					placeholder_ready = true
					return
				}

				log.Printf("Failed to load Placeholder URL, %v", err)
			}
		}
	}()

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		if placeholder_ready {
			return
		}

		if placeholder_error != nil {
			gohttp.Error(rsp, placeholder_error.Error(), gohttp.StatusInternalServerError)
			return
		}

		gohttp.Error(rsp, "Service Unavailable", gohttp.StatusServiceUnavailable)
		return
	}

	return gohttp.HandlerFunc(fn), nil
}
