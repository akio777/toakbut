package svc

import (
	"fmt"
	"time"

	atdModel "toakbut/pkg/attendance/model"
	"toakbut/pkg/bot/enum"
	"toakbut/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func (dc *DiscordClient) CheckIn(args []string, msg *discordgo.MessageCreate) {
	ss := dc.Session
	lengthArgs := len(args)
	if IsInWorkTypes(args[1]) {
		inTime := time.Now()
		if lengthArgs == 3 {
			_t, err := time.Parse("15:04", args[2])
			if err != nil {
				ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf("invalid time format : %s", err.Error()))
				dc.Log.Error(err)
				return
			}
			inTime = time.Date(inTime.Year(), inTime.Month(), inTime.Day(), _t.Hour(), _t.Minute(), 0, 0, inTime.Location())
		}
		workType := enum.WFO
		if utils.CompareUpperString(args[1], enum.WFH) {
			workType = enum.WFH
		}
		_, err := dc.Attendance.Create(&atdModel.Attendance{
			UserID:   msg.Author.ID,
			CheckIn:  &inTime,
			WorkType: workType,
		})
		if err != nil {
			ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+err.Error())
			dc.Log.Error(err)
			return
		}
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+fmt.Sprintf(" checked in at : %s", utils.TimeFormat(inTime.Local())))
		return
	} else {
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+enum.INVALID_IN_COMMANDS)
		return
	}
}
