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
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// LeakedCredentialCheckDetectionService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialCheckDetectionService] method instead.
type LeakedCredentialCheckDetectionService struct {
	Options []option.RequestOption
}

// NewLeakedCredentialCheckDetectionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLeakedCredentialCheckDetectionService(opts ...option.RequestOption) (r *LeakedCredentialCheckDetectionService) {
	r = &LeakedCredentialCheckDetectionService{}
	r.Options = opts
	return
}

// Create user-defined detection pattern for Leaked Credential Checks
func (r *LeakedCredentialCheckDetectionService) New(ctx context.Context, params LeakedCredentialCheckDetectionNewParams, opts ...option.RequestOption) (res *LeakedCredentialCheckDetectionNewResponse, err error) {
	var env LeakedCredentialCheckDetectionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update user-defined detection pattern for Leaked Credential Checks
func (r *LeakedCredentialCheckDetectionService) Update(ctx context.Context, detectionID string, params LeakedCredentialCheckDetectionUpdateParams, opts ...option.RequestOption) (res *LeakedCredentialCheckDetectionUpdateResponse, err error) {
	var env LeakedCredentialCheckDetectionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", params.ZoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List user-defined detection patterns for Leaked Credential Checks
func (r *LeakedCredentialCheckDetectionService) List(ctx context.Context, query LeakedCredentialCheckDetectionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LeakedCredentialCheckDetectionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List user-defined detection patterns for Leaked Credential Checks
func (r *LeakedCredentialCheckDetectionService) ListAutoPaging(ctx context.Context, query LeakedCredentialCheckDetectionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LeakedCredentialCheckDetectionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Remove user-defined detection pattern for Leaked Credential Checks
func (r *LeakedCredentialCheckDetectionService) Delete(ctx context.Context, detectionID string, body LeakedCredentialCheckDetectionDeleteParams, opts ...option.RequestOption) (res *LeakedCredentialCheckDetectionDeleteResponse, err error) {
	var env LeakedCredentialCheckDetectionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", body.ZoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A custom set of username/password expressions to match Leaked Credential Checks
// on
type LeakedCredentialCheckDetectionNewResponse struct {
	// The unique ID for this custom detection
	ID string `json:"id"`
	// The ruleset expression to use in matching the password in a request
	Password string `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username string                                        `json:"username"`
	JSON     leakedCredentialCheckDetectionNewResponseJSON `json:"-"`
}

// leakedCredentialCheckDetectionNewResponseJSON contains the JSON metadata for the
// struct [LeakedCredentialCheckDetectionNewResponse]
type leakedCredentialCheckDetectionNewResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionNewResponseJSON) RawJSON() string {
	return r.raw
}

// A custom set of username/password expressions to match Leaked Credential Checks
// on
type LeakedCredentialCheckDetectionUpdateResponse struct {
	// The unique ID for this custom detection
	ID string `json:"id"`
	// The ruleset expression to use in matching the password in a request
	Password string `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username string                                           `json:"username"`
	JSON     leakedCredentialCheckDetectionUpdateResponseJSON `json:"-"`
}

// leakedCredentialCheckDetectionUpdateResponseJSON contains the JSON metadata for
// the struct [LeakedCredentialCheckDetectionUpdateResponse]
type leakedCredentialCheckDetectionUpdateResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A custom set of username/password expressions to match Leaked Credential Checks
// on
type LeakedCredentialCheckDetectionListResponse struct {
	// The unique ID for this custom detection
	ID string `json:"id"`
	// The ruleset expression to use in matching the password in a request
	Password string `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username string                                         `json:"username"`
	JSON     leakedCredentialCheckDetectionListResponseJSON `json:"-"`
}

// leakedCredentialCheckDetectionListResponseJSON contains the JSON metadata for
// the struct [LeakedCredentialCheckDetectionListResponse]
type leakedCredentialCheckDetectionListResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionListResponseJSON) RawJSON() string {
	return r.raw
}

type LeakedCredentialCheckDetectionDeleteResponse = interface{}

type LeakedCredentialCheckDetectionNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The ruleset expression to use in matching the password in a request
	Password param.Field[string] `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username param.Field[string] `json:"username"`
}

func (r LeakedCredentialCheckDetectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LeakedCredentialCheckDetectionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A custom set of username/password expressions to match Leaked Credential Checks
	// on
	Result LeakedCredentialCheckDetectionNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LeakedCredentialCheckDetectionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    leakedCredentialCheckDetectionNewResponseEnvelopeJSON    `json:"-"`
}

// leakedCredentialCheckDetectionNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [LeakedCredentialCheckDetectionNewResponseEnvelope]
type leakedCredentialCheckDetectionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LeakedCredentialCheckDetectionNewResponseEnvelopeSuccess bool

const (
	LeakedCredentialCheckDetectionNewResponseEnvelopeSuccessTrue LeakedCredentialCheckDetectionNewResponseEnvelopeSuccess = true
)

func (r LeakedCredentialCheckDetectionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LeakedCredentialCheckDetectionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LeakedCredentialCheckDetectionUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The ruleset expression to use in matching the password in a request
	Password param.Field[string] `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username param.Field[string] `json:"username"`
}

func (r LeakedCredentialCheckDetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LeakedCredentialCheckDetectionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A custom set of username/password expressions to match Leaked Credential Checks
	// on
	Result LeakedCredentialCheckDetectionUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    leakedCredentialCheckDetectionUpdateResponseEnvelopeJSON    `json:"-"`
}

// leakedCredentialCheckDetectionUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [LeakedCredentialCheckDetectionUpdateResponseEnvelope]
type leakedCredentialCheckDetectionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccess bool

const (
	LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccessTrue LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccess = true
)

func (r LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LeakedCredentialCheckDetectionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LeakedCredentialCheckDetectionListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LeakedCredentialCheckDetectionDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LeakedCredentialCheckDetectionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   LeakedCredentialCheckDetectionDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    leakedCredentialCheckDetectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// leakedCredentialCheckDetectionDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [LeakedCredentialCheckDetectionDeleteResponseEnvelope]
type leakedCredentialCheckDetectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LeakedCredentialCheckDetectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r leakedCredentialCheckDetectionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccess bool

const (
	LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccessTrue LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccess = true
)

func (r LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LeakedCredentialCheckDetectionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
