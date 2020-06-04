package main

import (
	"flag"
	"fmt"
	"github.com/xyclos/github-gpg-keys/client"
	"log"
	"time"
)

func main() {
	githubUser := flag.String("u", "", "Github username to get keys for")
	saveKeys := flag.Bool("s", false, "Save keys to current directory")
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	flag.Parse()

	githubClient := client.NewGithubClient()
	githubClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	keys, err := githubClient.Fetch(client.GithubUser(*githubUser), *saveKeys)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(keys.JSON())
}
