package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/caarlos0/env"

	envConf "github.com/qinhan-shu/study/go/config/env"
	jsonConf "github.com/qinhan-shu/study/go/config/json"
	s "github.com/qinhan-shu/study/go/config/struct"
)

func main() {
	// env
	envConf.SetEnv()
	var conf1 s.Config
	if err := env.Parse(&conf1); err != nil {
		log.Fatal("env parse error : ", err)
	}
	conf1.Show()

	// josn
	var conf2 s.Config
	bytes, err := jsonConf.GetJSON()
	if err != nil {
		log.Fatal("get json config error")
	}
	fmt.Printf("\n%s\n", string(bytes))
	fmt.Println()
	if err := json.Unmarshal(bytes, &conf2); err != nil {
		log.Fatal("json parse error : ", err)
	}
	conf2.Show()

	conf2.CheckStruct()
}
