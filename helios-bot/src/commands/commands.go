package commands

import (
	"github.com/bwmarrin/discordgo"
)

// integerOptionMinValue          = 1.0
// dmPermission                   = false
// defaultMemberPermissions int64 = discordgo.PermissionManageRoles

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "KickDate",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Assigns how old a user must be to be kicked",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "date",
				Description: "The date to kick users before",
				Required:    true,
			},
		},
	},
}
