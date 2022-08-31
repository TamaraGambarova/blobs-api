package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type WSConfig struct {
	Addr string `fig:"addr"`
}

func (c *config) WSConf() *WSConfig {
	return c.ws.Do(func() interface{} {
		var cfg WSConfig

		err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "ws")).Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*WSConfig)
}
