// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

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

// BucketDomainCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketDomainCustomService] method instead.
type BucketDomainCustomService struct {
	Options []option.RequestOption
}

// NewBucketDomainCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketDomainCustomService(opts ...option.RequestOption) (r *BucketDomainCustomService) {
	r = &BucketDomainCustomService{}
	r.Options = opts
	return
}

// Register a new custom domain for an existing R2 bucket.
func (r *BucketDomainCustomService) New(ctx context.Context, bucketName string, params BucketDomainCustomNewParams, opts ...option.RequestOption) (res *BucketDomainCustomNewResponse, err error) {
	var env BucketDomainCustomNewResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all custom domains registered with an existing R2 bucket.
func (r *BucketDomainCustomService) List(ctx context.Context, bucketName string, params BucketDomainCustomListParams, opts ...option.RequestOption) (res *BucketDomainCustomListResponse, err error) {
	var env BucketDomainCustomListResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketDomainCustomNewResponse struct {
	// Domain name of the affected custom domain
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain
	Enabled bool `json:"enabled,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomNewResponseMinTLS `json:"minTLS"`
	JSON   bucketDomainCustomNewResponseJSON   `json:"-"`
}

// bucketDomainCustomNewResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomNewResponse]
type bucketDomainCustomNewResponseJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	MinTLS      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomNewResponseMinTLS string

const (
	BucketDomainCustomNewResponseMinTLS1_0 BucketDomainCustomNewResponseMinTLS = "1.0"
	BucketDomainCustomNewResponseMinTLS1_1 BucketDomainCustomNewResponseMinTLS = "1.1"
	BucketDomainCustomNewResponseMinTLS1_2 BucketDomainCustomNewResponseMinTLS = "1.2"
	BucketDomainCustomNewResponseMinTLS1_3 BucketDomainCustomNewResponseMinTLS = "1.3"
)

func (r BucketDomainCustomNewResponseMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewResponseMinTLS1_0, BucketDomainCustomNewResponseMinTLS1_1, BucketDomainCustomNewResponseMinTLS1_2, BucketDomainCustomNewResponseMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomListResponse struct {
	Domains []BucketDomainCustomListResponseDomain `json:"domains,required"`
	JSON    bucketDomainCustomListResponseJSON     `json:"-"`
}

// bucketDomainCustomListResponseJSON contains the JSON metadata for the struct
// [BucketDomainCustomListResponse]
type bucketDomainCustomListResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomListResponseDomain struct {
	// Domain name of the custom domain to be added
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain
	Enabled bool                                        `json:"enabled,required"`
	Status  BucketDomainCustomListResponseDomainsStatus `json:"status,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS BucketDomainCustomListResponseDomainsMinTLS `json:"minTLS"`
	// Zone ID of the custom domain resides in
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in
	ZoneName string                                   `json:"zoneName"`
	JSON     bucketDomainCustomListResponseDomainJSON `json:"-"`
}

// bucketDomainCustomListResponseDomainJSON contains the JSON metadata for the
// struct [BucketDomainCustomListResponseDomain]
type bucketDomainCustomListResponseDomainJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	MinTLS      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseDomainJSON) RawJSON() string {
	return r.raw
}

type BucketDomainCustomListResponseDomainsStatus struct {
	// Ownership status of the domain
	Ownership BucketDomainCustomListResponseDomainsStatusOwnership `json:"ownership,required"`
	// SSL certificate status
	SSL  BucketDomainCustomListResponseDomainsStatusSSL  `json:"ssl,required"`
	JSON bucketDomainCustomListResponseDomainsStatusJSON `json:"-"`
}

// bucketDomainCustomListResponseDomainsStatusJSON contains the JSON metadata for
// the struct [BucketDomainCustomListResponseDomainsStatus]
type bucketDomainCustomListResponseDomainsStatusJSON struct {
	Ownership   apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseDomainsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseDomainsStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain
type BucketDomainCustomListResponseDomainsStatusOwnership string

const (
	BucketDomainCustomListResponseDomainsStatusOwnershipPending     BucketDomainCustomListResponseDomainsStatusOwnership = "pending"
	BucketDomainCustomListResponseDomainsStatusOwnershipActive      BucketDomainCustomListResponseDomainsStatusOwnership = "active"
	BucketDomainCustomListResponseDomainsStatusOwnershipDeactivated BucketDomainCustomListResponseDomainsStatusOwnership = "deactivated"
	BucketDomainCustomListResponseDomainsStatusOwnershipBlocked     BucketDomainCustomListResponseDomainsStatusOwnership = "blocked"
	BucketDomainCustomListResponseDomainsStatusOwnershipError       BucketDomainCustomListResponseDomainsStatusOwnership = "error"
	BucketDomainCustomListResponseDomainsStatusOwnershipUnknown     BucketDomainCustomListResponseDomainsStatusOwnership = "unknown"
)

func (r BucketDomainCustomListResponseDomainsStatusOwnership) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsStatusOwnershipPending, BucketDomainCustomListResponseDomainsStatusOwnershipActive, BucketDomainCustomListResponseDomainsStatusOwnershipDeactivated, BucketDomainCustomListResponseDomainsStatusOwnershipBlocked, BucketDomainCustomListResponseDomainsStatusOwnershipError, BucketDomainCustomListResponseDomainsStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status
type BucketDomainCustomListResponseDomainsStatusSSL string

const (
	BucketDomainCustomListResponseDomainsStatusSSLInitializing BucketDomainCustomListResponseDomainsStatusSSL = "initializing"
	BucketDomainCustomListResponseDomainsStatusSSLPending      BucketDomainCustomListResponseDomainsStatusSSL = "pending"
	BucketDomainCustomListResponseDomainsStatusSSLActive       BucketDomainCustomListResponseDomainsStatusSSL = "active"
	BucketDomainCustomListResponseDomainsStatusSSLDeactivated  BucketDomainCustomListResponseDomainsStatusSSL = "deactivated"
	BucketDomainCustomListResponseDomainsStatusSSLError        BucketDomainCustomListResponseDomainsStatusSSL = "error"
	BucketDomainCustomListResponseDomainsStatusSSLUnknown      BucketDomainCustomListResponseDomainsStatusSSL = "unknown"
)

func (r BucketDomainCustomListResponseDomainsStatusSSL) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsStatusSSLInitializing, BucketDomainCustomListResponseDomainsStatusSSLPending, BucketDomainCustomListResponseDomainsStatusSSLActive, BucketDomainCustomListResponseDomainsStatusSSLDeactivated, BucketDomainCustomListResponseDomainsStatusSSLError, BucketDomainCustomListResponseDomainsStatusSSLUnknown:
		return true
	}
	return false
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomListResponseDomainsMinTLS string

const (
	BucketDomainCustomListResponseDomainsMinTLS1_0 BucketDomainCustomListResponseDomainsMinTLS = "1.0"
	BucketDomainCustomListResponseDomainsMinTLS1_1 BucketDomainCustomListResponseDomainsMinTLS = "1.1"
	BucketDomainCustomListResponseDomainsMinTLS1_2 BucketDomainCustomListResponseDomainsMinTLS = "1.2"
	BucketDomainCustomListResponseDomainsMinTLS1_3 BucketDomainCustomListResponseDomainsMinTLS = "1.3"
)

func (r BucketDomainCustomListResponseDomainsMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseDomainsMinTLS1_0, BucketDomainCustomListResponseDomainsMinTLS1_1, BucketDomainCustomListResponseDomainsMinTLS1_2, BucketDomainCustomListResponseDomainsMinTLS1_3:
		return true
	}
	return false
}

type BucketDomainCustomNewParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the custom domain to be added
	Domain param.Field[string] `json:"domain,required"`
	// Whether to enable public bucket access at the custom domain. If undefined, the
	// domain will be enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Zone ID of the custom domain
	ZoneID param.Field[string] `json:"zoneId,required"`
	// Minimum TLS Version the custom domain will accept for incoming connections. If
	// not set, defaults to 1.0.
	MinTLS param.Field[BucketDomainCustomNewParamsMinTLS] `json:"minTLS"`
	// The bucket jurisdiction
	Jurisdiction param.Field[BucketDomainCustomNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketDomainCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Minimum TLS Version the custom domain will accept for incoming connections. If
// not set, defaults to 1.0.
type BucketDomainCustomNewParamsMinTLS string

const (
	BucketDomainCustomNewParamsMinTLS1_0 BucketDomainCustomNewParamsMinTLS = "1.0"
	BucketDomainCustomNewParamsMinTLS1_1 BucketDomainCustomNewParamsMinTLS = "1.1"
	BucketDomainCustomNewParamsMinTLS1_2 BucketDomainCustomNewParamsMinTLS = "1.2"
	BucketDomainCustomNewParamsMinTLS1_3 BucketDomainCustomNewParamsMinTLS = "1.3"
)

func (r BucketDomainCustomNewParamsMinTLS) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewParamsMinTLS1_0, BucketDomainCustomNewParamsMinTLS1_1, BucketDomainCustomNewParamsMinTLS1_2, BucketDomainCustomNewParamsMinTLS1_3:
		return true
	}
	return false
}

