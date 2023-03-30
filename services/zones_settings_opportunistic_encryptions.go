package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsOpportunisticEncryptionService struct {
	Options []options.RequestOption
}

func NewZonesSettingsOpportunisticEncryptionService(opts ...options.RequestOption) (r *ZonesSettingsOpportunisticEncryptionService) {
	r = &ZonesSettingsOpportunisticEncryptionService{}
	r.Options = opts
	return
}

// Changes Opportunistic Encryption setting.
func (r *ZonesSettingsOpportunisticEncryptionService) Update(ctx context.Context, zone_identifier string, body *requests.OpportunisticEncryptionUpdateParams, opts ...options.RequestOption) (res *responses.OpportunisticEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets Opportunistic Encryption setting.
func (r *ZonesSettingsOpportunisticEncryptionService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.OpportunisticEncryptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
