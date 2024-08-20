package receiver

const (
	TypeSlack        string = "slack"
	TypeLark         string = "lark"
	TypeLarkChannel  string = "lark_channel"
	TypeSlackChannel string = "slack_channel"
	TypeHTTP         string = "http"
	TypePagerDuty    string = "pagerduty"
	TypeFile         string = "file"
)

var SupportedTypes = []string{
	TypeSlack,
	TypeLark,
	TypeSlackChannel,
	TypeLarkChannel,
	TypeHTTP,
	TypePagerDuty,
	TypeFile,
}

func IsTypeSupported(receiverType string) bool {
	for _, st := range SupportedTypes {
		if st == receiverType {
			return true
		}
	}
	return false
}
