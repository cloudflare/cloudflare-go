package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type RocketLoaderListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
	// prioritises rendering your content while loading your site's Javascript
	// asynchronously. Turning on Rocket Loader will immediately improve a web page's
	// rendering time sometimes measured as Time to First Paint (TTFP), and also the
	// `window.onload` time (assuming there is JavaScript on the page). This can have a
	// positive impact on your Google search ranking. When turned on, Rocket Loader
	// will automatically defer the loading of all Javascript referenced in your HTML,
	// with no configuration required. Refer to
	// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
	// for more information.
	Result RocketLoader `json:"result"`
	JSON   RocketLoaderListResponseJSON
}

type RocketLoaderListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RocketLoaderListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RocketLoaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RocketLoaderUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
	// prioritises rendering your content while loading your site's Javascript
	// asynchronously. Turning on Rocket Loader will immediately improve a web page's
	// rendering time sometimes measured as Time to First Paint (TTFP), and also the
	// `window.onload` time (assuming there is JavaScript on the page). This can have a
	// positive impact on your Google search ranking. When turned on, Rocket Loader
	// will automatically defer the loading of all Javascript referenced in your HTML,
	// with no configuration required. Refer to
	// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
	// for more information.
	Result RocketLoader `json:"result"`
	JSON   RocketLoaderUpdateResponseJSON
}

type RocketLoaderUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RocketLoaderUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RocketLoaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
