package main

import (
  "github.com/go-chat-bot/bot/irc"
  _ "github.com/aburgd/plugins/gif"
  _ "github.com/aburgd/plugins/catgif"
  _ "github.com/aburgd/plugins/catfacts"
  _ "github.com/aburgd/plugins/treta"
  _ "github.com/aburgd/plugins/url"
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
