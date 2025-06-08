package main

import (
	"fmt"
	"log"

	"github.com/grokify/commonchat/examples"
	"github.com/grokify/commonchat/glip"
	"github.com/grokify/commonchat/glip/config"
	glipwebhook "github.com/grokify/go-glip"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	flags "github.com/jessevdk/go-flags"
	"github.com/valyala/fasthttp"
)

type Options struct {
	URL string `short:"u" long:"url" description:"url for webhook" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	logutil.FatalErr(err)

	cfg := config.DefaultConverterConfig()
	cfg.UseAttachments = true

	adapt := glip.NewGlipAdapter("", cfg)
	fmtutil.MustPrintJSON(adapt.CommonConverter.Config)

	msi := map[string]any{"useAttachments": false}

	cfg2, err := adapt.CommonConverter.Config.UpsertMSI(msi)
	logutil.FatalErr(err)
	fmtutil.MustPrintJSON(cfg2)

	glMsg := &glipwebhook.GlipWebhookMessage{}

	msg := examples.ExampleHookBodyAttachment()

	req, res, err := adapt.SendWebhook(opts.URL, msg, glMsg, map[string]any{})
	logutil.FatalErr(err)

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	req, res, err = adapt.SendWebhook(opts.URL, msg, glMsg,
		map[string]any{"useAttachments": false})
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	fmt.Println("DONE")
}
