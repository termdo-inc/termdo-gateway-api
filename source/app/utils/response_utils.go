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

func (rescp *ResponseCapture) WriteHeader(statusCode int) {
    rescp.Status = statusCode
    rescp.ResponseWriter.WriteHeader(statusCode)
}

func (rescp *ResponseCapture) Write(b []byte) (int, error) {
    return rescp.Buffer.Write(b)
}
