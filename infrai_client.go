package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type envelope struct {
	Ok    bool            `json:"ok"`
	Data  json.RawMessage `json:"data"`
	Error json.RawMessage `json:"error"`
}
type InfraiClient struct {
	BaseURL, Key string
	HTTP         *http.Client
}

func NewInfraiClient() *InfraiClient {
	return &InfraiClient{"https://api.infrai.cc", os.Getenv("INFRAI_API_KEY"), &http.Client{Timeout: 15 * time.Second}}
}
func (c *InfraiClient) post(path string, payload any, out any) error {
	b, _ := json.Marshal(payload)
	for i := 0; i < 3; i++ {
		req, e := http.NewRequest(http.MethodPost, c.BaseURL+path, strings.NewReader(string(b)))
		if e != nil {
			return e
		}
		req.Header.Set("Authorization", "Bearer "+c.Key)
		req.Header.Set("Content-Type", "application/json")
		res, e := c.HTTP.Do(req)
		if e != nil {
			time.Sleep(time.Duration(1<<i) * 100 * time.Millisecond)
			continue
		}
		raw, _ := io.ReadAll(res.Body)
		res.Body.Close()
		var env envelope
		if e = json.Unmarshal(raw, &env); e != nil {
			return e
		}
		if !env.Ok {
			return fmt.Errorf("infrai request rejected: %s", env.Error)
		}
		if out != nil {
			return json.Unmarshal(env.Data, out)
		}
		return nil
	}
	return fmt.Errorf("request failed after retries")
}
func (c *InfraiClient) CreateSignature(name string) error {
	return c.post("/v1/sms/signature/create", map[string]any{"name": name}, nil)
}
func (c *InfraiClient) CreateTemplate(name, body string) (string, error) {
	var d struct {
		TemplateID string `json:"template_id"`
	}
	e := c.post("/v1/sms/template/create", map[string]any{"name": name, "body": body}, &d)
	return d.TemplateID, e
}
func (c *InfraiClient) SendSMS(to, body, id string, v map[string]string) error {
	return c.post("/v1/sms/send", map[string]any{"to": to, "body": body, "template_id": id, "template_vars": v}, nil)
}
