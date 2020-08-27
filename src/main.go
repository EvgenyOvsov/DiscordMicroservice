package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Token string `json:"token"`
	To string `json:"to"`
	Text string `json:"text"`
}

func Parse (c *gin.Context) *Request{
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil{
		return nil
		c.AbortWithStatus(http.StatusBadRequest)
	}
	return &req
}

func main(){
	var discord Discord
	discord.Init("...Og3vCBGFx64orW8OjbI7ZOQIwKs")
	discord.client.AddHandler(messageCreate)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		req := Parse(c)
		if req!=nil && req.Token=="0x00-0xff"{
			discord.SendMessage(req.To, req.Text)
		}
	})
	r.Run(":5001")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	//DoNothing
	return
}
