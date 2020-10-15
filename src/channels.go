package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (d *DiscordClient)GetChannels(){
	channels, _ := d.client.GuildChannels(d.client.State.Guilds[0].ID)
	fmt.Printf("Searching for channels in %v:\n", d.client.State.Guilds[0].Name)
	for _,v := range channels{
		d.channels[v.Name] = v.ID
		fmt.Printf("%+v -> %+v\n", v.ID, v.Name)
	}
}

func (d *DiscordClient)GetChID(name string)(id string){
	for k,v := range d.channels{
		if k==name{id = v;return}
		if v==name{id = v;return}
	}
	return
}

func (d *DiscordClient)GetChName(id string)(name string){
	for k,v := range d.channels{
		if v==id{name = k;return}
	}
	return
}

func (d *DiscordClient)ChannelMessage(ch_name, text string)*discordgo.Message{
	d.client.Open()
	ch_id := d.GetChID(ch_name)
	if ch_id==""{return nil}
	m, _ := d.client.ChannelMessageSend(ch_id, text)
	return m
}
