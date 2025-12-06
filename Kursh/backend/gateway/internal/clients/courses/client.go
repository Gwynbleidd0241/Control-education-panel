package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	base string
	http *http.Client
}

func New(base string) *Client {
	return &Client{
		base: base,
		http: &http.Client{Timeout: 5 * time.Second},
	}
}

type CourseInfo struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Level string  `json:"level"`
}

func (c *Client) GetByID(ctx context.Context, id string) (*CourseInfo, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/courses/"+id, nil)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("courses-service status %d", resp.StatusCode)
	}

	var out CourseInfo
	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) List(ctx context.Context) ([]CourseInfo, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/courses", nil)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("courses-service status %d", resp.StatusCode)
	}
	var out []CourseInfo
	return out, json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) Base() string {
	return c.base
}
