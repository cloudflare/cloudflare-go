package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type AutomaticPlatformOptimizationListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                          `json:"success"`
	Result  AutomaticPlatformOptimization `json:"result"`
	JSON    AutomaticPlatformOptimizationListResponseJSON
}

type AutomaticPlatformOptimizationListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AutomaticPlatformOptimizationListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AutomaticPlatformOptimizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AutomaticPlatformOptimizationUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                          `json:"success"`
	Result  AutomaticPlatformOptimization `json:"result"`
	JSON    AutomaticPlatformOptimizationUpdateResponseJSON
}

type AutomaticPlatformOptimizationUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AutomaticPlatformOptimizationUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AutomaticPlatformOptimizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
