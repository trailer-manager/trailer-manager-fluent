package broker

import (
	"database/sql"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/trailer-manager/trailer-manager-common/logger"
	db "github.com/trailer-manager/trailer-manager-fluent/db/rdb"
	"github.com/trailer-manager/trailer-manager-fluent/model/api"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Published message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

var GpsLogHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	gpsLog := model.GpsLogRequest{}
	if err := json.Unmarshal(msg.Payload(), &gpsLog); err != nil {
		logger.Error(err.Error())
	}

	store := db.NewStore(db.RDB)
	_, err := store.CreateGpsLog(nil, db.CreateGpsLogParams{
		Sid:     sql.NullString{String: gpsLog.Sid},
		Lat:     gpsLog.Lat,
		Lon:     gpsLog.Lon,
		Speed:   sql.NullString{String: gpsLog.Speed, Valid: true},
		WifiLoc: gpsLog.WifiLoc,
		Battery: sql.NullInt32{Int32: int32(gpsLog.Battery), Valid: true},
	})
	if err != nil {
		logger.Error(err.Error())
	}

	return
}
