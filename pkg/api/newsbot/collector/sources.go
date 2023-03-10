package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type SourcesApiClient struct {
	apiServer string
	routeRoot string

	client *http.Client
	rest   RestClient
}

func NewSourcesApiClient(serverAddress string, client *http.Client) SourcesApi {
	c := SourcesApiClient{
		apiServer: serverAddress,
		routeRoot: "api/sources/",
		client:    client,
	}
	return c
}

type ListSourcesResult struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Payload []Source `json:"payload"`
}

func (c SourcesApiClient) List() (*ListSourcesResult, error) {
	var items ListSourcesResult

	uri := fmt.Sprintf("%v/%v", c.apiServer, c.routeRoot)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	//for _, i := range result {
	//	items = append(items, c.convertFromDto(i))
	//}

	return &items, nil
}

func (c SourcesApiClient) ListBySource(value string) (*ListSourcesResult, error) {
	var items ListSourcesResult

	uri := fmt.Sprintf("%v/%v/by/source?source=%v", c.apiServer, c.routeRoot, value)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	//for _, i := range result {
	//	items = append(items, c.convertFromDto(i))
	//}

	return &items, nil
}

type SingleSourcesResult struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Payload Source `json:"payload"`
}

func (c SourcesApiClient) GetById(ID uuid.UUID) (*SingleSourcesResult, error) {
	var items SingleSourcesResult

	uri := fmt.Sprintf("%v/%v/%v", c.apiServer, c.routeRoot, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}

func (c SourcesApiClient) GetBySourceAndName(SourceName string, Name string) (*SingleSourcesResult, error) {
	var items SingleSourcesResult

	uri := fmt.Sprintf("%v/%v/by/sourceAndName?source=%v&name=%v", c.apiServer, c.routeRoot, SourceName, Name)

	res, err := c.rest.Get(context.Background(), RestArgs{
		Url:        uri,
		StatusCode: 200,
	})
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}

func (c SourcesApiClient) NewReddit(name string, sourceUrl string) error {
	endpoint := fmt.Sprintf("%v/%v/new/reddit?name=%v&url=%v", c.apiServer, c.routeRoot, name, url.QueryEscape(sourceUrl))
	res, err := http.Post(endpoint, "application/json", nil)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) NewYouTube(Name string, Url string) error {
	endpoint := fmt.Sprintf("%v/%v/new/youtube?name=%v&url=%v", c.apiServer, c.routeRoot, Name, url.QueryEscape(Url))
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) NewTwitch(Name string) error {
	endpoint := fmt.Sprintf("%v/%v/new/twitch?name=%v", c.apiServer, c.routeRoot, Name)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) Delete(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v", c.apiServer, c.routeRoot, ID)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Disable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v/disable", c.apiServer, c.routeRoot, ID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Enable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v/enable", c.apiServer, c.routeRoot, ID)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}
