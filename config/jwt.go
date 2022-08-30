package config

import "time"

type JWT struct {
	SigningKey  string        `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	ExpiresHour time.Duration `mapstructure:"expires-hour" json:"expires-hour" yaml:"expires-hour"`
	BufferTime  int64         `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	Issuer      string        `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	TokenType   string        `mapstructure:"token-type" json:"token-type" yaml:"token-type"`
}
