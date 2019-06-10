package slack

import (
	cc "github.com/grokify/commonchat"
)

func WebhookInBodySlackToCc(slMsg Message) cc.Message {
	ccMsg := cc.Message{
		Activity:    slMsg.Username,
		Attachments: []cc.Attachment{},
		Text:        slMsg.Text,
		IconEmoji:   slMsg.IconEmoji,
		IconURL:     slMsg.IconURL}
	for _, slAtt := range slMsg.Attachments {
		ccMsg.Attachments = append(ccMsg.Attachments, attachmentSlackToCc(slAtt))
	}
	return ccMsg
}

func attachmentSlackToCc(slAtt Attachment) cc.Attachment {
	ccAtt := cc.Attachment{
		AuthorName:   slAtt.AuthorName,
		Color:        slAtt.Color,
		Fallback:     slAtt.Fallback,
		Fields:       []cc.Field{},
		MarkdownIn:   slAtt.MarkdownIn,
		Pretext:      slAtt.Pretext,
		Text:         slAtt.Text,
		ThumbnailURL: slAtt.ThumbnailURL,
		Title:        slAtt.Title,
	}
	for _, slField := range slAtt.Fields {
		ccAtt.Fields = append(ccAtt.Fields, fieldSlackToCc(slField))
	}
	return ccAtt
}

func fieldSlackToCc(slField Field) cc.Field {
	ccField := cc.Field{
		Short: slField.Short,
		Title: slField.Title,
		Value: slField.Value}
	return ccField
}
