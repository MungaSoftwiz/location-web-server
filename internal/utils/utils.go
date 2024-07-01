package utils

import (
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = r.Header.Get("X-Real-IP")
		if clientIP == "" {
			clientIP = r.RemoteAddr
		}
	}
	if strings.Contains(clientIP, ",") {
		clientIP = strings.Split(clientIP, ",")[0]
	}
	return clientIP
}