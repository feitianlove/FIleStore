package store

import (
	"github.com/feitianlove/FIleStore/config"
	"testing"
)

func TestNewStore(t *testing.T) {
	conf := config.Config{MysqlConf: &config.MysqlConf{
		User:     "root",
		Pass:     "fenghui",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "fileserver",
	}}
	_, err := NewStore(&conf)
	if err != nil {
		panic(err)
	}
}
