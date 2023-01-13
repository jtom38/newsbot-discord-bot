package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/jtom38/newsbot-discord-bot/pkg/domain"
)

type ConfigClient struct{}

func New() ConfigClient {
	c := ConfigClient{}
	c.refreshEnv()

	return c
}

// This will load the config from the ENV and return a struct with the values
func LoadConfig() domain.Config {
	c := New()
	cfg := domain.Config{
		BotToken:            c.String(domain.BotToken),
		GuildId:             c.String(domain.GuildId),
		RemoveCommands:      c.Bool(domain.RemoveCommands),
		NewsbotCollectorUri: c.String(domain.NewsbotCollectorUri),
	}
	return cfg
}

func (cc ConfigClient) String(key string) string {
	res, filled := os.LookupEnv(key)
	if !filled {
		return ""
	}
	return res
}

func (cc ConfigClient) Bool(flag string) bool {
	cc.refreshEnv()

	res, filled := os.LookupEnv(flag)
	if !filled {
		return false
	}

	b, err := strconv.ParseBool(res)
	if err != nil {
		return false
	}
	return b
}

// Use this when your ConfigClient has been opened for awhile and you want to ensure you have the most recent env changes.
func (cc *ConfigClient) refreshEnv() {
	// Check to see if we have the env file on the system
	_, err := os.Stat(".env")

	// We have the file, load it.
	if err == nil {
		_, err := os.Open(".env")
		if err == nil {
			loadEnvFile()
		}
	}
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}
