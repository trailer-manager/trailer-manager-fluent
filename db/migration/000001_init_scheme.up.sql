CREATE TABLE "trailer" (
                           "tid" uuid PRIMARY KEY,
                           "tnum" varchar NOT NULL,
                           "created_at" timestamp DEFAULT (now()),
                           "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "sensor" (
                          "sid" varchar PRIMARY KEY,
                          "uid" uuid,
                          "lat" varchar,
                          "lon" varchar,
                          "wifi_loc" varchar[],
                          "battery" int,
                          "created_at" timestamp DEFAULT (now()),
                          "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "trailer_sensor_map" (
                                      "seq" serial PRIMARY KEY,
                                      "tid" uuid NOT NULL,
                                      "sid" varchar NOT NULL,
                                      "created_at" timestamp DEFAULT (now()),
                                      "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "sensor_log" (
                              "seq" serial PRIMARY KEY,
                              "sid" varchar NOT NULL,
                              "uid" uuid NOT NULL,
                              "tid" uuid NOT NULL,
                              "real_creaetd_at" timestamp DEFAULT (now())
);

CREATE TABLE "gps_log" (
                           "sid" varchar PRIMARY KEY,
                           "lat" varchar NOT NULL,
                           "lon" varchar NOT NULL,
                           "speed" numeric,
                           "wifi_loc" varchar[],
                           "battery" int,
                           "real_creaetd_at" timestamp DEFAULT (now())
);

CREATE TABLE "user" (
                        "uid" uuid PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "email" varchar NOT NULL,
                        "mobile" varchar NOT NULL,
                        "birthdate" varchar NOT NULL,
                        "created_at" timestamp DEFAULT (now()),
                        "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_role" (
                             "uid" uuid PRIMARY KEY,
                             "role" varchar NOT NULL,
                             "created_at" timestamp DEFAULT (now()),
                             "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_auth" (
                             "uid" uuid PRIMARY KEY,
                             "salt" varchar NOT NULL,
                             "password" varchar NOT NULL,
                             "created_at" timestamp DEFAULT (now()),
                             "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_cert" (
                             "uid" uuid PRIMARY KEY,
                             "ci" varchar NOT NULL,
                             "di" varchar NOT NULL,
                             "created_at" timestamp DEFAULT (now()),
                             "updated_at" timestamp DEFAULT (now())
);

CREATE INDEX "sensor_log_index" ON "sensor_log" ("sid", "uid", "tid");

CREATE INDEX "gps_log_index" ON "gps_log" ("sid", "real_creaetd_at");

COMMENT ON COLUMN "trailer"."tid" IS '트레일러 ID';

COMMENT ON COLUMN "trailer"."tnum" IS '트레일러 번호';

COMMENT ON COLUMN "sensor"."sid" IS '센서 시리얼 넘버';

COMMENT ON COLUMN "sensor"."uid" IS '사용 중인 사람 유저 ID';

COMMENT ON COLUMN "sensor"."lat" IS '현재 위치 위도';

COMMENT ON COLUMN "sensor"."lon" IS '현재 위치 경도';

COMMENT ON COLUMN "sensor"."wifi_loc" IS 'WIFI 로케이션 (SSID)';

COMMENT ON COLUMN "trailer_sensor_map"."seq" IS '로그 이력 번호';

COMMENT ON COLUMN "trailer_sensor_map"."tid" IS '트레일러 ID';

COMMENT ON COLUMN "trailer_sensor_map"."sid" IS '센서 시리얼 넘버';

COMMENT ON COLUMN "sensor_log"."seq" IS '로그 이력 번호';

COMMENT ON COLUMN "sensor_log"."sid" IS '센서 시리얼 넘버';

COMMENT ON COLUMN "sensor_log"."uid" IS '사용 중인 사람 유저 ID';

COMMENT ON COLUMN "sensor_log"."tid" IS '사용 중인 트레일러 ID';

COMMENT ON COLUMN "gps_log"."sid" IS '센서 시리얼 넘버';

COMMENT ON COLUMN "gps_log"."lat" IS '위도';

COMMENT ON COLUMN "gps_log"."lon" IS '경도';

COMMENT ON COLUMN "gps_log"."speed" IS '속도';

COMMENT ON COLUMN "gps_log"."wifi_loc" IS 'WIFI 로케이션 (SSID)';

COMMENT ON COLUMN "user"."uid" IS '사용자 ID';

COMMENT ON COLUMN "user"."name" IS '사용자 명';

COMMENT ON COLUMN "user"."email" IS '사용자 이메일';

COMMENT ON COLUMN "user"."mobile" IS '사용자 모바일 번호';

COMMENT ON COLUMN "user"."birthdate" IS '사용자 생년월일';

COMMENT ON COLUMN "user_role"."uid" IS '사용자 ID';

COMMENT ON COLUMN "user_role"."role" IS '사용자 권한';

COMMENT ON COLUMN "user_auth"."uid" IS '사용자 ID';

COMMENT ON COLUMN "user_auth"."salt" IS '패스워드 Salt';

COMMENT ON COLUMN "user_auth"."password" IS '패스워드';

COMMENT ON COLUMN "user_cert"."uid" IS '사용자 ID';

COMMENT ON COLUMN "user_cert"."ci" IS '본인인증 CI';

COMMENT ON COLUMN "user_cert"."di" IS '본인인증 DI';

ALTER TABLE "trailer_sensor_map" ADD FOREIGN KEY ("tid") REFERENCES "trailer" ("tid");

ALTER TABLE "trailer_sensor_map" ADD FOREIGN KEY ("sid") REFERENCES "sensor" ("sid");

ALTER TABLE "sensor_log" ADD FOREIGN KEY ("sid") REFERENCES "sensor" ("sid");

ALTER TABLE "sensor_log" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");

ALTER TABLE "gps_log" ADD FOREIGN KEY ("sid") REFERENCES "sensor" ("sid");

ALTER TABLE "user_auth" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");

ALTER TABLE "user_cert" ADD FOREIGN KEY ("uid") REFERENCES "user" ("uid");
