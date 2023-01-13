package collector

import "net/http"

type QueueClient struct {
	apiServer string
	routeRoot string

	client *http.Client
	rest   RestClient
}

func NewQueueClient(serverAddress string, client *http.Client) QueueClient {
	c := QueueClient{
		apiServer: serverAddress,
		routeRoot: "api/queue",
		client:    client,
	}

	return c
}

func (c QueueClient) ListDiscordWebHooks() {

}
