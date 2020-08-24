package tasks

import (
	models "TwitterContest/types"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func getClient(flags models.FlagsAuth) *twitter.Client {

	if flags.ConsumerKey == "" || flags.ConsumerSecret == "" {
		log.Fatal("Application Access Token required")
	}
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     flags.ConsumerKey,
		ClientSecret: flags.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)
	// Twitter client
	client := twitter.NewClient(httpClient)

	return client
}
