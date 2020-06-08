package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/xyclos/github-gpg-keys/client"
)

func main() {
	githubUser := flag.String("u", "", "GitHub username to get keys for")
	email := flag.String("e", "", "Email to filter keys by")
	saveKeys := flag.Bool("s", false, "Save keys to current directory")
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	flag.Parse()

	githubClient := client.NewGithubClient()
	githubClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	keys, err := githubClient.Fetch(client.GithubUser(*githubUser), email)
	if err != nil {
		log.Println(err)
	}

	if *saveKeys {
		for _, key := range keys {
			if filePath, err := githubClient.SaveToDisk(key.KeyID, key.RawKey, "."); err != nil {
				fmt.Println("Failed to save key!")
			} else {
				fmt.Println("Saved key to file: ", filePath)
			}
		}
	} else {
		log.Println(keys.JSON())
	}
}
