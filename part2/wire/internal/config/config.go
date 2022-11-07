package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"io/ioutil"
	"os"
)

var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, error) {
	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	fmt.Println(cfg)
	return &cfg, nil
}
func main() {
	fp, _ := os.Open("../config/app.json")
	fmt.Println(fp)
	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		println(err)
	}

}
