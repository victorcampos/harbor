package config

import (
	"fmt"
	"github.com/victorcampos/harbor/commandline"
	"strings"
)

func SetEnv(cliConfigVars commandline.ConfigVarsMap, configString []byte) []byte {
	str := string(configString)

	// FIXME: Parallelize replace
	for key, value := range cliConfigVars {
		str = strings.Replace(str, fmt.Sprintf("${%s}", key), value, -1)
	}

	return []byte(str)
}
