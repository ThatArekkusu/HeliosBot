package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	ChannelID = "1332105742785708054"
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(userInfo)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func userInfo(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	member := m.Member
	user := member.User
	createdate, _ := discordgo.SnowflakeTimestamp(user.ID)

	haveAvatar := user.Avatar != ""

	currentTime := time.Now()
	dif := currentTime.Sub(createdate)
	accountAge := int(dif.Hours())

	var trigger int
	switch {
	case accountAge < 1:
		trigger = 1
	case accountAge < 24:
		trigger = 2
	case accountAge < 168:
		trigger = 3
	case accountAge < 336:
		trigger = 4
	default:
		trigger = 0
	}

	switch trigger {
	case 1:
		if haveAvatar {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is suspiciously new. Created less than an hour ago and has no profile picture.")
		} else {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is suspiciously new. Created less than an hour ago.")
		}
	case 2:
		if haveAvatar {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is new. Created less than a day ago and has no profile picture.")
		} else {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is new. Created less than a day ago.")
		}
	case 3:
		if haveAvatar {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is relatively new. Created less than a week ago and has no profile picture.")
		} else {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is relatively new. Created less than a week ago.")
		}
	case 4:
		if haveAvatar {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is new. Created less than two weeks ago and has no profile picture.")
		} else {
			s.ChannelMessageSend(ChannelID, user.Username+"'s account is new. Created less than two weeks ago.")
		}
	case 0:
		s.ChannelMessageSend(ChannelID, user.Username+"'s account is old.")
	}
}
