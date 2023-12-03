package utility

import (
	"SiverPineValley/trailer-manager/common"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type nvlTypes interface {
	~string | ~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Nvl[T nvlTypes](v *T, d ...T) (rtn T) {
	if v == nil {
		return
	}
	if len(d) > 0 {
		return d[0]
	}
	return *v
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func NewUuidGenerate() string {
	return uuid.New().String()
}

func GetContextFromEchoContext(c echo.Context) (ctx context.Context) {
	ctx = context.Background()

	if c != nil && c.Request() != nil && c.Request().Header != nil {
		h := c.Request().Header
		for k := range h {
			if k == http.CanonicalHeaderKey(common.HeaderTransactionId) {
				ctx = context.WithValue(ctx, common.HeaderTransactionId, h.Get(common.HeaderTransactionId))
			} else {
				ctx = context.WithValue(ctx, common.ContextKey+k, h.Get(k))
			}
		}
	}

	return
}

func Base64DecodeStruct(body interface{}) {
	value := reflect.ValueOf(body).Elem()
	typ := reflect.TypeOf(body).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("decode")

		if tag == "true" {
			fieldValue := value.Field(i)
			if fieldValue.Kind() == reflect.String && fieldValue.String() != "" {
				decoded, err := base64.StdEncoding.DecodeString(fieldValue.String())
				if err != nil {
					fmt.Printf("Failed to decode field %s: %v\n", field.Name, err)
				} else {
					fieldValue.SetString(string(decoded))
				}
			}
		}
	}
}
