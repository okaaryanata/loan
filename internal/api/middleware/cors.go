package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS env variables.
const (
	// CORSAllowedHeadersKeyName is env var name for allowed headers.
	CORSAllowedHeadersKeyName = "CORS_ALLOWED_HEADERS"
	// CORSAllowedMethodsKeyName is env var name for allowed methods.
	CORSAllowedMethodsKeyName = "CORS_ALLOWED_METHODS"
	// CORSAllowedOriginsKeyName is env var name for allowed origins.
	CORSAllowedOriginsKeyName = "CORS_ALLOWED_ORIGINS"
)

var (
	corsDefaultAllowedOrigins = []string{"*"}
	corsDefaultAllowedMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodPatch,
	}
	corsDefaultAllowedHeaders = []string{
		"Authorization",
		"Accept",
		"Origin",
		"DNT",
		"Keep-Alive",
		"User-Agent",
		"Content-Type",
		"Content-Range",
		"Range",
	}
)

func getAllowedCORSHeaders() []string {
	opts := strings.TrimSpace(os.Getenv(CORSAllowedHeadersKeyName))
	if opts == "" {
		return corsDefaultAllowedHeaders
	}

	allowedHeaders := strings.Split(opts, ",")
	if len(allowedHeaders) == 0 {
		allowedHeaders = corsDefaultAllowedHeaders
	}

	return allowedHeaders
}
func getAllowedCORSMethod() []string {
	opts := strings.TrimSpace(os.Getenv(CORSAllowedMethodsKeyName))
	if opts == "" {
		return corsDefaultAllowedMethods
	}

	allowedMethods := strings.Split(opts, ",")
	if len(allowedMethods) == 0 {
		allowedMethods = corsDefaultAllowedMethods
	}

	return allowedMethods
}

func getAllowedCORSOrigins() []string {
	opts := strings.TrimSpace(os.Getenv(CORSAllowedOriginsKeyName))
	// TODO: only allow internal access (from VPN)
	// Since default is "*", ensure the config is provided through env vars
	if opts == "" {
		return corsDefaultAllowedOrigins
	}

	allowedOrigins := strings.Split(opts, ",")
	if len(allowedOrigins) == 0 {
		allowedOrigins = corsDefaultAllowedOrigins
	}

	return allowedOrigins
}

func SetCORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     getAllowedCORSOrigins(),
		AllowMethods:     getAllowedCORSMethod(),
		AllowHeaders:     getAllowedCORSHeaders(),
		AllowCredentials: true,
	})
}
