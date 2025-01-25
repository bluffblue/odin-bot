package commands

import (
	"fmt"
	"discord-bot/src/modals"
	"github.com/bwmarrin/discordgo"
)

func checkAdmin(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	member := i.Member
	if member == nil {
		return false
	}
	
	perms := member.Permissions
	if perms&discordgo.PermissionAdministrator == 0 {
		_ = modals.SendEphemeralResponse(s, i, "You need administrator permissions to use this command.")
		modals.SendCommandLog(s, i, i.ApplicationCommandData().Name, false, "Insufficient permissions")
		return false
	}
	return true
}

func HandleBan(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	err := s.GuildBanCreate(i.GuildID, targetUser.ID, 0)
	if err != nil {
		_ = modals.SendEphemeralResponse(s, i, "Failed to ban user: "+err.Error())
		modals.SendCommandLog(s, i, "ban", false, fmt.Sprintf("Failed to ban %s: %v", targetUser.Username, err))
		return
	}

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully banned %s. Reason: %s", targetUser.Username, reason))
	modals.SendCommandLog(s, i, "ban", true, fmt.Sprintf("Banned %s. Reason: %s", targetUser.Username, reason))
}

func HandleUnban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)

	err := s.GuildBanDelete(i.GuildID, targetUser.ID)
	if err != nil {
		_ = modals.SendEphemeralResponse(s, i, "Failed to unban user: "+err.Error())
		modals.SendCommandLog(s, i, "unban", false, fmt.Sprintf("Failed to unban %s: %v", targetUser.Username, err))
		return
	}

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully unbanned %s", targetUser.Username))
	modals.SendCommandLog(s, i, "unban", true, fmt.Sprintf("Unbanned %s", targetUser.Username))
}

func HandleKick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	err := s.GuildMemberDelete(i.GuildID, targetUser.ID)
	if err != nil {
		_ = modals.SendEphemeralResponse(s, i, "Failed to kick user: "+err.Error())
		modals.SendCommandLog(s, i, "kick", false, fmt.Sprintf("Failed to kick %s: %v", targetUser.Username, err))
		return
	}

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully kicked %s. Reason: %s", targetUser.Username, reason))
	modals.SendCommandLog(s, i, "kick", true, fmt.Sprintf("Kicked %s. Reason: %s", targetUser.Username, reason))
}

func HandleWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)
	reason := ""
	if len(options) > 1 {
		reason = options[1].StringValue()
	}

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully warned %s. Reason: %s", targetUser.Username, reason))
	modals.SendCommandLog(s, i, "warn", true, fmt.Sprintf("Warned %s. Reason: %s", targetUser.Username, reason))
}

func HandleUnwarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	targetUser := options[0].UserValue(s)

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully removed warning from %s", targetUser.Username))
	modals.SendCommandLog(s, i, "unwarn", true, fmt.Sprintf("Removed warning from %s", targetUser.Username))
}

func HandleClearMessages(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkAdmin(s, i) {
		return
	}

	options := i.ApplicationCommandData().Options
	amount := int(options[0].IntValue())

	messages, err := s.ChannelMessages(i.ChannelID, amount, "", "", "")
	if err != nil {
		_ = modals.SendEphemeralResponse(s, i, "Failed to fetch messages: "+err.Error())
		modals.SendCommandLog(s, i, "clearmsg", false, fmt.Sprintf("Failed to fetch %d messages: %v", amount, err))
		return
	}

	messageIDs := make([]string, len(messages))
	for i, msg := range messages {
		messageIDs[i] = msg.ID
	}

	err = s.ChannelMessagesBulkDelete(i.ChannelID, messageIDs)
	if err != nil {
		_ = modals.SendEphemeralResponse(s, i, "Failed to delete messages: "+err.Error())
		modals.SendCommandLog(s, i, "clearmsg", false, fmt.Sprintf("Failed to delete %d messages: %v", amount, err))
		return
	}

	_ = modals.SendEphemeralResponse(s, i, fmt.Sprintf("Successfully deleted %d messages", len(messageIDs)))
	modals.SendCommandLog(s, i, "clearmsg", true, fmt.Sprintf("Deleted %d messages in channel %s", len(messageIDs), i.ChannelID))
}
