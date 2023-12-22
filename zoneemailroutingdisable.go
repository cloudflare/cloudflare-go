// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneEmailRoutingDisableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneEmailRoutingDisableService] method instead.
type ZoneEmailRoutingDisableService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingDisableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingDisableService(opts ...option.RequestOption) (r *ZoneEmailRoutingDisableService) {
	r = &ZoneEmailRoutingDisableService{}
	r.Options = opts
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *ZoneEmailRoutingDisableService) EmailRoutingSettingsDisableEmailRouting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
