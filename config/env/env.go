package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// 	slog.Error("Error loading .env file")
	// }

	ENV.PORT = os.Getenv("PORT")
	ENV.GOENV = os.Getenv("GO_ENV")
	ENV.OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	ENV.JAEGER_TRACEPROVIDER = os.Getenv("JAEGER_TRACEPROVIDER")

	fmt.Println("Environment: " + ENV.GOENV)
	fmt.Println("Starting WebServer on PORT: " + ENV.PORT)

	return nil
}
