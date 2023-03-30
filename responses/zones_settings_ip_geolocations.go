package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type IpGeolocationListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result IpGeolocation `json:"result"`
	JSON   IpGeolocationListResponseJSON
}

type IpGeolocationListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into IpGeolocationListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *IpGeolocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type IpGeolocationUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result IpGeolocation `json:"result"`
	JSON   IpGeolocationUpdateResponseJSON
}

type IpGeolocationUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into IpGeolocationUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *IpGeolocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
