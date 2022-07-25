package main

import (
	"fmt"


	"net/http"

	"github.com/go-playground/webhooks/v6/github


	"github.com/go-git/go-git/plumbing"
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Example of how to:
// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func main() {

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)


	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		hook, _ := github.New()
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch value = payload.(type) {

		case github.PushPayload:
			release := value
			// Do whatever you want from here...
			fmt.Printf("%+v", release)
			fmt.Printf("%+v", release)

			fmt.Printf("%+v", release)


		}
	})
	http.ListenAndServe(":8000", nil)

}
