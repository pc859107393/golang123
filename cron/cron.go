package cron

import (
	"github.com/robfig/cron"
	"golang123/config"
	"golang123/model"
)

var cronMap = map[string]func(){}

func init() {
	if config.ServerConfig.Env != model.DevelopmentMode {
		cronMap["0 0 3 * * *"] = yesterdayCron
	} else {
		// go func() {
		// 	time.Sleep(1 * time.Second)
		// 	yesterdayCron()
		// }()
	}
}

// New 构造cron
func New() *cron.Cron {
	c := cron.New()
	for spec, cmd := range cronMap {
		c.AddFunc(spec, cmd)
	}
	return c
}
