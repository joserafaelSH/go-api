package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/joserafaelSH/go-api/config/logger"
)

type Env struct {
	GOENV                       string
	PORT                        string
	OTEL_EXPORTER_OTLP_ENDPOINT string
	JAEGER_TRACEPROVIDER        string
}

var ENV Env

func LoadEnv() error {
	fmt.Println("Loading Environment Variables")
	err := godotenv.Load()
	if err != nil {
		logger.Log.Writter.Error().Msg("Error loading .env file, getting environment variables from OS")
	}

	ENV.PORT = os.Getenv("PORT")
	ENV.GOENV = os.Getenv("GO_ENV")
	ENV.OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	ENV.JAEGER_TRACEPROVIDER = os.Getenv("JAEGER_TRACEPROVIDER")

	err = validateEnv()
	if err != nil {
		logger.Log.Writter.Error().Msg(err.Error())
		return err
	}

	return nil
}

func validateEnv() error {
	if ENV.PORT == "" {
		return fmt.Errorf("PORT is not set")
	}
	if ENV.GOENV == "" {
		return fmt.Errorf("GO_ENV is not set")
	}
	if ENV.OTEL_EXPORTER_OTLP_ENDPOINT == "" {
		return fmt.Errorf("OTEL_EXPORTER_OTLP_ENDPOINT is not set")
	}
	if ENV.JAEGER_TRACEPROVIDER == "" {
		return fmt.Errorf("JAEGER_TRACEPROVIDER is not set")
	}
	return nil
}
