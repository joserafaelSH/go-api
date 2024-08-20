package env

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GOENV string
	PORT  string
}

var ENV Env

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		slog.Error("Error loading .env file")
		return err
	}

	ENV.PORT = os.Getenv("PORT")
	ENV.GOENV = os.Getenv("GO_ENV")
	return nil
}
