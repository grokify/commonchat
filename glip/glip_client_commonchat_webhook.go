package glip

import (
	"encoding/json"

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
	WebhookURLOrUID string
}

func NewGlipAdapter(webhookURLOrUID string, cfg *config.ConverterConfig) *GlipAdapter {
	glip := glipwebhook.NewGlipWebhookClient(webhookURLOrUID, 1)
	return &GlipAdapter{
		GlipClient:      glip,
		WebhookURLOrUID: glip.WebhookURL,
		CommonConverter: classic.NewGlipMessageConverter(cfg)}
}

func NewGlipAdapterMSI(webhookURLOrUID string, cfg map[string]interface{}) (*GlipAdapter, error) {
	ccfg, err := config.NewConverterConfigMSI(cfg)
	if err != nil {
		return nil, err
	}
	return NewGlipAdapter(webhookURLOrUID, ccfg), nil
}

func (adapter *GlipAdapter) Clone() *GlipAdapter {
	return NewGlipAdapter(
		adapter.WebhookURLOrUID,
		adapter.CommonConverter.Config.Clone())
}

func (adapter *GlipAdapter) SendWebhook(urlOrUID string, message commonchat.Message, glipmsg interface{}, cfg map[string]interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
	if len(cfg) > 0 {
		newCfg, err := adapter.CommonConverter.Config.UpsertMSI(cfg)
		if err != nil {
			return nil, nil, err
		}
		thisAdapter := NewGlipAdapter(adapter.WebhookURLOrUID, newCfg)
		return thisAdapter.SendWebhook(urlOrUID, message, glipmsg, map[string]interface{}{})
	}
	glipMessage := adapter.CommonConverter.ConvertCommonMessage(message)
	glipmsg = &glipMessage

	glipMessageBytes, err := json.Marshal(glipMessage)
	if err != nil {
		return nil, nil, err
	}
	log.Info().
		Str("event", "outgoing.webhook.glip").
		Str("handler", "Glip Adapter").
		Msg(string(glipMessageBytes))
	return adapter.GlipClient.PostWebhookGUIDFast(urlOrUID, glipMessage)
}

func (adapter *GlipAdapter) SendMessage(message commonchat.Message, glipmsg interface{}, opts map[string]interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
	return adapter.SendWebhook(adapter.WebhookURLOrUID, message, glipmsg, opts)
}

func (adapter *GlipAdapter) WebhookUID(ctx *fasthttp.RequestCtx) (string, error) {
	webhookUID := adapter.WebhookURLOrUID
	return webhookUID, nil
}
