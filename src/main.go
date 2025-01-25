package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	cmdhandlers "discord-bot/src/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BotToken string
	GuildID  string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	BotToken = os.Getenv("BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")
}

var (
	slashCommands = []*discordgo.ApplicationCommand{
		{
			Name:        "ban",
			Description: "Ban a user from the server",
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
			Name:        "unban",
			Description: "Unban a user from the server",
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
			Name:        "kick",
			Description: "Kick a user from the server",
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
			Name:        "warn",
			Description: "Warn a user",
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
			Name:        "unwarn",
			Description: "Remove a warning from a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to remove warning from",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ban":    cmdhandlers.HandleBan,
		"unban":  cmdhandlers.HandleUnban,
		"kick":   cmdhandlers.HandleKick,
		"warn":   cmdhandlers.HandleWarn,
		"unwarn": cmdhandlers.HandleUnwarn,
	}
)

func main() {
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(slashCommands))
	for i, cmd := range slashCommands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, GuildID, cmd)
		if err != nil {
			log.Printf("Error creating command %v: %v", slashCommands[i].Name, err)
			continue
		}
		registeredCommands[i] = cmd
	}

	fmt.Println("Bot is running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	for _, cmd := range registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, GuildID, cmd.ID)
		if err != nil {
			log.Printf("Error deleting command %v: %v", cmd.Name, err)
		}
	}

	dg.Close()
}
