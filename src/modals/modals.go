package modals

import "github.com/bwmarrin/discordgo"

var SlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:                     "ban",
		Description:             "Ban a user from the server",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to ban",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "Reason for the ban",
				Required:    false,
			},
		},
	},
	{
		Name:                     "unban",
		Description:             "Unban a user from the server",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to unban",
				Required:    true,
			},
		},
	},
	{
		Name:                     "kick",
		Description:             "Kick a user from the server",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to kick",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "Reason for the kick",
				Required:    false,
			},
		},
	},
	{
		Name:                     "warn",
		Description:             "Warn a user",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to warn",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "Reason for the warning",
				Required:    false,
			},
		},
	},
	{
		Name:                     "unwarn",
		Description:             "Remove a warning from a user",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to remove warning from",
				Required:    true,
			},
		},
	},
	{
		Name:                     "clearmsg",
		Description:             "Clear messages in the current channel",
		DefaultMemberPermissions: &AdminPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "amount",
				Description: "Number of messages to delete (max 100)",
				Required:    true,
				MinValue:    &MinMessageAmount,
				MaxValue:    MaxMessageAmount,
			},
		},
	},
}

var (
	AdminPermission   int64 = discordgo.PermissionAdministrator
	MinMessageAmount  = float64(1)
	MaxMessageAmount  = float64(100)
)
