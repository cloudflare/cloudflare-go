package cloudflare

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// CreateZone creates a zone on an account.
// API reference:
// 	https://api.cloudflare.com/#zone-create-a-zone
//	POST /zones
func (api *API) CreateZone(z Zone) {
	// res, err := api.makeRequest("POST", "/zones", z)
}

// ListZones liist zones on an account. Optionally takes a list of zone names
// to filter against.
// API reference:
//	https://api.cloudflare.com/#zone-list-zones
//	GET /zones
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
				return []Zone{}, errors.Wrap(err, errMakeRequestError)
			}
			err = json.Unmarshal(res, &r)
			if err != nil {
				return []Zone{}, errors.Wrap(err, errUnmarshalError)
			}
			if !r.Success {
				// TODO: Provide an actual error message instead of always returning nil
				return []Zone{}, err
			}
			for zi := range r.Result {
				zones = append(zones, r.Result[zi])
			}
		}
	} else {
		// TODO: Paginate here. We only grab the first page of results.
		// Could do this concurrently after the first request by creating a
		// sync.WaitGroup or just a channel + workers.
		res, err = api.makeRequest("GET", "/zones", nil)
		if err != nil {
			return []Zone{}, errors.Wrap(err, errMakeRequestError)
		}
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []Zone{}, errors.Wrap(err, errUnmarshalError)
		}
		zones = r.Result
	}

	return zones, nil
}

// ZoneDetails fetches information about a zone.
// API reference:
// 	https://api.cloudflare.com/#zone-zone-details
// 	GET /zones/:id
func (api *API) ZoneDetails(z Zone) {
	// TODO: This should either accept a *Zone (and update it), or return a
	// (Zone, error).

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

// EditZone edits the given zone.
// API reference:
// 	https://api.cloudflare.com/#zone-edit-zone-properties
// 	PATCH /zones/:id
func EditZone() {
}

// PurgeEverything purges the cache for the given zone.
// Note: this will substantially increase load on the origin server for that
// zone if there is a high cached vs. uncached request ratio.
// API reference:
// https://api.cloudflare.com/#zone-purge-all-files
// DELETE /zones/:id/purge_cache
func (api *API) PurgeEverything(zoneID string) (PurgeCacheResponse, error) {
	uri := "/zones/" + zoneID + "/purge_cache"
	res, err := api.makeRequest("DELETE", uri, PurgeCacheRequest{true, nil, nil})
	if err != nil {
		return PurgeCacheResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PurgeCacheResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PurgeCacheResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// PurgeCache purges the cache using the given PurgeCacheRequest (zone/url/tag).
// API reference:
// 	https://api.cloudflare.com/#zone-purge-individual-files-by-url-and-cache-tags
// 	DELETE /zones/:id/purge_cache
func (api *API) PurgeCache(zoneID string, pcr PurgeCacheRequest) (PurgeCacheResponse, error) {
	uri := "/zones/" + zoneID + "/purge_cache"
	res, err := api.makeRequest("DELETE", uri, pcr)
	if err != nil {
		return PurgeCacheResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PurgeCacheResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PurgeCacheResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteZone deletes the given zone.
// API reference:
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
