package confmap

import (
	"fmt"

	"haldi/internal/alias"
)

type ConfigMap struct {
	Aliases []alias.Alias
}

func (cnfmp *ConfigMap) ToSh() string {
	shFile := ""

	for _, cnfg := range cnfmp.Aliases {
		shFile = fmt.Sprintf("%s \n %s", shFile, cnfg.ToSh())
	}

	return shFile
}
