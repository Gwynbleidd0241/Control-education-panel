package students

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

type StudentBase struct {
	ID          string     `json:"id"`
	FullName    string     `json:"fullName"`
	Email       string     `json:"email"`
	Age         int        `json:"age"`
	Performance string     `json:"performance"`
	PhotoURL    string     `json:"photoUrl"`
	City        string     `json:"city"`
	Phone       string     `json:"phone"`
	Format      string     `json:"format"`
	Progress    int        `json:"progress"`
	CourseID    *string    `json:"courseId,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
}

func (c *Client) List(ctx context.Context) ([]StudentBase, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/api/students", nil)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("students-service status %d", resp.StatusCode)
	}
	var out []StudentBase
	return out, json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) GetByID(ctx context.Context, id string) (*StudentBase, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/api/students/"+id, nil)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("students-service status %d", resp.StatusCode)
	}
	var out StudentBase
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (c *Client) Base() string {
	return c.base
}
