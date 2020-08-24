package main

import (
	"TwitterContest/db"
	"TwitterContest/tasks"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	db.ConnectDb()
	defer db.DB.Close()
	r.GET("/twitter/retweets/:user_handle/max", tasks.Winner)      // to get winner
	r.GET("/twitter/tweet/:user_handle/latest", tasks.LatestTweet) // to get latest tweet of a particular user handle
	r.Run()                                                        // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}

