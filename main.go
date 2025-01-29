package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	ChannelID = os.Getenv("ChannelID")
)

func init() {
	flag.StringVar(&Token, "t", "", os.Getenv("Token"))
	flag.Parse()
}

func main() {
	// Create a new Discord session with the necessary intents
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Enable the GUILD_MEMBERS intent
	dg.Identify.Intents = discordgo.IntentsGuildMembers

	// Register the event handler for GuildMemberAdd
	dg.AddHandler(userInfo)

	// Open the WebSocket connection to Discord
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

	hasNoAvatar := user.Avatar == ""

	currentTime := time.Now()
	dif := currentTime.Sub(createdate)
	accountAge := int(dif.Hours())

	var ageDesc string
	switch {
	case accountAge < 1:
		ageDesc = "suspiciously new. Created less than an hour ago"
	case accountAge < 24:
		ageDesc = "new. Created less than a day ago"
	case accountAge < 168:
		ageDesc = "relatively new. Created less than a week ago"
	case accountAge < 336:
		ageDesc = "new. Created less than two weeks ago"
	default:
	}

	msg := fmt.Sprintf("%s's account is %s", user.Username, ageDesc)
	if hasNoAvatar {
		msg += " and has no profile picture"
	}

	msg += "."

	s.ChannelMessageSend(ChannelID, msg)
}
