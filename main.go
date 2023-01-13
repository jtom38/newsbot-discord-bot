package main

import (
	"github.com/jtom38/newsbot-discord-bot/pkg/bot"
	"github.com/jtom38/newsbot-discord-bot/pkg/config"
)

func main() {
	//	integerOptionMinValue          = 1.0
	//	dmPermission                   = false
	//	defaultMemberPermissions int64 = discordgo.PermissionManageServer
	cfg := config.LoadConfig()

	bot.Boot(cfg)
}
