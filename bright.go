package main

import (
  "github.com/joho/godotenv"
  "github.com/go-chat-bot/bot/irc"
  _ "github.com/aburgd/plugins/gif"
  _ "github.com/aburgd/plugins/catgif"
  _ "github.com/aburgd/plugins/catfacts"
  _ "github.com/aburgd/plugins/treta"
  _ "github.com/aburgd/plugins/puppet"
  // custom plugins
  _ "github.com/aburgd/plugins/hello"
  _ "github.com/aburgd/plugins/about"
  _ "github.com/aburgd/plugins/doctors"
  _ "github.com/aburgd/plugins/botgod"
  // system
  "os"
  "log"
  "strings"
)

func main()  {
  err := godotenv.Load()
  if err != nil {
   log.Fatal("Error loading .env")
  }

  // load the config
  config := brightConfig()
  log.Printf("connected to server: %s\n channels: %s", config.Server, config.Channels)
  irc.Run(config)
}

func brightConfig() *irc.Config {
  server := os.Getenv("IRC_SERVER")
  channels := strings.Split("#" + os.Getenv("IRC_CHANNELS"), ",")

  return &irc.Config{
		Server:   server, //os.Getenv("IRC_SERVER"),
		Channels: channels, //strings.Split("#" + os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   false,
		Debug:    os.Getenv("DEBUG") != "",
	}
}
