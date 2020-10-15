package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xanzy/go-gitlab"
	"io/ioutil"
)

func WebhookHandler(c *gin.Context){
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	event, err := gitlab.ParseHook(gitlab.HookEventType(c.Request), data)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch event := event.(type) {
		case *gitlab.MergeEvent:
			if event.ObjectAttributes.Action=="open" && event.ObjectAttributes.TargetBranch=="master"{

				m := Discord.ChannelMessage(Discord.GetChID("analyzer"), //TODO разные репозитории
				fmt.Sprintf("%v создал новый Merge Request из ветки **%v** в ветку **master**",
					event.User.Name,
					event.ObjectAttributes.SourceBranch))
				Discord.client.MessageReactionAdd(m.ChannelID, m.ID, "🚀")
			}
	}
}
