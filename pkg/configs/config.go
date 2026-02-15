package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env string

const (
	Dev  Env = "development"
	Prod Env = "production"
)

type Config struct {
	Host     string
	TCPPort  string
	GRPCPort string
	MaxConn  int
	Env      Env
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load("../../.env")

	maxConn, err := strconv.Atoi(os.Getenv("MAX_CONN"))
	if err != nil {
		return nil, err
	}
	return &Config{
		Host:     os.Getenv("HOST"),
		TCPPort:  os.Getenv("TCP_PORT"),
		GRPCPort: os.Getenv("GRPC_PORT"),
		MaxConn:  maxConn,
		Env:      Env(os.Getenv("ENV")),
	}, nil
}
