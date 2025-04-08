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

// ZoneTransferForceAXFRService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferForceAXFRService] method instead.
type ZoneTransferForceAXFRService struct {
	Options []option.RequestOption
}

// NewZoneTransferForceAXFRService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferForceAXFRService(opts ...option.RequestOption) (r *ZoneTransferForceAXFRService) {
	r = &ZoneTransferForceAXFRService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ZoneTransferForceAXFRService) New(ctx context.Context, params ZoneTransferForceAXFRNewParams, opts ...option.RequestOption) (res *ForceAXFR, err error) {
	var env ZoneTransferForceAXFRNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/force_axfr", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ForceAXFR = string

type ZoneTransferForceAXFRNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r ZoneTransferForceAXFRNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneTransferForceAXFRNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferForceAXFRNewResponseEnvelopeSuccess `json:"success,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result ForceAXFR                                    `json:"result"`
	JSON   zoneTransferForceAXFRNewResponseEnvelopeJSON `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferForceAXFRNewResponseEnvelope]
type zoneTransferForceAXFRNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferForceAXFRNewResponseEnvelopeSuccess bool

const (
	ZoneTransferForceAXFRNewResponseEnvelopeSuccessTrue ZoneTransferForceAXFRNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferForceAXFRNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferForceAXFRNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
