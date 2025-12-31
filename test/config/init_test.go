package config

import (
	"fmt"
	"testing"
	"Zchat/config"
)

func TestInit(t *testing.T) {
	conf := config.GetConfig()
	fmt.Println(conf.MainConfig)
	fmt.Println(conf.MysqlConfig)
	fmt.Println(conf.RedisConfig)
	fmt.Println(conf.AuthCodeConfig)
}