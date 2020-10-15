package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func (d *DiscordClient) GetUsers(){
	fmt.Println("\nSearching for users:")
	d.client.Open()
	members, err := d.client.GuildMembers(d.client.State.Guilds[0].ID, "0", 1000)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _,v := range members{
		d.users[v.User.Username] = v.User.ID
		fmt.Printf("%v -> %v\n",v.User.ID, v.User.Username)
	}
}

func (d *DiscordClient) GetUser(name string)(id string){
	for k,v := range d.users{
		if k==name{id = v;return}
		if v==name{id = v;return}
	}
	return
}

func (d *DiscordClient) MessageUser(username, text string) *discordgo.Message{
	d.client.Open()
	id := d.GetUser(username)
	ch, err := d.client.UserChannelCreate(id)
	if err != nil {
		fmt.Println(err)
	}
	m, _ := d.client.ChannelMessageSend(ch.ID, text)
	return m
}
