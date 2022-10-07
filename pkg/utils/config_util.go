package utils

import "fmt"

// DefaultConfigStrategy describes loading order of config files, later config file overrides earlier configs
func DefaultConfigStrategy(env string) []string {
	if env == "" {
		return []string{"conf"}
	}
	return []string{fmt.Sprintf("conf-%s", env)}
}
