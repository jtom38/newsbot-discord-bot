package collector

import "github.com/google/uuid"

type CollectorApi interface {
	Articles() ArticlesApi
	Sources() SourcesApi
	Outputs() OutputsApi
	Subscriptions() SubscriptionsApi
}

type ArticlesApi interface {
	List() ([]Article, error)
	ListPointers() (*[]Article, error)
	Get(ID uuid.UUID) (*Article, error)
	ListBySourceId(ID uuid.UUID) (*[]Article, error)
}

type SourcesApi interface {
	List() (*ListSourcesResult, error)
	ListBySource(value string) (*ListSourcesResult, error)

	GetById(ID uuid.UUID) (*SingleSourcesResult, error)
	GetBySourceAndName(SourceName string, Name string) (*SingleSourcesResult, error)

	NewReddit(name string, sourceUrl string) error
	NewYouTube(name string, url string) error
	NewTwitch(Name string) error
	Delete(ID uuid.UUID) error
	Disable(ID uuid.UUID) error
	Enable(ID uuid.UUID) error
}

type OutputsApi interface {
	DiscordWebHook() OutputDiscordWebHookApi
}

type OutputDiscordWebHookApi interface {
	List() (*[]Discordwebhook, error)
	Get(id uuid.UUID) (*Discordwebhook, error)
	Delete(id uuid.UUID) error
	Disable(id uuid.UUID) error
	Enable(id uuid.UUID) error
	New(server string, channel string, url string) error
	GetByServerAndChannel(server string, channel string) ([]Discordwebhook, error)
}

type SubscriptionsApi interface {
	List() (*[]Subscription, error)
	GetByDiscordID(ID uuid.UUID) (*[]Subscription, error)
	GetBySourceID(ID uuid.UUID) (*[]Subscription, error)
	New(DiscordID uuid.UUID, SourceID uuid.UUID) error
	Delete(ID uuid.UUID) error
}

type QueueApi interface {
	ListDiscordWebHooks()
}
