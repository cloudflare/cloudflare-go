package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsAdvancedDdoService struct {
	Options []options.RequestOption
}

func NewZonesSettingsAdvancedDdoService(opts ...options.RequestOption) (r *ZonesSettingsAdvancedDdoService) {
	r = &ZonesSettingsAdvancedDdoService{}
	r.Options = opts
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
func (r *ZonesSettingsAdvancedDdoService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.AdvancedDdoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
