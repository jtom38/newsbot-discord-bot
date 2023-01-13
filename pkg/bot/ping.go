package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	SlashCommandPing = "ping"
)

func init() {
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        SlashCommandPing,
		Description: "Basic response to validate the bot is online.",
		Type:        discordgo.ChatApplicationCommand,
	})

	Interactions[SlashCommandPing] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Printf("'%v' requested %v", i.Interaction.Member.User.Username, SlashCommandPing)
		s.InteractionRespond(i.Interaction, PingResponse(i))
	}
}

func PingResponse(i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong! 👋",
		},
	}
}
