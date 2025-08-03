package utils

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type ResponseCapture struct {
	gin.ResponseWriter
	Buffer *bytes.Buffer
	Status int
}

func (r *ResponseCapture) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *ResponseCapture) Write(b []byte) (int, error) {
	return r.Buffer.Write(b)
}
