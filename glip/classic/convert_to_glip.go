package classic

import (
	"fmt"
	"regexp"
	"strings"

	glipwebhook "github.com/grokify/go-glip"
	"github.com/grokify/mogo/text/emoji"
	"github.com/grokify/mogo/text/markdown"

	cc "github.com/grokify/commonchat"
	"github.com/grokify/commonchat/glip/config"
)

var rxTripleBackTick *regexp.Regexp = regexp.MustCompile(`(^|\n)` + "```([^`]*?)```" + `(\n|$)`)

type GlipMessageConverter struct {
	Config         *config.ConverterConfig
	EmojiConverter emoji.Converter
}

func NewGlipMessageConverter(cfg *config.ConverterConfig) GlipMessageConverter {
	if cfg == nil {
		cfg = config.DefaultConverterConfig()
	}
	cfg.ConvertTripleBacktick = true
	return GlipMessageConverter{
		Config:         cfg,
		EmojiConverter: emoji.NewConverter()}
}

func (cv *GlipMessageConverter) ConvertCommonMessage(commonMessage cc.Message) glipwebhook.GlipWebhookMessage {
	glip := glipwebhook.GlipWebhookMessage{
		Activity: cv.EmojiConverter.ConvertShortcodesString(commonMessage.Activity, emoji.Unicode),
		Title:    cv.EmojiConverter.ConvertShortcodesString(commonMessage.Title, emoji.Unicode),
		Icon:     commonMessage.IconURL}
	if cv.Config.ActivityIncludeIntegrationName {
		glip.Activity += cv.integrationActivitySuffix(cv.Config.IntegrationName)
	}

	if len(commonMessage.IconURL) > 0 {
		glip.Icon = commonMessage.IconURL
	} else if len(commonMessage.IconEmoji) > 0 {
		iconURL, err := cc.EmojiToURL(cv.Config.EmojiURLFormat, commonMessage.IconEmoji)
		if err == nil {
			glip.Icon = iconURL
		}
	}
	bodyLines := []string{}
	if len(commonMessage.Text) > 0 {
		bodyLines = append(bodyLines, commonMessage.Text)
	}

	if len(commonMessage.Attachments) > 0 {
		if cv.Config.UseAttachments {
			glip.Attachments = convertAttachments(&cv.EmojiConverter, cv.Config.ConvertTripleBacktick, commonMessage.Attachments)
		} else {
			attachmentText := cv.renderAttachmentsAsMarkdown(commonMessage.Attachments)
			if len(attachmentText) > 0 {
				bodyLines = append(bodyLines, attachmentText)
			}
		}
	}

	if len(bodyLines) > 0 {
		glip.Body = strings.Join(bodyLines, "\n")
	}
	return glip
}

func (cv *GlipMessageConverter) getMarkdownBodyPrefix() string {
	if cv.Config.UseMarkdownQuote {
		return "> "
	}
	return ""
}

func convertAttachments(emoconv *emoji.Converter, convertBacktick3ToCode bool, commonAttachments []cc.Attachment) []glipwebhook.Attachment {
	glipAttachments := []glipwebhook.Attachment{}
	for _, commonAttachment := range commonAttachments {
		glipAttachments = append(glipAttachments, convertAttachment(emoconv, convertBacktick3ToCode, commonAttachment))
	}
	return glipAttachments
}

func convertAttachment(emoconv *emoji.Converter, convertBacktick3ToCode bool, commonAttachment cc.Attachment) glipwebhook.Attachment {
	glipAttachment := glipwebhook.Attachment{
		AuthorIcon: commonAttachment.AuthorIcon,
		AuthorLink: commonAttachment.AuthorLink,
		AuthorName: commonAttachment.AuthorName,
		Color:      commonAttachment.Color,
		Fields:     convertFields(emoconv, convertBacktick3ToCode, commonAttachment.Fields),
		Pretext:    emoconv.ConvertShortcodesString(commonAttachment.Pretext, emoji.Unicode),
		Text:       emoconv.ConvertShortcodesString(commonAttachment.Text, emoji.Unicode),
		Title:      emoconv.ConvertShortcodesString(commonAttachment.Title, emoji.Unicode),
		Type:       "Card"}
	if convertBacktick3ToCode {
		glipAttachment.Pretext = TripleBacktickToCode(glipAttachment.Pretext)
		glipAttachment.Text = TripleBacktickToCode(glipAttachment.Text)
	}
	return glipAttachment
}

func convertFields(emoconv *emoji.Converter, convertBacktick3ToCode bool, commonFields []cc.Field) []glipwebhook.Field {
	glipFields := []glipwebhook.Field{}
	for _, commonField := range commonFields {
		glipFields = append(glipFields, convertField(emoconv, convertBacktick3ToCode, commonField))
	}
	return glipFields
}

