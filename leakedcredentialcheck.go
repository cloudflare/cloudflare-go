// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// LeakedCredentialCheckService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialCheckService] method instead.
type LeakedCredentialCheckService struct {
	Options    []option.RequestOption
	Detections *LeakedCredentialCheckDetectionService
}

// NewLeakedCredentialCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLeakedCredentialCheckService(opts ...option.RequestOption) (r *LeakedCredentialCheckService) {
	r = &LeakedCredentialCheckService{}
	r.Options = opts
	r.Detections = NewLeakedCredentialCheckDetectionService(opts...)
	return
}

// Updates the current status of Leaked Credential Checks
func (r *LeakedCredentialCheckService) New(ctx context.Context, params LeakedCredentialCheckNewParams, opts ...option.RequestOption) (res *LeakedCredentialCheckNewResponse, err error) {
	var env LeakedCredentialCheckNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the current status of Leaked Credential Checks
func (r *LeakedCredentialCheckService) List(ctx context.Context, query LeakedCredentialCheckListParams, opts ...option.RequestOption) (res *LeakedCredentialCheckListResponse, err error) {
	var env LeakedCredentialCheckListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The overall status for Leaked Credential Checks
type LeakedCredentialCheckNewResponse struct {
	// Whether or not Leaked Credential Checks are enabled
	Enabled bool                                 `json:"enabled"`
	JSON    leakedCredentialCheckNewResponseJSON `json:"-"`
}

// leakedCredentialCheckNewResponseJSON contains the JSON metadata for the struct
// [LeakedCredentialCheckNewResponse]
type leakedCredentialCheckNewResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckNewResponseJSON) RawJSON() string {
	return r.raw
}

// The overall status for Leaked Credential Checks
type LeakedCredentialCheckListResponse struct {
	// Whether or not Leaked Credential Checks are enabled
	Enabled bool                                  `json:"enabled"`
	JSON    leakedCredentialCheckListResponseJSON `json:"-"`
}

// leakedCredentialCheckListResponseJSON contains the JSON metadata for the struct
// [LeakedCredentialCheckListResponse]
type leakedCredentialCheckListResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckListResponseJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialCheckNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether or not Leaked Credential Checks are enabled
	Enabled param.Field[bool] `json:"enabled"`
}

func (r LeakedCredentialCheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LeakedCredentialCheckNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The overall status for Leaked Credential Checks
	Result LeakedCredentialCheckNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LeakedCredentialCheckNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    leakedCredentialCheckNewResponseEnvelopeJSON    `json:"-"`
}

// leakedCredentialCheckNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LeakedCredentialCheckNewResponseEnvelope]
type leakedCredentialCheckNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LeakedCredentialCheckNewResponseEnvelopeSuccess bool

const (
	LeakedCredentialCheckNewResponseEnvelopeSuccessTrue LeakedCredentialCheckNewResponseEnvelopeSuccess = true
)

func (r LeakedCredentialCheckNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LeakedCredentialCheckNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LeakedCredentialCheckListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LeakedCredentialCheckListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The overall status for Leaked Credential Checks
	Result LeakedCredentialCheckListResponse `json:"result,required"`
	// Whether the API call was successful
	Success LeakedCredentialCheckListResponseEnvelopeSuccess `json:"success,required"`
	JSON    leakedCredentialCheckListResponseEnvelopeJSON    `json:"-"`
}

// leakedCredentialCheckListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LeakedCredentialCheckListResponseEnvelope]
type leakedCredentialCheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LeakedCredentialCheckListResponseEnvelopeSuccess bool

const (
	LeakedCredentialCheckListResponseEnvelopeSuccessTrue LeakedCredentialCheckListResponseEnvelopeSuccess = true
)

func (r LeakedCredentialCheckListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LeakedCredentialCheckListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
