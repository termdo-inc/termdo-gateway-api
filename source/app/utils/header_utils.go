package utils

import "net/http"

func CopyHeaders(src http.Header, dst http.ResponseWriter) {
	for key, values := range src {
		for _, v := range values {
			dst.Header().Add(key, v)
		}
	}
}
