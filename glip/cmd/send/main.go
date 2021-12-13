package main

import (
	"fmt"
	"log"

	"github.com/grokify/commonchat/examples"
	"github.com/grokify/commonchat/glip"
	"github.com/grokify/commonchat/glip/config"
	glipwebhook "github.com/grokify/go-glip"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
	"github.com/valyala/fasthttp"
)

type Options struct {
	URL string `short:"u" long:"url" description:"url for webhook" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.DefaultConverterConfig()
	cfg.UseAttachments = true

	ad, err := glip.NewGlipAdapter("", cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(ad.CommonConverter.Config)

	msi := map[string]interface{}{"useAttachments": false}

	cfg2, err := ad.CommonConverter.Config.UpsertMSI(msi)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(cfg2)

	glMsg := &glipwebhook.GlipWebhookMessage{}

	msg := examples.ExampleHookBodyAttachment()

	req, res, err := ad.SendWebhook(opts.URL, msg, glMsg, map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	if 1 == 1 {
		req, res, err := ad.SendWebhook(opts.URL, msg, glMsg,
			map[string]interface{}{"useAttachments": false})
		if err != nil {
			log.Fatal(err)
		}
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}

	fmt.Println("DONE")
}
