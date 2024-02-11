package main

import (
	"flag"
	"fmt"
	"github.com/trailer-manager/trailer-manager-common/config"
	"github.com/trailer-manager/trailer-manager-common/logger"
	"github.com/trailer-manager/trailer-manager-common/server"
	"github.com/trailer-manager/trailer-manager-fluent/router"

	//"github.com/trailer-manager/trailer-manager-fluent/broker"
	"github.com/trailer-manager/trailer-manager-common/utility"
	db "github.com/trailer-manager/trailer-manager-fluent/db/rdb"
	"log"
)

func main() {
	mode := utility.Nvl(flag.String("mode", "local", "서버 모드 (로컬: local, 개발: dev, 검증: stg, 운영: prd"))

	// 1. Init Config
	if err := config.InitConfig(mode); err != nil {
		log.Fatal(err)
	}

	// 2. Init Logger
	if err := logger.InitLogger(mode); err != nil {
		log.Fatal(err)
	}

	// 3. Init Broker
	// if err := broker.InitBroker(); err != nil {
	// 	logger.Fatal(err.Error())
	// }

	// 4. Init Database
	if err := db.InitRdb(); err != nil {
		logger.Fatal(err.Error())
	}

	// 5. Init Server
	e := server.NewServer()
	router.InitRouter(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.GetConfig().Port)))
	return
}