// The bucket jurisdiction
type BucketDomainCustomNewParamsCfR2Jurisdiction string

const (
	BucketDomainCustomNewParamsCfR2JurisdictionDefault BucketDomainCustomNewParamsCfR2Jurisdiction = "default"
	BucketDomainCustomNewParamsCfR2JurisdictionEu      BucketDomainCustomNewParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomNewParamsCfR2JurisdictionFedramp BucketDomainCustomNewParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewParamsCfR2JurisdictionDefault, BucketDomainCustomNewParamsCfR2JurisdictionEu, BucketDomainCustomNewParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   BucketDomainCustomNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success BucketDomainCustomNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomNewResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomNewResponseEnvelope]
type bucketDomainCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketDomainCustomNewResponseEnvelopeSuccess bool

const (
	BucketDomainCustomNewResponseEnvelopeSuccessTrue BucketDomainCustomNewResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketDomainCustomListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction
	Jurisdiction param.Field[BucketDomainCustomListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type BucketDomainCustomListParamsCfR2Jurisdiction string

const (
	BucketDomainCustomListParamsCfR2JurisdictionDefault BucketDomainCustomListParamsCfR2Jurisdiction = "default"
	BucketDomainCustomListParamsCfR2JurisdictionEu      BucketDomainCustomListParamsCfR2Jurisdiction = "eu"
	BucketDomainCustomListParamsCfR2JurisdictionFedramp BucketDomainCustomListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketDomainCustomListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketDomainCustomListParamsCfR2JurisdictionDefault, BucketDomainCustomListParamsCfR2JurisdictionEu, BucketDomainCustomListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketDomainCustomListResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []string                       `json:"messages,required"`
	Result   BucketDomainCustomListResponse `json:"result,required"`
	// Whether the API call was successful
	Success BucketDomainCustomListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDomainCustomListResponseEnvelopeJSON    `json:"-"`
}

// bucketDomainCustomListResponseEnvelopeJSON contains the JSON metadata for the
// struct [BucketDomainCustomListResponseEnvelope]
type bucketDomainCustomListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDomainCustomListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDomainCustomListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketDomainCustomListResponseEnvelopeSuccess bool

const (
	BucketDomainCustomListResponseEnvelopeSuccessTrue BucketDomainCustomListResponseEnvelopeSuccess = true
)

func (r BucketDomainCustomListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDomainCustomListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
