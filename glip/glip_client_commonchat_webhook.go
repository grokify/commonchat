package glip

import (
	"encoding/json"
	"fmt"

	glipwebhook "github.com/grokify/go-glip"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"

	"github.com/grokify/commonchat"
	"github.com/grokify/commonchat/glip/classic"
	"github.com/grokify/commonchat/glip/config"
)

type GlipAdapter struct {
	GlipClient      glipwebhook.GlipWebhookClient
	CommonConverter classic.GlipMessageConverter
	EmojiURLFormat  string
	WebhookURLOrUID string
}

func NewGlipAdapter(webhookURLOrUID string, cfg *config.ConverterConfig) (*GlipAdapter, error) {
	glip, err := glipwebhook.NewGlipWebhookClient(webhookURLOrUID)
	if err != nil {
		return nil, err
	}
	return &GlipAdapter{
		GlipClient:      glip,
		WebhookURLOrUID: webhookURLOrUID,
		CommonConverter: classic.NewGlipMessageConverter(cfg)}, err
}

func (adapter *GlipAdapter) SendWebhook(urlOrUid string, message commonchat.Message, glipmsg interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
	glipMessage := adapter.CommonConverter.ConvertCommonMessage(message)
	glipmsg = &glipMessage

	glipMessageBytes, err := json.Marshal(glipMessage)
	if err == nil {
		log.Info().
			Str("event", "outgoing.webhook.glip").
			Str("handler", "Glip Adapter").
			Msg(string(glipMessageBytes))

	}
	if 1 == 0 {
		fmt.Println(string(glipMessageBytes))
		glipMessageJson, err := json.MarshalIndent(glipMessage, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(glipMessageJson))
	}
	return adapter.GlipClient.PostWebhookGUIDFast(urlOrUid, glipMessage)
}

func (adapter *GlipAdapter) SendMessage(message commonchat.Message, glipmsg interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
	return adapter.SendWebhook(adapter.WebhookURLOrUID, message, glipmsg)
}

func (adapter *GlipAdapter) WebhookUID(ctx *fasthttp.RequestCtx) (string, error) {
	webhookUID := adapter.WebhookURLOrUID
	return webhookUID, nil
}
