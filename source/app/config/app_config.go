package config

import (
	"os"
	"strconv"
)

var AppPort int

func LoadAppConfig() {
	AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
}
