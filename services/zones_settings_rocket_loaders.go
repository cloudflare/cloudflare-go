package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsRocketLoaderService struct {
	Options []options.RequestOption
}

func NewZonesSettingsRocketLoaderService(opts ...options.RequestOption) (r *ZonesSettingsRocketLoaderService) {
	r = &ZonesSettingsRocketLoaderService{}
	r.Options = opts
	return
}

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
func (r *ZonesSettingsRocketLoaderService) Update(ctx context.Context, zone_identifier string, body *requests.RocketLoaderUpdateParams, opts ...options.RequestOption) (res *responses.RocketLoaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

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
func (r *ZonesSettingsRocketLoaderService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.RocketLoaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/rocket_loader", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
