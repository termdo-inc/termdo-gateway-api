package config

import (
	"os"
)

var AuthApiURL string
var TasksApiURL string

func LoadApiConfig() {
	AuthApiURL = os.Getenv("AUTH_API_URL")
	TasksApiURL = os.Getenv("TASKS_API_URL")
}
