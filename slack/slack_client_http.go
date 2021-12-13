package slack

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/grokify/mogo/net/httputilmore"
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
	HttpClient *http.Client
	FastClient fasthttp.Client
	WebhookUrl string
	UrlPrefix  *regexp.Regexp
}

func NewSlackWebhookClient(urlOrUid string, clientType ClientType) (SlackWebhookClient, error) {
	log.Debug().
		Str("lib", "slack_client.go").
		Str("request_url_client_init", urlOrUid).
		Msg("NewSlackWebhookClient init")

	client := SlackWebhookClient{UrlPrefix: regexp.MustCompile(`^https:`)}
	client.WebhookUrl = client.BuildWebhookURL(urlOrUid)
	if clientType == FastHTTP {
		client.FastClient = fasthttp.Client{}
	} else {
		client.HttpClient = httputilmore.NewHttpClient()
	}
	return client, nil
}

func (client *SlackWebhookClient) BuildWebhookURL(urlOrUid string) string {
	rx := regexp.MustCompile(`^https:`)
	rs := rx.FindString(urlOrUid)
	if len(rs) > 0 {
		log.Debug().
			Str("lib", "slack_client.go").
			Str("request_url_http_match", urlOrUid).
			Msg("BuildWebhookURL")
		return urlOrUid
	}
	return strings.Join([]string{WebhookBaseURL, urlOrUid}, "")
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

	req.Header.Set(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)

	err = client.FastClient.Do(req, resp)
	return req, resp, err
}

func (client *SlackWebhookClient) PostWebhookGUIDFast(urlOrUid string, message Message) (*fasthttp.Request, *fasthttp.Response, error) {
	return client.PostWebhookFast(client.BuildWebhookURL(urlOrUid), message)
}
