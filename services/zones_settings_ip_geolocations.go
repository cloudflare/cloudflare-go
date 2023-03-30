package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsIpGeolocationService struct {
	Options []options.RequestOption
}

func NewZonesSettingsIpGeolocationService(opts ...options.RequestOption) (r *ZonesSettingsIpGeolocationService) {
	r = &ZonesSettingsIpGeolocationService{}
	r.Options = opts
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZonesSettingsIpGeolocationService) Update(ctx context.Context, zone_identifier string, body *requests.IpGeolocationUpdateParams, opts ...options.RequestOption) (res *responses.IpGeolocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZonesSettingsIpGeolocationService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.IpGeolocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
