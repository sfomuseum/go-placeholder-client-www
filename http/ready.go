package http

import (
	"context"
	"errors"
	"log"
	gohttp "net/http"
	"time"
)

func PlaceholderReadyHandler(placeholder_endpoint string, ttl time.Time) (gohttp.Handler, error) {

	placeholder_ready := false
	var placeholder_error error

	go func() {

		ctx := context.Background()

		ctx, cancel := context.WithDeadline(ctx, ttl)
		defer cancel()

		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				placeholder_error = errors.New("Placeholder ready check timed out")
				log.Println(placeholder_error.Error())
				return
			case <-ticker.C:

				log.Println("Check Placeholder status (%s)", placeholder_endpoint)

				rsp, err := gohttp.Get(placeholder_endpoint)

				if err == nil {
					rsp.Body.Close()
					placeholder_ready = true

					log.Printf("Placeholder appears to running and accepting connections")
					return
				}

				log.Printf("Failed to determine Placeholder status, %v", err)
			}
		}
	}()

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		if placeholder_ready {
			rsp.Write([]byte("OK"))
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
