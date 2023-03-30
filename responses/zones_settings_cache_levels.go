package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type CacheLevelListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result CacheLevel `json:"result"`
	JSON   CacheLevelListResponseJSON
}

type CacheLevelListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CacheLevelListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CacheLevelListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CacheLevelUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result CacheLevel `json:"result"`
	JSON   CacheLevelUpdateResponseJSON
}

type CacheLevelUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CacheLevelUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *CacheLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
