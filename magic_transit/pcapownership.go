// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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

// PCAPOwnershipService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPCAPOwnershipService] method instead.
type PCAPOwnershipService struct {
	Options []option.RequestOption
}

// NewPCAPOwnershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPCAPOwnershipService(opts ...option.RequestOption) (r *PCAPOwnershipService) {
	r = &PCAPOwnershipService{}
	r.Options = opts
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *PCAPOwnershipService) New(ctx context.Context, params PCAPOwnershipNewParams, opts ...option.RequestOption) (res *Ownership, err error) {
	var env PCAPOwnershipNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes buckets added to the packet captures API.
func (r *PCAPOwnershipService) Delete(ctx context.Context, ownershipID string, body PCAPOwnershipDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ownershipID == "" {
		err = errors.New("missing required ownership_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", body.AccountID, ownershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *PCAPOwnershipService) Get(ctx context.Context, query PCAPOwnershipGetParams, opts ...option.RequestOption) (res *[]Ownership, err error) {
	var env PCAPOwnershipGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates buckets added to the packet captures API.
func (r *PCAPOwnershipService) Validate(ctx context.Context, params PCAPOwnershipValidateParams, opts ...option.RequestOption) (res *Ownership, err error) {
	var env PCAPOwnershipValidateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Ownership struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status OwnershipStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string        `json:"validated"`
	JSON      ownershipJSON `json:"-"`
}

// ownershipJSON contains the JSON metadata for the struct [Ownership]
type ownershipJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Ownership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipJSON) RawJSON() string {
	return r.raw
}

// The status of the ownership challenge. Can be pending, success or failed.
type OwnershipStatus string

const (
	OwnershipStatusPending OwnershipStatus = "pending"
	OwnershipStatusSuccess OwnershipStatus = "success"
	OwnershipStatusFailed  OwnershipStatus = "failed"
)

func (r OwnershipStatus) IsKnown() bool {
	switch r {
	case OwnershipStatusPending, OwnershipStatusSuccess, OwnershipStatusFailed:
		return true
	}
	return false
}

type PCAPOwnershipNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r PCAPOwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPOwnershipNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Ownership             `json:"result,required"`
	// Whether the API call was successful
	Success PCAPOwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipNewResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPOwnershipNewResponseEnvelope]
type pcapOwnershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapOwnershipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PCAPOwnershipNewResponseEnvelopeSuccess bool

const (
	PCAPOwnershipNewResponseEnvelopeSuccessTrue PCAPOwnershipNewResponseEnvelopeSuccess = true
)

func (r PCAPOwnershipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PCAPOwnershipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PCAPOwnershipDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PCAPOwnershipGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PCAPOwnershipGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Ownership           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PCAPOwnershipGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PCAPOwnershipGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapOwnershipGetResponseEnvelopeJSON       `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPOwnershipGetResponseEnvelope]
type pcapOwnershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapOwnershipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PCAPOwnershipGetResponseEnvelopeSuccess bool

const (
	PCAPOwnershipGetResponseEnvelopeSuccessTrue PCAPOwnershipGetResponseEnvelopeSuccess = true
)

func (r PCAPOwnershipGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PCAPOwnershipGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PCAPOwnershipGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       pcapOwnershipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PCAPOwnershipGetResponseEnvelopeResultInfo]
type pcapOwnershipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapOwnershipGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PCAPOwnershipValidateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r PCAPOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPOwnershipValidateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Ownership             `json:"result,required"`
	// Whether the API call was successful
	Success PCAPOwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeJSON contains the JSON metadata for the
// struct [PCAPOwnershipValidateResponseEnvelope]
type pcapOwnershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPOwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapOwnershipValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PCAPOwnershipValidateResponseEnvelopeSuccess bool

const (
	PCAPOwnershipValidateResponseEnvelopeSuccessTrue PCAPOwnershipValidateResponseEnvelopeSuccess = true
)

func (r PCAPOwnershipValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PCAPOwnershipValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
