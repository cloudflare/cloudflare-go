// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// R2SippyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewR2SippyService] method instead.
type R2SippyService struct {
	Options []option.RequestOption
}

// NewR2SippyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewR2SippyService(opts ...option.RequestOption) (r *R2SippyService) {
	r = &R2SippyService{}
	r.Options = opts
	return
}

// Sets configuration for Sippy for an existing R2 bucket.
func (r *R2SippyService) Update(ctx context.Context, bucketName string, params R2SippyUpdateParams, opts ...option.RequestOption) (res *R2Sippy, err error) {
	opts = append(r.Options[:], opts...)
	var env R2SippyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables Sippy on this bucket
func (r *R2SippyService) Delete(ctx context.Context, bucketName string, body R2SippyDeleteParams, opts ...option.RequestOption) (res *R2SippyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env R2SippyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets configuration for Sippy for an existing R2 bucket.
func (r *R2SippyService) Get(ctx context.Context, bucketName string, query R2SippyGetParams, opts ...option.RequestOption) (res *R2Sippy, err error) {
	opts = append(r.Options[:], opts...)
	var env R2SippyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

type R2SippySourceProvider string

const (
	R2SippySourceProviderAws R2SippySourceProvider = "aws"
	R2SippySourceProviderGcs R2SippySourceProvider = "gcs"
)

type R2SippyDeleteResponse struct {
	Enabled R2SippyDeleteResponseEnabled `json:"enabled"`
	JSON    r2SippyDeleteResponseJSON    `json:"-"`
}

// r2SippyDeleteResponseJSON contains the JSON metadata for the struct
// [R2SippyDeleteResponse]
type r2SippyDeleteResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SippyDeleteResponseEnabled bool

const (
	R2SippyDeleteResponseEnabledFalse R2SippyDeleteResponseEnabled = false
)

type R2SippyUpdateParams struct {
	// Account ID
	AccountID   param.Field[string]                         `path:"account_id,required"`
	Destination param.Field[R2SippyUpdateParamsDestination] `json:"destination"`
	Source      param.Field[R2SippyUpdateParamsSource]      `json:"source"`
}

func (r R2SippyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2SippyUpdateParamsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                                 `json:"accessKeyId"`
	Provider    param.Field[R2SippyUpdateParamsDestinationProvider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r R2SippyUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2SippyUpdateParamsDestinationProvider string

const (
	R2SippyUpdateParamsDestinationProviderR2 R2SippyUpdateParamsDestinationProvider = "r2"
)

type R2SippyUpdateParamsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the GCS bucket
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket)
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket)
	PrivateKey param.Field[string]                            `json:"privateKey"`
	Provider   param.Field[R2SippyUpdateParamsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r R2SippyUpdateParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2SippyUpdateParamsSourceProvider string

const (
	R2SippyUpdateParamsSourceProviderGcs R2SippyUpdateParamsSourceProvider = "gcs"
)

type R2SippyUpdateResponseEnvelope struct {
	Errors   []R2SippyUpdateResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   R2Sippy                               `json:"result,required"`
	// Whether the API call was successful
	Success R2SippyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2SippyUpdateResponseEnvelopeJSON    `json:"-"`
}

// r2SippyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2SippyUpdateResponseEnvelope]
type r2SippyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SippyUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    r2SippyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// r2SippyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2SippyUpdateResponseEnvelopeErrors]
type r2SippyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2SippyUpdateResponseEnvelopeSuccess bool

const (
	R2SippyUpdateResponseEnvelopeSuccessTrue R2SippyUpdateResponseEnvelopeSuccess = true
)

type R2SippyDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2SippyDeleteResponseEnvelope struct {
	Errors   []R2SippyDeleteResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   R2SippyDeleteResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success R2SippyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2SippyDeleteResponseEnvelopeJSON    `json:"-"`
}

// r2SippyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2SippyDeleteResponseEnvelope]
type r2SippyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SippyDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    r2SippyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// r2SippyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2SippyDeleteResponseEnvelopeErrors]
type r2SippyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2SippyDeleteResponseEnvelopeSuccess bool

const (
	R2SippyDeleteResponseEnvelopeSuccessTrue R2SippyDeleteResponseEnvelopeSuccess = true
)

type R2SippyGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2SippyGetResponseEnvelope struct {
	Errors   []R2SippyGetResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   R2Sippy                            `json:"result,required"`
	// Whether the API call was successful
	Success R2SippyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2SippyGetResponseEnvelopeJSON    `json:"-"`
}

// r2SippyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2SippyGetResponseEnvelope]
type r2SippyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SippyGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    r2SippyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// r2SippyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [R2SippyGetResponseEnvelopeErrors]
type r2SippyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2SippyGetResponseEnvelopeSuccess bool

const (
	R2SippyGetResponseEnvelopeSuccessTrue R2SippyGetResponseEnvelopeSuccess = true
)
