package tasks

import (
	"TwitterContest/db"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	models "TwitterContest/types"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var user models.User

func retweets(id int, tweetID int64, client *twitter.Client, count *map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	retweets, _, _ := client.Statuses.Retweets(tweetID, &twitter.StatusRetweetsParams{Count: 100})
	fmt.Println(len(retweets), " retweets")
	for _, retweet := range retweets {
		(*count)[retweet.User.Name] = (*count)[retweet.User.Name] + 1
	}
}

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

func Winner(c *gin.Context) {

	var flags models.FlagsAuth
	if err := c.ShouldBindHeader(&flags); err != nil { // will get the key and secret through postman header
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := getClient(flags) //get client from getClient() Function

	user_handle := c.Param("user_handle")
	// user timeline
	bool1 := false
	bool2 := true
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: user_handle, ExcludeReplies: &bool2, IncludeRetweets: &bool1, Count: 100}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
	// for i, user := range tweets {
	//  fmt.Printf("Tweet %d : \n%+v\n", i, user.Text)
	fmt.Println(len(tweets), " tweets")
	// }
	usertweets := make([]string, 0, len(tweets))
	for _, tweet := range tweets {
		usertweets = append(usertweets, tweet.IDStr)
	}
	//fmt.Println(allUsernames)
	count := make(map[string]int)
	var wg sync.WaitGroup
	for i, tweet := range usertweets {
		wg.Add(1)
		temp, _ := strconv.ParseInt(tweet, 10, 64)
		go retweets(i, temp, client, &count, &wg) // function call
	}
	wg.Wait()

	// connecting to mysql database using gorm

	user = db.DatabaseSQL(&count)

	c.JSON(200, gin.H{
		"winner": user.Username,
		"Count":  user.Count,
	})
}

func LatestTweet(c *gin.Context) {

	var flags models.FlagsAuth
	if err := c.ShouldBindHeader(&flags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := getClient(flags) //get client from getClient() Function

	user_handle := c.Param("user_handle")
	// user timeline
	bool1 := false
	bool2 := true
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: user_handle, ExcludeReplies: &bool2, IncludeRetweets: &bool1, Count: 1}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
	// for i, user := range tweets {
	//  fmt.Printf("Tweet %d : \n%+v\n", i, user.Text)
	// }

	// usertweets := make([]string, 0, len(tweets))
	// for _, tweet := range tweets {
	// 	usertweets = append(usertweets, tweet.Text)
	// }

	//fmt.Println(allUsernames)

	c.JSON(200, gin.H{
		"LatestTweet": tweets[0].Text,
	})
}
