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
