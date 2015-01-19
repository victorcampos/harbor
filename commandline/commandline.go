package commandline

import (
	"errors"
	"fmt"
	"strings"
)

type ConfigVarsMap map[string]string

func (s *ConfigVarsMap) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *ConfigVarsMap) Set(value string) error {
	valueTuple := strings.Split(value, "=")

	if len(valueTuple) == 2 {
		(*s)[valueTuple[0]] = valueTuple[1]
	} else {
		return errors.New("missing '='")
	}

	return nil
}
