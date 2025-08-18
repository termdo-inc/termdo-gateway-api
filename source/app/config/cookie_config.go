package config

import (
	"log"
	"os"
	"strconv"
)

var CookieSecure bool

func LoadCookieConfig() {
	CookieSecure = parseBool(os.Getenv("COOKIE_IS_SECURE"))
}

func parseBool(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Fatalf("Invalid environment variable '%q': %v", v, err)
	}
	return b
}
