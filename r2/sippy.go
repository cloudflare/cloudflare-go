// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.getAccountID(), bucketName)
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
	Bucket   string                                           `json:"bucket"`
	Provider UnnamedSchemaRef6430970563db310f19d39aafe3debd27 `json:"provider"`
	JSON     r2SippyDestinationJSON                           `json:"-"`
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

func (r R2SippySourceProvider) IsKnown() bool {
	switch r {
	case R2SippySourceProviderAws, R2SippySourceProviderGcs:
		return true
	}
	return false
}

type UnnamedSchemaRef6430970563db310f19d39aafe3debd27 string

const (
	UnnamedSchemaRef6430970563db310f19d39aafe3debd27R2 UnnamedSchemaRef6430970563db310f19d39aafe3debd27 = "r2"
)

func (r UnnamedSchemaRef6430970563db310f19d39aafe3debd27) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef6430970563db310f19d39aafe3debd27R2:
		return true
	}
	return false
}

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

func (r SippyDeleteResponseEnabled) IsKnown() bool {
	switch r {
	case SippyDeleteResponseEnabledFalse:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [SippyUpdateParamsR2EnableSippyAws], [SippyUpdateParamsR2EnableSippyGcs].
type SippyUpdateParams interface {
	ImplementsSippyUpdateParams()

	getAccountID() param.Field[string]
}

type SippyUpdateParamsR2EnableSippyAws struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// R2 bucket to copy objects to
	Destination param.Field[SippyUpdateParamsR2EnableSippyAwsDestination] `json:"destination"`
	// AWS S3 bucket to copy objects from
	Source param.Field[SippyUpdateParamsR2EnableSippyAwsSource] `json:"source"`
}

func (r SippyUpdateParamsR2EnableSippyAws) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SippyUpdateParamsR2EnableSippyAws) getAccountID() param.Field[string] {
	return r.AccountID
}

func (SippyUpdateParamsR2EnableSippyAws) ImplementsSippyUpdateParams() {

}

// R2 bucket to copy objects to
type SippyUpdateParamsR2EnableSippyAwsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                                           `json:"accessKeyId"`
	Provider    param.Field[UnnamedSchemaRef6430970563db310f19d39aafe3debd27] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsR2EnableSippyAwsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AWS S3 bucket to copy objects from
type SippyUpdateParamsR2EnableSippyAwsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the AWS S3 bucket
	Bucket   param.Field[string]                                          `json:"bucket"`
	Provider param.Field[SippyUpdateParamsR2EnableSippyAwsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsR2EnableSippyAwsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsR2EnableSippyAwsSourceProvider string

const (
	SippyUpdateParamsR2EnableSippyAwsSourceProviderAws SippyUpdateParamsR2EnableSippyAwsSourceProvider = "aws"
)

func (r SippyUpdateParamsR2EnableSippyAwsSourceProvider) IsKnown() bool {
	switch r {
	case SippyUpdateParamsR2EnableSippyAwsSourceProviderAws:
		return true
	}
	return false
}

type SippyUpdateParamsR2EnableSippyGcs struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// R2 bucket to copy objects to
	Destination param.Field[SippyUpdateParamsR2EnableSippyGcsDestination] `json:"destination"`
	// GCS bucket to copy objects from
	Source param.Field[SippyUpdateParamsR2EnableSippyGcsSource] `json:"source"`
}

func (r SippyUpdateParamsR2EnableSippyGcs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SippyUpdateParamsR2EnableSippyGcs) getAccountID() param.Field[string] {
	return r.AccountID
}

func (SippyUpdateParamsR2EnableSippyGcs) ImplementsSippyUpdateParams() {

}

// R2 bucket to copy objects to
type SippyUpdateParamsR2EnableSippyGcsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]                                           `json:"accessKeyId"`
	Provider    param.Field[UnnamedSchemaRef6430970563db310f19d39aafe3debd27] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsR2EnableSippyGcsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GCS bucket to copy objects from
type SippyUpdateParamsR2EnableSippyGcsSource struct {
	// Name of the GCS bucket
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket)
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket)
	PrivateKey param.Field[string]                                          `json:"privateKey"`
	Provider   param.Field[SippyUpdateParamsR2EnableSippyGcsSourceProvider] `json:"provider"`
}

func (r SippyUpdateParamsR2EnableSippyGcsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsR2EnableSippyGcsSourceProvider string

const (
	SippyUpdateParamsR2EnableSippyGcsSourceProviderGcs SippyUpdateParamsR2EnableSippyGcsSourceProvider = "gcs"
)

func (r SippyUpdateParamsR2EnableSippyGcsSourceProvider) IsKnown() bool {
	switch r {
	case SippyUpdateParamsR2EnableSippyGcsSourceProviderGcs:
		return true
	}
	return false
}

type SippyUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []string                                                  `json:"messages,required"`
	Result   R2Sippy                                                   `json:"result,required"`
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

// Whether the API call was successful
type SippyUpdateResponseEnvelopeSuccess bool

const (
	SippyUpdateResponseEnvelopeSuccessTrue SippyUpdateResponseEnvelopeSuccess = true
)

func (r SippyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SippyDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []string                                                  `json:"messages,required"`
	Result   SippyDeleteResponse                                       `json:"result,required"`
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

// Whether the API call was successful
type SippyDeleteResponseEnvelopeSuccess bool

const (
	SippyDeleteResponseEnvelopeSuccessTrue SippyDeleteResponseEnvelopeSuccess = true
)

func (r SippyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SippyGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []string                                                  `json:"messages,required"`
	Result   R2Sippy                                                   `json:"result,required"`
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

// Whether the API call was successful
type SippyGetResponseEnvelopeSuccess bool

const (
	SippyGetResponseEnvelopeSuccessTrue SippyGetResponseEnvelopeSuccess = true
)

func (r SippyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
