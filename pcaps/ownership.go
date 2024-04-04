// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pcaps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// OwnershipService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewOwnershipService] method instead.
type OwnershipService struct {
	Options []option.RequestOption
}

// NewOwnershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOwnershipService(opts ...option.RequestOption) (r *OwnershipService) {
	r = &OwnershipService{}
	r.Options = opts
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *OwnershipService) New(ctx context.Context, params OwnershipNewParams, opts ...option.RequestOption) (res *MagicVisibilityPCAPsOwnership, err error) {
	opts = append(r.Options[:], opts...)
	var env OwnershipNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes buckets added to the packet captures API.
func (r *OwnershipService) Delete(ctx context.Context, ownershipID string, body OwnershipDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", body.AccountID, ownershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *OwnershipService) Get(ctx context.Context, query OwnershipGetParams, opts ...option.RequestOption) (res *[]MagicVisibilityPCAPsOwnership, err error) {
	opts = append(r.Options[:], opts...)
	var env OwnershipGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates buckets added to the packet captures API.
func (r *OwnershipService) Validate(ctx context.Context, params OwnershipValidateParams, opts ...option.RequestOption) (res *MagicVisibilityPCAPsOwnership, err error) {
	opts = append(r.Options[:], opts...)
	var env OwnershipValidateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicVisibilityPCAPsOwnership struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status MagicVisibilityPCAPsOwnershipStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                            `json:"validated"`
	JSON      magicVisibilityPCAPsOwnershipJSON `json:"-"`
}

// magicVisibilityPCAPsOwnershipJSON contains the JSON metadata for the struct
// [MagicVisibilityPCAPsOwnership]
type magicVisibilityPCAPsOwnershipJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicVisibilityPCAPsOwnership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicVisibilityPCAPsOwnershipJSON) RawJSON() string {
	return r.raw
}

// The status of the ownership challenge. Can be pending, success or failed.
type MagicVisibilityPCAPsOwnershipStatus string

const (
	MagicVisibilityPCAPsOwnershipStatusPending MagicVisibilityPCAPsOwnershipStatus = "pending"
	MagicVisibilityPCAPsOwnershipStatusSuccess MagicVisibilityPCAPsOwnershipStatus = "success"
	MagicVisibilityPCAPsOwnershipStatusFailed  MagicVisibilityPCAPsOwnershipStatus = "failed"
)

func (r MagicVisibilityPCAPsOwnershipStatus) IsKnown() bool {
	switch r {
	case MagicVisibilityPCAPsOwnershipStatusPending, MagicVisibilityPCAPsOwnershipStatusSuccess, MagicVisibilityPCAPsOwnershipStatusFailed:
		return true
	}
	return false
}

type OwnershipNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r OwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172  `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172  `json:"messages,required"`
	Result   MagicVisibilityPCAPsOwnership `json:"result,required"`
	// Whether the API call was successful
	Success OwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipNewResponseEnvelopeJSON    `json:"-"`
}

// ownershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelope]
type ownershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipNewResponseEnvelopeSuccess bool

const (
	OwnershipNewResponseEnvelopeSuccessTrue OwnershipNewResponseEnvelopeSuccess = true
)

func (r OwnershipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type OwnershipGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type OwnershipGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172    `json:"messages,required"`
	Result   []MagicVisibilityPCAPsOwnership `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    OwnershipGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo OwnershipGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ownershipGetResponseEnvelopeJSON       `json:"-"`
}

// ownershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipGetResponseEnvelope]
type ownershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipGetResponseEnvelopeSuccess bool

const (
	OwnershipGetResponseEnvelopeSuccessTrue OwnershipGetResponseEnvelopeSuccess = true
)

func (r OwnershipGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       ownershipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ownershipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [OwnershipGetResponseEnvelopeResultInfo]
type ownershipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r OwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipValidateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172  `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172  `json:"messages,required"`
	Result   MagicVisibilityPCAPsOwnership `json:"result,required"`
	// Whether the API call was successful
	Success OwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipValidateResponseEnvelopeJSON    `json:"-"`
}

// ownershipValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipValidateResponseEnvelope]
type ownershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipValidateResponseEnvelopeSuccess bool

const (
	OwnershipValidateResponseEnvelopeSuccessTrue OwnershipValidateResponseEnvelopeSuccess = true
)

func (r OwnershipValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
