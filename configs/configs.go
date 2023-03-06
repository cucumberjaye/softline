package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	SigningKey  string
	DataBaseDSN string
)

func InitConfigs() error {
	err := godotenv.Load(".env")
	if err != nil {
		err = nil
	}

	SigningKey = mustEnvStr("SIGNING_KEY")
	DataBaseDSN = mustEnvStr("DATA_BASE_DSN")

	return nil
}

func mustEnvStr(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		log.Error().Err(fmt.Errorf("environment variable %v must be set", key))
		return ""
	}
}
