package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesRateLimitService struct {
	Options []options.RequestOption
}

func NewZonesRateLimitService(opts ...options.RequestOption) (r *ZonesRateLimitService) {
	r = &ZonesRateLimitService{}
	r.Options = opts
	return
}

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
func (r *ZonesRateLimitService) New(ctx context.Context, zone_identifier string, body *map[string]interface{}, opts ...options.RequestOption) (res *responses.RatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Fetches the details of a rate limit.
func (r *ZonesRateLimitService) Get(ctx context.Context, id string, params *requests.RateLimitsGetParams, opts ...options.RequestOption) (res *responses.RatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", params.ZoneIdentifier, id)
	err = options.ExecuteNewRequest(ctx, "GET", path, params, &res, opts...)
	return
}

// Fetches the rate limits for a zone.
func (r *ZonesRateLimitService) List(ctx context.Context, zone_identifier string, query *requests.RateLimitListParams, opts ...options.RequestOption) (res *responses.RatelimitCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}

// Deletes an existing rate limit.
func (r *ZonesRateLimitService) Delete(ctx context.Context, id string, params *requests.RateLimitsDeleteParams, opts ...options.RequestOption) (res *responses.RateLimitDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", params.ZoneIdentifier, id)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, params, &res, opts...)
	return
}
