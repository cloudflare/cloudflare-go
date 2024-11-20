// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// DiscoveryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscoveryService] method instead.
type DiscoveryService struct {
	Options    []option.RequestOption
	Operations *DiscoveryOperationService
}

// NewDiscoveryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiscoveryService(opts ...option.RequestOption) (r *DiscoveryService) {
	r = &DiscoveryService{}
	r.Options = opts
	r.Operations = NewDiscoveryOperationService(opts...)
	return
}

// Retrieve the most up to date view of discovered operations, rendered as OpenAPI
// schemas
func (r *DiscoveryService) Get(ctx context.Context, query DiscoveryGetParams, opts ...option.RequestOption) (res *DiscoveryGetResponse, err error) {
	var env DiscoveryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DiscoveryGetResponse struct {
	Schemas   []interface{}            `json:"schemas,required"`
	Timestamp time.Time                `json:"timestamp,required" format:"date-time"`
	JSON      discoveryGetResponseJSON `json:"-"`
}

// discoveryGetResponseJSON contains the JSON metadata for the struct
// [DiscoveryGetResponse]
type discoveryGetResponseJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryGetResponseJSON) RawJSON() string {
	return r.raw
}

type DiscoveryGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DiscoveryGetResponseEnvelope struct {
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   DiscoveryGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success DiscoveryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    discoveryGetResponseEnvelopeJSON    `json:"-"`
}

// discoveryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DiscoveryGetResponseEnvelope]
type discoveryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscoveryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discoveryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DiscoveryGetResponseEnvelopeSuccess bool

const (
	DiscoveryGetResponseEnvelopeSuccessTrue DiscoveryGetResponseEnvelopeSuccess = true
)

func (r DiscoveryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DiscoveryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
