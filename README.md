# Listcord.go

An official wrapper for the listcord api in golang!

> View https://listcord.xyz/apidocs to view the raw api documentation!

## Installation

```sh
go get github.com/listcordteam/listcord.go
```

## Getting started

> Get your api token from https://listcord.xyz/me. The Listcord api is currently only available only for those who have bots registered in our botlist! After getting your token, make sure your token is kept somewhere secret!

```go
package main

import "fmt"
import listcord "github.com/listcordteam/listcord.go"

func main(){
    client := listcord.Client("Replace this with your token")
    bot, err := client.GetBot("some bot id")
}
```

## Methods

There are only some methods within the package itself!

```go
bot, err := client.GetBot("bot id") // Returns the listcord.Bot which contains the bot information!
reviews, err := client.GetBotReviews("bot id") // Returns the []listcord.BotReview containing array of bot reviews!
review, err := client.GetReview("user id", "bot id") // Returns listcord.BotReview containing the review information by the user id who reviewed and the bot id where it was reviewed!
voted, err := client.HasVoted("user id", "bot id") // Returns listcord.VoteData containing the upvote data of the user upvote to the bot!
```

## Contact

- [Join our discord server](https://discord.gg/cMGAyhZXwW)
- [Website](https://listcord.xyz)