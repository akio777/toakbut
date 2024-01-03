package svc

import (
	"context"
	"fmt"
	"strings"
	atdSvc "toakbut/pkg/attendance/svc"
	"toakbut/pkg/bot/enum"
	breaksSvc "toakbut/pkg/breaks/svc"
	"toakbut/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type DiscordClient struct {
	Ctx        context.Context
	Db         *bun.DB
	Log        *logrus.Logger
	Session    *discordgo.Session
	ChannelID  string
	Attendance *atdSvc.Attendance
	Breaks     *breaksSvc.Breaks
}

var (
	dcc *DiscordClient
)

func Dcc() *DiscordClient {
	return dcc
}

func Init(token string, client DiscordClient) *discordgo.Session {
	dcc = &client
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		panic(err)
	}
	dcc.Session = dg
	dcc.Session.AddHandler(dcc.MessageHandler)
	err = dcc.Session.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		panic(err)
	}
	return dg
}

func (dc *DiscordClient) MessageHandler(ss *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == ss.State.User.ID || msg.ChannelID != dc.ChannelID {
		return
	}

	args := strings.Fields(msg.Content)
	lengthArgs := len(args)
	if lengthArgs > 0 {
		if IsInCommands(args[0]) {
			if utils.CompareUpperString(args[0], enum.PrefixCommand(enum.IN)) && lengthArgs >= 2 {
				// ! CHECK-IN
				dc.CheckIn(args, msg)
				return
			} else if utils.CompareUpperString(args[0], enum.PrefixCommand(enum.OUT)) && lengthArgs <= 2 {
				// ! CHECK-OUT
				dc.CheckOut(args, msg)
				return
			} else if utils.CompareUpperString(args[0], enum.PrefixCommand(enum.BREAK)) && lengthArgs <= 2 {
				// ! BREAK
				dc.Break(args, msg)
				return
			} else if utils.CompareUpperString(args[0], enum.PrefixCommand(enum.BACK)) && lengthArgs <= 2 {
				// ! BREAK
				dc.Back(args, msg)
				return
			}

		} else {
			ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+enum.INVALID_COMMANDS)
			return
		}
		ss.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+enum.EMPTY_COMMANDS)
		return
	}

}
