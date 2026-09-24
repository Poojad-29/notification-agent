package notification

func Route(channel string) string {
	switch channel {
	case "email":
		return "email handler"
	case "sms":
		return "sms handler"
	case "inapp":
		return "inapp handler"
	default:
		return "unknown channel"
	}
}
