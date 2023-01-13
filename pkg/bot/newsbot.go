package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jtom38/newsbot-discord-bot/pkg/api/newsbot/collector"
)

const (
	SlashCommandNewsbot = "newsbot"
)

func init() {
	// adds the command to the array to be sent
	log.Printf("loading command /newsbot")
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "newsbot",
		Description: "Interfaces with the newsbot API",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Description: "Returns the newest articles from the server",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	})

	// when a user requests a command, how do we respond?
	// This uses an map to find the command quickly rather then over an array.
	Interactions[SlashCommandNewsbot] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Printf("'%v' requested %v", i.Interaction.Member.User.Username, SlashCommandNewsbot)
		s.InteractionRespond(i.Interaction, NewsbotListResponse(i))
	}
}

func NewsbotListResponse(i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	var embeds []*discordgo.MessageEmbed

	msg := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
	}

	items, err := NewsbotCollector.Articles().List()
	if err != nil {
		msg.Data.Content = err.Error()
		return &msg
	}

	//	source, err := nbc.Sources().GetById(items[0].Sourceid)
	//	if err != nil {
	//		msg.Data.Content = err.Error()
	//	}

	embeds = append(embeds, generateArticleEmbed(items[0]))

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	}
}

func generateArticleEmbed(item collector.Article) *discordgo.MessageEmbed {
	e := &discordgo.MessageEmbed{
		Title: item.Title,
		URL:   item.Url,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    item.Authorname,
			IconURL: item.Authorimage,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Newsbot",
		},
	}

	// if the content collect was a picture, display it
	if strings.Contains(item.Description, ".jpg") {
		e.Image = &discordgo.MessageEmbedImage{
			URL: item.Description,
		}
	} else {
		e.Description = item.Description
	}

	return e
}
