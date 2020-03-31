package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var Data map[string]interface{}

func init() {
	once.Do(func() {
		confPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		confPath = confPath + "/conf"
		if _, err := toml.DecodeFile(confPath+"config.toml", &Data); err != nil {
			log.Fatal(err)
		}
	})
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
