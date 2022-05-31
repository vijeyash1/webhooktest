package main

import (
	"fmt"

	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		hook, _ := github.New()
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.PushPayload:
			release := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		}
	})
	http.ListenAndServe(":8000", nil)
}
