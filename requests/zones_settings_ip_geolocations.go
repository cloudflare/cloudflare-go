package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type IpGeolocationListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result fields.Field[IpGeolocation] `json:"result"`
}

// MarshalJSON serializes IpGeolocationListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *IpGeolocationListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r IpGeolocationListResponse) String() (result string) {
	return fmt.Sprintf("&IpGeolocationListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type IpGeolocationUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result fields.Field[IpGeolocation] `json:"result"`
}

// MarshalJSON serializes IpGeolocationUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *IpGeolocationUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r IpGeolocationUpdateResponse) String() (result string) {
	return fmt.Sprintf("&IpGeolocationUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type IpGeolocationUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[IpGeolocationValue] `json:"value,required"`
}

// MarshalJSON serializes IpGeolocationUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *IpGeolocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r IpGeolocationUpdateParams) String() (result string) {
	return fmt.Sprintf("&IpGeolocationUpdateParams{Value:%s}", r.Value)
}
