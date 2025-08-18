package utils

import "net/http"

func CopyHeaders(src http.Header, dst http.ResponseWriter) {
	for key, values := range src {
		dst.Header().Set(key, values[0])
		for i := 1; i < len(values); i++ {
			dst.Header().Add(key, values[i])
		}
	}
}
