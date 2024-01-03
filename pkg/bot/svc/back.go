package svc

import (
	"fmt"
	"time"

	breakModel "toakbut/pkg/breaks/model"
	"toakbut/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func (dc *DiscordClient) Back(args []string, msg *discordgo.MessageCreate) {
	ss := dc.Session
	lengthArgs := len(args)
	backTime := time.Now()
	if lengthArgs == 2 {
		_t, err := time.Parse("15:04", args[1])
		if err != nil {
			ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf("invalid time format : %s", err.Error()))
			dc.Log.Error(err)
			return
		}
		backTime = time.Date(backTime.Year(), backTime.Month(), backTime.Day(), _t.Hour(), _t.Minute(), 0, 0, backTime.Location())
	}
	_, err := dc.Breaks.Update(&breakModel.Breaks{
		UserID:   msg.Author.ID,
		BreakOut: &backTime,
	})
	if err != nil {
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+err.Error())
		dc.Log.Error(err)
		return
	}
	ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf(" break in at : %s", utils.TimeFormat(backTime.Local())))
}
