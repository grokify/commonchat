package commonchat

import (
	"fmt"
	"strings"

	cc "github.com/commonchat/commonchat-go"
	"github.com/grokify/glip-go-webhook"
)

type GlipMessageConverter struct {
	EmojiURLFormat                 string
	UseMarkdownQuote               bool
	UseShortFields                 bool
	UseFieldExtraSpacing           bool
	ActivityIncludeIntegrationName bool
}

func NewGlipMessageConverter() GlipMessageConverter {
	cv := GlipMessageConverter{}
	return cv
}

func (cv *GlipMessageConverter) ConvertCommonMessage(commonMessage cc.Message) glipwebhook.GlipWebhookMessage {
	glip := glipwebhook.GlipWebhookMessage{
		Activity: commonMessage.Activity,
		Title:    commonMessage.Title,
		Icon:     commonMessage.IconURL}
	if len(commonMessage.IconURL) > 0 {
		glip.Icon = commonMessage.IconURL
	} else if len(commonMessage.IconEmoji) > 0 {
		iconURL, err := cc.EmojiToURL(cv.EmojiURLFormat, commonMessage.IconEmoji)
		if err == nil {
			glip.Icon = iconURL
		}
	}
	bodyLines := []string{}
	if len(commonMessage.Text) > 0 {
		bodyLines = append(bodyLines, commonMessage.Text)
	}
	if len(commonMessage.Attachments) > 0 {
		attachmentText := cv.RenderAttachments(commonMessage.Attachments)
		if len(attachmentText) > 0 {
			bodyLines = append(bodyLines, attachmentText)
		}
	}
	if len(bodyLines) > 0 {
		glip.Body = strings.Join(bodyLines, "\n")
	}
	return glip
}

func (cv *GlipMessageConverter) GetGlipMarkdownBodyPrefix() string {
	if cv.UseMarkdownQuote {
		return "> "
	}
	return ""
}

func (cv *GlipMessageConverter) RenderAttachments(attachments []cc.Attachment) string {
	lines := []string{}
	prefix := cv.GetGlipMarkdownBodyPrefix()
	shortFields := []cc.Field{}
	for _, att := range attachments {
		if len(att.Title) > 0 {
			lines = append(lines, fmt.Sprintf("%s**%s**", prefix, att.Title))
		}
		if len(att.Text) > 0 {
			lines = append(lines, fmt.Sprintf("%s%s", prefix, att.Text))
		}
		for _, field := range att.Fields {
			if !cv.UseShortFields {
				field.Short = false
			}
			if field.Short {
				shortFields = append(shortFields, field)
				if len(shortFields) == 2 {
					fieldLines := cv.BuildShortFieldLines(shortFields)
					if len(fieldLines) > 0 {
						lines = cv.AppendEmptyLine(lines)
						lines = append(lines, fieldLines...)
					}
					shortFields = []cc.Field{}
				}
				continue
			} else {
				if len(shortFields) > 0 {
					fieldLines := cv.BuildShortFieldLines(shortFields)
					if len(fieldLines) > 0 {
						lines = cv.AppendEmptyLine(lines)
						lines = append(lines, fieldLines...)
					}
				}
				shortFields = []cc.Field{}
			}
			if len(field.Title) > 0 || len(field.Value) > 0 {
				lines = cv.AppendEmptyLine(lines)
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

func (cv *GlipMessageConverter) BuildShortFieldLines(shortFields []cc.Field) []string {
	lines := []string{}
	prefix := cv.GetGlipMarkdownBodyPrefix()
	for len(shortFields) > 0 {
		if len(shortFields) >= 2 {
			lines = cv.AppendEmptyLine(lines)
			field1 := shortFields[0]
			field2 := shortFields[1]
			if len(field2.Title) > 0 || len(field2.Title) > 0 {
				lines = append(lines, fmt.Sprintf("%s| **%v** | **%v** |", prefix, field1.Title, field2.Title))
			}
			if len(field2.Value) > 0 || len(field2.Value) > 0 {
				lines = append(lines, fmt.Sprintf("%s| %v | %v |", prefix, field1.Value, field2.Value))
			}
			shortFields = shortFields[2:]
		} else {
			lines = cv.AppendEmptyLine(lines)
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

func (cv *GlipMessageConverter) AppendEmptyLine(lines []string) []string {
	if cv.UseFieldExtraSpacing {
		if len(lines) > 0 {
			if len(lines[len(lines)-1]) > 0 {
				lines = append(lines, "")
			}
		}
	}
	return lines
}

func (cv *GlipMessageConverter) RenderMessage(message cc.Message) string {
	lines := []string{}
	attachments := cv.RenderAttachments(message.Attachments)
	if len(attachments) > 0 {
		lines = append(lines, attachments)
	}
	return strings.Join(lines, "\n")
}

func (cv *GlipMessageConverter) IntegrationActivitySuffix(displayName string) string {
	if cv.ActivityIncludeIntegrationName {
		return fmt.Sprintf(" (%v)", displayName)
	}
	return ""
}
