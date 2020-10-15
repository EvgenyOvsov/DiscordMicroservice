package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	discordtoken = "Njk4ODE2MzA4MDk5NDE[DELETED]GFx64orW8OjbI7ZOQIwKs"
	password = "0x00-0xff"
)
var(
	Discord DiscordClient
	Version = "0.2"
)
func Parse (c *gin.Context) *Request{
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return nil
	}
	if req.Token==password{return &req}
	return nil
}

func main(){
	Discord.Init(discordtoken)
	Discord.client.AddHandler(MessageHandler)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", message)

	r.Run(":5001")
}

func message(c *gin.Context){
	if c.GetHeader("X-Gitlab-Event")!=""{WebhookHandler(c);return}

	req := Parse(c)
	if req==nil{return}
	if (req.Color!="" || req.Image!=""){
		var color int
		switch req.Color {
		default:
				color = embed_color_default
			case "red":
				color = embed_color_red
			case "yellow":
				color = embed_color_yellow
			case "green":
				color = embed_color_green
		}
		e := Embed{
			Color:       color,
			Title:       req.Title,
			Description: req.Text,
		}
		if req.Image!=""{
			e.AddThumbnail(req.Image)
		}
		if req.Channel!=""{
			e.SendEmbed(Discord.GetChID(req.Channel))
			return
		}
		ch, _ := Discord.client.UserChannelCreate(Discord.GetUser(req.To))
		e.SendEmbed(ch.ID)
		return
	}
	if req.Channel!=""{
		Discord.ChannelMessage(req.Channel, req.Text)
		return
	}
	Discord.MessageUser(req.To, req.Text)
	return

}