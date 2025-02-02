package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageRoles

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "KickDate",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Ban a user based on their account creation date",
		},
	}
)
