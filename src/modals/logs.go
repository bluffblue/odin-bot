package modals

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var LogChannelID string

func InitLogChannel(channelID string) {
	LogChannelID = channelID
}

func SendCommandLog(s *discordgo.Session, i *discordgo.InteractionCreate, commandName string, success bool, details string) {
	if LogChannelID == "" {
		fmt.Println("Warning: LOG_CHANNEL_ID not set")
		return
	}

	logMessage := fmt.Sprintf("```\nCommand: %s\nUser: %s (%s)\nChannel: %s\nTime: %s\nStatus: %s\nDetails: %s\n```",
		commandName,
		i.Member.User.Username,
		i.Member.User.ID,
		i.ChannelID,
		time.Now().Format("2006-01-02 15:04:05"),
		getStatus(success),
		details,
	)

	_, err := s.ChannelMessageSend(LogChannelID, logMessage)
	if err != nil {
		fmt.Printf("Error sending log message: %v\n", err)
	}
}

func SendEphemeralResponse(s *discordgo.Session, i *discordgo.InteractionCreate, content string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func getStatus(success bool) string {
	if success {
		return "✅ Success"
	}
	return "❌ Failed"
}
