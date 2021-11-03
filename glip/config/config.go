package config

type ConverterConfig struct {
	EmojiURLFormat                 string
	ActivityIncludeIntegrationName bool
	UseAttachments                 bool // overrides other 'use' options
	UseMarkdownQuote               bool
	UseShortFields                 bool
	UseFieldExtraSpacing           bool
	ConvertTripleBacktick          bool
}

func DefaultConverterConfig() ConverterConfig {
	return ConverterConfig{
		UseAttachments:        true,
		UseFieldExtraSpacing:  true,
		ConvertTripleBacktick: true,
	}
}
