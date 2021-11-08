package slack

import (
	"encoding/json"
	"net/url"
)

func ExampleMessageAttachmentURLValues() url.Values {
	msg := ExampleMessageAttachment()
	bytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	qry := url.Values{}
	qry.Add(ParamNamePayload, string(bytes))
	return qry
}

func ExampleMessageAttachment() Message {
	return Message{
		IconURL:  "https://i.imgur.com/9yILi61.png",
		Mrkdwn:   true,
		Text:     "Text of the post ♠♥♣♦",
		Username: "Username of the post ♠♥♣♦",
		Attachments: []Attachment{
			{
				AuthorIcon: "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/000080_Navy_Blue_Square.svg/1200px-000080_Navy_Blue_Square.svg.png",
				AuthorLink: "https://example.com/author_link",
				AuthorName: "Attachment.AuthorName ♠♥♣♦",
				Color:      "#00ff2a",
				Fallback:   "Atttachment.Fallback ♠♥♣♦",
				Fields: []Field{
					{
						Title: "Field 1 ♠♥♣♦",
						Value: "A short field ♠♥♣♦",
						Short: true},
					{
						Title: "Field 2",
						Value: "This is [a linked short field](https://example.com)",
						Short: true},
					{
						Title: "Field 3 ♠♥♣♦",
						Value: "A long, full-width field with *formatting* and [a link](https://example.com) \n\n ♠♥♣♦",
						Short: false},
				},
				MarkdownIn:   []string{},
				Pretext:      "Attachment.Pretext ♠♥♣♦",
				Text:         "Attachment.Text ♠♥♣♦",
				ThumbnailURL: "https://raw.githubusercontent.com/grokify/go-glip/master/docs/example_thumbnail-url.png",
				Title:        "Attachment.Title ♠♥♣♦",
			},
		},
	}
}

/*
func ExampleHookBodyAttachmentGlip() glipwebhook.GlipWebhookMessage {
	return glipwebhook.GlipWebhookMessage{
		Icon:     "https://i.imgur.com/9yILi61.png",
		Activity: "Activity of the post ♠♥♣♦",
		Title:    "**Title of the post ♠♥♣♦**",
		Body:     "Body of the post ♠♥♣♦",
		Attachments: []glipwebhook.Attachment{
			{
				Title:        "Attachment Title ♠♥♣♦",
				TitleLink:    "https://example.com/title_link",
				Color:        "#00ff2a",
				AuthorName:   "Author Name ♠♥♣♦",
				AuthorLink:   "https://example.com/author_link",
				AuthorIcon:   "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/000080_Navy_Blue_Square.svg/1200px-000080_Navy_Blue_Square.svg.png",
				Text:         "Attachment text ♠♥♣♦",
				Pretext:      "Attachment pretext appears before the attachment block ♠♥♣♦",
				ImageURL:     "https://media3.giphy.com/media/l4FssTixISsPStXRC/giphy.gif",
				ThumbnailURL: "https://raw.githubusercontent.com/grokify/go-glip/master/docs/example_thumbnail-url.png",
				Fields: []glipwebhook.Field{
					{
						Title: "Field 1 ♠♥♣♦",
						Value: "A short field ♠♥♣♦",
						Short: true},
					{
						Title: "Field 2",
						Value: "This is [a linked short field](https://example.com)",
						Short: true},
					{
						Title: "Field 3 ♠♥♣♦",
						Value: "A long, full-width field with *formatting* and [a link](https://example.com) \n\n ♠♥♣♦",
						Short: false},
				},
				Footer:     "Attachment footer and timestamp ♠♥♣♦",
				FooterIcon: "https://raw.githubusercontent.com/grokify/go-glip/master/docs/example_footer-icon.png",
				TS:         time.Now().Unix(),
			},
		},
	}
}
*/
