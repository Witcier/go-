package config

import "time"

type Captcha struct {
	KeyLong   int           `mapstructure:"key-long" json:"key-long" yaml:"key-long"`
	ImgWidth  int           `mapstructure:"img-width" json:"img-width" yaml:"img-width"`
	ImgHeight int           `mapstructure:"img-height" json:"img-height" yaml:"img-height"`
	Expires   time.Duration `mapstructure:"expires" json:"expires" yaml:"expires"`
	PrefixKey string        `mapstructure:"prefix-key" json:"prefix-key" yaml:"prefix-key"`
	MaxSkew   float64       `mapstructure:"max-skew" json:"max-skew" yaml:"max-skew"`
	DotCount  int           `mapstructure:"dot-count" json:"dot-count" yaml:"dot-count"`
}
