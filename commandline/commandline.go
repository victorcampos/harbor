package commandline

import (
	"errors"
	"strings"
)

type ConfigVarsMap map[string]string

func NewConfigVarsMap(keyValues []string) (ConfigVarsMap, error) {
	configVarsMap := make(ConfigVarsMap)

	for _, value := range keyValues {
		valueTuple := strings.Split(value, "=")

		if len(valueTuple) == 2 {
			configVarsMap[valueTuple[0]] = valueTuple[1]
		} else {
			return nil, errors.New("missing '='")
		}
	}

	return configVarsMap, nil
}
