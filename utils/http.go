package utils

import "net/http"

// GetIpAddress gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIpAddress(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
