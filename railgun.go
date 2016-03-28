package cloudflare

// CreateRailgun creates a new Railgun.
// API reference:
// 	https://api.cloudflare.com/#railgun-create-railgun
// 	POST /railguns
func (c *API) CreateRailgun() {
}

// API reference:
// 	https://api.cloudflare.com/#railgun-railgun-details
// 	GET /railguns/:identifier

// API reference:
// 	https://api.cloudflare.com/#railgun-get-zones-connected-to-a-railgun
// 	GET /railguns/:identifier/zones

// API reference:
// 	https://api.cloudflare.com/#railgun-enable-or-disable-a-railgun
// 	PATCH /railguns/:identifier

// API reference:
// 	https://api.cloudflare.com/#railgun-delete-railgun
// 	DELETE /railguns/:identifier

// 	Zone railgun info

// Railguns returns the available Railguns for a zone.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-get-available-railguns
// 	GET /zones/:zone_identifier/railguns
func (c *API) Railguns() {
}

// Railgun returns the configuration for a given Railgun.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-get-railgun-details
// 	GET /zones/:zone_identifier/railguns/:identifier
func (c *API) Railgun() {
}

// ZoneRailgun connects (true) or disconnects (false) a Railgun for a given zone.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-connect-or-disconnect-a-railgun
// 	PATCH /zones/:zone_identifier/railguns/:identifier
func (c *API) ZoneRailgun(connected bool) {
}
