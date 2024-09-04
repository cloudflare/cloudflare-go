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

// DomainCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainCustomService] method instead.
type DomainCustomService struct {
	Options []option.RequestOption
}

// NewDomainCustomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainCustomService(opts ...option.RequestOption) (r *DomainCustomService) {
	r = &DomainCustomService{}
	r.Options = opts
	return
}

// Register a new custom domain for an existing R2 bucket.
func (r *DomainCustomService) New(ctx context.Context, bucketName string, params DomainCustomNewParams, opts ...option.RequestOption) (res *DomainCustomNewResponse, err error) {
	var env DomainCustomNewResponseEnvelope
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

// Edit the configuration for a custom domain on an existing R2 bucket.
func (r *DomainCustomService) Update(ctx context.Context, bucketName string, domainName string, params DomainCustomUpdateParams, opts ...option.RequestOption) (res *DomainCustomUpdateResponse, err error) {
	var env DomainCustomUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", params.AccountID, bucketName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all custom domains registered with an existing R2 bucket.
func (r *DomainCustomService) List(ctx context.Context, bucketName string, query DomainCustomListParams, opts ...option.RequestOption) (res *DomainCustomListResponse, err error) {
	var env DomainCustomListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove custom domain registration from an existing R2 bucket
func (r *DomainCustomService) Delete(ctx context.Context, bucketName string, domainName string, body DomainCustomDeleteParams, opts ...option.RequestOption) (res *DomainCustomDeleteResponse, err error) {
	var env DomainCustomDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/domains/custom/%s", body.AccountID, bucketName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DomainCustomNewResponse struct {
	// Domain name of the affected custom domain
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain
	Enabled bool                        `json:"enabled,required"`
	JSON    domainCustomNewResponseJSON `json:"-"`
}

// domainCustomNewResponseJSON contains the JSON metadata for the struct
// [DomainCustomNewResponse]
type domainCustomNewResponseJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

type DomainCustomUpdateResponse struct {
	// Domain name of the affected custom domain
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain
	Enabled bool                           `json:"enabled"`
	JSON    domainCustomUpdateResponseJSON `json:"-"`
}

// domainCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DomainCustomUpdateResponse]
type domainCustomUpdateResponseJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DomainCustomListResponse struct {
	Domains []DomainCustomListResponseDomain `json:"domains,required"`
	JSON    domainCustomListResponseJSON     `json:"-"`
}

// domainCustomListResponseJSON contains the JSON metadata for the struct
// [DomainCustomListResponse]
type domainCustomListResponseJSON struct {
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomListResponseJSON) RawJSON() string {
	return r.raw
}

type DomainCustomListResponseDomain struct {
	// Domain name of the custom domain to be added
	Domain string `json:"domain,required"`
	// Whether this bucket is publicly accessible at the specified custom domain
	Enabled bool                                  `json:"enabled,required"`
	Status  DomainCustomListResponseDomainsStatus `json:"status,required"`
	// Zone ID of the custom domain resides in
	ZoneID string `json:"zoneId"`
	// Zone that the custom domain resides in
	ZoneName string                             `json:"zoneName"`
	JSON     domainCustomListResponseDomainJSON `json:"-"`
}

// domainCustomListResponseDomainJSON contains the JSON metadata for the struct
// [DomainCustomListResponseDomain]
type domainCustomListResponseDomainJSON struct {
	Domain      apijson.Field
	Enabled     apijson.Field
	Status      apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomListResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomListResponseDomainJSON) RawJSON() string {
	return r.raw
}

type DomainCustomListResponseDomainsStatus struct {
	// Ownership status of the domain
	Ownership DomainCustomListResponseDomainsStatusOwnership `json:"ownership,required"`
	// SSL certificate status
	SSL  DomainCustomListResponseDomainsStatusSSL  `json:"ssl,required"`
	JSON domainCustomListResponseDomainsStatusJSON `json:"-"`
}

// domainCustomListResponseDomainsStatusJSON contains the JSON metadata for the
// struct [DomainCustomListResponseDomainsStatus]
type domainCustomListResponseDomainsStatusJSON struct {
	Ownership   apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomListResponseDomainsStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomListResponseDomainsStatusJSON) RawJSON() string {
	return r.raw
}

// Ownership status of the domain
type DomainCustomListResponseDomainsStatusOwnership string

const (
	DomainCustomListResponseDomainsStatusOwnershipPending     DomainCustomListResponseDomainsStatusOwnership = "pending"
	DomainCustomListResponseDomainsStatusOwnershipActive      DomainCustomListResponseDomainsStatusOwnership = "active"
	DomainCustomListResponseDomainsStatusOwnershipDeactivated DomainCustomListResponseDomainsStatusOwnership = "deactivated"
	DomainCustomListResponseDomainsStatusOwnershipBlocked     DomainCustomListResponseDomainsStatusOwnership = "blocked"
	DomainCustomListResponseDomainsStatusOwnershipError       DomainCustomListResponseDomainsStatusOwnership = "error"
	DomainCustomListResponseDomainsStatusOwnershipUnknown     DomainCustomListResponseDomainsStatusOwnership = "unknown"
)

func (r DomainCustomListResponseDomainsStatusOwnership) IsKnown() bool {
	switch r {
	case DomainCustomListResponseDomainsStatusOwnershipPending, DomainCustomListResponseDomainsStatusOwnershipActive, DomainCustomListResponseDomainsStatusOwnershipDeactivated, DomainCustomListResponseDomainsStatusOwnershipBlocked, DomainCustomListResponseDomainsStatusOwnershipError, DomainCustomListResponseDomainsStatusOwnershipUnknown:
		return true
	}
	return false
}

// SSL certificate status
type DomainCustomListResponseDomainsStatusSSL string

const (
	DomainCustomListResponseDomainsStatusSSLInitializing DomainCustomListResponseDomainsStatusSSL = "initializing"
	DomainCustomListResponseDomainsStatusSSLPending      DomainCustomListResponseDomainsStatusSSL = "pending"
	DomainCustomListResponseDomainsStatusSSLActive       DomainCustomListResponseDomainsStatusSSL = "active"
	DomainCustomListResponseDomainsStatusSSLDeactivated  DomainCustomListResponseDomainsStatusSSL = "deactivated"
	DomainCustomListResponseDomainsStatusSSLError        DomainCustomListResponseDomainsStatusSSL = "error"
	DomainCustomListResponseDomainsStatusSSLUnknown      DomainCustomListResponseDomainsStatusSSL = "unknown"
)

func (r DomainCustomListResponseDomainsStatusSSL) IsKnown() bool {
	switch r {
	case DomainCustomListResponseDomainsStatusSSLInitializing, DomainCustomListResponseDomainsStatusSSLPending, DomainCustomListResponseDomainsStatusSSLActive, DomainCustomListResponseDomainsStatusSSLDeactivated, DomainCustomListResponseDomainsStatusSSLError, DomainCustomListResponseDomainsStatusSSLUnknown:
		return true
	}
	return false
}

type DomainCustomDeleteResponse struct {
	// Name of the removed custom domain
	Domain string                         `json:"domain,required"`
	JSON   domainCustomDeleteResponseJSON `json:"-"`
}

// domainCustomDeleteResponseJSON contains the JSON metadata for the struct
// [DomainCustomDeleteResponse]
type domainCustomDeleteResponseJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DomainCustomNewParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the custom domain to be added
	Domain param.Field[string] `json:"domain,required"`
	// Zone ID of the custom domain
	ZoneID param.Field[string] `json:"zoneId,required"`
	// Whether to enable public bucket access at the custom domain. If undefined, the
	// domain will be enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DomainCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []string                `json:"messages,required"`
	Result   DomainCustomNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainCustomNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainCustomNewResponseEnvelopeJSON    `json:"-"`
}

// domainCustomNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainCustomNewResponseEnvelope]
type domainCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainCustomNewResponseEnvelopeSuccess bool

const (
	DomainCustomNewResponseEnvelopeSuccessTrue DomainCustomNewResponseEnvelopeSuccess = true
)

func (r DomainCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainCustomUpdateParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to enable public bucket access at the specified custom domain
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DomainCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainCustomUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []string                   `json:"messages,required"`
	Result   DomainCustomUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainCustomUpdateResponseEnvelope]
type domainCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainCustomUpdateResponseEnvelopeSuccess bool

const (
	DomainCustomUpdateResponseEnvelopeSuccessTrue DomainCustomUpdateResponseEnvelopeSuccess = true
)

func (r DomainCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainCustomListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainCustomListResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []string                 `json:"messages,required"`
	Result   DomainCustomListResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainCustomListResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainCustomListResponseEnvelopeJSON    `json:"-"`
}

// domainCustomListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainCustomListResponseEnvelope]
type domainCustomListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainCustomListResponseEnvelopeSuccess bool

const (
	DomainCustomListResponseEnvelopeSuccessTrue DomainCustomListResponseEnvelopeSuccess = true
)

func (r DomainCustomListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainCustomListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainCustomDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainCustomDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []string                   `json:"messages,required"`
	Result   DomainCustomDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success DomainCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// domainCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainCustomDeleteResponseEnvelope]
type domainCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainCustomDeleteResponseEnvelopeSuccess bool

const (
	DomainCustomDeleteResponseEnvelopeSuccessTrue DomainCustomDeleteResponseEnvelopeSuccess = true
)

func (r DomainCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
