package cloudflare

import "os"

// GetAPIUrl Returns CLOUDFLARE_API_HOSTNAME or default api hostname
func GetAPIUrl() (apiURL string) {
	var apiHost string
	if apiHost = os.Getenv("CLOUDFLARE_API_HOSTNAME"); apiHost == "" {
		apiHost = "api.cloudflare.com"
	}
	return "https://" + apiHost + "/client/v4"
}
