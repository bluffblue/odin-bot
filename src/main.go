package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	cmdhandlers "discord-bot/src/commands"
	"discord-bot/src/modals"
	"discord-bot/src/cache"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BotToken    string
	GuildID     string
	LogChannelID string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	BotToken = os.Getenv("BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")
	LogChannelID = os.Getenv("LOG_CHANNEL_ID")

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	channels := []string{"1332680730739740746", "1332681078892400674", "1332680730739740749"}
	cache.InitializeAutoCacheCleanup(dg, GuildID, channels)
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ban":      cmdhandlers.HandleBan,
	"unban":    cmdhandlers.HandleUnban,
	"kick":     cmdhandlers.HandleKick,
	"warn":     cmdhandlers.HandleWarn,
	"unwarn":   cmdhandlers.HandleUnwarn,
	"clearmsg": cmdhandlers.HandleClearMessages,
}

func main() {
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	modals.InitLogChannel(LogChannelID)

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

	registeredCommands := make([]*discordgo.ApplicationCommand, len(modals.SlashCommands))
	for i, cmd := range modals.SlashCommands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, GuildID, cmd)
		if err != nil {
			log.Printf("Error creating command %v: %v", modals.SlashCommands[i].Name, err)
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
