package alert

type Config struct {
	GroupBy []string `mapstructure:"group_by" yaml:"group_by"`
	// Experimental, might change in the future
	NotificationVerboseEnabled bool `mapstructure:"notification_verbose_enabled" yaml:"notification_verbose_enabled" default:"false"`
}
