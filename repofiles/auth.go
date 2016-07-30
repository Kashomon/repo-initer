package repofiles

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func Auth() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	_, _, err := client.Repositories.List("", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
