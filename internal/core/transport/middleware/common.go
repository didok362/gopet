package core_http_middleware

import (
	"context"
	core_logger "gopet/internal/core/logger"
	core_http_respose "gopet/internal/core/transport/response"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDheader = "X-Request-ID"
)

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:5050":   {},
				"http://172.25.87.4:5050": {},
			}

			origin := r.Header.Get("Origin")

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Acces-Control-Allowed-Origin", origin)
				w.Header().Set("Acces-Control-Allowed-Method", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Acces-Control-Allowed-Method", "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDheader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDheader, requestID)
			w.Header().Set(requestIDheader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDheader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_respose.NewHTTPResponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicReponse(p, "handle http requset panic")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_respose.NewResponseWriter(w)
			start := time.Now()

			log.Debug(
				">>> incmonig HTTP request",
				zap.String("http_method", r.Method),
				zap.Time("time", start.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_Code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(start)),
			)
		})
	}
}
