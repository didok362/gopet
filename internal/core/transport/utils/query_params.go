package core_http_utils

import (
	"fmt"
	core_errors "gopet/internal/core/errors"
	"net/http"
	"strconv"
	"time"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("param='%s' by key='%s' not a valid integer: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgumnet,
		)

	}

	return &val, err
}

func GetDateQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	layout := "2006-01-02"

	date, err := time.Parse(layout, param)
	if err != nil {
		return nil, fmt.Errorf(
			"param = %s by key = %s not a valid date: %v :%w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgumnet,
		)
	}

	return &date, nil
}
