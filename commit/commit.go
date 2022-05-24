package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

func main() {

	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://gitlab.com/vijeyash/AIML.git",
	})

	// retrieves the branch pointed by HEAD
	ref, _ := r.Head()

	// get the commit object, pointed by ref
	commit, _ := r.CommitObject(ref.Hash())

	// retrieves the commit history
	author := commit.Author.Name

	// iterates over the commits and print each

	fmt.Println(author)

}
