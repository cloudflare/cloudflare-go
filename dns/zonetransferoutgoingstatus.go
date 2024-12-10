// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ZoneTransferOutgoingStatusService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferOutgoingStatusService] method instead.
type ZoneTransferOutgoingStatusService struct {
	Options []option.RequestOption
}

// NewZoneTransferOutgoingStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneTransferOutgoingStatusService(opts ...option.RequestOption) (r *ZoneTransferOutgoingStatusService) {
	r = &ZoneTransferOutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *ZoneTransferOutgoingStatusService) Get(ctx context.Context, query ZoneTransferOutgoingStatusGetParams, opts ...option.RequestOption) (res *EnableTransfer, err error) {
	var env ZoneTransferOutgoingStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/status", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneTransferOutgoingStatusGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferOutgoingStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result EnableTransfer                                    `json:"result"`
	JSON   zoneTransferOutgoingStatusGetResponseEnvelopeJSON `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingStatusGetResponseEnvelope]
type zoneTransferOutgoingStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingStatusGetResponseEnvelopeSuccessTrue ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
