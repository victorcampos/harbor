package config

import (
	"strings"
)

func SetEnv(environment string, configString []byte) []byte {
	str := string(configString)
	str = strings.Replace(str, "$ENVIRONMENT", environment, -1)

	return []byte(str)
}
