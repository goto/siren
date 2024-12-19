package alert

import "time"

type Config struct {
	GroupBy       []string      `mapstructure:"group_by" yaml:"group_by"`
	ValidDuration time.Duration `mapstructure:"valid_duration" yaml:"valid_duration" default:"0s"`
}
