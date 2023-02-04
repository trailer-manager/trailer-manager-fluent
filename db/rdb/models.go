// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type GpsLog struct {
	// 센서 시리얼 넘버
	Sid string `json:"sid"`
	// 위도
	Lat string `json:"lat"`
	// 경도
	Lon string `json:"lon"`
	// 속도
	Speed sql.NullString `json:"speed"`
	// WIFI 로케이션 (SSID)
	WifiLoc       []string      `json:"wifi_loc"`
	Battery       sql.NullInt32 `json:"battery"`
	RealCreaetdAt sql.NullTime  `json:"real_creaetd_at"`
}

type Sensor struct {
	// 센서 시리얼 넘버
	Sid string `json:"sid"`
	// 사용 중인 사람 유저 ID
	Uid uuid.UUID `json:"uid"`
	// 현재 위치 위도
	Lat sql.NullString `json:"lat"`
	// 현재 위치 경도
	Lon sql.NullString `json:"lon"`
	// WIFI 로케이션 (SSID)
	WifiLoc   []string      `json:"wifi_loc"`
	Battery   sql.NullInt32 `json:"battery"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type SensorLog struct {
	// 로그 이력 번호
	Seq int32 `json:"seq"`
	// 센서 시리얼 넘버
	Sid string `json:"sid"`
	// 사용 중인 사람 유저 ID
	Uid uuid.UUID `json:"uid"`
	// 사용 중인 트레일러 ID
	Tid           uuid.UUID    `json:"tid"`
	RealCreaetdAt sql.NullTime `json:"real_creaetd_at"`
}

type Trailer struct {
	// 트레일러 ID
	Tid uuid.UUID `json:"tid"`
	// 트레일러 번호
	Tnum      string       `json:"tnum"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type TrailerSensorMap struct {
	// 로그 이력 번호
	Seq int32 `json:"seq"`
	// 트레일러 ID
	Tid uuid.UUID `json:"tid"`
	// 센서 시리얼 넘버
	Sid       string       `json:"sid"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type User struct {
	// 사용자 ID
	Uid uuid.UUID `json:"uid"`
	// 사용자 명
	Name string `json:"name"`
	// 사용자 이메일
	Email string `json:"email"`
	// 사용자 모바일 번호
	Mobile string `json:"mobile"`
	// 사용자 생년월일
	Birthdate string       `json:"birthdate"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type UserAuth struct {
	// 사용자 ID
	Uid uuid.UUID `json:"uid"`
	// 패스워드 Salt
	Salt string `json:"salt"`
	// 패스워드
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type UserCert struct {
	// 사용자 ID
	Uid uuid.UUID `json:"uid"`
	// 본인인증 CI
	Ci string `json:"ci"`
	// 본인인증 DI
	Di        string       `json:"di"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type UserRole struct {
	// 사용자 ID
	Uid uuid.UUID `json:"uid"`
	// 사용자 권한
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
