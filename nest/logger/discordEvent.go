package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type DiscordEventConfig struct {
	WebhookURL string
	Error      bool
	Warn       bool
	Info       bool
	Debug      bool
}

func DiscordEvent(config DiscordEventConfig) Events {
	events := Events{}

	if config.Error {
		events.Error = func(ctx context.Context, r Record) {
			payload, _ := json.Marshal(getEmbedPayload(r))
			http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(payload))
		}
	}

	if config.Warn {
		events.Warn = func(ctx context.Context, r Record) {
			payload, _ := json.Marshal(getEmbedPayload(r))
			http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(payload))
		}
	}

	if config.Info {
		events.Info = func(ctx context.Context, r Record) {
			payload, _ := json.Marshal(getEmbedPayload(r))
			http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(payload))
		}
	}

	if config.Debug {
		events.Debug = func(ctx context.Context, r Record) {
			payload, _ := json.Marshal(getEmbedPayload(r))
			http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(payload))
		}
	}
	return events

}

func getEmbedPayload(r Record) map[string]interface{} {

	var color int

	switch r.Level {
	case LevelInfo:
		color = 0x00ff00
	case LevelWarn:
		color = 0xffd700
	case LevelError:
		color = 0xff0000
	case LevelDebug:
		color = 0x0000ff
	}

	return map[string]interface{}{

		"embeds": []map[string]interface{}{
			{
				"title":       r.Level.String(),
				"description": r.Message,
				"color":       color,
				"fields": append(r.Attributes, Attribute{
					Name:  "Time",
					Value: r.Time.String(),
				}),
			},
		},
	}
}
