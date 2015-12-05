package cloudflare

import (
	"encoding/json"
	"net/url"
)

/*
Creates a zone on an account.

API reference: https://api.cloudflare.com/#zone-create-a-zone
*/
func (api *API) CreateZone(z Zone) {
	// res, err := api.makeRequest("POST", "/zones", z)
}

/*
List zones on an account. Optionally takes a list of zones to filter results.

API reference: https://api.cloudflare.com/#zone-list-zones
*/
func (api *API) ListZones(z ...string) ([]Zone, error) {
	v := url.Values{}
	var res []byte
	var r ZoneResponse
	var zones []Zone
	var err error
	if len(z) > 0 {
		for _, zone := range z {
			v.Set("name", zone)
			res, err = api.makeRequest("GET", "/zones?"+v.Encode(), nil)
			if err != nil {
				return []Zone{}, err
			}
			err = json.Unmarshal(res, &r)
			if err != nil {
				return []Zone{}, err
			}
			if !r.Success {
				return []Zone{}, err //errors.New(r.Messages)
			}
			for zi, _ := range r.Result {
				zones = append(zones, r.Result[zi])
			}
		}
	} else {
		res, err = api.makeRequest("GET", "/zones", nil)
		if err != nil {
			return []Zone{}, err
		}
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []Zone{}, err
		}
		zones = r.Result
	}

	return zones, nil
}

/*
Fetches information about a zone.


  https://api.cloudflare.com/#zone-zone-details
  GET /zones/:id
*/
func (api *API) ZoneDetails(z Zone) {
	// XXX: Should we make the user get the zone ID themselves with ListZones, or do the hard work here?
	// ListZones gives the same information as this endpoint anyway so perhaps this is of limited use?
	// Maybe for users who already know the ID or fetched it in another call.
	type result struct {
		Response
		Result Zone `json:"result"`
	}
	// If z has an ID then query for that directly, else call ListZones to
	// fetch by name.
	// var zone Zone
	if z.ID != "" {
		// res, _ := makeRequest(c, "GET", "/zones/"+z.ID, nil)
		// zone = res.Result
	} else {
		// zones, err := ListZones(c, z.Name)
		// if err != nil {
		// return
		// }
		// Only one zone should have been returned
		// zone := zones[0]
	}
}

// https://api.cloudflare.com/#zone-edit-zone-properties
// PATCH /zones/:id
func EditZone() {
}

// https://api.cloudflare.com/#zone-purge-all-files
// DELETE /zones/:id/purge_all
func PurgeAll() {
}

// https://api.cloudflare.com/#zone-purge-individual-files
// DELETE /zones/:id/purge_cache
func PurgeFile() {
}

// https://api.cloudflare.com/#zone-delete-a-zone
// DELETE /zones/:id
func DeleteZone() {
}

// Zone Plan
// https://api.cloudflare.com/#zone-plan-available-plans
// https://api.cloudflare.com/#zone-plan-plan-details

// Zone Settings
// https://api.cloudflare.com/#zone-settings-for-a-zone-get-all-zone-settings
// e.g.
// https://api.cloudflare.com/#zone-settings-for-a-zone-get-always-online-setting
// https://api.cloudflare.com/#zone-settings-for-a-zone-change-always-online-setting
