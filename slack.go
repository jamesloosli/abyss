package main

import (
	"github.com/nlopes/slack"
	"github.com/wunderlist/ttlcache"

	"flag"
	//"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)
	
var (
	token = flag.String("token", os.Getenv("SLACK_TOKEN"), "Slack API token")
	targetChannel = flag.String("channel", os.Getenv("SLACK_CHANNEL_ID"), "Slack channel to watch")
	verbose = flag.Bool("verbose", false, "Turn on logging")
)


func getUsername(id string, c *ttlcache.Cache, api *slack.Client) string {
	n, f := c.Get(id)
	if f {
		return n
	} else {
		u, err := api.GetUserInfo(id)
		if err != nil {
			return "error"
		} else {
			c.Set(id, u.Name)
			return u.Name
		}

	}
}


func readSlack() {
	logger := log.New(os.Stdout, "[read][slack]:  ", log.LstdFlags)
	flag.Parse()

	if *verbose {
		sarama.Logger = logger
		slack.SetLogger(logger)
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerlist := strings.Split(*brokers, ",")

	spew.Dump(brokerlist)

	api := slack.New(*token)
	channels, err := api.GetChannels(false)
	if err != nil {
		log.Fatalln("Couldn't list channels:", err)
	}

	f := false
	
	for _, channel := range channels {
		if *targetChannel == channel.Name {
			f = true
		}
	}
	
	if f == false{ 
		_, err := api.JoinChannel(*targetChannel)
		if err != nil {
			log.Fatalln("Couldn't join target channel", err)
		}
	}

	userCache := ttlcache.NewCache(time.Hour)
	kafkaProducer := newProducer(brokerlist)
	
	spew.Dump(userCache, kafkaProducer)

	/*
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				u := getUsername(ev.User, userCache, api)
				m := ev.Text
				t := "slack" + *targetChannel 
				writeToKafka(m, u, t, *kafkaProducer)
				fmt.Printf("%+v: %+v\n", u,m)

			default:
				// Do nothing.
		}
	}

	/**/
}

func main() {
	readSlack()
}
