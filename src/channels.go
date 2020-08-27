package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	client *discordgo.Session
	channels map[string]string
}

func (d *Discord)Init(token string){
	conn, _ := discordgo.New("Bot " + token)
	d.client = conn
	d.client.Open()
	gld, _ := d.client.Guild(d.client.State.Guilds[0].ID)
	d.channels = make(map[string]string)
	fmt.Print("Searching for channels:\n")
	for _,v := range gld.Channels{
		d.channels[v.Name] = v.ID
		fmt.Printf("%+v -> %+v\n", v.ID, v.Name)
	}
	d.client.Close()
}

func (d *Discord)GetChID(name string) string{
	for k,v := range d.channels{
		if k==name{return v}
	}
	return ""
}

func (d *Discord)SendMessage(ch_name, text string){
	d.client.Open()
	ch_id := d.GetChID(ch_name)
	if ch_id==""{return}
	d.client.ChannelMessageSend(ch_id, text)
	d.client.Close()
}