func convertField(emoconv *emoji.Converter, convertBacktick3ToCode bool, commonField cc.Field) glipwebhook.Field {
	glipField := glipwebhook.Field{
		Title: commonField.Title,
		Value: emoconv.ConvertShortcodesString(commonField.Value, emoji.Unicode),
		Short: commonField.Short}
	if convertBacktick3ToCode {
		glipField.Value = TripleBacktickToCode(glipField.Value)
	}
	return glipField
}

func (cv *GlipMessageConverter) renderAttachmentsAsMarkdown(attachments []cc.Attachment) string {
	lines := []string{}
	prefix := cv.getMarkdownBodyPrefix()
	shortFields := []cc.Field{}
	for _, att := range attachments {
		lines = append(lines, "")
		if len(att.AuthorName) > 0 || len(strings.TrimSpace(att.AuthorLink)) > 0 {
			lines = append(lines, fmt.Sprintf("%s%s", prefix,
				markdown.Linkify(att.AuthorLink, att.AuthorName)))
		}
		if len(att.Title) > 0 {
			lines = append(lines, fmt.Sprintf("%s**%s**", prefix, att.Title))
		}
		if len(att.Text) > 0 {
			lines = append(lines, fmt.Sprintf("%s%s", prefix, att.Text))
		}
		for _, field := range att.Fields {
			if !cv.Config.UseShortFields {
				field.Short = false
			}
			if field.Short {
				shortFields = append(shortFields, field)
				if len(shortFields) == 2 {
					fieldLines := cv.buildMarkdownShortFieldLines(shortFields)
					if len(fieldLines) > 0 {
						lines = cv.appendEmptyLine(lines)
						lines = append(lines, fieldLines...)
					}
					shortFields = []cc.Field{}
				}
				continue
			} else {
				if len(shortFields) > 0 {
					fieldLines := cv.buildMarkdownShortFieldLines(shortFields)
					if len(fieldLines) > 0 {
						lines = cv.appendEmptyLine(lines)
						lines = append(lines, fieldLines...)
					}
				}
				shortFields = []cc.Field{}
			}
			if len(field.Title) > 0 || len(field.Value) > 0 {
				lines = cv.appendEmptyLine(lines)
				if len(field.Title) > 0 {
					lines = append(lines, fmt.Sprintf("%s**%s**", prefix, field.Title))
				}
				if len(field.Value) > 0 {
					lines = append(lines, fmt.Sprintf("%s%s", prefix, field.Value))
				}
			}
		}
	}
	return strings.Join(lines, "\n")
}

func (cv *GlipMessageConverter) buildMarkdownShortFieldLines(shortFields []cc.Field) []string {
	lines := []string{}
	prefix := cv.getMarkdownBodyPrefix()
	for len(shortFields) > 0 {
		if len(shortFields) >= 2 {
			lines = cv.appendEmptyLine(lines)
			field1 := shortFields[0]
			field2 := shortFields[1]
			if len(field1.Title) > 0 || len(field2.Title) > 0 {
				lines = append(lines, fmt.Sprintf("%s| **%v** | **%v** |", prefix, field1.Title, field2.Title))
			}
			if len(field1.Value) > 0 || len(field2.Value) > 0 {
				lines = append(lines, fmt.Sprintf("%s| %v | %v |", prefix, field1.Value, field2.Value))
			}
			shortFields = shortFields[2:]
		} else {
			lines = cv.appendEmptyLine(lines)
			field1 := shortFields[0]
			if len(field1.Title) > 0 {
				lines = append(lines, fmt.Sprintf("%s**%s**", prefix, field1.Title))
			}
			if len(field1.Value) > 0 {
				lines = append(lines, fmt.Sprintf("%s%s", prefix, field1.Value))
			}
			shortFields = shortFields[1:]
		}
	}
	return lines
}

func (cv *GlipMessageConverter) appendEmptyLine(lines []string) []string {
	if cv.Config.UseFieldExtraSpacing {
		if len(lines) > 0 {
			if len(lines[len(lines)-1]) > 0 {
				lines = append(lines, "")
			}
		}
	}
	return lines
}

/*
func (cv *GlipMessageConverter) RenderMessage(message cc.Message) string {
	lines := []string{}
	attachments := cv.RenderAttachments(message.Attachments)
	if len(attachments) > 0 {
		lines = append(lines, attachments)
	}
	return strings.Join(lines, "\n")
}
*/

func (cv *GlipMessageConverter) integrationActivitySuffix(displayName string) string {
	displayName = strings.TrimSpace(displayName)
	if len(displayName) > 0 {
		return fmt.Sprintf(" (%v)", displayName)
	}
	return ""
}

// TripleBacktickToCode converts markdown triple backticks to Glip code blocks.
func TripleBacktickToCode(input string) string {
	return rxTripleBackTick.ReplaceAllString(input, "$1\n[code]\n$2\n[/code]\n$3")
}
