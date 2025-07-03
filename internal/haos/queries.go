package haos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	client        *http.Client
	token         string
	haosUrl       string
	alarmEntityId string
}

func (c *Client) GetAlarmState() (State, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/states/%s", c.haosUrl, c.alarmEntityId), nil)

	if err != nil {
		return "", fmt.Errorf("error when created request for HAOS: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error happened while fetching alarm status: %w", err)
	}

	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("couldn't read response from HAOS query: %w", err)
	}

	var e Entity

	if err := json.Unmarshal(resp_body, &e); err != nil {
		return "", fmt.Errorf("could not unmarshal haos response: %w", err)
	}

	return e.State, nil
}

func NewClient() *Client {
	return &Client{
		client:        &http.Client{},
		token:         os.Getenv("HAOS_TOKEN"),
		haosUrl:       os.Getenv("HAOS_URL"),
		alarmEntityId: os.Getenv("HAOS_ALARM"),
	}
}
