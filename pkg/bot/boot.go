package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/jtom38/newsbot-discord-bot/pkg/api/newsbot/collector"
	"github.com/jtom38/newsbot-discord-bot/pkg/domain"
)

var (
	s            *discordgo.Session
	Commands     []*discordgo.ApplicationCommand
	Interactions map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

	NewsbotCollector collector.CollectorApi
)

// If this init does not run first, things will fail
func init() {
	Interactions = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
}

// This exists as a staged object for cobra
type BootParams struct {
	Token          string
	GuildId        string
	RemoveCommands bool
}

func Boot(params domain.Config) {
	s, err := discordgo.New("Bot " + params.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	NewsbotCollector = collector.New(params.NewsbotCollectorUri)

	// when a command is sent to the bot, this is how to figure out what to do with it.
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// This extracts information sent by the user
		req := i.ApplicationCommandData()
		// query the interactions map and find the one that matches the name
		if h, ok := Interactions[req.Name]; ok {
			// redirect to that func
			h(s, i)
		}
	})

	// add the request to the session to login
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// open the connection
	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	gui := params.GuildId
	for i, v := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, gui, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if params.RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, params.GuildId, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
