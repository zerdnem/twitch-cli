package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/catsby/go-twitch/service/kraken"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

//audio_only, 160p, 360p, 480p, 720p, 720p60, 1080p
var quality string = "480p"

var chat bool = true

//var streamer string

func main() {
	client := kraken.DefaultClient(nil)

	out, err := client.GetFollowedStreams(&kraken.GetFollowedStreamsInput{
		StreamType: kraken.StreamTypeLive,
	})
	if err != nil {
		log.Fatalf("Error getting followed streams: %s", err)
	}

	if len(out.Streams) == 0 {
		fmt.Println("None of your followed streams are live right now, or you have none at all")
	}

	var streaming []string
	for _, s := range out.Streams {
		//streaming = append(streaming, fmt.Sprintf("%s - Playing %s", s.Channel.DisplayName, s.Game))
		streaming = append(streaming, s.Channel.DisplayName)
	}

	var streamers = []*survey.Question{{
		Name: "streamer",

		Prompt: &survey.Select{
			Message:  "Streaming now:",
			Options:  streaming,
			PageSize: 20,
		},
		Validate: survey.Required,
	}}

	answers := struct {
		Streamer string
	}{}

	survey.Ask(streamers, &answers)
	streamer := strings.ToLower(answers.Streamer)
	//r, _ := regexp.Compile(streamer)
	//answer := r.FindString(answers.Streamer)
	livestreamer := exec.Command("livestreamer", "twitch.tv/"+streamer, quality, "--http-header", "Client-ID=jzkbprff40iqj646a697cyrvl0zt2m6", "--player=mpv")
	if chat {
		cmd := exec.Command("google-chrome", "--chrome-frame", "--window-size=400,600", "--window-position=580,240", "--app=https://www.twitch.tv/popout/"+streamer+"/chat?popout=")
		cmd.Start()
	}

	livestreamer.Run()
}
