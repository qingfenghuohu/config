package src

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var once sync.Once
var confPath = "./conf/"
var Data map[string]interface{}

func init() {
	once.Do(func() {
		if !Exists(confPath) {
			if Exists("../../conf/") {
				confPath = "../../conf/"
			}
		}
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
