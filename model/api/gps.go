package model

type GpsLogRequest struct {
	Sid     string   `json:"sid" form:"sid"`
	Lat     string   `json:"lat" form:"lat"`
	Lon     string   `json:"lon" form:"lon"`
	Speed   string   `json:"speed" form:"speed"`
	WifiLoc []string `json:"wifiLoc" form:"wifiLoc"`
	Battery int      `json:"battery" form:"battery"`
}
