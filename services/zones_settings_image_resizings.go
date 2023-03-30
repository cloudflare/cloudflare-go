package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsImageResizingService struct {
	Options []options.RequestOption
}

func NewZonesSettingsImageResizingService(opts ...options.RequestOption) (r *ZonesSettingsImageResizingService) {
	r = &ZonesSettingsImageResizingService{}
	r.Options = opts
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *ZonesSettingsImageResizingService) Update(ctx context.Context, zone_identifier string, body *requests.ImageResizingUpdateParams, opts ...options.RequestOption) (res *responses.ImageResizingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *ZonesSettingsImageResizingService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ImageResizingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
