package api

import (
	"SiverPineValley/trailer-manager/common"
	"SiverPineValley/trailer-manager/logger"
	model "SiverPineValley/trailer-manager/model/api"
	"SiverPineValley/trailer-manager/utility"
	"bytes"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func transactionIdHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		transactionId := req.Header.Get(common.HeaderTransactionId)
		if transactionId == "" {
			tid := utility.NewUuidGenerate()
			req.Header.Set(common.HeaderTransactionId, tid)
		}
		_ = next(c)
		return nil
	}
}

func errHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()

	if e, ok := err.(*echo.HTTPError); ok {
		code = e.Code
		if m, ok := e.Message.(string); ok {
			msg = m
		}
	}

	var errCode string = strconv.Itoa(code * 10)
	if e, ok := err.(*common.Error); ok {
		errCode = e.Code
	}

	ctx := utility.GetContextFromEchoContext(c)
	var tid string
	tid, _ = ctx.Value(common.HeaderTransactionId).(string)
	c.JSON(code, model.Response{
		Result: model.Result{
			StatusCode:    code,
			ErrorCode:     errCode,
			TransactionId: tid,
			Message:       msg,
		},
	})
}

func notFoundHandler(c echo.Context) error {
	c.Error(echo.ErrNotFound)
	return echo.ErrNotFound
}

func methodNotAllowdHandler(c echo.Context) error {
	c.Error(echo.ErrMethodNotAllowed)
	return echo.ErrMethodNotAllowed
}

func httpLogHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		httpBefore(c)

		_ = next(c)

		latency := time.Since(start)

		httpAfter(c, latency)

		return nil
	}
}

func httpBefore(c echo.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.ErrorContext(utility.GetContextFromEchoContext(c), string(debug.Stack()))
		}
	}()

	req := c.Request()
	realIp := req.RemoteAddr
	if req.Header.Get("X-Real-IP") != "" {
		realIp = req.Header.Get("X-Real-IP")
	}

	ctx := utility.GetContextFromEchoContext(c)
	ctx = context.WithValue(ctx, common.ContextLogType, common.ContextLogTypeStart)
	logger.InfoContext(ctx, fmt.Sprintf("[METHOD=%s] [URL=%s] [HOST=%s] [REMOTE_ADDR=%s]", strings.ToUpper(req.Method), req.URL.String(), req.Host, realIp))
	ctx = context.WithValue(ctx, common.ContextLogType, common.ContextLogTypeNormal)
	headers := make([]string, 0)
	for key, val := range req.Header {
		headers = append(headers, fmt.Sprintf("%s=%v", key, val))
	}
	if len(headers) > 0 {
		logger.InfoContext(ctx, "[HEADER] ", strings.Join(headers, ", "))
	}

	contentType := req.Header.Get("Content-Type")
	contentLength := req.Header.Get("Content-Length")
	if contentLength != "" && strings.Contains(contentType, echo.MIMEApplicationJSON) {
		reqBody := make([]byte, 0)
		if req.Body != nil {
			reqBody, _ = io.ReadAll(req.Body)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		logger.InfoContext(ctx, "[Body] ", string(reqBody))
	}
}

func httpAfter(c echo.Context, latency time.Duration) {
	defer func() {
		if r := recover(); r != nil {
			logger.ErrorContext(utility.GetContextFromEchoContext(c), string(debug.Stack()))
		}
	}()
	req := c.Request()
	ctx := utility.GetContextFromEchoContext(c)
	ctx = context.WithValue(ctx, common.ContextLogType, common.ContextLogTypeEnd)

	logger.InfoContext(ctx, fmt.Sprintf("[METHOD=%s] [URL=%s] [STAUTS=%d] [LATENCY=%v]", strings.ToUpper(req.Method), req.URL.String(), c.Response().Status, latency))
}
