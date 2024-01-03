package svc

import (
	"fmt"
	"time"

	atdModel "toakbut/pkg/attendance/model"
	"toakbut/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func (dc *DiscordClient) CheckOut(args []string, msg *discordgo.MessageCreate) {
	ss := dc.Session
	lengthArgs := len(args)
	outTime := time.Now()
	if lengthArgs == 2 {
		_t, err := time.Parse("15:04", args[1])
		if err != nil {
			ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf("invalid time format : %s", err.Error()))
			dc.Log.Error(err)
			return
		}
		outTime = time.Date(outTime.Year(), outTime.Month(), outTime.Day(), _t.Hour(), _t.Minute(), 0, 0, outTime.Location())
	}
	_, err := dc.Attendance.Update(&atdModel.Attendance{
		UserID:   msg.Author.ID,
		CheckOut: &outTime,
	})
	if err != nil {
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+err.Error())
		dc.Log.Error(err)
		return
	}
	ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf(" checked out at : %s", utils.TimeFormat(outTime.Local())))
}
