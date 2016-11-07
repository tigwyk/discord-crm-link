package main

import "fmt"
import "github.com/bwmarrin/discordgo"

func main() {
	const token = "MjQ0OTA3NTYyMzQ3MTM0OTc3.CwEXhw.VYdUT4dllK2wVm2UVsucWKFQ6FY"
	if token == "" {
		fmt.Println("No bot token specified. Exiting...")
		return
	}
	discord, err := discordgo.New("Bot " + token)
	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	fmt.Println("Discord-CRM-Link is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
