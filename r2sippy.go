// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *R2SippyService) Update(ctx context.Context, bucketName string, params R2SippyUpdateParams, opts ...option.RequestOption) (res *R2SippyUpdateResponse, err error) {
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
func (r *R2SippyService) Get(ctx context.Context, bucketName string, query R2SippyGetParams, opts ...option.RequestOption) (res *R2SippyGetResponse, err error) {
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

type R2SippyUpdateResponse struct {
	// Details about the configured destination bucket
	Destination R2SippyUpdateResponseDestination `json:"destination"`
	// State of Sippy for this bucket
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket
	Source R2SippyUpdateResponseSource `json:"source"`
	JSON   r2SippyUpdateResponseJSON   `json:"-"`
}

// r2SippyUpdateResponseJSON contains the JSON metadata for the struct
// [R2SippyUpdateResponse]
type r2SippyUpdateResponseJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket
type R2SippyUpdateResponseDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider
	Bucket   string                                   `json:"bucket"`
	Provider R2SippyUpdateResponseDestinationProvider `json:"provider"`
	JSON     r2SippyUpdateResponseDestinationJSON     `json:"-"`
}

// r2SippyUpdateResponseDestinationJSON contains the JSON metadata for the struct
// [R2SippyUpdateResponseDestination]
type r2SippyUpdateResponseDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyUpdateResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyUpdateResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type R2SippyUpdateResponseDestinationProvider string

const (
	R2SippyUpdateResponseDestinationProviderR2 R2SippyUpdateResponseDestinationProvider = "r2"
)

// Details about the configured source bucket
type R2SippyUpdateResponseSource struct {
	// Name of the bucket on the provider
	Bucket   string                              `json:"bucket"`
	Provider R2SippyUpdateResponseSourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only)
	Region string                          `json:"region,nullable"`
	JSON   r2SippyUpdateResponseSourceJSON `json:"-"`
}

// r2SippyUpdateResponseSourceJSON contains the JSON metadata for the struct
// [R2SippyUpdateResponseSource]
type r2SippyUpdateResponseSourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyUpdateResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyUpdateResponseSourceJSON) RawJSON() string {
	return r.raw
}

type R2SippyUpdateResponseSourceProvider string

const (
	R2SippyUpdateResponseSourceProviderAws R2SippyUpdateResponseSourceProvider = "aws"
	R2SippyUpdateResponseSourceProviderGcs R2SippyUpdateResponseSourceProvider = "gcs"
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

func (r r2SippyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type R2SippyDeleteResponseEnabled bool

const (
	R2SippyDeleteResponseEnabledFalse R2SippyDeleteResponseEnabled = false
)

type R2SippyGetResponse struct {
	// Details about the configured destination bucket
	Destination R2SippyGetResponseDestination `json:"destination"`
	// State of Sippy for this bucket
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket
	Source R2SippyGetResponseSource `json:"source"`
	JSON   r2SippyGetResponseJSON   `json:"-"`
}

// r2SippyGetResponseJSON contains the JSON metadata for the struct
// [R2SippyGetResponse]
type r2SippyGetResponseJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket
type R2SippyGetResponseDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider
	Bucket   string                                `json:"bucket"`
	Provider R2SippyGetResponseDestinationProvider `json:"provider"`
	JSON     r2SippyGetResponseDestinationJSON     `json:"-"`
}

// r2SippyGetResponseDestinationJSON contains the JSON metadata for the struct
// [R2SippyGetResponseDestination]
type r2SippyGetResponseDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyGetResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyGetResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type R2SippyGetResponseDestinationProvider string

const (
	R2SippyGetResponseDestinationProviderR2 R2SippyGetResponseDestinationProvider = "r2"
)

// Details about the configured source bucket
type R2SippyGetResponseSource struct {
	// Name of the bucket on the provider
	Bucket   string                           `json:"bucket"`
	Provider R2SippyGetResponseSourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only)
	Region string                       `json:"region,nullable"`
	JSON   r2SippyGetResponseSourceJSON `json:"-"`
}

// r2SippyGetResponseSourceJSON contains the JSON metadata for the struct
// [R2SippyGetResponseSource]
type r2SippyGetResponseSourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SippyGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2SippyGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

type R2SippyGetResponseSourceProvider string

const (
	R2SippyGetResponseSourceProviderAws R2SippyGetResponseSourceProvider = "aws"
	R2SippyGetResponseSourceProviderGcs R2SippyGetResponseSourceProvider = "gcs"
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
	Result   R2SippyUpdateResponse                 `json:"result,required"`
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

func (r r2SippyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r r2SippyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r r2SippyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r r2SippyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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
	Result   R2SippyGetResponse                 `json:"result,required"`
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

func (r r2SippyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r r2SippyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type R2SippyGetResponseEnvelopeSuccess bool

const (
	R2SippyGetResponseEnvelopeSuccessTrue R2SippyGetResponseEnvelopeSuccess = true
)
