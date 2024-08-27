// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DomainManagedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainManagedService] method instead.
type DomainManagedService struct {
	Options []option.RequestOption
}

// NewDomainManagedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainManagedService(opts ...option.RequestOption) (r *DomainManagedService) {
	r = &DomainManagedService{}
	r.Options = opts
	return
}

// Updates state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *DomainManagedService) Update(ctx context.Context, bucketName string, params DomainManagedUpdateParams, opts ...option.RequestOption) (res *DomainManagedUpdateResponse, err error) {
	var env DomainManagedUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets state of public access over the bucket's R2-managed (r2.dev) domain.
func (r *DomainManagedService) List(ctx context.Context, bucketName string, query DomainManagedListParams, opts ...option.RequestOption) (res *DomainManagedListResponse, err error) {
	var env DomainManagedListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/managed", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DomainManagedUpdateResponse struct {
	// Bucket ID
	BucketID string `json:"bucketId,required"`
	// Domain name of the bucket's r2.dev domain
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the r2.dev domain
	Enabled bool                            `json:"enabled,required"`
	JSON    domainManagedUpdateResponseJSON `json:"-"`
}

// domainManagedUpdateResponseJSON contains the JSON metadata for the struct
// [DomainManagedUpdateResponse]
type domainManagedUpdateResponseJSON struct {
	BucketID    apijson.Field
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainManagedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainManagedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DomainManagedListResponse struct {
	// Bucket ID
	BucketID string `json:"bucketId,required"`
	// Domain name of the bucket's r2.dev domain
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the r2.dev domain
	Enabled bool                          `json:"enabled,required"`
	JSON    domainManagedListResponseJSON `json:"-"`
}

// domainManagedListResponseJSON contains the JSON metadata for the struct
// [DomainManagedListResponse]
type domainManagedListResponseJSON struct {
	BucketID    apijson.Field
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainManagedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainManagedListResponseJSON) RawJSON() string {
	return r.raw
}

type DomainManagedUpdateParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to enable public bucket access at the r2.dev domain
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r DomainManagedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainManagedUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []string                    `json:"messages,required"`
	Result   DomainManagedUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainManagedUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainManagedUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainManagedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DomainManagedUpdateResponseEnvelope]
type domainManagedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainManagedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainManagedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainManagedUpdateResponseEnvelopeSuccess bool

const (
	DomainManagedUpdateResponseEnvelopeSuccessTrue DomainManagedUpdateResponseEnvelopeSuccess = true
)

func (r DomainManagedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainManagedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainManagedListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainManagedListResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []string                  `json:"messages,required"`
	Result   DomainManagedListResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainManagedListResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainManagedListResponseEnvelopeJSON    `json:"-"`
}

// domainManagedListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainManagedListResponseEnvelope]
type domainManagedListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainManagedListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainManagedListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainManagedListResponseEnvelopeSuccess bool

const (
	DomainManagedListResponseEnvelopeSuccessTrue DomainManagedListResponseEnvelopeSuccess = true
)

func (r DomainManagedListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainManagedListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
