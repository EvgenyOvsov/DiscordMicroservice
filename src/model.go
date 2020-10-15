package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)
type DiscordClient struct {
	client *discordgo.Session
	channels map[string]string
	users map[string]string
}

type Request struct {
	Token string `json:"token"`
	To string `json:"to"`
	Title string `json:"title"`
	Channel string `json:"channel"`
	Text string `json:"text"`
	Image string `json:"image"`
	Color string `json:"color"`
}

func (d *DiscordClient)Init(token string){
	conn, _ := discordgo.New("Bot " + token)
	d.client = conn
	err := d.client.Open()
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	d.channels = make(map[string]string)
	d.users = make(map[string]string)

	d.GetChannels()
	d.GetUsers()

}
