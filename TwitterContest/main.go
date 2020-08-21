package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type flagsAuth struct {
	ConsumerKey    string `header:“ConsumerKey”`
	ConsumerSecret string `header:“ConsumerSecret”`
}

type User struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:"size:255"`
	Count    int
}

func getClient(flags flagsAuth) *twitter.Client {

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

func retweets(id int, tweetID int64, client *twitter.Client, count *map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	retweets, _, _ := client.Statuses.Retweets(tweetID, &twitter.StatusRetweetsParams{Count: 100})
	fmt.Println(len(retweets), " retweets")
	for _, retweet := range retweets {
		(*count)[retweet.User.Name] = (*count)[retweet.User.Name] + 1
	}
}

func database(count *map[string]int) User {
	db, err := gorm.Open("mysql", "admin:9160266544135@tcp(127.0.0.1:3306)/testDb?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.Debug().DropTableIfExists(&User{})
	//Drops table if already exists

	db.Debug().AutoMigrate(&User{})
	//Auto create table based on

	fmt.Println(len(*count))

	// Inserting data into database
	var max = 0
	for key, value := range *count {
		user := &User{Username: key, Count: value}
		//fmt.Println(user.Username, " ", user.Count)
		db.Create(user)
		if max < value {
			max = value
		}
	}

	var user User
	db.Where("Count = ?", max).First(&user)
	//SELECT * FROM users WHERE count = max;
	fmt.Println("Hello")

	return user
}

func winner(c *gin.Context) {

	var flags flagsAuth
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
	user := database(&count)

	c.JSON(200, gin.H{
		"winner": user.Username,
		"Count":  user.Count,
	})
}

func latestTweet(c *gin.Context) {

	var flags flagsAuth
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

func main() {
	r := gin.Default()
	r.GET("/twitter/retweets/:user_handle/max", winner)      // to get winner
	r.GET("/twitter/tweet/:user_handle/latest", latestTweet) // to get latest tweet of a particular user handle
	r.Run()                                                  // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}
