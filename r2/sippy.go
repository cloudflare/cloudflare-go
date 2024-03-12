// File generated from our OpenAPI spec by Stainless.

package r2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SippyService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSippyService] method instead.
type SippyService struct {
	Options []option.RequestOption
}

// NewSippyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSippyService(opts ...option.RequestOption) (r *SippyService) {
	r = &SippyService{}
	r.Options = opts
	return
}

// Sets configuration for Sippy for an existing R2 bucket.
func (r *SippyService) Update(ctx context.Context, bucketName string, params SippyUpdateParams, opts ...option.RequestOption) (res *R2Sippy, err error) {
	opts = append(r.Options[:], opts...)
	var env SippyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables Sippy on this bucket
func (r *SippyService) Delete(ctx context.Context, bucketName string, body SippyDeleteParams, opts ...option.RequestOption) (res *SippyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SippyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets configuration for Sippy for an existing R2 bucket.
func (r *SippyService) Get(ctx context.Context, bucketName string, query SippyGetParams, opts ...option.RequestOption) (res *R2Sippy, err error) {
	opts = append(r.Options[:], opts...)
	var env SippyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type R2Sippy struct {
	// Details about the configured destination bucket
	Destination R2SippyDestination `json:"destination"`
	// State of Sippy for this bucket
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket
	Source R2SippySource `json:"source"`
	JSON   r2SippyJSON   `json:"-"`
}

// r2SippyJSON contains the JSON metadata for the struct [R2Sippy]
type r2SippyJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2Sippy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket
type R2SippyDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider
	Bucket   string                     `json:"bucket"`
	Provider R2SippyDestinationProvider `json:"provider"`
	JSON     r2SippyDestinationJSON     `json:"-"`
}

// r2SippyDestinationJSON contains the JSON metadata for the struct
// [R2SippyDestination]
type r2SippyDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyDestinationJSON) RawJSON() string {
	return r.raw
}

type R2SippyDestinationProvider string

const (
	R2SippyDestinationProviderR2 R2SippyDestinationProvider = "r2"
)

// Details about the configured source bucket
type R2SippySource struct {
	// Name of the bucket on the provider
	Bucket   string                `json:"bucket"`
	Provider R2SippySourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only)
	Region string            `json:"region,nullable"`
	JSON   r2SippySourceJSON `json:"-"`
}

// r2SippySourceJSON contains the JSON metadata for the struct [R2SippySource]
type r2SippySourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippySource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippySourceJSON) RawJSON() string {
	return r.raw
}

type R2SippySourceProvider string

const (
	R2SippySourceProviderAws R2SippySourceProvider = "aws"
	R2SippySourceProviderGcs R2SippySourceProvider = "gcs"
)

type SippyDeleteResponse struct {
	Enabled SippyDeleteResponseEnabled `json:"enabled"`
	JSON    sippyDeleteResponseJSON    `json:"-"`
}

// sippyDeleteResponseJSON contains the JSON metadata for the struct
// [SippyDeleteResponse]
type sippyDeleteResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SippyDeleteResponseEnabled bool

const (
	SippyDeleteResponseEnabledFalse SippyDeleteResponseEnabled = false
)

type SippyUpdateParams struct {
	// Account ID
	AccountID   param.Field[string]                       `path:"account_id,required"`
	Destination param.Field[SippyUpdateParamsDestination] `json:"destination"`
	Source      param.Field[SippyUpdateParamsSource]      `json:"source"`
}

func (r SippyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                               `json:"accessKeyId"`
	Provider    param.Field[SippyUpdateParamsDestinationProvider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsDestinationProvider string

const (
	SippyUpdateParamsDestinationProviderR2 SippyUpdateParamsDestinationProvider = "r2"
)

type SippyUpdateParamsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the GCS bucket
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket)
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket)
	PrivateKey param.Field[string]                          `json:"privateKey"`
	Provider   param.Field[SippyUpdateParamsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsSourceProvider string

const (
	SippyUpdateParamsSourceProviderGcs SippyUpdateParamsSourceProvider = "gcs"
)

type SippyUpdateResponseEnvelope struct {
	Errors   []SippyUpdateResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	Result   R2Sippy                             `json:"result,required"`
	// Whether the API call was successful
	Success SippyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyUpdateResponseEnvelopeJSON    `json:"-"`
}

// sippyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyUpdateResponseEnvelope]
type sippyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SippyUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    sippyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// sippyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SippyUpdateResponseEnvelopeErrors]
type sippyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyUpdateResponseEnvelopeSuccess bool

const (
	SippyUpdateResponseEnvelopeSuccessTrue SippyUpdateResponseEnvelopeSuccess = true
)

type SippyDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyDeleteResponseEnvelope struct {
	Errors   []SippyDeleteResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	Result   SippyDeleteResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success SippyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyDeleteResponseEnvelopeJSON    `json:"-"`
}

// sippyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyDeleteResponseEnvelope]
type sippyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SippyDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    sippyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// sippyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SippyDeleteResponseEnvelopeErrors]
type sippyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyDeleteResponseEnvelopeSuccess bool

const (
	SippyDeleteResponseEnvelopeSuccessTrue SippyDeleteResponseEnvelopeSuccess = true
)

type SippyGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyGetResponseEnvelope struct {
	Errors   []SippyGetResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                         `json:"messages,required"`
	Result   R2Sippy                          `json:"result,required"`
	// Whether the API call was successful
	Success SippyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyGetResponseEnvelopeJSON    `json:"-"`
}

// sippyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyGetResponseEnvelope]
type sippyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SippyGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    sippyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// sippyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SippyGetResponseEnvelopeErrors]
type sippyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyGetResponseEnvelopeSuccess bool

const (
	SippyGetResponseEnvelopeSuccessTrue SippyGetResponseEnvelopeSuccess = true
)
