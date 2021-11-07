package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/simplego/encoding/jsonutil"
)

const (
	PropNameEmojiURLFormat                 = "emojiURLFormat"
	PropNameActivityIncludeIntegrationName = "activityIncludeIntegrationName"
	PropNameUseAttachments                 = "useAttachments"
	PropNameUseMarkdownQuote               = "useMarkdownQuote"
	PropNameUseShortFields                 = "useShortFields"
	PropNameUseFieldExtraSpacing           = "useFieldExtraSpacing"
	PropNameConvertTripleBacktick          = "convertTripleBacktick"
)

type ConverterConfig struct {
	EmojiURLFormat                 string `json:"emojiURLFormat,omitempty"`
	ActivityIncludeIntegrationName bool   `json:"activityIncludeIntegrationName,omitempty"`
	UseAttachments                 bool   `json:"useAttachments,omitempty"`
	UseMarkdownQuote               bool   `json:"useMarkdownQuote,omitempty"`
	UseShortFields                 bool   `json:"useShortFields,omitempty"`
	UseFieldExtraSpacing           bool   `json:"useFieldExtraSpacing,omitempty"`
	ConvertTripleBacktick          bool   `json:"convertTripleBacktick,omitempty"`
}

func NewConverterConfigMSI(cfg map[string]interface{}) (*ConverterConfig, error) {
	ccfg := &ConverterConfig{}
	return ccfg, jsonutil.UnmarshalMSI(cfg, ccfg)
}

func DefaultConverterConfig() *ConverterConfig {
	return &ConverterConfig{
		UseAttachments:        true,
		UseFieldExtraSpacing:  true,
		ConvertTripleBacktick: true,
	}
}

func (cfg *ConverterConfig) Clone() *ConverterConfig {
	return &ConverterConfig{
		EmojiURLFormat:                 cfg.EmojiURLFormat,
		ActivityIncludeIntegrationName: cfg.ActivityIncludeIntegrationName,
		UseAttachments:                 cfg.UseAttachments,
		UseMarkdownQuote:               cfg.UseMarkdownQuote,
		UseShortFields:                 cfg.UseShortFields,
		UseFieldExtraSpacing:           cfg.UseFieldExtraSpacing,
		ConvertTripleBacktick:          cfg.ConvertTripleBacktick}
}

func (cfg *ConverterConfig) UpsertMSI(data map[string]interface{}) (*ConverterConfig, error) {
	newCfg := cfg.Clone()
	errMsgs := []string{}
	for k, v := range data {
		switch k {
		case PropNameEmojiURLFormat:
			switch t := v.(type) {
			case string:
				newCfg.EmojiURLFormat = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameActivityIncludeIntegrationName:
			switch t := v.(type) {
			case bool:
				newCfg.ActivityIncludeIntegrationName = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameUseAttachments:
			switch t := v.(type) {
			case bool:
				newCfg.UseAttachments = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameUseMarkdownQuote:
			switch t := v.(type) {
			case bool:
				newCfg.UseMarkdownQuote = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameUseShortFields:
			switch t := v.(type) {
			case bool:
				newCfg.UseShortFields = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameUseFieldExtraSpacing:
			switch t := v.(type) {
			case bool:
				newCfg.UseFieldExtraSpacing = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		case PropNameConvertTripleBacktick:
			switch t := v.(type) {
			case bool:
				newCfg.ConvertTripleBacktick = t
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("`%s` is not `bool`", k))
			}
		}
	}
	if len(errMsgs) > 0 {
		return newCfg, errors.New(strings.Join(errMsgs, ";"))
	}
	return newCfg, nil
}
