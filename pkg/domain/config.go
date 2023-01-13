package domain

const (
	BotToken            = "BotToken"
	GuildId             = "GuildId"
	RemoveCommands      = "RemoveCommands"
	NewsbotCollectorUri = "NewsbotCollectorUri"
)

type Config struct {
	BotToken            string
	GuildId             string
	RemoveCommands      bool
	NewsbotCollectorUri string
}
