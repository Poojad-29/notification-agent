package config

import "os"

type Config struct {
	SendGridAPIKey string
	TwilioSID      string
	TwilioToken    string
	GRPCToken      string
}

func Load() Config {
	return Config{
		SendGridAPIKey: os.Getenv("SENDGRID_API_KEY"),
		TwilioSID:      os.Getenv("TWILIO_SID"),
		TwilioToken:    os.Getenv("TWILIO_TOKEN"),
		GRPCToken:      os.Getenv("GRPC_TOKEN"),
	}
}
