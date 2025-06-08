package slack

import (
	"fmt"

	"github.com/grokify/commonchat"
	"github.com/valyala/fasthttp"
)

type SlackAdapter struct {
	SlackClient     SlackWebhookClient
	EmojiURLFormat  string
	WebhookURLOrUID string
}

func NewSlackAdapter(webhookURLOrUID string) (*SlackAdapter, error) {
	client, err := NewSlackWebhookClient(webhookURLOrUID, FastHTTP)
	return &SlackAdapter{
		SlackClient:     client,
		WebhookURLOrUID: webhookURLOrUID}, err
}

func (adapter *SlackAdapter) SendWebhook(urlOrUID string, ccMsg commonchat.Message, slackmsg any, opts map[string]any) (*fasthttp.Request, *fasthttp.Response, error) {
	slackMessage := ConvertCommonMessage(ccMsg)
	slackmsg = &slackMessage //nolint:ineffassign // slackmsg is meant to be a pointer
	return adapter.SlackClient.PostWebhookFast(urlOrUID, slackMessage)
}

func (adapter *SlackAdapter) SendMessage(message commonchat.Message, slackmsg any, opts map[string]any) (*fasthttp.Request, *fasthttp.Response, error) {
	return adapter.SendWebhook(adapter.WebhookURLOrUID, message, slackmsg, opts)
}

func (adapter *SlackAdapter) WebhookUID(ctx *fasthttp.RequestCtx) (string, error) {
	webhookUID := fmt.Sprintf("%s", ctx.UserValue("webhookuid"))
	return webhookUID, nil
}
