package config

import (
	"os"
)

type Config struct {
	Port        string
	GRPCPort    string
	DynamoTable string
	AWSRegion   string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8281"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
		DynamoTable: getEnv("DYNAMO_TABLE", "WydTable"),
		AWSRegion:   getEnv("AWS_REGION", "us-east-1"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
