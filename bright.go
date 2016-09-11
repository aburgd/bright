package main

import (
  "github.com/joho/godotenv"
  "github.com/go-chat-bot/bot/irc"
  _ "github.com/go-chat-bot/plugins/gif"
  _ "github.com/go-chat-bot/plugins/catgif"
  _ "github.com/go-chat-bot/plugins/catfacts"
  _ "github.com/go-chat-bot/plugins/treta"
  _ "github.com/go-chat-bot/plugins/url"
  _ "github.com/go-chat-bot/plugins/puppet"
  // custom plugins
  _ "github.com/go-chat-bot/plugins/hello"
  _ "github.com/go-chat-bot/plugins/about"
  _ "github.com/go-chat-bot/plugins/doctors"
  _ "github.com/go-chat-bot/plugins/botgod"
  // system
  "os"
  "log"
  "strings"
)

func main()  {
  // error handling for env loadings
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  // load the config
  config := brightConfig()
  log.Printf("%v\n", config)
  irc.Run(config)
}

func brightConfig() *irc.Config {
  return &irc.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: strings.Split("#" + os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   false,
		Debug:    os.Getenv("DEBUG") != "",
	}
}
