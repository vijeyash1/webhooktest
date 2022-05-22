package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func main() {
	repo, _ := git.PlainClone("/tmp/webhooktest", false, &git.CloneOptions{
		URL: "https://github.com/vijeyash1/webhooktest",
	})
	head, _ := repo.Head()
	commit, _ := repo.CommitObject(head.Hash())
	stats, _ := commit.Stats()

	for _, stat := range stats {
		fmt.Printf("add: %d\tdel: %d\tfile: %s\n", stat.Addition, stat.Deletion, stat.Name)
		

	}
}
