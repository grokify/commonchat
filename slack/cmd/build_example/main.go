package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/grokify/commonchat/glip/classic"
	"github.com/grokify/commonchat/glip/config"
	"github.com/grokify/commonchat/slack"
	"github.com/grokify/mogo/fmt/fmtutil"
)

const (
	filenameSlackOrig      = "example_attachment_orig_slack.txt"
	filenameSlackJSON      = "example_attachment_orig_slack_sp2.json"
	filenameGlipSimpleJSON = "example_attachment_conv_glip_simple.json"
	filenameGlipAttachJSON = "example_attachment_conv_glip_attach.json"
)

func main() {
	qry := slack.ExampleMessageAttachmentURLValues()
	err := os.WriteFile(filenameSlackOrig, []byte(qry.Encode()), 0644)
	if err != nil {
		log.Fatal(err)
	}

	slMsg, err := slack.ParseMessageURLValues(qry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SLMSG")
	fmtutil.PrintJSON(slMsg)
	ccMsg := slack.WebhookInBodySlackToCc(slMsg)
	fmt.Println("CCMSG")
	fmtutil.PrintJSON(ccMsg)

	glCfg := config.DefaultConverterConfig()
	glCfg.UseAttachments = false
	glConv := classic.NewGlipMessageConverter(glCfg)
	glMsg := glConv.ConvertCommonMessage(ccMsg)
	fmt.Println("GLMSG_SIMP")
	fmtutil.PrintJSON(glMsg)
	glJson, err := json.MarshalIndent(glMsg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filenameGlipSimpleJSON, []byte(glJson), 0644)
	if err != nil {
		log.Fatal(err)
	}

	glCfg.UseAttachments = true
	glConv2 := classic.NewGlipMessageConverter(glCfg)
	glMsg2 := glConv2.ConvertCommonMessage(ccMsg)
	fmt.Println("GLMSG_SIMP")
	fmtutil.PrintJSON(glMsg)
	glJson2, err := json.MarshalIndent(glMsg2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filenameGlipAttachJSON, []byte(glJson2), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
