package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)
const (
	embed_color_red = 16711680
	embed_color_green = 65280
	embed_color_yellow = 16776960
	embed_color_default = 0
)
type Embed struct {
	Color int
	Title string
	Description string
	Footer *discordgo.MessageEmbedFooter
	Thumbnail *discordgo.MessageEmbedThumbnail
	Fields []*discordgo.MessageEmbedField
	Image *discordgo.MessageEmbedImage
}
func(e *Embed) AddThumbnail (url string){
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
}
func(e *Embed) AddImage (url string){
	e.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}
}
func(e *Embed) AddFields (fileds map[string]interface{}){
	for k,v := range fileds{
		e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
			Name: k,
			Value: fmt.Sprintf("%+v", v),
		})
	}
	fmt.Printf("%+v", e.Fields[0])
}
func (e *Embed)SendEmbed(id string){
	Discord.client.ChannelMessageSendEmbed(id, &discordgo.MessageEmbed{
		Title:       e.Title,
		Description: e.Description,
		Color:       e.Color,
		Footer: 	 e.Footer,
		Image:       e.Image,
		Thumbnail:   e.Thumbnail,
		Fields: e.Fields,
	})
}
