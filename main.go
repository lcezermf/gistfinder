package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var GetGistsError = errors.New("Cannot get user gists")
var UnsupportedPlatform = errors.New("Unsupported platform")

func main() {
	client := getClient()

	gists, _, err := client.Gists.List("", nil)

	if err != nil {
		panic(GetGistsError)
	}

	fmt.Print("\nYou Gists: \n\n")

	gistsUrls := make(map[string]string)

	for i := 0; i < len(gists); i++ {
		indexToString := strconv.Itoa(i)

		if gists[i].Description == nil || *gists[i].Description == "" {
			var filesName []string
			for gistFilename := range gists[i].Files {
				filesName = append(filesName, string(gistFilename))
			}
			fmt.Printf("%s - %s \n", indexToString, filesName[0])
		} else {
			fmt.Printf("%s - %v \n", indexToString, *gists[i].Description)
		}

		gistsUrls[indexToString] = *gists[i].HTMLURL
	}

	fmt.Print("\nSelect the number of gist that you want to open in browser: \n\n")
	var input string
	fmt.Scan(&input)

	if gistUrl, ok := gistsUrls[input]; ok {
		openBrowser(gistUrl)
	} else {
		fmt.Print("\nGist not found\n")
	}
}

func openBrowser(gistUrl string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", gistUrl).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", gistUrl).Start()
	case "darwin":
		err = exec.Command("open", gistUrl).Start()
	default:
		err = UnsupportedPlatform
	}

	if err != nil {
		fmt.Print(err)
	}
}

func getClient() *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tokenClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

	return github.NewClient(tokenClient)
}
