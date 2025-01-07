// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

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
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// PrefixAdvertisementStatusService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixAdvertisementStatusService] method instead.
type PrefixAdvertisementStatusService struct {
	Options []option.RequestOption
}

// NewPrefixAdvertisementStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPrefixAdvertisementStatusService(opts ...option.RequestOption) (r *PrefixAdvertisementStatusService) {
	r = &PrefixAdvertisementStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *PrefixAdvertisementStatusService) Edit(ctx context.Context, prefixID string, params PrefixAdvertisementStatusEditParams, opts ...option.RequestOption) (res *PrefixAdvertisementStatusEditResponse, err error) {
	var env PrefixAdvertisementStatusEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the current advertisement state for a prefix.
func (r *PrefixAdvertisementStatusService) Get(ctx context.Context, prefixID string, query PrefixAdvertisementStatusGetParams, opts ...option.RequestOption) (res *PrefixAdvertisementStatusGetResponse, err error) {
	var env PrefixAdvertisementStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixAdvertisementStatusEditResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                 `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixAdvertisementStatusEditResponseJSON `json:"-"`
}

// prefixAdvertisementStatusEditResponseJSON contains the JSON metadata for the
// struct [PrefixAdvertisementStatusEditResponse]
type prefixAdvertisementStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixAdvertisementStatusGetResponseJSON `json:"-"`
}

// prefixAdvertisementStatusGetResponseJSON contains the JSON metadata for the
// struct [PrefixAdvertisementStatusGetResponse]
type prefixAdvertisementStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r PrefixAdvertisementStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixAdvertisementStatusEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixAdvertisementStatusEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixAdvertisementStatusEditResponse                `json:"result"`
	JSON    prefixAdvertisementStatusEditResponseEnvelopeJSON    `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [PrefixAdvertisementStatusEditResponseEnvelope]
type prefixAdvertisementStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixAdvertisementStatusEditResponseEnvelopeSuccess bool

const (
	PrefixAdvertisementStatusEditResponseEnvelopeSuccessTrue PrefixAdvertisementStatusEditResponseEnvelopeSuccess = true
)

func (r PrefixAdvertisementStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixAdvertisementStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixAdvertisementStatusGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixAdvertisementStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixAdvertisementStatusGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixAdvertisementStatusGetResponse                `json:"result"`
	JSON    prefixAdvertisementStatusGetResponseEnvelopeJSON    `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [PrefixAdvertisementStatusGetResponseEnvelope]
type prefixAdvertisementStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixAdvertisementStatusGetResponseEnvelopeSuccess bool

const (
	PrefixAdvertisementStatusGetResponseEnvelopeSuccessTrue PrefixAdvertisementStatusGetResponseEnvelopeSuccess = true
)

func (r PrefixAdvertisementStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixAdvertisementStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
