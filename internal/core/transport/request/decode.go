package core_http_request

import (
	"encoding/json"
	"fmt"
	core_errors "gopet/internal/core/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		return fmt.Errorf("decoder json err: %v: %w", err, core_errors.ErrInvalidArgumnet)
	}
	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf("decoder validator err: %v: %w", err, core_errors.ErrInvalidArgumnet)
	}

	return nil
}
