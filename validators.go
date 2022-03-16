package cloudflare

import "regexp"

func isValidZoneIdentifier(s string) bool {
	matches, _ := regexp.MatchString(`^[0-9a-fA-F]{32}$`, s)
	return matches
}
