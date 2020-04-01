package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var once sync.Once
var Data map[string]interface{}

func init() {
	once.Do(func() {
		confPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		confPath = confPath + "/conf"
		if _, err := toml.DecodeFile(confPath+"config.toml", &Data); err != nil {
			log.Fatal(err)
		}
	})
}

func main() {

}
