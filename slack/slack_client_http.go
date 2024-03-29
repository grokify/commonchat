package slack

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

var (
	WebhookBaseURL = "https://hooks.slack.com/services/"
)

type ClientType int

const (
	HTTP ClientType = iota
	FastHTTP
)

type SlackWebhookClient struct {
	HTTPClient *http.Client
	FastClient *fasthttp.Client
	WebhookURL string
	URLPrefix  *regexp.Regexp
}

func NewSlackWebhookClient(urlOrUID string, clientType ClientType) (SlackWebhookClient, error) {
	log.Debug().
		Str("lib", "slack_client.go").
		Str("request_url_client_init", urlOrUID).
		Msg("NewSlackWebhookClient init")

	client := SlackWebhookClient{URLPrefix: regexp.MustCompile(`^https:`)}
	client.WebhookURL = client.BuildWebhookURL(urlOrUID)
	if clientType == FastHTTP {
		client.FastClient = &fasthttp.Client{}
	} else {
		client.HTTPClient = httputilmore.NewHTTPClient()
	}
	return client, nil
}

func (client *SlackWebhookClient) BuildWebhookURL(urlOrUID string) string {
	rx := regexp.MustCompile(`^https:`)
	rs := rx.FindString(urlOrUID)
	if len(rs) > 0 {
		log.Debug().
			Str("lib", "slack_client.go").
			Str("request_url_http_match", urlOrUID).
			Msg("BuildWebhookURL")
		return urlOrUID
	}
	return strings.Join([]string{WebhookBaseURL, urlOrUID}, "")
}

func (client *SlackWebhookClient) PostWebhookFast(url string, message Message) (*fasthttp.Request, *fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	bytes, err := json.Marshal(message)
	if err != nil {
		return req, resp, err
	}
	req.SetBody(bytes)

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetRequestURI(url)

	req.Header.Set(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)

	err = client.FastClient.Do(req, resp)
	return req, resp, err
}

func (client *SlackWebhookClient) PostWebhookGUIDFast(urlOrUID string, message Message) (*fasthttp.Request, *fasthttp.Response, error) {
	return client.PostWebhookFast(client.BuildWebhookURL(urlOrUID), message)
}
