package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var buffer = make([][]byte, 0)

func main() {
	const token = "MjQ0OTA3NTYyMzQ3MTM0OTc3.CwEXhw.VYdUT4dllK2wVm2UVsucWKFQ6FY"
	if token == "" {
		fmt.Println("No bot token specified. Exiting...")
		return
	}
	discord, err := discordgo.New("Bot " + token)

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	fmt.Println("Discord-CRM-Link is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	_ = s.UpdateStatus(0, "!airhorn")
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable != nil {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!airhorn") {
		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			return
		}
		fmt.Println(c.Name, g.Name)
	}
}
