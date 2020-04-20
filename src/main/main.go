package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
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
	fmt.Println("Bot inited.\n Channels: ")
	d.channels = make(map[string]string)
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
	discord.Init("Njk4ODE2MzA4MDk...        ...kbZZFOiy7RI-hp6CuTk")
	//discord.client.AddHandler(messageCreate)
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		req := Parse(c)
		if req!=nil && req.Token=="0x00-0xff"{
			discord.SendMessage(req.To, req.Text)
		}

	})
	r.Run(":5001")
}

//func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
//	//todo some reactions on messages
//}
