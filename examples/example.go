package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/nickname32/discordhook"
)

func main() {
	for {
		embed := discordhook.Embed{
			Title:       "title",
			Color:       0x00ff00,
			Description: "val",
		}

		webhookurl, _ := url.Parse("https://discord.com/api/webhooks/991684601611300965/8t_QjeOBmH_7oOGjvvufHMBQyt3ZV8L_dWEmGRZ8M0HRF9gaiZIwXsEXavCtcsoJxSGX")
		hook := discordhook.WebhookAPI{
			URL:    webhookurl,
			Client: &http.Client{},
		}
		hookpre := discordhook.WebhookExecuteParams{
			Embeds: []*discordhook.Embed{&embed},
		}
		hook.Execute(context.Background(), &hookpre, nil, "")
	}
}
