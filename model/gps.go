package model

type GpsLog struct {
	Sid     string   `json:"sid"`
	Lat     string   `json:"lat"`
	Lon     string   `json:"lon"`
	Speed   string   `json:"speed"`
	WifiLoc []string `json:"wifi_loc"`
	Battery int      `json:"battery"`
}
