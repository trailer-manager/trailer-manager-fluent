package main

import (
	"SiverPineValley/trailer-manager/broker"
	"SiverPineValley/trailer-manager/config"
	db "SiverPineValley/trailer-manager/db/rdb"
	"SiverPineValley/trailer-manager/logger"
	"SiverPineValley/trailer-manager/utility"
	"flag"
	"log"
)

func main() {
	mode := utility.NvlString(flag.String("mode", "dev", "서버 모드 (개발: dev, 검증: stg, 운영: prd"))

	// 1. Init Config
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	// 2. Init Logger
	if err := logger.InitLogger(mode); err != nil {
		log.Fatal(err)
	}

	// 3. Init Broker
	if err := broker.InitBroker(); err != nil {
		logger.Fatal(err.Error())
	}

	// 4. Init Database
	if err := db.InitRdb(); err != nil {
		logger.Fatal(err.Error())
	}

	// 5. Init Server
	//server, err := api.NewServer(config, store)
	//if err != nil {
	//	log.Fatal("cannot create server: ", err)
	//}

	return
}
