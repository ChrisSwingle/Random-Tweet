package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
    // "reflect"
    // "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
    ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}

var router *gin.Engine

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
    // Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    // Pass in your Access Token and your Access Token Secret
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    // Verify Credentials
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

    // we can retrieve the user and verify if the credentials
    // we have used successfully allow us to log in!
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }
    //
    // log.Printf("User's ACCOUNT:\n%+v\n", user)
    return client, nil
}

func main() {

    router = gin.Default()

    router.LoadHTMLGlob("templates/*")

    router.GET("/", showIndexPage)

    router.Run()

}

func getHaiku(c *gin.Context) string {
    fmt.Println("twitter bot v0.01")
    creds := Credentials{
        AccessToken:        "2432940775-DQlE5yukwxKoJC1Xz2V1qcmLBxV9yHCrbrZKbOF",
        AccessTokenSecret:  "5CoEJdYcVZ6UFIvzOXfI31shZn3ktbwaETJefxCSkf5Zu",
        ConsumerKey:        "x9xtwuXqXTNJ9Rti1zGLL2FUo",
        ConsumerSecret:     "Dy1YS8wlR4D77qh5kvHcYBhb22u1SXvUaVb5m3Ei91gTVCSYJE",
    }

    client, err := getClient(&creds)
    if err != nil {
        log.Println("Error getting Twitter Client")
        log.Println(err)
    }

    // Print out the pointer to our client
    // for now so it doesn't throw errors
    fmt.Printf("%+v\n", client)

    //searcing tweets
    search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
        Query: "#USA",
    })

    if err != nil {
        log.Print(err)
    }

    var ss []string
    for _, element := range search.Statuses {
        ss = append(ss,string(element.Text))
    }

    rand.Seed(time.Now().Unix())
    var rndtweet = string(ss[rand.Intn(len(ss))])

    return rndtweet



    // byteArray,err := json.MarshalIndent(search, "", "   ")
    // fmt.Println(string(byteArray))
}
func render (c *gin.Context, data gin.H, templateName string) {
    switch c.Request.Header.Get("Accept") {
    case "application/json":
        // Respond with JSON
        c.JSON(http.StatusOK, data["payload"])
    case "application/xml":
        // Respond with XML
        c.XML(http.StatusOK, data["payload"])
    default:
        // Respond with HTML
        c.HTML(http.StatusOK, templateName, data)
  }
}
