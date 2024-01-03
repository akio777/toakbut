package svc

import (
	"fmt"
	"time"

	breakModel "toakbut/pkg/breaks/model"
	"toakbut/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func (dc *DiscordClient) Break(args []string, msg *discordgo.MessageCreate) {
	ss := dc.Session
	lengthArgs := len(args)
	inTime := time.Now()
	if lengthArgs == 2 {
		_t, err := time.Parse("15:04", args[1])
		if err != nil {
			ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf("invalid time format : %s", err.Error()))
			dc.Log.Error(err)
			return
		}
		inTime = time.Date(inTime.Year(), inTime.Month(), inTime.Day(), _t.Hour(), _t.Minute(), 0, 0, inTime.Location())
	}
	_, err := dc.Breaks.Create(&breakModel.Breaks{
		UserID:  msg.Author.ID,
		BreakIn: &inTime,
	})
	if err != nil {
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+err.Error())
		dc.Log.Error(err)
		return
	}
	ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf(" break in at : %s", utils.TimeFormat(inTime.Local())))
}
