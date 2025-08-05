package config

import (
	"os"
	"strconv"
)

var AuthApiProtocol string
var AuthApiHost string
var AuthApiPort int
var AuthApiURL string

var TasksApiProtocol string
var TasksApiHost string
var TasksApiPort int
var TasksApiURL string

func LoadApiConfig() {
	AuthApiProtocol = os.Getenv("AUTH_API_PROTOCOL")
	AuthApiHost = os.Getenv("AUTH_API_HOST")
	AuthApiPort, _ = strconv.Atoi(os.Getenv("AUTH_API_PORT"))
	AuthApiURL = AuthApiProtocol + "://" + AuthApiHost + ":" +
		strconv.Itoa(AuthApiPort)

	TasksApiProtocol = os.Getenv("TASKS_API_PROTOCOL")
	TasksApiHost = os.Getenv("TASKS_API_HOST")
	TasksApiPort, _ = strconv.Atoi(os.Getenv("TASKS_API_PORT"))
	TasksApiURL = TasksApiProtocol + "://" + TasksApiHost + ":" +
		strconv.Itoa(TasksApiPort)
}
