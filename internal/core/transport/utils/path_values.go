package core_http_utils

import (
	"fmt"
	core_errors "gopet/internal/core/errors"
	"net/http"
	"strconv"
)

func GetIntPathValues(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("emty key '%s'path value: %w", key, core_errors.ErrInvalidArgumnet)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value = %s, by key %s, not a valid integer: %v: %w",
			pathValue,
			key,
			err,
			core_errors.ErrInvalidArgumnet,
		)
	}

	return val, nil
}
