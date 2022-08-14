package conf

import (
	"github.com/derekAHua/goLib/base"
)

type (
	ApiConf struct {
		Demo base.ApiClient
	}

	AppConf struct {
		SingKey string `yaml:"singKey"`
		Limit   map[string]struct {
			Switch int `yaml:"switch"`
			MaxNum int `yaml:"maxNum"`
		}
	}
)
