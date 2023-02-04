package main

import (
	broker "SiverPineValley/trailer-manager/broker"
	"SiverPineValley/trailer-manager/config"
	"SiverPineValley/trailer-manager/logger"
	"SiverPineValley/trailer-manager/utility"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	mode := utility.NvlString(flag.String("mode", "dev", "서버 모드 (개발: dev, 검증: stg, 운영: prd"))
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := logger.InitLogger(mode); err != nil {
		panic(err)
	}


}
