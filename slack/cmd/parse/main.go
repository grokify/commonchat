package main

import (
	"fmt"
	"log"

	"github.com/grokify/commonchat/slack"
	gitlab "github.com/grokify/gowebhooks/gitlab.com"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/text/markdown"
)

func main() {
	data, err := gitlab.ReadExampleFile("event-example_push_slack.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(data)

	slMsg, err := slack.ParseMessageURLEncoded(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n" + slMsg.Text)

	fmt.Println("\n\n>>> MANUAL CONVERSION <<<")
	mkdn := markdown.SkypeToMarkdown(slMsg.Text, true)
	fmt.Println("\n\n[" + mkdn + "]\n")

	ccMsg := slack.WebhookInBodySlackToCc(slMsg)
	fmtutil.PrintJSON(ccMsg)

	fmt.Println("DONE")
}

/*
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/grokify/mogo/text/markdown"
)
*/
/*
func main2() {
	qry, err := url.ParseQuery(ExampleBody)
	if err != nil {
		log.Fatal(err)
	}
	data := qry.Get("payload")
	msg := Message{}
	err = json.Unmarshal([]byte(data), &msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n" + msg.Text)

	mkdn := markdown.SkypeToMarkdown(msg.Text, true)
	fmt.Println("\n\n" + mkdn)

	fmt.Println("DONE")
}

type Message struct {
	Text string `json:"text,omitempty"`
}

const (
	ExampleBody = `payload=%7B%22username%22%3A%22%22%2C%22fallback%22%3A%22John%20Wang%20pushed%20to%20branch%20%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%2Fcommits%2Fmain%7Cmain%5Cu003e%20of%20%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%7CJohn%20Wang%20%2F%20Example%20Project%5Cu003e%20%28%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%2Fcompare%2F90c523f219e3d41b098c74678ecc50ac581428b6...8b2c5d00884b08d12a9d8f97e3ce81de30d78b8b%7CCompare%20changes%5Cu003e%29%22%2C%22text%22%3A%22John%20Wang%20pushed%20to%20branch%20%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%2Fcommits%2Fmain%7Cmain%5Cu003e%20of%20%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%7CJohn%20Wang%20%2F%20Example%20Project%5Cu003e%20%28%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%2Fcompare%2F90c523f219e3d41b098c74678ecc50ac581428b6...8b2c5d00884b08d12a9d8f97e3ce81de30d78b8b%7CCompare%20changes%5Cu003e%29%22%2C%22attachments%22%3A%5B%7B%22text%22%3A%22%5Cu003chttps%3A%2F%2Fgitlab.com%2Fgrokify%2Fexample-project%2F-%2Fcommit%2F8b2c5d00884b08d12a9d8f97e3ce81de30d78b8b%7C8b2c5d00%5Cu003e%3A%20Add%20LICENSE%20-%20John%20Wang%22%2C%22color%22%3A%22%23345%22%7D%5D%7D`
)
*/
