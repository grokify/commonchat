package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/commonchat/glip/classic"
	"github.com/grokify/commonchat/glip/config"
	"github.com/grokify/commonchat/slack"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
)

const (
	filenameSlackOrig      = "example_attachment_orig_slack.txt"
	filenameSlackJSON      = "example_attachment_orig_slack_sp2.json"
	filenameGlipSimpleJSON = "example_attachment_conv_glip_simple.json"
	filenameGlipAttachJSON = "example_attachment_conv_glip_attach.json"
)

func main() {
	qry := slack.ExampleMessageAttachmentURLValues()
	err := os.WriteFile(filenameSlackOrig, []byte(qry.Encode()), 0600)
	logutil.FatalErr(err)

	slMsg, err := slack.ParseMessageURLValues(qry)
	logutil.FatalErr(err)

	fmt.Println("SLMSG")
	fmtutil.MustPrintJSON(slMsg)

	ccMsg := slack.WebhookInBodySlackToCc(slMsg)
	fmt.Println("CCMSG")
	fmtutil.MustPrintJSON(ccMsg)

	glCfg := config.DefaultConverterConfig()
	glCfg.UseAttachments = false
	glConv := classic.NewGlipMessageConverter(glCfg)
	glMsg := glConv.ConvertCommonMessage(ccMsg)
	fmt.Println("GLMSG_SIMP")
	fmtutil.MustPrintJSON(glMsg)

	glJson, err := json.MarshalIndent(glMsg, "", "  ")
	logutil.FatalErr(err)

	err = os.WriteFile(filenameGlipSimpleJSON, glJson, 0600)
	logutil.FatalErr(err)

	glCfg.UseAttachments = true
	glConv2 := classic.NewGlipMessageConverter(glCfg)
	glMsg2 := glConv2.ConvertCommonMessage(ccMsg)
	fmt.Println("GLMSG_SIMP")
	fmtutil.MustPrintJSON(glMsg)

	glJson2, err := json.MarshalIndent(glMsg2, "", "  ")
	logutil.FatalErr(err)

	err = os.WriteFile(filenameGlipAttachJSON, glJson2, 0600)
	logutil.FatalErr(err)

	fmt.Println("DONE")
}
