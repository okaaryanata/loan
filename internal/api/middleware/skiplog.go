package middleware

import (
	"fmt"

	"github.com/okaaryanata/loan/internal/domain"
)

const (
	HomePath   = "/"
	HealthPath = "/health"
)

var (
	SkipPaths = []string{
		HomePath,
		HealthPath,
	}
)

func GetListSkipLogPath() []string {
	template := "%s%s"

	var listSkip []string
	for _, r := range SkipPaths {
		route := fmt.Sprintf(template, domain.MainRoute, r)
		listSkip = append(listSkip, route)
	}

	return listSkip
}
