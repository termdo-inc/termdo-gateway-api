package config

import (
	"os"
	"strconv"
)

var AppHost string
var AppPort int
var AppHostname string

func LoadAppConfig() {
	AppHost = os.Getenv("APP_HOST")
	AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	AppHostname, _ = os.Hostname()
}
