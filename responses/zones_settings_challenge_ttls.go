package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ChallengeTtlListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ChallengeTtl `json:"result"`
	JSON   ChallengeTtlListResponseJSON
}

type ChallengeTtlListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ChallengeTtlListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ChallengeTtlListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ChallengeTtlUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ChallengeTtl `json:"result"`
	JSON   ChallengeTtlUpdateResponseJSON
}

type ChallengeTtlUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ChallengeTtlUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ChallengeTtlUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
