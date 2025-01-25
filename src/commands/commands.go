package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HandleBan(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	err := s.GuildBanCreate(i.GuildID, targetUser.ID, 0)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to ban user: " + err.Error(),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: targetUser.Username + " has been banned. Reason: " + reason,
		},
	})
}

func HandleUnban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)

	err := s.GuildBanDelete(i.GuildID, targetUser.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to unban user: " + err.Error(),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: targetUser.Username + " has been unbanned",
		},
	})
}

func HandleKick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	err := s.GuildMemberDelete(i.GuildID, targetUser.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to kick user: " + err.Error(),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: targetUser.Username + " has been kicked. Reason: " + reason,
		},
	})
}

func HandleWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: targetUser.Username + " has been warned. Reason: " + reason,
		},
	})
}

func HandleUnwarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Warning has been removed for " + targetUser.Username,
		},
	})
}
