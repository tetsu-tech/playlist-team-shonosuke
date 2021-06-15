package error

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type severity string

const (
	SeverityDefault = "DEFAULT"
	SeverityInfo    = "INFO"
	SeverityWarn    = "WARNING"
	SeverityError   = "ERROR"
)

const TimeGap = 9 * 60 * 60

type AppError struct {
	Severity severity `json:"severity"`
	Err      error    `json:"error_message"`
	Caller   string   `json:"-"`
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func New(err error, s severity) error {
	caller := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fname := filepath.Base(file)
		caller = fmt.Sprintf("file: %s, line: %d, func: %s`",
			fname, line, runtime.FuncForPC(pc).Name())
	}
	return &AppError{
		Err:      err,
		Severity: s,
		Caller:   caller,
	}
}

var (
	logger   = log.New(os.Stdout, "", 0)
	location = time.FixedZone("Asia/Tokyo", TimeGap)
)

func Logging(ctx context.Context, err error) {
	now := time.Now().In(location)
	var entry []byte
	var appErr *AppError
	var userCtx string

	firebaseUID := ctx.Value("firebaseUID")
	if firebaseUID != "" {
		userCtx = "user"
	} else {
		userCtx = "guest"
	}

	if errors.As(err, &appErr) {
		entry, err = json.Marshal(map[string]interface{}{
			"severity":           appErr.Severity,
			"message":            appErr.Error(),
			"caller":             appErr.Caller,
			"user_ctx":           userCtx,
			"users_firebase_uid": firebaseUID,
			"time":               now,
		})

		if err != nil {
			log.Println(err)
		}
	} else {
		entry, err = json.Marshal(map[string]interface{}{
			"severity":           SeverityDefault,
			"message":            "not formatted error to AppErr: " + err.Error(),
			"user_ctx":           userCtx,
			"users_firebase_uid": firebaseUID,
			"time":               now,
		})
		if err != nil {
			log.Println(err)
		}
	}

	logger.Println(string(entry))
}
