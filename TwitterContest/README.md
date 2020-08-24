# Twitter Contest

## Setup Instructions:

 - Packages required to run the app: 
  ```go
   go get  "github.com/dghubble/go-twitter/twitter"
   go get  "github.com/gin-gonic/gin"
   go get  "github.com/go-sql-driver/mysql"
   go get  "github.com/jinzhu/gorm"
   go get  "github.com/jinzhu/gorm/dialects/mysql"
   go get  "golang.org/x/oauth2"
   go get  "golang.org/x/oauth2/clientcredentials"
  ```
  
## Prerequisites

 You need to install the following software before using this API

 ```
 Mysql
 Golang
 Postman
 ```
## Endpoints:
- ```sh
  GET /twitter/retweets/:user_handle/max  
  ```
  Picks up winners who has retweeted the most number of times across the last 100 tweets of a given Twitter Handle and then prints out the winner's username and the   count of retweets
  
- ```sh
  GET /twitter/tweet/:user_handle/latest  
  ```
  Prints out the latest tweets of a given Twitter Handle

## Sending request to the Api
 Use Postman to send the request by providing the ConsumerKey and ConsumerSecret as header in Postman

## Project Structure
 The project has three packages apart from main
 - **types**  
  The types.go file in the types package consists of the definitions of various structs
 - **tasks**  
  The tasks package consists of logic to handle various endpoints and authentication(client.go)
 - **db**  
  This package is used to handle database connection and to post or retrieve data from MySQl
  
  
## To run the application using docker
 * Requirements : `Docker`
 * ### RUN `docker build --tag twittercontest .`
 * ### RUN `docker run -it -p 8080:8080 twittercontest`
 * The Application is up and is running at [localhost:8080](http://localhost:8080)   

## To run the application locally
 * Requirements : `Golang`

 * `cd` to the root directory
 * ### RUN  `go mod download`
 * ### RUN  `go run main.go`
 * The Application is up and is running at [localhost:8080](http://localhost:8080) 
