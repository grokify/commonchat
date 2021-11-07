package config

import (
	"encoding/json"

	"github.com/grokify/simplego/encoding/jsonutil"
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

func UnmarshalMSI(cfg map[string]interface{}) (*ConverterConfig, error) {
	if len(cfg) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	ccfg := &ConverterConfig{}
	return ccfg, json.Unmarshal(bytes, ccfg)
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
