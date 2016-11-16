package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	client := getClient()

	gists, _, err := client.Gists.List("", nil)

	if err != nil {
		panic(err)
	}

	fmt.Print("Which gist do you wanna open ? \n\n")

	for i := 0; i < len(gists); i++ {
		if gists[i].Description == nil || *gists[i].Description == "" {
			var filesName []string
			for _, gistFile := range gists[i].Files {
				filesName = append(filesName, *gistFile.Filename)
			}
			fmt.Printf("%d - %s \n", i+1, filesName[0])
		} else {
			fmt.Printf("%d - %v \n", i+1, *gists[i].Description)
		}
	}
}

func getClient() *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tokenClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

	return github.NewClient(tokenClient)
}
