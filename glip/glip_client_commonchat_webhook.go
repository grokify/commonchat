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

func NewGlipAdapter(webhookURLOrUID string, cfg *config.ConverterConfig) (*GlipAdapter, error) {
	glip, err := glipwebhook.NewGlipWebhookClient(webhookURLOrUID, 1)
	if err != nil {
		return nil, err
	}
	return &GlipAdapter{
		GlipClient:      glip,
		WebhookURLOrUID: glip.WebhookUrl,
		CommonConverter: classic.NewGlipMessageConverter(cfg)}, err
}

func NewGlipAdapterMSI(webhookURLOrUID string, cfg map[string]interface{}) (*GlipAdapter, error) {
	ccfg, err := config.NewConverterConfigMSI(cfg)
	if err != nil {
		return nil, err
	}
	return NewGlipAdapter(webhookURLOrUID, ccfg)
}

func (adapter *GlipAdapter) Clone() (*GlipAdapter, error) {
	return NewGlipAdapter(
		adapter.WebhookURLOrUID,
		adapter.CommonConverter.Config.Clone())
}

func (adapter *GlipAdapter) SendWebhook(urlOrUid string, message commonchat.Message, glipmsg interface{}, cfg map[string]interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
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

	if len(cfg) == 0 {
		return adapter.GlipClient.PostWebhookGUIDFast(urlOrUid, glipMessage)
	}
	newCfg, err := adapter.CommonConverter.Config.UpsertMSI(cfg)
	if err != nil {
		return nil, nil, err
	}
	thisAdapter, err := NewGlipAdapter(adapter.WebhookURLOrUID, newCfg)
	if err != nil {
		return nil, nil, err
	}
	return thisAdapter.GlipClient.PostWebhookGUIDFast(urlOrUid, glipMessage)
}

func (adapter *GlipAdapter) SendMessage(message commonchat.Message, glipmsg interface{}, opts map[string]interface{}) (*fasthttp.Request, *fasthttp.Response, error) {
	return adapter.SendWebhook(adapter.WebhookURLOrUID, message, glipmsg, opts)
}

func (adapter *GlipAdapter) WebhookUID(ctx *fasthttp.RequestCtx) (string, error) {
	webhookUID := adapter.WebhookURLOrUID
	return webhookUID, nil
}
