package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZoneService struct {
	Options    []options.RequestOption
	Dnssecs    *ZonesDnssecService
	RateLimits *ZonesRateLimitService
	Settings   *ZonesSettingService
}

func NewZoneService(opts ...options.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	r.Dnssecs = NewZonesDnssecService(opts...)
	r.RateLimits = NewZonesRateLimitService(opts...)
	r.Settings = NewZonesSettingService(opts...)
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body *requests.ZoneNewParams, opts ...options.RequestOption) (res *responses.ZoneCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = options.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, identifier string, opts ...options.RequestOption) (res *responses.ZoneRetrieveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Editz a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Update(ctx context.Context, identifier string, body *requests.ZoneUpdateParams, opts ...options.RequestOption) (res *responses.ZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Lists, searches, sorts, and filters your zones.
func (r *ZoneService) List(ctx context.Context, query *requests.ZoneListParams, opts ...options.RequestOption) (res *responses.ZoneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = options.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}

// Deletes an existing zone.
func (r *ZoneService) Delete(ctx context.Context, identifier string, opts ...options.RequestOption) (res *responses.APIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s", identifier)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
