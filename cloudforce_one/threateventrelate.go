// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ThreatEventRelateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventRelateService] method instead.
type ThreatEventRelateService struct {
	Options []option.RequestOption
}

// NewThreatEventRelateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventRelateService(opts ...option.RequestOption) (r *ThreatEventRelateService) {
	r = &ThreatEventRelateService{}
	r.Options = opts
	return
}

// Creates event references for a event
func (r *ThreatEventRelateService) Update(ctx context.Context, accountID float64, eventID string, body ThreatEventRelateUpdateParams, opts ...option.RequestOption) (res *ThreatEventRelateUpdateResponse, err error) {
	var env ThreatEventRelateUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/relate/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes an event refence
func (r *ThreatEventRelateService) Delete(ctx context.Context, accountID float64, eventID string, opts ...option.RequestOption) (res *ThreatEventRelateDeleteResponse, err error) {
	var env ThreatEventRelateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/relate/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventRelateUpdateResponse struct {
	Success bool                                `json:"success,required"`
	JSON    threatEventRelateUpdateResponseJSON `json:"-"`
}

// threatEventRelateUpdateResponseJSON contains the JSON metadata for the struct
// [ThreatEventRelateUpdateResponse]
type threatEventRelateUpdateResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRelateDeleteResponse struct {
	Success bool                                `json:"success,required"`
	JSON    threatEventRelateDeleteResponseJSON `json:"-"`
}

// threatEventRelateDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventRelateDeleteResponse]
type threatEventRelateDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRelateUpdateParams struct {
	Events param.Field[[]string] `json:"events,required"`
}

func (r ThreatEventRelateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventRelateUpdateResponseEnvelope struct {
	Result  ThreatEventRelateUpdateResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    threatEventRelateUpdateResponseEnvelopeJSON `json:"-"`
}

// threatEventRelateUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventRelateUpdateResponseEnvelope]
type threatEventRelateUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRelateDeleteResponseEnvelope struct {
	Result  ThreatEventRelateDeleteResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    threatEventRelateDeleteResponseEnvelopeJSON `json:"-"`
}

// threatEventRelateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventRelateDeleteResponseEnvelope]
type threatEventRelateDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRelateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
