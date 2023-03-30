package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type RocketLoaderListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
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
	Result fields.Field[RocketLoader] `json:"result"`
}

// MarshalJSON serializes RocketLoaderListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RocketLoaderListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RocketLoaderListResponse) String() (result string) {
	return fmt.Sprintf("&RocketLoaderListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type RocketLoaderUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
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
	Result fields.Field[RocketLoader] `json:"result"`
}

// MarshalJSON serializes RocketLoaderUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RocketLoaderUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RocketLoaderUpdateResponse) String() (result string) {
	return fmt.Sprintf("&RocketLoaderUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type RocketLoaderUpdateParams struct {
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
	Value fields.Field[RocketLoader] `json:"value,required"`
}

// MarshalJSON serializes RocketLoaderUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RocketLoaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RocketLoaderUpdateParams) String() (result string) {
	return fmt.Sprintf("&RocketLoaderUpdateParams{Value:%s}", r.Value)
}
